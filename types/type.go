package types

import (
	"encoding/json"
	"reflect"

	"github.com/francoispqt/gojay"
)

type Marshaler interface {
	Marshal(encoder *gojay.Encoder, v interface{}) error
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
	marshal   func(encoder *gojay.Encoder, v interface{}) error
	unmarshal func(decoder *json.Decoder, v reflect.Value) error
}

func (c combinedType) Marshal(encoder *gojay.Encoder, v interface{}) error {
	return c.marshal(encoder, v)
}

func (c combinedType) Unmarshal(decoder *json.Decoder, v reflect.Value) error {
	return c.unmarshal(decoder, v)
}
