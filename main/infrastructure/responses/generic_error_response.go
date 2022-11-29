package responses

import "gitlab.com/intruderlabs/toolbox/intrusearch.git/main/domain/errors"

type GenericErrorResponse struct {
	Status int `json:"status"`
	Error  struct {
		Type      string                `json:"type"`
		Reason    string                `json:"reason"`
		RootCause []errors.GenericError `json:"root_cause"`
	} `json:"error"`
}
