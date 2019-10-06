package types

import (
	"encoding/json"
	"reflect"

	"github.com/francoispqt/gojay"
)

type dictType struct {
	inner Type
}

func (dictType) Marshal(encoder *gojay.Encoder, v interface{}) error {
	panic("not implemented")
}

func (dictType) Unmarshal(decoder *json.Decoder, v reflect.Value) error {
	panic("not implemented")
}
