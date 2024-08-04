package duql

import (
	"fmt"

	"gopkg.in/yaml.v3"
)

type Take struct {
	Number int    `yaml:"number" json:"number" mapstructure:"number"`
	Range  string `yaml:"range" json:"range" mapstructure:",squash"`
}

func (t *Take) Type() string {
	return "take"
}

func (t *Take) Validate() error {
	// Add validation logic here
	return nil
}

func (t *Take) UnmarshalYAML(value *yaml.Node) error {
	var tmp interface{}
	if err := value.Decode(&tmp); err != nil {
		return err
	}

	switch v := tmp.(type) {
	case int:
		t.Number = v
	case string:
		t.Range = v
	default:
		return fmt.Errorf("invalid take specification")
	}
	return nil
}
