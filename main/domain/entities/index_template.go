package entities

import "fmt"

type indexTemplateDynamicTemplate struct {
	StringsAsKeyword struct {
		Mapping struct {
			IgnoreAbove int    `json:"ignore_above"`
			Type        string `json:"type"`
		} `json:"mapping"`
		MatchMappingType string `json:"match_mapping_type"`
	} `json:"strings_as_keyword"`
}

type indexTemplate struct {
	ComposedOf    []string `json:"composed_of,omitempty"`
	Priority      int      `json:"priority"`
	IndexPatterns []string `json:"index_patterns"`
	Template      struct {
		// default settings + component
		Settings struct {
			Index struct {
				Mapping struct {
					TotalFields struct {
						Limit int `json:"limit"`
					} `json:"total_fields"`
				} `json:"mapping"`
			} `json:"index"`
			RefreshInterval  string `json:"refresh_interval"`
			NumberOfShards   int    `json:"number_of_shards"`
			NumberOfReplicas int    `json:"number_of_replicas"`
			Plugins          struct {
				IndexStateManagement struct {
					RolloverAlias string `json:"rollover_alias"`
				} `json:"index_state_management"`
			} `json:"plugins"`
		} `json:"settings"`
		// template + component
		Mappings struct {
			DateDetection    bool                           `json:"date_detection"`
			DynamicTemplates []indexTemplateDynamicTemplate `json:"dynamic_templates"`
			Properties       IndexTemplateMappingProperties `json:"properties"`
		} `json:"mappings"`
	} `json:"template"`
}

func NewIndexTemplate(indexName string, properties IndexTemplateMappingProperties) *indexTemplate {
	template := indexTemplate{}
	template.Priority = 100
	template.IndexPatterns = []string{fmt.Sprintf("%s*", indexName)}

	// default settings
	template.Template.Settings.Index.Mapping.TotalFields.Limit = 3072
	template.Template.Settings.RefreshInterval = "30s"
	template.Template.Settings.NumberOfShards = 2
	template.Template.Settings.NumberOfReplicas = 1

	// component
	template.Template.Settings.Plugins.IndexStateManagement.RolloverAlias = indexName

	// template
	dynamicTemplate := indexTemplateDynamicTemplate{}
	dynamicTemplate.StringsAsKeyword.Mapping.IgnoreAbove = 1024
	dynamicTemplate.StringsAsKeyword.Mapping.Type = "keyword"
	dynamicTemplate.StringsAsKeyword.MatchMappingType = "string"

	template.Template.Mappings.DateDetection = false
	template.Template.Mappings.DynamicTemplates = []indexTemplateDynamicTemplate{dynamicTemplate}

	//component
	template.Template.Mappings.Properties = properties

	return &template
}
