package entities

type index struct {
	Aliases map[string]interface{} `json:"aliases"`
}

func NewIndex(indexName string) *index {
	idx := index{}
	idx.Aliases = map[string]interface{}{}
	idx.Aliases[indexName] = struct{}{}
	return &idx
}
