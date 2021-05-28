package api

// PipelineTemplateV2 defines the schema of the pipeline template JSON.
type PipelineTemplateV2 struct {
	ID        string                       `json:"id,omitempty"`
	Metadata  PipelineTemplateV2Metadata   `json:"metadata"`
	Pipeline  interface{}                  `json:"pipeline"`
	Protect   *bool                        `json:"protect,omitempty"`
	Schema    string                       `json:"schema"`
	Variables []PipelineTemplateV2Variable `json:"variables,omitempty"`
}

type PipelineTemplateV2Metadata struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Owner       *string  `json:"owner,omitempty"`
	Scopes      []string `json:"scopes,omitempty"`
}

type PipelineTemplateV2Variable struct {
	Name         string      `json:"name"`
	Description  *string     `json:"description,omitempty"`
	DefaultValue interface{} `json:"defaultValue,omitempty"`
	Type         string      `json:"type"`
}

type PipelineTemplateV2Version struct {
	ID     string `json:"id"`
	Digest string `json:"digest"`
	Tag    string `json:"tag"`
}
