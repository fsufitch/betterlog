package main

import (
	"fmt"
	"net/http"

	"github.com/fsufitch/betterlog"
)

type requestLog struct {
	Method string `json:"method"`
	Path   string `json:"path"`
}

type ExampleHTTPServer struct {
	logger *betterlog.LogEntryEmitter[any]
}

func (server ExampleHTTPServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	server.logger.InfoDataf("got request", requestLog{Method: r.Method, Path: r.URL.Path})

	if r.URL.Path == "/error" {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(400)
		fmt.Fprintln(w, "You asked for an error and you got it!")
		server.logger.Errorf("uh oh, the user asked for an error")
		return
	}

	if r.URL.Path == "/critical" {
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(500)
		fmt.Fprintln(w, "Oh no, this is a server crash! See the logs for a traceback, if DEBUG is set.")
		server.logger.Criticalf("oh no, a crash")
		return
	}

	server.logger.Debugf("starting plain success response")

	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(200)

	fmt.Fprintln(w, "<html><head><title>Betterlog Example</title></head>")
	fmt.Fprintln(w, "<body>")
	fmt.Fprintln(w, "<h1>It works!</h1>")
	fmt.Fprintf(w, "<p>You queried: <code>%s %s</code></p>\n", r.Method, r.URL.Path)
	fmt.Fprintln(w, "<p>To see other behavior, query <code>/error</code>, or <code>/critical</code>")
	fmt.Fprintln(w, "</body></html>")

	server.logger.Debugf("plain success response complete")

}
