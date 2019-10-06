package types

import (
	"encoding/json"
	"reflect"

	"github.com/francoispqt/gojay"
)

type listType struct {
	inner Type
}

func (listType) Marshal(encoder *gojay.Encoder, v interface{}) error {
	panic("not implemented")
}

func (listType) Unmarshal(decoder *json.Decoder, v reflect.Value) error {
	panic("not implemented")
}
