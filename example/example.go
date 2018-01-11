// Package main provides a sample example to use the httplogger API
// This example can accept a URl as input. If not, it assumes a default Url and logs Req/Resp for that Url
//          The Command line usage guide. Copyright (c) Mohd Anwar Jamal Faiz
//          Toughjamy@yahoo.com

package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"httplogger"
)

func main() {
	client := http.Client{
		Transport: httplogger.NewLoggedTransport(http.DefaultTransport, newLogger()),
	}

	argsWithoutProg := os.Args[1:]
	if len(argsWithoutProg)==1 {
			client.Get(argsWithoutProg[0])
		} else if len(argsWithoutProg)>1 {
			println("just one parameter required")
		} else {
			defaultUrl := "http://www.w3lc.com"
			client.Get(defaultUrl)
		}
}

type httpLogger struct {
	log *log.Logger
}

func newLogger() *httpLogger {
	return &httpLogger{
		log: log.New(os.Stderr, "log - ", log.LstdFlags),
	}
}

func (l *httpLogger) LogRequest(req *http.Request) {
	l.log.Printf(
		"Request %s %s",
		req.Method,
		req.URL.String(),
	)
}

func (l *httpLogger) LogResponse(req *http.Request, res *http.Response, err error, duration time.Duration) {
	duration /= time.Millisecond
	l.log.Printf(
		"Response method=%s status=%d durationMs=%d %s",
		req.Method,
		res.StatusCode,
		duration,
		req.URL.String(),
	)
}
