package intrusearch

import (
	"fmt"
	"github.com/intruderlabs/intrusearch/main/domain/entities"
	"github.com/intruderlabs/intrusearch/main/domain/errors"
	"github.com/intruderlabs/intrusearch/main/domain/helpers"
	"github.com/intruderlabs/intrusearch/main/infrastructure/requests"
	"github.com/opensearch-project/opensearch-go/opensearchapi"
	logger "github.com/sirupsen/logrus"
)

func (itself Client) CreateIndex(indexName string) (bool, []errors.GenericError) {
	logger.Infof("Creating the index '%s'...", indexName)

	body := entities.NewIndex(indexName)
	wrapper, mapped := requests.DoRequest(itself.client, opensearchapi.IndicesCreateRequest{
		Index: fmt.Sprintf("%s-000001", indexName),
		Body:  helpers.NewSerializationHelper().ToReader(*body),
	})

	return wrapper.Success, mapped
}
