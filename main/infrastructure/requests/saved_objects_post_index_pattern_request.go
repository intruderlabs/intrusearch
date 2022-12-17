package requests

import (
	"context"
	"github.com/opensearch-project/opensearch-go/opensearchapi"
	"io"
	"net/http"
	"strings"
)

type SavedObjectsPostIndexPatternRequest struct {
	Body io.Reader
}

func (r SavedObjectsPostIndexPatternRequest) Do(
	ctx context.Context,
	transport opensearchapi.Transport,
) (*opensearchapi.Response, error) {
	var (
		method string
		path   strings.Builder
	)

	method = "POST"

	// TODO: this is a OpenSearch Dashboards API. Without it, it doesn't work
	relative := "_dashboards/api/saved_objects/index-pattern"
	path.Grow(1 + len(relative))
	path.WriteString("/")
	path.WriteString(relative)

	req, err := http.NewRequest(method, path.String(), r.Body)
	if err != nil {
		return nil, err
	}

	if r.Body != nil {
		req.Header["osd-xsrf"] = []string{"true"}
		req.Header[headerContentType] = headerContentTypeJSON
	}

	if ctx != nil {
		req = req.WithContext(ctx)
	}

	res, err := transport.Perform(req)
	if err != nil {
		return nil, err
	}

	response := opensearchapi.Response{
		StatusCode: res.StatusCode,
		Body:       res.Body,
		Header:     res.Header,
	}

	return &response, nil
}
