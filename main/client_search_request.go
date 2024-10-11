package intrusearch

import (
	"fmt"
	"github.com/intruderlabs/intrusearch/main/domain/errors"
	"github.com/intruderlabs/intrusearch/main/domain/helpers"
	"github.com/intruderlabs/intrusearch/main/infrastructure/requests"
	"github.com/intruderlabs/intrusearch/main/infrastructure/responses"
	"github.com/opensearch-project/opensearch-go/opensearchapi"
	logger "github.com/sirupsen/logrus"
	"io"
)

func (itself Client) ClientSearchRequest(
	queryPaginationRequest requests.OsSearchRequest,
) (
	responses.OsResponse,
	[]errors.GenericError,
) {
	logger.Info("initialize search request in db ")

	wrapper, mapped := requests.DoRequest(itself.client, opensearchapi.SearchRequest{
		Size:  &queryPaginationRequest.Size,
		From:  &queryPaginationRequest.From,
		Query: queryPaginationRequest.QueryString,
		Index: queryPaginationRequest.Index,
	})

	response := responses.OsResponse{}

	if wrapper.Success {
		helpers.NewSerializationHelper().FromBytes(wrapper.Body, &response)
	}
	return response, mapped
}

func (itself Client) ClientIdSearchRequest(
	index string,
	id string,
) (
	responses.OsResponse,
	[]errors.GenericError,
) {
	logger.Info("initialize ID search request in db!")

	var mappedErrors []errors.GenericError
	indexResponse, err := itself.client.Get(index, id)

	if err != nil {
		logger.Errorln("ClientIdSearchRequest()->Get():", err)
		mappedErrors = append(mappedErrors, errors.GenericError{Type: "osd_error", Reason: fmt.Sprintf("%s", err)})
		return responses.OsResponse{}, make([]errors.GenericError, 0)
	}

	response := responses.OsResponse{}
	indexResponseBytes, err := io.ReadAll(indexResponse.Body)

	if err != nil {
		logger.Errorln("ClientIdSearchRequest()->ReadAll():", err)
		mappedErrors = append(mappedErrors, errors.GenericError{Type: "response_error", Reason: fmt.Sprintf("%s", err)})
	}

	logger.Infoln("ClientIdSearchRequest HTTP status code:", indexResponse.StatusCode)

	helpers.NewSerializationHelper().FromBytes(indexResponseBytes, &response)

	response.Hits.Total.Value = indexResponse.StatusCode
	return response, mappedErrors
}
