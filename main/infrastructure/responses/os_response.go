package responses

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type Source struct {
	URL       string    `json:"url"`
	Content   string    `json:"content"`
	Body      string    `json:"body"`
	Links     []string  `json:"links"`
	Timestamp time.Time `json:"@timestamp"`
}

func (s *Source) UnmarshalJSON(data []byte) error {
	type Alias Source
	aux := &struct {
		Content json.RawMessage `json:"content"`
		*Alias
	}{
		Alias: (*Alias)(s),
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	var content string
	if err := json.Unmarshal(aux.Content, &content); err == nil {
		s.Content = content
		return nil
	}

	var contents []string
	if err := json.Unmarshal(aux.Content, &contents); err == nil {
		s.Content = strings.Join(contents, ",")
		return nil
	}

	return fmt.Errorf("content field is neither a string nor an array of strings")
}

type OsResponse struct {
	Took     int  `json:"took"`
	TimedOut bool `json:"timed_out"`
	Shards   struct {
		Total      int `json:"total"`
		Successful int `json:"successful"`
		Skipped    int `json:"skipped"`
		Failed     int `json:"failed"`
	} `json:"_shards"`

	Hits struct {
		Total struct {
			Value    int    `json:"value"`
			Relation string `json:"relation"`
		} `json:"total"`
		MaxScore float64 `json:"max_score"`

		Hits []struct {
			Index  string  `json:"_index"`
			Type   string  `json:"_type"`
			ID     string  `json:"_id"`
			Source Source  `json:"_source"`
			Score  float64 `json:"_score"`
		} `json:"hits"`
	} `json:"hits"`
}
