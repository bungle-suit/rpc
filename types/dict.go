package types

import (
	"encoding/json"
	"reflect"
)

type dictType struct {
	inner Type
}

func (dictType) Marshal(encoder *json.Encoder, v interface{}) error {
	panic("not implemented")
}

func (dictType) Unmarshal(decoder *json.Decoder, v reflect.Value) error {
	panic("not implemented")
}
