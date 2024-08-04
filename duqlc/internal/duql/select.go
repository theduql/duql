package duql

import (
	"fmt"

	"gopkg.in/yaml.v3"
)

type Select struct {
	Column  string   `yaml:"column" json:"column" mapstructure:"column"`
	Columns []string `yaml:",inline" json:",inline" mapstructure:",squash"`
}

type SelectNot struct {
	Column  string   `yaml:"column" json:"column" mapstructure:"column"`
	Columns []string `yaml:",inline" json:",inline" mapstructure:",squash"`
}

func (s *Select) Type() string {
	return "select"
}

func (s *Select) Validate() error {
	// Add validation logic here
	return nil
}

func (s *SelectNot) Type() string {
	return "select!"
}

func (s *SelectNot) Validate() error {
	// Add validation logic here
	return nil
}

func (s *Select) UnmarshalYAML(value *yaml.Node) error {
	var tmp interface{}
	if err := value.Decode(&tmp); err != nil {
		return err
	}

	switch v := tmp.(type) {
	case string:
		s.Column = v
	case []interface{}:
		s.Columns = make([]string, len(v))
		for i, col := range v {
			s.Columns[i] = col.(string)
		}
	default:
		return fmt.Errorf("invalid select specification")
	}
	return nil
}

func (s *SelectNot) UnmarshalYAML(value *yaml.Node) error {
	var tmp interface{}
	if err := value.Decode(&tmp); err != nil {
		return err
	}

	switch v := tmp.(type) {
	case string:
		s.Column = v
	case []interface{}:
		s.Columns = make([]string, len(v))
		for i, col := range v {
			s.Columns[i] = col.(string)
		}
	default:
		return fmt.Errorf("invalid select not specification")
	}
	return nil
}
