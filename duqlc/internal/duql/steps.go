package duql

type Step interface {
	Type() string
	Validate() error
}
