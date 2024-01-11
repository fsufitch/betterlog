package betterlog

type LevelFilter[T any] struct {
	MinLevel Level
	MaxLevel Level
	Child    Logger[T]
}

func (l LevelFilter[T]) Close() error {
	return l.Child.Close()
}

func (l LevelFilter[T]) Emit(entry Entry[T]) error {
	if l.MinLevel != nil && (entry.Level().Value() < l.MinLevel.Value()) {
		return nil
	}
	if l.MaxLevel != nil && (entry.Level().Value() > l.MaxLevel.Value()) {
		return nil
	}
	if l.Child == nil {
		return LoggingError("tried to log to an empty child; did you initialize your logging right?")
	}
	return l.Child.Emit(entry)
}
