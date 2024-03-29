package handler

import (
	"bufio"
	"log"
	"os"

	"logparser/internal/adapter/dto"
	"logparser/internal/adapter/mapper"
	"logparser/internal/adapter/utils"
	"logparser/internal/core/domain"
	"logparser/internal/core/port"
)

type LogFileHandler struct {
	service port.CreateMatchService
}

func NewLogFileHandler(createMatchService port.CreateMatchService) LogFileHandler {
	return LogFileHandler{
		service: createMatchService,
	}
}

func (h LogFileHandler) CreateMatchesFromLogFile(source *os.File) *dto.ProcessResult {
	logEntriesByMatchID := h.filterAndGroupLogEntriesByMatchID(source)

	matchHistoryList := make([]*domain.MatchHistory, 0, len(logEntriesByMatchID))

	for matchID, entries := range logEntriesByMatchID {
		matchHistory := mapper.FromLogEntriesDtoToMatchHistory(matchID, entries)
		matchHistoryList = append(matchHistoryList, matchHistory)
	}

	processedMatches := h.service.BulkCreate(matchHistoryList)

	return mapper.FromBulkCreationResultToProcessResult(processedMatches)
}

func (h LogFileHandler) filterAndGroupLogEntriesByMatchID(source *os.File) map[string][]dto.LogEntry {
	currentMatchID := 0
	matchesMap := make(map[string][]dto.LogEntry)

	scanner := bufio.NewScanner(source)
	for scanner.Scan() {
		entry := dto.LogEntry(scanner.Text())

		if entry.IsGameInitialization() {
			currentMatchID++
		}

		if entry.IsValid() {
			id := utils.FormatMatchID(currentMatchID)
			matchesMap[id] = append(matchesMap[id], entry)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return matchesMap
}
