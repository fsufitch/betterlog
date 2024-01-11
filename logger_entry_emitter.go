package betterlog

import (
	"time"
)

type LogEntryEmitter[T any] struct {
	Child            Logger[T]
	DisableTraceback bool
}

func (emt LogEntryEmitter[T]) Close() error {
	return emt.Child.Close()
}

func (emt LogEntryEmitter[T]) Emit(entry Entry[T]) error {
	if emt.Child == nil {
		return LoggingError("tried to log to an empty child; did you initialize your logging right?")
	}
	return emt.Child.Emit(entry)
}

// log is common logic for the other utility functions
func (emt LogEntryEmitter[T]) log(level Level, format string, formatArgs []any, data *T) error {
	ts := time.Now()
	tbMax := 0
	if emt.DisableTraceback {
		tbMax = 1
	}
	tb, overflow := GetTraceback(1, tbMax)
	return emt.Emit(NewEntry[T](level, ts, format, formatArgs, data, tb, overflow))
}

func (emt LogEntryEmitter[T]) Emitf(level Level, format string, args ...any) error {
	return emt.log(level, format, args, nil)
}

func (emt LogEntryEmitter[T]) Debugf(format string, args ...any) error {
	return emt.log(LevelDebug, format, args, nil)
}

func (emt LogEntryEmitter[T]) Infof(format string, args ...any) error {
	return emt.log(LevelInfo, format, args, nil)
}

func (emt LogEntryEmitter[T]) Warningf(format string, args ...any) error {
	return emt.log(LevelWarning, format, args, nil)
}

func (emt LogEntryEmitter[T]) Errorf(format string, args ...any) error {
	return emt.log(LevelError, format, args, nil)
}

func (emt LogEntryEmitter[T]) Criticalf(format string, args ...any) error {
	return emt.log(LevelCritical, format, args, nil)
}

func (emt LogEntryEmitter[T]) EmitDataf(level Level, format string, data T, args ...any) error {
	return emt.log(level, format, args, &data)
}

func (emt LogEntryEmitter[T]) DebugDataf(format string, data T, args ...any) error {
	return emt.log(LevelDebug, format, args, &data)
}

func (emt LogEntryEmitter[T]) InfoDataf(format string, data T, args ...any) error {
	return emt.log(LevelInfo, format, args, &data)
}

func (emt LogEntryEmitter[T]) WarningDataf(format string, data T, args ...any) error {
	return emt.log(LevelWarning, format, args, &data)
}

func (emt LogEntryEmitter[T]) ErrorDataf(format string, data T, args ...any) error {
	return emt.log(LevelError, format, args, &data)
}

func (emt LogEntryEmitter[T]) CriticalDataf(format string, data T, args ...any) error {
	return emt.log(LevelCritical, format, args, &data)
}
