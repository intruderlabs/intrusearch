package responses

import "github.com/intruderlabs/intrusearch/main/domain/entities"

type SavedObjectsFindResponse struct {
	Page         int                    `json:"page"`
	PerPage      int                    `json:"per_page"`
	Total        int                    `json:"total"`
	SavedObjects []entities.SavedObject `json:"saved_objects"`
}
