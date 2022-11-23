package intrusearch

import (
	"fmt"
	"github.com/opensearch-project/opensearch-go/opensearchapi"
	logger "github.com/sirupsen/logrus"
	"gitlab.com/intruderlabs/toolbox/intrusearch/main/domain/entities"
	"gitlab.com/intruderlabs/toolbox/intrusearch/main/domain/errors"
	"gitlab.com/intruderlabs/toolbox/intrusearch/main/domain/helpers"
	"gitlab.com/intruderlabs/toolbox/intrusearch/main/infrastructure/requests"
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
