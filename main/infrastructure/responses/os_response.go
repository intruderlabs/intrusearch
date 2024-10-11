package responses

import "time"

type Source struct {
	URL       string    `json:"url"`
	Content   string    `json:"content"`
	Body      string    `json:"body"`
	Channel   string    `json:"channel"`
	Links     []string  `json:"links"`
	FoundAt   string    `json:"foundAt"`
	Timestamp time.Time `json:"@timestamp"`
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
