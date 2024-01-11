package main

import (
	"io"
	"net/http"
	"os"

	"github.com/fsufitch/betterlog"
)

func main() {
	logDestinations := []io.WriteCloser{os.Stderr}
	if logFilePath := os.Getenv("LOG_FILE"); logFilePath != "" {
		fp, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			panic(err)
		}
		defer fp.Close()
		logDestinations = append(logDestinations, fp)
	}

	_, debugMode := os.LookupEnv("DEBUG")

	logger := betterlog.SimpleLogging(logDestinations, debugMode)
	logger.Debugf("started logging, with debug mode=%v", debugMode)

	server := &ExampleHTTPServer{logger: logger}

	logger.Infof("serving on localhost:9001")
	err := http.ListenAndServe("localhost:9001", server)
	logger.Warningf("server exited with error: %s", err)

	defer func() {
		// Handle any panics
		recover()
		logger.Criticalf("recovered from a panic crash")
	}()
}
