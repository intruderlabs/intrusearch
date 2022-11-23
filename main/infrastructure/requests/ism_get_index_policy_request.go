package requests

import (
	"context"
	"github.com/opensearch-project/opensearch-go/opensearchapi"
	"net/http"
	"strings"
)

type IsmGetIndexPolicyRequest struct {
	Name string
}

func (r IsmGetIndexPolicyRequest) Do(
	ctx context.Context,
	transport opensearchapi.Transport,
) (*opensearchapi.Response, error) {
	var (
		method string
		path   strings.Builder
	)

	method = "GET"

	relative := "_plugins/_ism/policies"
	path.Grow(1 + len(relative) + 1 + len(r.Name))
	path.WriteString("/")
	path.WriteString(relative)
	path.WriteString("/")
	path.WriteString(r.Name)

	req, err := http.NewRequest(method, path.String(), nil)
	if err != nil {
		return nil, err
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
