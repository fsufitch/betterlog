package betterlog

import (
	"io"
)

func SimpleLogging(outputs []io.WriteCloser, debug bool) *LogEntryEmitter[any] {
	outputChain := func(level Level, templateName string, includeCaller bool, includeTraceback bool) Logger[any] {
		template := BasicLogTemplate(templateName, includeCaller, includeTraceback)
		outputMux := Multiplexer[string]{}

		for _, writer := range outputs {
			outputMux = append(outputMux, TemplateWriter[string]{Writer: writer, Template: template})
		}

		return LevelFilter[any]{
			MinLevel: level,
			MaxLevel: level,
			Child:    DataToJSON[any]{Child: outputMux},
		}
	}

	mux := Multiplexer[any]{
		outputChain(LevelInfo, "info", debug, false),
		outputChain(LevelWarning, "warning", debug, false),
		outputChain(LevelError, "error", debug, false),
		outputChain(LevelCritical, "critical", true, true),
	}

	if debug {
		mux = append(mux, outputChain(LevelDebug, "debug", debug, false))
	}

	return &LogEntryEmitter[any]{Child: mux, DisableTraceback: !debug}
}
