package httplogger

import (
	"log"
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

// DefaultLogger is an http logger that will use the standard logger in the log package to provide basic information about http responses
type DefaultLogger struct {
}

// LogRequest doens't do anything since we'll be logging replies only
func (dl DefaultLogger) LogRequest(*http.Request) {
}

// LogResponse logs path, host, status code and duration in milliseconds
func (dl DefaultLogger) LogResponse(res *http.Response, err error, duration time.Duration) {
	duration /= time.Millisecond
	log.Printf("HTTP Request host=%s method=%s status=%d durationMs=%d", res.Request.Host, res.Request.Method, res.StatusCode, duration)
}

// DefaultLoggedTransport wraps http.DefaultTransport to log using DefaultLogger
var DefaultLoggedTransport = NewLoggedTransport(http.DefaultTransport, DefaultLogger{})
