package dto

type MatchDetails struct {
	ID         string          `json:"id"`
	TotalKills int             `json:"total_kills"`
	Players    []Player        `json:"players"`
	Kills      map[string]Kill `json:"kills"`
}

type Player string

type Kill int
