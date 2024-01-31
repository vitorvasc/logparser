package dto

type MatchDetails struct {
	TotalKills   int            `json:"total_kills"`
	Players      []Player       `json:"players"`
	Kills        map[string]int `json:"kills"`
	KillsByMeans map[string]int `json:"kills_by_means,omitempty"`
}

type Player string
