package types

import (
	"encoding/json"
	"reflect"
)

type listType struct {
	inner Type
}

func (listType) Marshal(encoder *json.Encoder, v interface{}) error {
	panic("not implemented")
}

func (listType) Unmarshal(decoder *json.Decoder, v reflect.Value) error {
	panic("not implemented")
}
