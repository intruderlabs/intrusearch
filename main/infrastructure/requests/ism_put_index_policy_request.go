package requests

import (
	"context"
	"github.com/opensearch-project/opensearch-go/opensearchapi"
	"io"
	"net/http"
	"strings"
)

type IsmPutIndexPolicyRequest struct {
	Name string
	Body io.Reader
}

func (r IsmPutIndexPolicyRequest) Do(
	ctx context.Context,
	transport opensearchapi.Transport,
) (*opensearchapi.Response, error) {
	var (
		method string
		path   strings.Builder
	)

	method = "PUT"

	relative := "_plugins/_ism/policies"
	path.Grow(1 + len(relative) + 1 + len(r.Name))
	path.WriteString("/")
	path.WriteString(relative)
	path.WriteString("/")
	path.WriteString(r.Name)

	req, err := http.NewRequest(method, path.String(), r.Body)
	if err != nil {
		return nil, err
	}

	if r.Body != nil {
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
