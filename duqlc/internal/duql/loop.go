package duql

import (
	"errors"
	"fmt"

	"gopkg.in/yaml.v3"
)

type Loop struct {
	Steps Steps `yaml:"loop" json:"loop" mapstructure:"loop"`
}

func (l *Loop) Type() string {
	return "loop"
}

func (l *Loop) Validate() error {
	if len(l.Steps) == 0 {
		return errors.New("loop must contain at least one step")
	}
	for _, step := range l.Steps {
		if err := step.Validate(); err != nil {
			return fmt.Errorf("invalid step in loop: %w", err)
		}
	}
	return nil
}

func (l *Loop) UnmarshalYAML(value *yaml.Node) error {
	var steps Steps
	if err := value.Decode(&steps); err != nil {
		return err
	}
	l.Steps = steps
	return nil
}
