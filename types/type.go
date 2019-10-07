package types

import (
	"encoding/json"
	"reflect"

	myjson "github.com/bungle-suit/json"
)

type Marshaler interface {
	Marshal(w *myjson.Writer, v interface{})
}

type Unmarshaler interface {
	Unmarshal(decoder *json.Decoder, v reflect.Value) error
}

// Type interface to marshal values of rpc type system.
type Type interface {
	Marshaler
	Unmarshaler
}
