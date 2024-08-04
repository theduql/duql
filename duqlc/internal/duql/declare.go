package duql

import (
	"errors"
	"fmt"
	"regexp"

	"gopkg.in/yaml.v3"
)

type Declare map[string]DeclareValue

type DeclareValue struct {
	Pipeline   *Steps                 `yaml:",inline,omitempty" json:"pipeline,omitempty" mapstructure:"pipeline,omitempty"`
	Expression *Expression            `yaml:",inline,omitempty" json:"expression,omitempty" mapstructure:"expression,omitempty"`
	Tuple      map[string]interface{} `yaml:",inline,omitempty" json:"tuple,omitempty" mapstructure:"tuple,omitempty"`
	Function   *FunctionDefinition    `yaml:",inline,omitempty" json:"function,omitempty" mapstructure:"function,omitempty"`
}

type FunctionDefinition struct {
	Parameters []FunctionParameter `yaml:"parameters" json:"parameters" mapstructure:"parameters"`
	Expression Expression          `yaml:"expression" json:"expression" mapstructure:"expression"`
}

type FunctionParameter struct {
	Name    string      `yaml:"name" json:"name" mapstructure:"name"`
	Default interface{} `yaml:"default,omitempty" json:"default,omitempty" mapstructure:"default,omitempty"`
}

func (d *Declare) UnmarshalYAML(value *yaml.Node) error {
	var raw map[string]interface{}
	if err := value.Decode(&raw); err != nil {
		return err
	}

	*d = make(Declare)
	for key, val := range raw {
		if !isValidVariableName(key) {
			return fmt.Errorf("invalid variable name: %s", key)
		}

		var declareValue DeclareValue
		valBytes, err := yaml.Marshal(val)
		if err != nil {
			return err
		}

		if err := yaml.Unmarshal(valBytes, &declareValue); err != nil {
			return err
		}

		(*d)[key] = declareValue
	}

	return nil
}

func isValidVariableName(name string) bool {
	return regexp.MustCompile(`^[a-zA-Z_][a-zA-Z0-9_]*$`).MatchString(name)
}

func (d *Declare) Validate() error {
	for key, value := range *d {
		if !isValidVariableName(key) {
			return fmt.Errorf("invalid variable name: %s", key)
		}
		if err := value.Validate(); err != nil {
			return fmt.Errorf("invalid value for %s: %w", key, err)
		}
	}
	return nil
}

func (dv *DeclareValue) Validate() error {
	count := 0
	if dv.Pipeline != nil {
		count++
	}
	if dv.Expression != nil {
		count++
	}
	if dv.Tuple != nil {
		count++
	}
	if dv.Function != nil {
		count++
	}

	if count != 1 {
		return errors.New("declare value must be exactly one of: pipeline, expression, tuple, or function")
	}

	if dv.Function != nil {
		if len(dv.Function.Parameters) == 0 {
			return errors.New("function must have at least one parameter")
		}
		if err := dv.Function.Expression.Validate(); err != nil {
			return fmt.Errorf("invalid function expression: %w", err)
		}
	}

	return nil
}
