package types

import (
	"encoding/json"
	"reflect"
)

type directType struct{}

func (directType) Marshal(encoder *json.Encoder, v interface{}) error {
	return encoder.Encode(v)
}

func (directType) Unmarshal(decoder *json.Decoder, v reflect.Value) error {
	return decoder.Decode(v.Interface())
}
