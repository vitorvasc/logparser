package domain

import "reflect"

type Player struct {
	ID          int
	Name        string
	NameHistory []string
	Kills       int
	Deaths      int
}

func NewPlayer(id int, name string) *Player {
	return &Player{
		ID:          id,
		Name:        name,
		NameHistory: make([]string, 0),
		Kills:       0,
		Deaths:      0,
	}
}

func (p *Player) ChangeName(name string) {
	if p.Name == name {
		return
	}
	p.NameHistory = append(p.NameHistory, p.Name)
	p.Name = name
}

func (p *Player) AddKill() {
	p.Kills++
}

func (p *Player) AddDeath() {
	p.Deaths++
}

func (p *Player) Equals(target *Player) bool {
	return reflect.DeepEqual(p, target)
}
