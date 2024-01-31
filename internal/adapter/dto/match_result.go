package dto

type Match struct {
	ID         string `json:"id"`
	TotalKills int    `json:"total_kills,omitempty"`
}

type Players struct {
}
