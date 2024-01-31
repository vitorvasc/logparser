package mapper

import (
	"logparser/internal/adapter/dto"
	"logparser/internal/core/domain"
)

func FromLogEntriesDtoToMatchHistory(matchID string, logEntries []dto.LogEntry) *domain.MatchHistory {
	matchHistory := new(domain.MatchHistory)
	matchHistory.ID = matchID

	for _, entry := range logEntries {
		matchHistory.Logs = append(matchHistory.Logs, &domain.Log{
			Type:  entry.GetType(),
			Value: string(entry),
		})
	}

	return matchHistory
}

func FromBulkCreationResultToProcessResult(resultList []*domain.BulkCreationResult) *dto.ProcessResult {
	var processResult = new(dto.ProcessResult)
	processResult.TotalProcessedMatches = len(resultList)
	processResult.Failures = make([]*dto.ProcessFailure, 0)

	for _, result := range resultList {
		if !result.Success {
			processResult.Failures = append(processResult.Failures, &dto.ProcessFailure{
				MatchID: result.MatchID,
				Reason:  result.ErrorMessage,
			})
		}
	}

	return processResult
}
