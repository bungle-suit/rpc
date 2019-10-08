package types

import (
	"reflect"

	"github.com/bungle-suit/json"
)

// Type interface to marshal values of rpc type system.
type Type interface {
	Marshal(w *json.Writer, v interface{}) error
	Unmarshal(r *json.Reader, v reflect.Value) error

	// Create empty value of this type
	New() interface{}
}
