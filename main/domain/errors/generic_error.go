package errors

import "github.com/intruderlabs/intrusearch/main/domain/helpers"

type GenericError struct {
	Type   string `json:"type"`
	Reason string `json:"reason"`
}

func SerializeErrors(errors []GenericError) string {
	return helpers.NewSerializationHelper().ToString(errors)
}
