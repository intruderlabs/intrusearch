package intrusearch

import (
	"github.com/intruderlabs/intrusearch/main/domain/errors"
	"github.com/intruderlabs/intrusearch/main/domain/helpers"
	"github.com/intruderlabs/intrusearch/main/infrastructure/requests"
	"github.com/intruderlabs/intrusearch/main/infrastructure/responses"
	"github.com/opensearch-project/opensearch-go/opensearchapi"
	logger "github.com/sirupsen/logrus"
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
	})

	response := responses.OsResponse{}

	if wrapper.Success {
		helpers.NewSerializationHelper().FromBytes(wrapper.Body, &response)
	}
	return response, mapped
}

func (itself Client) ClientIdSearchRequest(
	id []string,
) (
	responses.OsResponse,
	[]errors.GenericError,
) {
	logger.Info("initialize search request in db ")

	wrapper, mapped := requests.DoRequest(itself.client, opensearchapi.SearchRequest{
		Index: id,
	})

	response := responses.OsResponse{}

	if wrapper.Success {
		helpers.NewSerializationHelper().FromBytes(wrapper.Body, &response)
	}
	return response, mapped
}
