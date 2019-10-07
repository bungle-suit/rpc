package types

import (
	"encoding/json"
	"reflect"

	myjson "github.com/bungle-suit/json"
)

type Marshaler interface {
	Marshal(w *myjson.Writer, v interface{}) error
}

type Unmarshaler interface {
	Unmarshal(decoder *json.Decoder, v reflect.Value) error
}

// Type interface to marshal values of rpc type system.
type Type interface {
	Marshaler
	Unmarshaler
}

type combinedType struct {
	marshal   func(w *myjson.Writer, v interface{}) error
	unmarshal func(decoder *json.Decoder, v reflect.Value) error
}

func (c combinedType) Marshal(w *myjson.Writer, v interface{}) error {
	return c.marshal(w, v)
}

func (c combinedType) Unmarshal(decoder *json.Decoder, v reflect.Value) error {
	return c.unmarshal(decoder, v)
}
