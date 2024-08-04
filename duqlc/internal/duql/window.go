package duql

import "gopkg.in/yaml.v3"

type Window struct {
	Rows      string `yaml:"rows,omitempty" json:"rows,omitempty" mapstructure:"rows,omitempty"`
	Range     string `yaml:"range,omitempty" json:"range,omitempty" mapstructure:"range,omitempty"`
	Expanding bool   `yaml:"expanding,omitempty" json:"expanding,omitempty" mapstructure:"expanding,omitempty"`
	Rolling   int    `yaml:"rolling,omitempty" json:"rolling,omitempty" mapstructure:"rolling,omitempty"`
	Steps     []Step `yaml:"steps" json:"steps" mapstructure:"steps"`
}

func (w *Window) Type() string {
	return "window"
}

func (w *Window) Validate() error {
	// Add validation logic here
	return nil
}

func (w *Window) UnmarshalYAML(value *yaml.Node) error {
	type rawWindow Window
	if err := value.Decode((*rawWindow)(w)); err != nil {
		return err
	}
	return nil
}
