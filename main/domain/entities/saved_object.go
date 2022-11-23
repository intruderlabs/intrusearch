package entities

import "time"

type SavedObject struct {
	Id               string                 `json:"id"`
	Type             string                 `json:"type"`
	Attributes       indexPatternAttributes `json:"attributes"`
	References       []interface{}          `json:"references"`
	MigrationVersion struct {
		IndexPattern string `json:"index-pattern"`
	} `json:"migrationVersion"`
	UpdatedAt  time.Time `json:"updated_at"`
	Version    string    `json:"version"`
	Namespaces []string  `json:"namespaces"`
	Score      int       `json:"score"`
}
