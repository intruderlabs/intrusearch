package intrusearch

import (
	"crypto/tls"
	"github.com/opensearch-project/opensearch-go"
	"gitlab.com/intruderlabs/toolbox/intrusearch/main/infrastructure/loggers"
	"net/http"
)

const maxRetries = 5

type Client struct {
	client *opensearch.Client
}

func NewClient(address string, coloredLogger bool) Client {
	client, _ := opensearch.NewClient(opensearch.Config{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
		Addresses:  []string{address},
		MaxRetries: maxRetries,
		RetryOnStatus: []int{
			http.StatusBadGateway,
			http.StatusServiceUnavailable,
			http.StatusGatewayTimeout,
		},
		Logger: loggers.NewLogrusLogger(coloredLogger),
	})
	return Client{client}
}
