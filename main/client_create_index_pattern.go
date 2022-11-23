package intrusearch

import (
	"fmt"
	logger "github.com/sirupsen/logrus"
	"gitlab.com/intruderlabs/toolbox/intrusearch/main/domain/entities"
	"gitlab.com/intruderlabs/toolbox/intrusearch/main/domain/errors"
	"gitlab.com/intruderlabs/toolbox/intrusearch/main/domain/helpers"
	domain "gitlab.com/intruderlabs/toolbox/intrusearch/main/domain/responses"
	"gitlab.com/intruderlabs/toolbox/intrusearch/main/infrastructure/requests"
	infra "gitlab.com/intruderlabs/toolbox/intrusearch/main/infrastructure/responses"
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
		helpers.NewSerializationHelper().Deserialize(wrapper.Body, &response)

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
