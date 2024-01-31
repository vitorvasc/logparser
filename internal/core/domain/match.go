package domain

type Match struct {
	ID          string
	TotalKills  int
	Players     []*Player
	KillHistory []*Kill
}

type MatchHistory struct {
	ID   string
	Logs []*Log
}

func NewMatch(id string) *Match {
	return &Match{
		ID:          id,
		TotalKills:  0,
		Players:     make([]*Player, 0),
		KillHistory: make([]*Kill, 0),
	}
}

func (m *Match) InsertOrUpdatePlayer(player *Player) {
	// if player already exists, update it
	for _, p := range m.Players {
		if p.ID == player.ID {
			p.ChangeName(player.Name)
			return
		}
	}
	// if player doesn't exist, add it
	m.Players = append(m.Players, player)
}

func (m *Match) NoticeKill(kill *Kill) {
	// increase match total kills and insert kill into history
	m.TotalKills++
	m.KillHistory = append(m.KillHistory, kill)

	// increase target deaths count
	target := m.GetPlayerByID(kill.TargetID)
	if target != nil {
		target.AddDeath()
	}

	// if kill was made by world OR player killed itself, we should end by here
	if kill.KillerEqualsWorld() || kill.KillerEqualsTarget() {
		return
	}

	// increase killer kills count
	killer := m.GetPlayerByID(kill.KillerID)
	if killer != nil {
		killer.AddKill()
	}
}

func (m *Match) GetPlayerByID(id int) *Player {
	// if player already exists, return it
	for _, player := range m.Players {
		if player.ID == id {
			return player
		}
	}
	// if player doesn't exist, return nil
	return nil
}
