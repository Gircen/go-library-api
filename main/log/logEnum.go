package log

import "strings"

type Level int

const (
	INFO Level = iota
	DEBUG
	WARN
	ERROR
)

var stateLevel = map[Level]string{
	INFO:  "INFO",
	DEBUG: "DEBUG",
	WARN:  "WARN",
	ERROR: "ERROR",
}

func valueLevel(state string) Level {
	switch strings.ToUpper(state) {
	case INFO.String():
		return INFO
	case DEBUG.String():
		return DEBUG
	case WARN.String():
		return WARN
	case ERROR.String():
		return ERROR
	default:
		return ERROR
	}

}

func (level Level) IsLog(state string) bool {
	return level >= valueLevel(state)
}

func (level Level) toInt() int {
	return int(level)
}

func (level Level) String() string {
	return stateLevel[level]
}
