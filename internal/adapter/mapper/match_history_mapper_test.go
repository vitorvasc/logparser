package mapper

import (
	"reflect"
	"testing"

	"logparser/internal/adapter/dto"
	"logparser/internal/config/defines"
	"logparser/internal/core/domain"
)

func TestFromLogEntriesDtoToMatchHistory(t *testing.T) {
	testCases := []struct {
		name            string
		inputMatchID    string
		inputLogEntries []dto.LogEntry
		expected        *domain.MatchHistory
	}{
		{
			name:         "given valid match id and slice of log entries, should create valid match history",
			inputMatchID: "1",
			inputLogEntries: []dto.LogEntry{
				"0:00 InitGame: \\sv_floodProtect\\1\\sv_maxPing\\0\\sv_minPing\\0\\sv_maxRate\\10000\\sv_minRate\\0\\sv_hostname\\Code Miner Server\\g_gametype\\0\\sv_privateClients\\2\\sv_maxclients\\16\\sv_allowDownload\\0\\dmflags\\0\\fraglimit\\20\\timelimit\\15\\g_maxGameClients\\0\\capturelimit\\8\\version\\ioq3 1.36 linux-x86_64 Apr 12 2009\\protocol\\68\\mapname\\q3dm17\\gamename\\baseq3\\g_needpass\\0",
				"20:34 ClientUserinfoChanged: 2 n\\Isgalamido\\t\\0\\model\\xian/default\\hmodel\\xian/default\\g_redteam\\\\g_blueteam\\\\c1\\4\\c2\\5\\hc\\100\\w\\0\\l\\0\\tt\\0\\tl\\0",
				"20:37 ShutdownGame:",
			},
			expected: &domain.MatchHistory{
				ID: "1",
				Logs: []*domain.Log{
					{
						Type:  defines.LogTypeStartMatch,
						Value: "0:00 InitGame: \\sv_floodProtect\\1\\sv_maxPing\\0\\sv_minPing\\0\\sv_maxRate\\10000\\sv_minRate\\0\\sv_hostname\\Code Miner Server\\g_gametype\\0\\sv_privateClients\\2\\sv_maxclients\\16\\sv_allowDownload\\0\\dmflags\\0\\fraglimit\\20\\timelimit\\15\\g_maxGameClients\\0\\capturelimit\\8\\version\\ioq3 1.36 linux-x86_64 Apr 12 2009\\protocol\\68\\mapname\\q3dm17\\gamename\\baseq3\\g_needpass\\0",
					},
					{
						Type:  defines.LogTypeClientUserInfoChanged,
						Value: "20:34 ClientUserinfoChanged: 2 n\\Isgalamido\\t\\0\\model\\xian/default\\hmodel\\xian/default\\g_redteam\\\\g_blueteam\\\\c1\\4\\c2\\5\\hc\\100\\w\\0\\l\\0\\tt\\0\\tl\\0",
					},
					{
						Type:  defines.LogTypeEndMatch,
						Value: "20:37 ShutdownGame:",
					},
				},
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			result := FromLogEntriesDtoToMatchHistory(testCase.inputMatchID, testCase.inputLogEntries)
			if !reflect.DeepEqual(result, testCase.expected) {
				t.Errorf("expected result %v, got %v", testCase.expected, result)
			}
		})
	}
}

func TestFromBulkCreationResultToProcessResult(t *testing.T) {
	testCases := []struct {
		name     string
		input    []*domain.BulkCreationResult
		expected *dto.ProcessResult
	}{
		{
			name: "given valid bulk creation result, should convert into process result",
			input: []*domain.BulkCreationResult{
				{
					MatchID:      "1",
					Success:      true,
					ErrorMessage: "",
				},
				{
					MatchID:      "2",
					Success:      false,
					ErrorMessage: "error parsing player info: INVALID_LOG",
				},
				{
					MatchID:      "4",
					Success:      true,
					ErrorMessage: "",
				},
			},
			expected: &dto.ProcessResult{
				TotalProcessedMatches: 3,
				Failures: []*dto.ProcessFailure{
					{
						MatchID: "2",
						Reason:  "error parsing player info: INVALID_LOG",
					},
				},
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			result := FromBulkCreationResultToProcessResult(testCase.input)
			if !reflect.DeepEqual(result, testCase.expected) {
				t.Errorf("expected result %v, got %v", testCase.expected, result)
			}
		})
	}
}
