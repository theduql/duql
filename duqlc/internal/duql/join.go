package duql

import "gopkg.in/yaml.v3"

type JoinType string

const (
	Inner JoinType = "inner"
	Left  JoinType = "left"
	Right JoinType = "right"
	Full  JoinType = "full"
)

type Join struct {
	Dataset Dataset    `yaml:"dataset" json:"dataset" mapstructure:"dataset"`
	Where   Expression `yaml:"where" json:"where" mapstructure:"where"`
	Retain  JoinType   `yaml:"retain,omitempty" json:"retain,omitempty" mapstructure:"retain,omitempty"`
}

func (j *Join) Type() string {
	return "join"
}

func (j *Join) Validate() error {
	// Add validation logic here
	return nil
}

func (j *Join) UnmarshalYAML(value *yaml.Node) error {
	type rawJoin Join
	if err := value.Decode((*rawJoin)(j)); err != nil {
		return err
	}
	if j.Retain == "" {
		j.Retain = Inner
	}
	return nil
}
