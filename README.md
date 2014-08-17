Small golang library useful for logging API requests.

It wraps any http.Transport to log its requests and responses,
including the duration time.

# Usage

See [example/example.go](example/example.go)

```go
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
```

Output:

```
% go run example/example.go
log - 2014/08/17 02:19:19 Request GET http://google.com
log - 2014/08/17 02:19:19 Response method=GET status=302
durationMs=85 http://google.com
log - 2014/08/17 02:19:19 Request GET
http://www.google.co.uk/?gfe_rd=cr&ei=GwPwU4GtPMKo8we3koKwDg
log - 2014/08/17 02:19:20 Response method=GET status=200
durationMs=138
http://www.google.co.uk/?gfe_rd=cr&ei=GwPwU4GtPMKo8we3koKwDg
```
