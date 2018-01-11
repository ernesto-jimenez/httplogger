A GoLang library useful for logging API requests - HttpLogger

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

	"github.com/Anwar-Faiz/httplogger"
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

func (l *httpLogger) LogResponse(req *http.Request, res *http.Response, err error, duration time.Duration) {
	duration /= time.Millisecond
	if err != nil {
		l.log.Println(err)
	} else {
		l.log.Printf(
			"Response method=%s status=%d durationMs=%d %s",
			req.Method,
			res.StatusCode,
			duration,
			req.URL.String(),
		)
	}
}
```

Output:

```
% go run example/example.go
log - 2018/01/11 15:42:27 Request GET http://google.com
log - 2018/01/11 15:42:28 Response method=GET status=302 durationMs=530 http://google.com
log - 2018/01/11 15:42:28 Request GET http://www.google.co.in/?gfe_rd=cr&dcr=0&ei=izhXWoXZOq_SXvmFoNgF
log - 2018/01/11 15:42:28 Response method=GET status=200 durationMs=594 http://www.google.co.in/?gfe_rd=cr&dcr=0&ei=izhXWoXZOq_SXvmFoNgF
```

# LICENSE

Authored - 2017 - Anwar Faiz

Many parts of the Software after fork is written By Mohd Anwar Jamal Faiz

Permission is granted to use/copy/educate with or without mentioning the author of those code. Although giving back credits is hugely appreciated. 

Original Copyright Notice (c) 2015 Ernesto Jimenez. Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.


```

Different use cases of the executable:

```
Case 1:

$ example.exe
log - 2018/01/11 22:41:55 Request GET http://www.w3lc.com
log - 2018/01/11 22:41:57 Response method=GET status=200 durationMs=2070 http://www.w3lc.com


Case 2:
$ example.exe kk ll
just one parameter required


Case 3:
$ example.exe http://facebook.com
log - 2018/01/11 22:42:19 Request GET http://facebook.com
log - 2018/01/11 22:42:20 Response method=GET status=301 durationMs=684 http://facebook.com
log - 2018/01/11 22:42:20 Request GET https://facebook.com/
log - 2018/01/11 22:42:22 Response method=GET status=301 durationMs=1427 https://facebook.com/
log - 2018/01/11 22:42:22 Request GET https://www.facebook.com/
log - 2018/01/11 22:42:23 Response method=GET status=200 durationMs=1192 https://www.facebook.com/



Cheers ;)

Anwar Jamal Faiz

Toughjamy@yahoo.com

http://www.w3LC.com  [ W3LC : World Wide Web Learners Consortium ]