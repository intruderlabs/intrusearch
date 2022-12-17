package requests

import (
	"context"
	"github.com/opensearch-project/opensearch-go/opensearchapi"
	"net/http"
	"strings"
)

type SavedObjectsGetIndexPatternRequest struct {
}

func (r SavedObjectsGetIndexPatternRequest) Do(
	ctx context.Context,
	transport opensearchapi.Transport,
) (*opensearchapi.Response, error) {
	var (
		method string
		path   strings.Builder
	)

	method = "GET"

	// TODO: this is a OpenSearch Dashboards API. Without it, it doesn't work
	relative := "_dashboards/api/saved_objects/_find?type=index-pattern"
	path.Grow(1 + len(relative))
	path.WriteString("/")
	path.WriteString(relative)

	req, err := http.NewRequest(method, path.String(), nil)
	if err != nil {
		return nil, err
	}

	req.Header["osd-xsrf"] = []string{"true"}

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
