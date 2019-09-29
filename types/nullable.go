package types

import (
	"encoding/json"
	"reflect"
)

type nullType struct {
	inner Type
}

func (nullType) Marshal(encoder *json.Encoder, v interface{}) error {
	panic("not implemented")
}

func (nullType) Unmarshal(decoder *json.Decoder, v reflect.Value) error {
	panic("not implemented")
}
