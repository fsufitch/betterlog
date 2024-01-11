package betterlog

import (
	"errors"
)

type Multiplexer[T any] []Logger[T]

func (mp Multiplexer[T]) Close() error {
	errs := []error{}
	for _, inner := range mp {
		if err := inner.Close(); err != nil {
			errs = append(errs, err)
		}
	}
	if len(errs) > 0 {
		joinErrs := append([]error{LoggingError("multiplex close failed")}, errs...)
		return errors.Join(joinErrs...)
	}
	return nil
}

func (mp Multiplexer[T]) Emit(entry Entry[T]) error {
	errs := []error{}
	for _, inner := range mp {
		if err := inner.Emit(entry); err != nil {
			errs = append(errs, err)
		}
	}
	if len(errs) > 0 {
		joinErrs := append([]error{LoggingError("multiplex emit failed")}, errs...)
		return errors.Join(joinErrs...)
	}
	return nil

}
