package intrusearch

import (
	"github.com/intruderlabs/intrusearch/main/domain/entities"
	"github.com/intruderlabs/intrusearch/main/domain/errors"
	dresponses "github.com/intruderlabs/intrusearch/main/domain/responses"
	"github.com/intruderlabs/intrusearch/main/infrastructure/requests"
	"github.com/intruderlabs/intrusearch/main/infrastructure/responses"
)

type ClientInterface interface {
	CreateDocuments(indexName string, documents []entities.Document) dresponses.CreateDocumentsResponse
	CreateIndex(indexName string) (bool, []errors.GenericError)
	CreateIndexPattern(indexName string) (bool, []errors.GenericError)
	CreateIndexPolicy(indexName string) (bool, []errors.GenericError)
	CreateIndexTemplate(indexName string, properties entities.IndexTemplateMappingProperties) (bool, []errors.GenericError)
	Initialize(indexName string, properties entities.IndexTemplateMappingProperties) error
	ClientSearchRequest(queryPaginationRequest requests.OsSearchRequest) (responses.OsResponse, []errors.GenericError)
	ClientIdSearchRequest(id string) (responses.OsResponse, []errors.GenericError)
}
