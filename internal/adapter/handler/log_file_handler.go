package handler

import (
	"bufio"
	"log"
	"os"
	"strconv"

	"logparser/internal/adapter/dto"
	"logparser/internal/adapter/mapper"
	"logparser/internal/core/domain"
	"logparser/internal/core/port"
)

type Handler struct {
	service port.CreateMatchHistoryService
}

func NewLogFileHandler(processMatchHistoryService port.CreateMatchHistoryService) Handler {
	return Handler{
		service: processMatchHistoryService,
	}
}

func (h Handler) CreateMatchesFromLogFile(source *os.File) *dto.ProcessResult {
	logEntriesByMatchID := h.filterAndGroupLogEntriesByMatchID(source)

	matchHistoryList := make([]*domain.MatchHistory, 0, len(logEntriesByMatchID))

	for matchID, entries := range logEntriesByMatchID {
		matchHistory := mapper.FromLogEntriesDtoToMatchHistory(matchID, entries)
		matchHistoryList = append(matchHistoryList, matchHistory)
	}

	processedMatches := h.service.BulkCreate(matchHistoryList)

	return mapper.FromBulkCreationResultToProcessResult(processedMatches)
}

func (h Handler) filterAndGroupLogEntriesByMatchID(source *os.File) map[string][]dto.LogEntry {
	currentMatchID := 0
	matchesMap := make(map[string][]dto.LogEntry)

	scanner := bufio.NewScanner(source)
	for scanner.Scan() {
		entry := dto.LogEntry(scanner.Text())

		if entry.IsGameInitialization() {
			currentMatchID++
		}

		if entry.IsValid() {
			id := h.formatID(currentMatchID)
			matchesMap[id] = append(matchesMap[id], entry)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return matchesMap
}

func (h Handler) formatID(gameID int) string {
	return "game_" + strconv.Itoa(gameID)
}
