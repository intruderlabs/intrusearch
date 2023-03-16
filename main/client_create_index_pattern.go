package intrusearch

import (
	"fmt"
	"github.com/intruderlabs/intrusearch/main/domain/entities"
	"github.com/intruderlabs/intrusearch/main/domain/errors"
	"github.com/intruderlabs/intrusearch/main/domain/helpers"
	domain "github.com/intruderlabs/intrusearch/main/domain/responses"
	"github.com/intruderlabs/intrusearch/main/infrastructure/requests"
	infra "github.com/intruderlabs/intrusearch/main/infrastructure/responses"
	logger "github.com/sirupsen/logrus"
	"strings"
)

func (itself Client) CreateIndexPattern(indexName string) (bool, []errors.GenericError) {
	logger.Infof("Checking if the index pattern for the '%s' exists...", indexName)

	var mapped []errors.GenericError
	wrapper := domain.GenericResponse{Success: true}

	if _, errors := itself.getIndexPattern(indexName); len(errors) <= 0 {
		logger.Infof("The index pattern for the index '%s' already exists...", indexName)
	} else {
		logger.Infof("Creating the index pattern for the index '%s'...", indexName)

		body := entities.NewIndexPattern(indexName)
		wrapper, mapped = requests.DoRequest(itself.client, requests.SavedObjectsPostIndexPatternRequest{
			Body: helpers.NewSerializationHelper().ToReader(*body),
		})
	}

	return wrapper.Success, mapped
}

func (itself Client) getIndexPattern(name string) (entities.SavedObject, []errors.GenericError) {
	wrapper, mapped := requests.DoRequest(itself.client, requests.SavedObjectsGetIndexPatternRequest{})

	var pattern entities.SavedObject
	if wrapper.Success { // if there's no HTTP error at all
		response := infra.SavedObjectsFindResponse{}
		helpers.NewSerializationHelper().FromBytes(wrapper.Body, &response)

		for _, value := range response.SavedObjects {
			if strings.Contains(value.Attributes.Title, name) {
				pattern = value
				break
			}
		}

		if pattern.Id == "" {
			mapped = append(mapped, errors.GenericError{"not_found", fmt.Sprintf("index pattern '%s' not found", name)})
		}
	}

	return pattern, mapped
}
