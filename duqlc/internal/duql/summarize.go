// duqlc/internal/duql/summarize.go

package duql

import (
	"fmt"

	"gopkg.in/yaml.v3"
)

type Summarize struct {
	Aggregations map[string]Expression `yaml:"summarize"`
}

func (s *Summarize) UnmarshalYAML(value *yaml.Node) error {
	var raw map[string]interface{}
	if err := value.Decode(&raw); err != nil {
		return err
	}

	s.Aggregations = make(map[string]Expression)
	for key, val := range raw {
		expr := Expression{}
		valBytes, err := yaml.Marshal(val)
		if err != nil {
			return err
		}
		if err := yaml.Unmarshal(valBytes, &expr); err != nil {
			return err
		}
		s.Aggregations[key] = expr
	}

	return nil
}

func (j *Summarize) Type() string {
	return "summarize"
}

func (s *Summarize) Validate() error {
	if len(s.Aggregations) == 0 {
		return fmt.Errorf("summarize must contain at least one aggregation")
	}

	return nil
}
