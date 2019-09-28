package extvals

// Nullable is the common interface implemented by all NullXXX types.
//
// For easier reflect operation, implementation struct should have shape:
//
//  V: wrapped value
//  Valid: bool
//
// Unlike database/sql package NullXXX types, such as NullInt32 have fields:
//
//  Int32 int32
//  Valid bool
type Nullable interface {
	// IsNull returns true if value is null, i.e. .Value field is false.
	IsNull() bool

	// Val returns wrapped value.
	Val() interface{}
}
