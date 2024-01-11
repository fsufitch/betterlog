package betterlog

import (
	"encoding/json"
	"errors"
	"strings"
)

type DataToJSON[T any] struct {
	Child Logger[string]
}

func (dtj DataToJSON[T]) Close() error {
	return dtj.Child.Close()
}

func (dtj DataToJSON[T]) Emit(entry Entry[T]) error {

	stringifiedJSON := ""

	if entry.Data() != nil {
		buf := &strings.Builder{}
		err := json.NewEncoder(buf).Encode(entry.Data())
		if err != nil {
			return errors.Join(LoggingError("json encode failed"), err)
		}
		stringifiedJSON = strings.TrimSpace(buf.String())
	}

	newEntry := NewEntry[string](
		entry.Level(),
		entry.Time(),
		entry.Format(),
		entry.FormatArgs(),
		&stringifiedJSON,
		entry.Traceback(),
		entry.TracebackOverflow(),
	)

	if dtj.Child == nil {
		return LoggingError("tried to log to an empty child; did you initialize your logging right?")
	}

	return dtj.Child.Emit(newEntry)
}
