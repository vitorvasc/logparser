package dto

import (
	"strings"

	"logparser/internal/config/defines"
)

const (
	InitGamePrefix              = "InitGame:"
	KillPrefix                  = "Kill:"
	ShutdownPrefix              = "ShutdownGame:"
	ClientUserInfoChangedPrefix = "ClientUserinfoChanged:"
)

var ValidPrefixes = []string{
	InitGamePrefix,
	KillPrefix,
	ShutdownPrefix,
	ClientUserInfoChangedPrefix,
}

var ValidPrefixesMap = map[string]string{
	InitGamePrefix:              defines.LogTypeStartMatch,
	KillPrefix:                  defines.LogTypeKill,
	ShutdownPrefix:              defines.LogTypeEndMatch,
	ClientUserInfoChangedPrefix: defines.LogTypeClientUserInfoChanged,
}

type LogEntry string

func (e *LogEntry) IsValid() bool {
	str := string(*e)
	for _, prefix := range ValidPrefixes {
		if strings.Contains(str, prefix) {
			return true
		}
	}
	return false
}

func (e *LogEntry) GetType() string {
	str := string(*e)
	for prefix, logType := range ValidPrefixesMap {
		if strings.Contains(str, prefix) {
			return logType
		}
	}
	return defines.LogTypeUnknown
}

func (e *LogEntry) IsGameInitialization() bool {
	return strings.Contains(string(*e), InitGamePrefix)
}
