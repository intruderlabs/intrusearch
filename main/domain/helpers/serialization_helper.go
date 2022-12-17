package helpers

import (
	"bytes"
	"encoding/json"
	logger "github.com/sirupsen/logrus"
)

type SerializationHelper struct {
}

func NewSerializationHelper() *SerializationHelper {
	return &SerializationHelper{}
}

func (itself SerializationHelper) ToString(entity interface{}) string {
	return string(itself.serialize(entity))
}

func (itself SerializationHelper) ToReader(entity interface{}) *bytes.Reader {
	return bytes.NewReader(itself.serialize(entity))
}

func (itself SerializationHelper) FromString(serialized string, entity interface{}) {
	itself.deserialize([]byte(serialized), entity)
}

func (itself SerializationHelper) FromBytes(serialized []byte, entity interface{}) {
	itself.deserialize(serialized, entity)
}

func (itself SerializationHelper) deserialize(serialized []byte, entity interface{}) {
	err := json.Unmarshal(serialized, entity)
	if err != nil {
		logger.Errorf("Couldn't deserialize entity. Here's why: '%s'.", err)
	}
}

func (itself SerializationHelper) serialize(entity interface{}) []byte {
	serialized, err := json.Marshal(entity)
	if err != nil {
		logger.Errorf("Couldn't serialize entity. Here's why: '%s'.", err)
	}
	return serialized
}
