package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/ernesto-jimenez/httplogger"
)

func main() {
	client := http.Client{
		Transport: httplogger.NewLoggedTransport(http.DefaultTransport, newLogger()),
	}

	client.Get("http://google.com")
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

func (l *httpLogger) LogResponse(res *http.Response, err error, duration time.Duration) {
	duration /= time.Millisecond
	l.log.Printf(
		"Response method=%s status=%d durationMs=%d %s",
		res.Request.Method,
		res.StatusCode,
		duration,
		res.Request.URL.String(),
	)
}
