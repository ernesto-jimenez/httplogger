package httplogger

import (
	"net/http"
	"time"
)

type loggedRoundTripper struct {
	rt  http.RoundTripper
	log HTTPLogger
}

func (c *loggedRoundTripper) RoundTrip(request *http.Request) (*http.Response, error) {
	c.log.LogRequest(request)
	startTime := time.Now()
	response, err := c.rt.RoundTrip(request)
	duration := time.Since(startTime)
	c.log.LogResponse(response, err, duration)
	return response, err
}

// NewLoggedTransport takes an http.RoundTripper and returns a new one that caches requests and responses
func NewLoggedTransport(rt http.RoundTripper, log HTTPLogger) http.RoundTripper {
	return &loggedRoundTripper{rt: rt, log: log}
}

// HTTPLogger defines the interface to log http request and responses
type HTTPLogger interface {
	LogRequest(*http.Request)
	LogResponse(response *http.Response, err error, duration time.Duration)
}
