package betterlog

import (
	"fmt"
	"time"
)

type Entry[T any] struct {
	level         Level
	time          time.Time
	fmt           string
	fmtArgs       []any
	trace         Traceback
	traceOverflow bool
	data          *T
	messageCached *string
}

func NewEntry[T any](level Level, time time.Time, format string, formatArgs []any, data *T, trace Traceback, tbOverflow bool) Entry[T] {
	return Entry[T]{
		level:   level,
		time:    time,
		fmt:     format,
		fmtArgs: formatArgs,
		data:    data,
		trace:   trace,
	}
}

func NewEntryFromFormat[T any](level Level, format string, formatArgs []any) Entry[T] {
	ts := time.Now()
	tb, overflow := GetTraceback(0, 1)
	return NewEntry[T](level, ts, format, formatArgs, nil, tb, overflow)
}

func NewEntryFromData[T any](level Level, data *T) Entry[T] {
	ts := time.Now()
	tb, overflow := GetTraceback(0, 1)
	return NewEntry(level, ts, "", nil, data, tb, overflow)
}

func (e Entry[T]) Level() Level {
	return e.level
}

func (e Entry[T]) Time() time.Time {
	return e.time
}

func (e Entry[T]) Format() string {
	return e.fmt
}

func (e Entry[T]) FormatArgs() []any {
	return append([]any{}, e.fmtArgs...)
}

func (e Entry[T]) Message() string {
	if e.messageCached == nil {
		msg := fmt.Sprintf(e.fmt, e.fmtArgs...)
		e.messageCached = &msg
	}
	return *e.messageCached
}

func (e Entry[T]) Traceback() Traceback {
	return append(Traceback{}, e.trace...)
}

func (e Entry[T]) TracebackOverflow() bool {
	return e.traceOverflow
}

func (e *Entry[T]) ResetTraceback(offset int, max int) {
	e.trace, e.traceOverflow = GetTraceback(offset, max)
}

func (e Entry[T]) Data() *T {
	return e.data
}
