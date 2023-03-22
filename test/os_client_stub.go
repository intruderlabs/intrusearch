package intrusearch

import (
	"github.com/intruderlabs/intrusearch/main/domain/entities"
	"github.com/intruderlabs/intrusearch/main/domain/errors"
	dresponses "github.com/intruderlabs/intrusearch/main/domain/responses"
	"github.com/intruderlabs/intrusearch/main/infrastructure/requests"
	"github.com/intruderlabs/intrusearch/main/infrastructure/responses"
)

type OsClientStub struct {
	Error        error
	OsResponse   responses.OsResponse
	GenericError []errors.GenericError
	Documents    dresponses.CreateDocumentsResponse
}

func (itself OsClientStub) CreateDocuments(indexName string, documents []entities.Document) dresponses.CreateDocumentsResponse {
	return itself.Documents
}

func (itself OsClientStub) CreateIndex(indexName string) (bool, []errors.GenericError) {
	return false, itself.GenericError
}

func (itself OsClientStub) CreateIndexPattern(indexName string) (bool, []errors.GenericError) {
	return false, itself.GenericError
}

func (itself OsClientStub) CreateIndexPolicy(indexName string) (bool, []errors.GenericError) {
	return false, itself.GenericError
}

func (itself OsClientStub) CreateIndexTemplate(indexName string, properties entities.IndexTemplateMappingProperties) (bool, []errors.GenericError) {
	return false, itself.GenericError
}

func (itself OsClientStub) Initialize(indexName string, properties entities.IndexTemplateMappingProperties) error {
	return itself.Error
}

func (itself OsClientStub) ClientSearchRequest(queryPaginationRequest requests.OsSearchRequest) (responses.OsResponse, []errors.GenericError) {
	return itself.OsResponse, itself.GenericError
}
