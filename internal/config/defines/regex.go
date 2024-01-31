package defines

const (
	GetPlayerInfoRegex = `ClientUserinfoChanged: (?P<id>\d+) n\\(?P<name>.+?)\\`
	GetKillInfoRegex   = `Kill: (?P<killer_id>\d+) (?P<killed_id>\d+) (?P<weapon_id>\d+): (?P<killer_name>.+?) killed (?P<killed_name>.+?) by (?P<weapon_name>.+)$`
)
