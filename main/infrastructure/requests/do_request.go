package requests

import (
	"context"
	"fmt"
	"github.com/intruderlabs/intrusearch/main/domain/errors"
	"github.com/intruderlabs/intrusearch/main/domain/helpers"
	domain "github.com/intruderlabs/intrusearch/main/domain/responses"
	"github.com/intruderlabs/intrusearch/main/infrastructure/responses"
	"github.com/opensearch-project/opensearch-go/opensearchapi"
	logger "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"reflect"
)

const headerContentType = "Content-Type"

var headerContentTypeJSON = []string{"application/json"}

func DoRequest(transport opensearchapi.Transport, request opensearchapi.Request) (domain.GenericResponse, []errors.GenericError) {
	response, err := request.Do(context.Background(), transport)

	logger.Infoln("=> OSD Request: ", request)

	acceptedCodes := map[int]bool{
		http.StatusOK:        true,
		http.StatusCreated:   true,
		http.StatusNoContent: true,
	}

	bodyBytes, errReadAll := ioutil.ReadAll(response.Body)
	if errReadAll != nil {
		logger.Errorf("Couldn't read all the body content. Here's why: '%s'.", errReadAll)
	}
	defer response.Body.Close()

	return MapFromRequestError(domain.GenericResponse{
		Success: err == nil && acceptedCodes[response.StatusCode],
		Status:  response.StatusCode,
		Body:    bodyBytes,
	}, err)
}

// TODO: this code needs to be changed to be better
func MapFromRequestError(wrapper domain.GenericResponse, err error) (domain.GenericResponse, []errors.GenericError) {
	var mapped []errors.GenericError

	if err != nil {
		mapped = append(mapped, errors.GenericError{"http_error", err.Error()})
	} else {
		unmappedResponse := map[string]interface{}{}
		helpers.NewSerializationHelper().FromBytes(wrapper.Body, &unmappedResponse)

		rawValue, exists := unmappedResponse["error"]
		if exists && reflect.ValueOf(rawValue).Kind() == reflect.String {
			mapped = append(mapped, errors.GenericError{"http_error", fmt.Sprintf("%s", rawValue)})
		} else {
			errorResponse := responses.GenericErrorResponse{}
			helpers.NewSerializationHelper().FromBytes(wrapper.Body, &errorResponse)

			// TODO: this needs to be moved to other place
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
	}

	return wrapper, mapped
}
