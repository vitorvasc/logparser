package domain

import "logparser/internal/config/defines"

type Kill struct {
	KillerID   int
	KillerName string
	TargetID   int
	TargetName string
	Weapon     Weapon
	WeaponID   int
}

func NewKill(killerID int, killerName string, targetID int, targetName string, weapon Weapon, weaponID int) *Kill {
	return &Kill{
		KillerID:   killerID,
		KillerName: killerName,
		TargetID:   targetID,
		TargetName: targetName,
		Weapon:     weapon,
		WeaponID:   weaponID,
	}
}

func (k *Kill) KillerEqualsWorld() bool {
	return k.KillerName == defines.WorldPlayerName
}

func (k *Kill) KillerEqualsTarget() bool {
	return k.KillerID == k.TargetID
}
