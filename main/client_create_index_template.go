package intrusearch

import (
	"github.com/intruderlabs/intrusearch/main/domain/entities"
	"github.com/intruderlabs/intrusearch/main/domain/errors"
	"github.com/intruderlabs/intrusearch/main/domain/helpers"
	"github.com/intruderlabs/intrusearch/main/infrastructure/requests"
	"github.com/opensearch-project/opensearch-go/opensearchapi"
	logger "github.com/sirupsen/logrus"
)

func (itself Client) CreateIndexTemplate(
	indexName string,
	properties entities.IndexTemplateMappingProperties,
) (bool, []errors.GenericError) {
	logger.Infof("Creating the template for the index '%s'...", indexName)

	body := entities.NewIndexTemplate(indexName, properties)
	wrapper, mapped := requests.DoRequest(itself.client, opensearchapi.IndicesPutIndexTemplateRequest{
		Name: indexName,
		Body: helpers.NewSerializationHelper().ToReader(*body),
	})

	return wrapper.Success, mapped
}
