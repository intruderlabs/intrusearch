package responses

import "gitlab.com/intruderlabs/toolbox/intrusearch/main/domain/entities"

type SavedObjectsFindResponse struct {
	Page         int                    `json:"page"`
	PerPage      int                    `json:"per_page"`
	Total        int                    `json:"total"`
	SavedObjects []entities.SavedObject `json:"saved_objects"`
}
