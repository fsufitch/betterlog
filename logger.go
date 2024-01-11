package betterlog

type Logger[T any] interface {
	Close() error
	Emit(entry Entry[T]) error
}
