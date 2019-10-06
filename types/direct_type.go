package types

import (
	"encoding/json"
	"reflect"

	"github.com/francoispqt/gojay"
)

type directType struct{}

func (directType) Marshal(encoder *gojay.Encoder, v interface{}) error {
	return encoder.Encode(v)
}

func (directType) Unmarshal(decoder *json.Decoder, v reflect.Value) error {
	return decoder.Decode(v.Interface())
}
