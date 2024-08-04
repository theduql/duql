package duql

import (
	"fmt"

	"gopkg.in/yaml.v3"
)

type Sort struct {
	Column  string   `yaml:"column" json:"column" mapstructure:"column"`
	Columns []string `yaml:",inline" json:",inline" mapstructure:",squash"`
}

func (s *Sort) Type() string {
	return "sort"
}

func (s *Sort) Validate() error {
	// Add validation logic here
	return nil
}

func (s *Sort) UnmarshalYAML(value *yaml.Node) error {
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
		return fmt.Errorf("invalid sort specification")
	}
	return nil
}
