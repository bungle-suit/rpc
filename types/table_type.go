package types

import (
	"encoding/json"
	"reflect"

	"github.com/francoispqt/gojay"
)

type tableType struct{}

func (t tableType) Marshal(encoder *gojay.Encoder, v interface{}) error {
	panic("not implemented")
}

func (t tableType) Unmarshal(decoder *json.Decoder, v reflect.Value) error {
	panic("not implemented")
}
