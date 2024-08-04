package duql

import (
	"errors"
	"fmt"

	"gopkg.in/yaml.v3"
)

type Generate struct {
	Expressions map[string]Expression `yaml:"generate" json:"generate" mapstructure:"generate"`
}

func (g *Generate) Type() string {
	return "generate"
}

func (g *Generate) Validate() error {
	if len(g.Expressions) == 0 {
		return errors.New("generate must contain at least one expression")
	}
	for name, expr := range g.Expressions {
		if err := expr.Validate(); err != nil {
			return fmt.Errorf("invalid expression for %s: %w", name, err)
		}
	}
	return nil
}

func (g *Generate) UnmarshalYAML(value *yaml.Node) error {
	var raw map[string]Expression
	if err := value.Decode(&raw); err != nil {
		return err
	}
	g.Expressions = raw
	return nil
}
