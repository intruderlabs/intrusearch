package requests

import (
	"context"
	"github.com/opensearch-project/opensearch-go/opensearchapi"
	"gitlab.com/intruderlabs/toolbox/intrusearch.git/main/domain/errors"
	"gitlab.com/intruderlabs/toolbox/intrusearch.git/main/domain/helpers"
	domain "gitlab.com/intruderlabs/toolbox/intrusearch.git/main/domain/responses"
	infra "gitlab.com/intruderlabs/toolbox/intrusearch.git/main/infrastructure/responses"
	"net/http"
)

const headerContentType = "Content-Type"

var headerContentTypeJSON = []string{"application/json"}

func DoRequest(transport opensearchapi.Transport, request opensearchapi.Request) (domain.GenericResponse, []errors.GenericError) {
	response, err := request.Do(context.Background(), transport)

	acceptedCodes := map[int]bool{
		http.StatusOK:        true,
		http.StatusCreated:   true,
		http.StatusNoContent: true,
	}

	return mapFromRequestError(domain.GenericResponse{
		err == nil && acceptedCodes[response.StatusCode],
		response.StatusCode,
		response.Body,
	}, err)
}

func mapFromRequestError(wrapper domain.GenericResponse, err error) (domain.GenericResponse, []errors.GenericError) {
	var mapped []errors.GenericError

	if err != nil {
		mapped = append(mapped, errors.GenericError{"http_error", err.Error()})
	} else {
		errorResponse := infra.GenericErrorResponse{}
		helpers.NewSerializationHelper().Deserialize(wrapper.Body, &errorResponse)

		if errorResponse.Status != 0 {
			if errorResponse.Error.Reason != "" {
				mapped = append(mapped, errors.GenericError{errorResponse.Error.Type, errorResponse.Error.Reason})
			} else {
				if len(errorResponse.Error.RootCause) > 0 {
					for _, value := range errorResponse.Error.RootCause {
						if value.Reason != "" {
							mapped = append(mapped, errors.GenericError{value.Type, value.Reason})
						}
					}
				}
			}
		}
	}

	return wrapper, mapped
}
