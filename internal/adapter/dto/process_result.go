package dto

type ProcessResult struct {
	TotalProcessedMatches int               `json:"total_processed_matches"`
	Failures              []*ProcessFailure `json:"failures,omitempty"`
}

type ProcessFailure struct {
	MatchID string `json:"match_id"`
	Reason  string `json:"reason"`
}
