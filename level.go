package betterlog

import "fmt"

type Level interface {
	Value() uint8
	Name() string
}

type Uint8Level uint8

const (
	LevelUndefined Uint8Level = 0
	LevelDebug     Uint8Level = 10
	LevelInfo      Uint8Level = 20
	LevelWarning   Uint8Level = 30
	LevelError     Uint8Level = 40
	LevelCritical  Uint8Level = 50
)

var LevelNames map[Uint8Level]string = map[Uint8Level]string{
	LevelUndefined: "LEVELUNDEFINED",
	LevelDebug:     "DEBUG",
	LevelInfo:      "INFO",
	LevelWarning:   "WARNING",
	LevelError:     "ERROR",
	LevelCritical:  "CRITICAL",
}

func (level Uint8Level) Value() uint8 {
	return uint8(level)
}

func (level Uint8Level) Name() string {
	if name, ok := LevelNames[level]; ok {
		return name
	}
	return fmt.Sprintf("LEVEL-%d", level)
}
