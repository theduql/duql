package duql

type Dataset struct {
	Simple  string
	Complex *DatasetComplex
}

type DatasetComplex struct {
	Name   string     `yaml:"name,omitempty" json:"name,omitempty" mapstructure:"name,omitempty"`
	Format DataFormat `yaml:"format,omitempty" json:"format,omitempty" mapstructure:"format,omitempty"`
}

type DataFormat string

const (
	Table   DataFormat = "table"
	CSV     DataFormat = "csv"
	JSON    DataFormat = "json"
	Parquet DataFormat = "parquet"
)

// UnmarshalYAML implements the yaml.Unmarshaler interface.
func (d *Dataset) UnmarshalYAML(unmarshal func(interface{}) error) error {
	// Try to unmarshal as a string first
	var s string
	if err := unmarshal(&s); err == nil {
		d.Simple = s
		return nil
	}

	// If that fails, try to unmarshal as a complex object
	var c DatasetComplex
	if err := unmarshal(&c); err != nil {
		return err
	}
	d.Complex = &c
	return nil
}

// MarshalYAML implements the yaml.Marshaler interface.
func (d Dataset) MarshalYAML() (interface{}, error) {
	if d.Simple != "" {
		return d.Simple, nil
	}
	return d.Complex, nil
}
