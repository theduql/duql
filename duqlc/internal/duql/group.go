package duql

import "gopkg.in/yaml.v3"

type Group struct {
	By    interface{} `yaml:"by" json:"by" mapstructure:"by"`
	Steps []Step      `yaml:"steps,omitempty" json:"steps,omitempty" mapstructure:"steps,omitempty"`
}

func (g *Group) Type() string {
	return "group"
}

func (g *Group) Validate() error {
	// Add validation logic here
	return nil
}

func (g *Group) UnmarshalYAML(value *yaml.Node) error {
	var tmp struct {
		By    interface{} `yaml:"by"`
		Steps Steps       `yaml:"steps,omitempty"`
	}
	if err := value.Decode(&tmp); err != nil {
		return err
	}
	g.By = tmp.By
	g.Steps = tmp.Steps
	return nil
}
