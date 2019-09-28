package extvals

// Nullable is the common interface implemented by all NullXXX types.
type Nullable interface {
	// IsNull returns true if value is null, i.e. .Value field is false.
	IsNull() bool

	// Val returns wrapped value.
	Val() interface{}
}
