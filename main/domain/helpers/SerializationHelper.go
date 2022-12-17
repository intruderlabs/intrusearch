package helpers

import (
	"bytes"
	"encoding/json"
	logger "github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
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

func (itself SerializationHelper) FromReader(body io.ReadCloser, entity any) {
	defer body.Close()

	bodyBytes, errReadAll := ioutil.ReadAll(body)
	if errReadAll != nil {
		logger.Errorf("Couldn't read all the body content. Here's why: '%s'.", errReadAll)
	}

	itself.Deserialize(string(bodyBytes), entity)
}

func (itself SerializationHelper) Deserialize(serialized string, entity interface{}) {
	errUnmarshal := json.Unmarshal([]byte(serialized), entity)
	if errUnmarshal != nil {
		logger.Errorf("Couldn't deserialize entity. Here's why: '%s'.", errUnmarshal)
	}
}

func (itself SerializationHelper) serialize(entity interface{}) []byte {
	serialized, err := json.Marshal(entity)
	if err != nil {
		logger.Errorf("Couldn't serialize entity. Here's why: '%s'.", err)
	}
	return serialized
}
