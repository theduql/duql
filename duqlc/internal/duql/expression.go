package duql

import (
	"errors"
	"fmt"

	"gopkg.in/yaml.v3"
)

type Expression struct {
	Value interface{} // This can be a string, map, or slice
}

func (e *Expression) UnmarshalYAML(value *yaml.Node) error {
	switch value.Kind {
	case yaml.ScalarNode:
		return value.Decode(&e.Value)
	case yaml.MappingNode:
		m := make(map[string]interface{})
		if err := value.Decode(&m); err != nil {
			return err
		}
		e.Value = m
	case yaml.SequenceNode:
		var s []interface{}
		if err := value.Decode(&s); err != nil {
			return err
		}
		e.Value = s
	default:
		return errors.New("unsupported YAML node type for Expression")
	}
	return nil
}

func (e *Expression) Validate() error {
	switch v := e.Value.(type) {
	case string:
		// Add any specific validation for string expressions
		return nil
	case map[string]interface{}:
		// Add any specific validation for map expressions
		return nil
	case []interface{}:
		// Add any specific validation for slice expressions
		return nil
	default:
		return fmt.Errorf("unsupported expression type: %T", v)
	}
}
