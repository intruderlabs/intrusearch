package intrusearch

import (
	"github.com/opensearch-project/opensearch-go/opensearchapi"
	logger "github.com/sirupsen/logrus"
	"gitlab.com/intruderlabs/toolbox/intrusearch.git/main/domain/errors"
	"gitlab.com/intruderlabs/toolbox/intrusearch.git/main/domain/helpers"
	"gitlab.com/intruderlabs/toolbox/intrusearch.git/main/infrastructure/requests"
	"gitlab.com/intruderlabs/toolbox/intrusearch.git/main/infrastructure/responses"
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
