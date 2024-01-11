package betterlog

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
	"text/template"
)

type TemplateWriter[T any] struct {
	Writer   io.WriteCloser
	Template *template.Template
	Silent   bool
}

func (logger TemplateWriter[T]) Close() error {
	return logger.Writer.Close()
}

func (logger TemplateWriter[T]) Emit(entry Entry[T]) error {
	if logger.Template == nil {
		return LoggingError("tried to render nil template; did you initialize your logging right?")
	}

	outText := strings.Builder{}
	err := logger.Template.Execute(&outText, entry)
	if err != nil {
		err = fmt.Errorf("template execution failed: %w", err)
		if !logger.Silent {
			fmt.Fprintf(os.Stderr, "WARNING: %s\n", err.Error())
		}
		return errors.Join(LoggingError("template execution failed"), err)
	}

	_, err = fmt.Fprintln(logger.Writer, outText.String())
	if err != nil {
		return errors.Join(LoggingError("write failed"), err)
	}
	return nil
}
