package duql

import (
	"fmt"

	"gopkg.in/yaml.v3"
)

type Query struct {
	Settings *Settings `yaml:"settings,omitempty" json:"settings,omitempty" mapstructure:"settings,omitempty"`
	Declare  Declare   `yaml:"declare,omitempty" json:"declare,omitempty" mapstructure:"declare,omitempty"`
	Dataset  Dataset   `yaml:"dataset" json:"dataset" mapstructure:"dataset"`
	Steps    Steps     `yaml:"steps,omitempty" json:"steps,omitempty" mapstructure:"steps,omitempty"`
	Into     string    `yaml:"into,omitempty" json:"into,omitempty" mapstructure:"into,omitempty"`
}

type Steps []Step

func (s *Steps) UnmarshalYAML(value *yaml.Node) error {
	var rawSteps []map[string]interface{}
	if err := value.Decode(&rawSteps); err != nil {
		return err
	}

	*s = make(Steps, len(rawSteps))
	for i, rawStep := range rawSteps {
		if len(rawStep) != 1 {
			return fmt.Errorf("each step must be a single-key object")
		}

		for stepType, stepValue := range rawStep {
			var step Step
			switch stepType {
			case "filter":
				step = &Filter{}
			case "join":
				step = &Join{}
			case "group":
				step = &Group{}
			case "generate":
				step = &Generate{}
			case "sort":
				step = &Sort{}
			case "take":
				step = &Take{}
			case "window":
				step = &Window{}
			case "select":
				step = &Select{}
			case "select!":
				step = &SelectNot{}
			case "loop":
				step = &Loop{}
			case "summarize":
				step = &Summarize{}
			default:
				return fmt.Errorf("unknown step type: %s", stepType)
			}

			valueBytes, err := yaml.Marshal(stepValue)
			if err != nil {
				return err
			}

			if err := yaml.Unmarshal(valueBytes, step); err != nil {
				return err
			}

			(*s)[i] = step
		}
	}

	return nil
}

func (q *Query) Validate() error {
	if q.Dataset == (Dataset{}) {
		return fmt.Errorf("dataset is required")
	}

	if q.Declare != nil {
		if err := q.Declare.Validate(); err != nil {
			return fmt.Errorf("invalid declare section: %w", err)
		}
	}

	if q.Steps != nil {
		for _, step := range q.Steps {
			if err := step.Validate(); err != nil {
				return fmt.Errorf("invalid step: %w", err)
			}
		}
	}

	return nil
}
