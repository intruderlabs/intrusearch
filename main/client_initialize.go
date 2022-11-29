package intrusearch

import (
	"errors"
	"github.com/opensearch-project/opensearch-go/opensearchapi"
	logger "github.com/sirupsen/logrus"
	"gitlab.com/intruderlabs/toolbox/intrusearch.git/main/domain/entities"
	derrors "gitlab.com/intruderlabs/toolbox/intrusearch.git/main/domain/errors"
	"gitlab.com/intruderlabs/toolbox/intrusearch.git/main/infrastructure/requests"
)

func (itself Client) Initialize(
	indexName string,
	properties entities.IndexTemplateMappingProperties,
) error {
	logger.Infof("Entering in the initialization routine for the index '%s'...", indexName)

	steps := []func(indexName string) (bool, []derrors.GenericError){
		itself.shouldInitialized,
		func(indexName string) (bool, []derrors.GenericError) {
			return itself.CreateIndexTemplate(indexName, properties)
		},
		itself.CreateIndexPolicy,
		itself.CreateIndex,
		itself.CreateIndexPattern,
	}

	var err error
	var success bool
	var mapped []derrors.GenericError

	for _, value := range steps {
		if success, mapped = value(indexName); !success {
			break
		}
	}

	if len(mapped) > 0 {
		err = errors.New(derrors.SerializeErrors(mapped))
	}

	return err
}

func (itself Client) shouldInitialized(indexName string) (bool, []derrors.GenericError) {
	logger.Infof("Checking if the index '%s' is initialized...", indexName)

	wrapper, mapped := requests.DoRequest(itself.client, opensearchapi.IndicesGetAliasRequest{
		Name: []string{indexName},
	})

	return !wrapper.Success, mapped
}
