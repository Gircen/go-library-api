package log

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

func (level Level) String() string {
	return stateLevel[level]
}
