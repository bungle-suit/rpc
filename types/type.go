package types

import (
	"reflect"

	"github.com/bungle-suit/json"
)

type Marshaler interface {
	Marshal(w *json.Writer, v interface{}) error
}

type Unmarshaler interface {
	Unmarshal(r *json.Reader, v reflect.Value) error
}

// Type interface to marshal values of rpc type system.
type Type interface {
	Marshaler
	Unmarshaler
}
