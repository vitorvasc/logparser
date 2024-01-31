package service

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"sync"

	"logparser/internal/config/defines"
	"logparser/internal/core/domain"
	"logparser/internal/core/errors"
	"logparser/internal/core/port"
)

type createMatchHistoryService struct {
	repository port.MatchRepository
}

func NewCreateMatchHistoryService(matchRepository port.MatchRepository) port.CreateMatchHistoryService {
	return &createMatchHistoryService{
		repository: matchRepository,
	}
}

func (service createMatchHistoryService) BulkCreate(matchHistoryList []*domain.MatchHistory) []*domain.BulkCreationResult {
	totalMatches := len(matchHistoryList)
	responseChannel := make(chan *domain.BulkCreationResult, totalMatches)
	for _, matchHistory := range matchHistoryList {
		go service.processRoutine(matchHistory, responseChannel)
	}

	result := make([]*domain.BulkCreationResult, totalMatches)
	for i := 0; i < totalMatches; i++ {
		result[i] = <-responseChannel
	}

	log.Printf("[INFO] Processed matches: %d", len(result))
	close(responseChannel)

	return result
}

func (service createMatchHistoryService) Create(matchHistory *domain.MatchHistory) (*domain.Match, error) {
	match := domain.NewMatch(matchHistory.ID)
	for _, logEntry := range matchHistory.Logs {
		switch logEntry.Type {
		case defines.LogTypeClientUserInfoChanged:
			player, err := service.getPlayerInfo(logEntry.Value)
			if err != nil {
				return nil, err
			}
			match.InsertOrUpdatePlayer(player)
		case defines.LogTypeKill:
			kill, err := service.getKillInfo(logEntry.Value)
			if err != nil {
				return nil, err
			}
			match.NoticeKill(kill)
		case defines.LogTypeStartMatch, defines.LogTypeEndMatch:
			// currently, any of these infos are being used.
			continue
		default:
			continue
		}
	}

	err := service.repository.SaveMatch(match)
	if err != nil {
		return nil, err
	}

	return match, nil
}

func (service createMatchHistoryService) getPlayerInfo(rawLog string) (*domain.Player, error) {
	playerInfo := regexp.MustCompile(defines.GetPlayerInfoRegex)
	if match := playerInfo.FindStringSubmatch(rawLog); len(match) > 0 {
		playerID, _ := strconv.Atoi(match[1])
		playerName := match[2]

		return domain.NewPlayer(playerID, playerName), nil
	}

	errorMsg := fmt.Sprintf("error parsing player info: %s", rawLog)
	log.Println(errorMsg)
	return nil, errors.NewError(defines.GetPlayerInfoErrorCode, errorMsg)
}

func (service createMatchHistoryService) getKillInfo(rawLog string) (*domain.Kill, error) {
	killInfo := regexp.MustCompile(defines.GetKillInfoRegex)
	if match := killInfo.FindStringSubmatch(rawLog); len(match) > 0 {
		killerID, _ := strconv.Atoi(match[1])
		targetID, _ := strconv.Atoi(match[2])
		weaponID, _ := strconv.Atoi(match[3])
		killerName, targetName, weaponName := match[4], match[5], match[6]
		return domain.NewKill(killerID, killerName, targetID, targetName, domain.Weapon(weaponName), weaponID), nil
	}

	errorMsg := fmt.Sprintf("error parsing kill info: %s", rawLog)
	log.Println(errorMsg)
	return nil, errors.NewError(defines.GetKillInfoErrorCode, errorMsg)
}

func (service createMatchHistoryService) processRoutine(matchHistory *domain.MatchHistory, responseChannel chan *domain.BulkCreationResult) {
	var mutex sync.Mutex
	mutex.Lock()
	defer mutex.Unlock()

	result := &domain.BulkCreationResult{
		MatchID: matchHistory.ID,
		Success: true,
	}

	_, err := service.Create(matchHistory)
	if err != nil {
		result.Success = false
		result.ErrorMessage = err.Error()
	}

	responseChannel <- result
}
