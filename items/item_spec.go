package items

import (
	"bytes"
	"fmt"

	"gopkg.in/yaml.v2"
	_ "embed"
)

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
	Definitions interface{} `yaml:"definitions,omitempty"` //They currently exist but are unused
}

type ItemSpec struct {
	ID string `yaml:"id"`
	Attributes []AttributeSpec `yarm:"attributes,omitempty"`
}

var (
	//go:embed data/super-mario.yml
	supermario []byte
)

type ItemSpecRepository interface {
	FindFor(target string) []*ItemSpec
}

type EmbedItemSpecRepository struct {
	specsByGame map[string][]*ItemSpec
}

func parseYamlToItemSpecs(content []byte) (result []*ItemSpec, err error) {
	allItems := bytes.Split(content, []byte("---"))
	for i, itemRawSpec := range allItems {
		var item *ItemSpec
		err = yaml.Unmarshal(itemRawSpec, &item)
		if err != nil {
			return nil, fmt.Errorf("Cannot read spec at index %v", i)
		}
		result = append(result, item)
	}
	return result, nil
}

func NewRepoItemSpecs() (*EmbedItemSpecRepository, error) {
	specsByGame := map[string][]*ItemSpec{} 
	if supermarioSpecs, err := parseYamlToItemSpecs(supermario); err == nil {
		specsByGame["supermario"] = supermarioSpecs
	}
	return &EmbedItemSpecRepository {
		specsByGame: specsByGame,
	}, nil
}

func (r *EmbedItemSpecRepository) FindFor(target string) []*ItemSpec {
	return r.specsByGame[target]
}