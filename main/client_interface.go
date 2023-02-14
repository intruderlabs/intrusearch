package intrusearch

import (
	"gitlab.com/intruderlabs/toolbox/intrusearch.git/main/domain/entities"
	"gitlab.com/intruderlabs/toolbox/intrusearch.git/main/domain/errors"
	dresponses "gitlab.com/intruderlabs/toolbox/intrusearch.git/main/domain/responses"
	"gitlab.com/intruderlabs/toolbox/intrusearch.git/main/infrastructure/requests"
	"gitlab.com/intruderlabs/toolbox/intrusearch.git/main/infrastructure/responses"
)

type ClientInterface interface {
	CreateDocuments(indexName string, documents []entities.Document) dresponses.CreateDocumentsResponse
	CreateIndex(indexName string) (bool, []errors.GenericError)
	CreateIndexPattern(indexName string) (bool, []errors.GenericError)
	CreateIndexPolicy(indexName string) (bool, []errors.GenericError)
	CreateIndexTemplate(indexName string, properties entities.IndexTemplateMappingProperties) (bool, []errors.GenericError)
	Initialize(indexName string, properties entities.IndexTemplateMappingProperties) error
	ClientSearchRequest(queryPaginationRequest requests.OsSearchRequest) (responses.OsResponse, []errors.GenericError)
}
