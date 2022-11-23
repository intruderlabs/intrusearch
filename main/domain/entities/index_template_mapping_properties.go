package entities

type indexTemplatePropertyProperties struct {
	Type        string `json:"type"`
	IgnoreAbove int    `json:"ignore_above,omitempty"`
}

type IndexTemplateMappingProperties map[string]indexTemplatePropertyProperties
