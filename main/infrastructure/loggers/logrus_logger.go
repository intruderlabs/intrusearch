package loggers

import (
	logger "github.com/sirupsen/logrus"
	"net/http"
	"net/url"
	"time"
)

type LogrusLogger struct {
	isColored          bool
	EnableRequestBody  bool
	EnableResponseBody bool
}

func NewLogrusLogger(isColored bool) *LogrusLogger {
	return &LogrusLogger{isColored: isColored}
}

func (itself LogrusLogger) RequestBodyEnabled() bool {
	return itself.EnableRequestBody
}

func (itself LogrusLogger) ResponseBodyEnabled() bool {
	return itself.EnableResponseBody
}

func (itself LogrusLogger) LogRoundTrip(req *http.Request, res *http.Response, err error, start time.Time, dur time.Duration) error {
	query, _ := url.QueryUnescape(req.URL.RawQuery)
	if query != "" {
		query = "?" + query
	}

	var (
		color  string
		status = res.Status
	)

	switch {
	case res.StatusCode > 0 && res.StatusCode < 300:
		color = "\x1b[32m"
	case res.StatusCode > 299 && res.StatusCode < 500:
		color = "\x1b[33m"
	case res.StatusCode > 499:
		color = "\x1b[31m"
	default:
		status = "ERROR"
		color = "\x1b[31;4m"
	}

	if itself.isColored {
		logger.Infof("%s \x1b[1;4m%s://%s%s\x1b[0m%s %s%s\x1b[0m \x1b[2m%s\x1b[0m\n",
			req.Method,
			req.URL.Scheme,
			req.URL.Host,
			req.URL.Path,
			query,
			color,
			status,
			dur.Truncate(time.Millisecond),
		)
	} else {
		logger.Infof("%s %s://%s%s%s %s %s",
			req.Method,
			req.URL.Scheme,
			req.URL.Host,
			req.URL.Path,
			query,
			status,
			dur.Truncate(time.Millisecond),
		)
	}

	return nil
}
