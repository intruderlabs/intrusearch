package responses

import "github.com/intruderlabs/intrusearch/main/domain/errors"

type BulkResponse struct {
	Took   int  `json:"took,omitempty"`
	Errors bool `json:"errors,omitempty"`
	Items  []map[string]struct {
		// success + error
		Index   string `json:"_index"`
		Type    string `json:"_type"`
		Id      string `json:"_id"`
		Version int    `json:"_version"`
		Result  string `json:"result"`
		Shards  struct {
			Successful int `json:"successful"`
			Failed     int `json:"failed"`
			Total      int `json:"total"`
		} `json:"_shards"`
		SeqNo       int `json:"_seq_no"`
		PrimaryTerm int `json:"_primary_term"`
		Status      int `json:"status"`

		// error
		Error struct {
			Type     string `json:"type"`
			Reason   string `json:"reason"`
			CausedBy struct {
				Type     string              `json:"type"`
				Reason   string              `json:"reason"`
				CausedBy errors.GenericError `json:"caused_by"`
			} `json:"caused_by"`
			Index     string `json:"index"`
			Shard     string `json:"shard"`
			IndexUuid string `json:"index_uuid"`
		} `json:"error,omitempty"`
	} `json:"items,omitempty"`
}
