package duql

type Filter struct {
	Expression `json:"filter" yaml:"filter"`
}

func (f *Filter) Type() string {
	return "filter"
}

func (f *Filter) Validate() error {
	return f.Expression.Validate()
}
