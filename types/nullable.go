package types

import (
	"encoding/json"
	"reflect"

	"github.com/francoispqt/gojay"
)

type nullType struct {
	inner Type
}

func (nullType) Marshal(encoder *gojay.Encoder, v interface{}) error {
	panic("not implemented")
}

func (nullType) Unmarshal(decoder *json.Decoder, v reflect.Value) error {
	panic("not implemented")
}
