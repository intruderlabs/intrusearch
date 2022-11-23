package intrusearch

import (
	"github.com/opensearch-project/opensearch-go/opensearchapi"
	logger "github.com/sirupsen/logrus"
	"gitlab.com/intruderlabs/toolbox/intrusearch/main/domain/entities"
	"gitlab.com/intruderlabs/toolbox/intrusearch/main/domain/errors"
	"gitlab.com/intruderlabs/toolbox/intrusearch/main/domain/helpers"
	"gitlab.com/intruderlabs/toolbox/intrusearch/main/infrastructure/requests"
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
