package item

type AttributeType string

const (
	STRING AttributeType = "STRING"
	INT AttributeType = "INT"
	FLOAT AttributeType = "FLOAT"
	BOOL AttributeType = "BOOL"
)

type AttributeSpec struct {
	ID string `yaml:"id"`
	Type AttributeType `yaml:"type"`
	Definitions map[string]interface{} `yaml:"definitions,omitempty"` //They currently exist but are unused
}

type ItemSpec struct {
	ID string `yaml:"id"`
	Attributes []AttributeSpec `yarm:"attributes,omitempty"`
}