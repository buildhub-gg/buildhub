package graph

import (
	"fmt"

	buildPackage "github.com/buildhub-gg/buildhub/build"
	"github.com/buildhub-gg/buildhub/graph/model"
	"github.com/buildhub-gg/buildhub/item"
)

func convertInputBuildToOutputBuild(ID string, in *model.InputBuild) *model.Build  {
	items := []*model.Item {}

	for _, item := range in.Items {
		items = append(items, &model.Item{
			ID: item.ID,
		})
	}

	return &model.Build{
		ID: ID,
		Name: in.Name,
		Items: items,
	}
}

func convertEntityBuildToOutputBuild(ID string, in *buildPackage.Build) *model.Build {
	items := []*model.Item {}

	for _, item := range in.Items {
		items = append(items, &model.Item{
			ID: item.ID,
		})
	}

	return &model.Build{
		ID: ID,
		Name: in.Name,
		Items: items,
	}
}

func convertInputBuildToEntityBuild(in *model.InputBuild) *buildPackage.Build {
	buildItems := []*buildPackage.Item {}

	for _, item := range in.Items {
		buildItems = append(buildItems, &buildPackage.Item{
			ID: item.ID,
		})
	}

	return &buildPackage.Build{
		Name: in.Name,
		Items: buildItems,
	}
}

func convertItemSpecToGraphql(in *item.ItemSpec) (*model.ItemSpec, error) {
	attributes := []model.AttributeSpec{}
	for _, attribute := range in.Attributes {
		convertedAttribute, err := convertAttributeSpecToGraphql(&attribute)
		if err != nil {
			return nil, err
		}
		attributes = append(attributes, convertedAttribute)
	}
	
	return &model.ItemSpec{
		ID: in.ID,
		Tags: in.Tags,
		Attributes: attributes,
	}, nil
}

func convertAttributeSpecToGraphql(in *item.AttributeSpec) (model.AttributeSpec, error) {
	switch itemType := in.Type; itemType {
	case item.BOOL:
		return convertBooleanAttributeSpecToGraphql(in), nil
	case item.STRING:
		return convertStringAttributeSpecToGraphql(in), nil
	case item.FLOAT:
		return convertFloatAttributeSpecToGraphql(in), nil
	case item.INT:
		return convertIntAttributeSpecToGraphql(in), nil
	default:
		return nil, fmt.Errorf("Cannot find spec of type %s", itemType)
	}
}

func convertBooleanAttributeSpecToGraphql(in *item.AttributeSpec) *model.BooleanAttributeSpec {
	out := &model.BooleanAttributeSpec {
		ID: in.ID,
	}
	
	if defaultValue, ok := in.Definitions["default"]; ok {
		defaultValueBool := defaultValue.(bool)
		out.Default = &defaultValueBool
	}

	return out
}

func convertStringAttributeSpecToGraphql(in *item.AttributeSpec) *model.StringAttributeSpec {
	out := &model.StringAttributeSpec {
		ID: in.ID,
	}
	
	if defaultValue, ok := in.Definitions["default"]; ok {
		defaultValueString := defaultValue.(string)
		out.Default = &defaultValueString
	}

	return out
}

func convertIntAttributeSpecToGraphql(in *item.AttributeSpec) *model.IntAttributeSpec {
	out := &model.IntAttributeSpec {
		ID: in.ID,
	}
	
	if defaultValue, ok := in.Definitions["default"]; ok {
		defaultValueInt := defaultValue.(int)
		out.Default = &defaultValueInt
	}
	
	if min, ok := in.Definitions["min"]; ok {
		minValue := min.(int)
		out.Min = &minValue
	}
	
	if max, ok := in.Definitions["max"]; ok {
		maxValue := max.(int)
		out.Max = &maxValue
	}

	return out
}

func convertFloatAttributeSpecToGraphql(in *item.AttributeSpec) *model.FloatAttributeSpec {
	out := &model.FloatAttributeSpec {
		ID: in.ID,
	}
	
	if defaultValue, ok := in.Definitions["default"]; ok {
		defaultValueFloat := defaultValue.(float64)
		out.Default = &defaultValueFloat
	}
	
	if min, ok := in.Definitions["min"]; ok {
		minValue := min.(float64)
		out.Min = &minValue
	}
	
	if max, ok := in.Definitions["max"]; ok {
		maxValue := max.(float64)
		out.Max = &maxValue
	}

	return out
}

