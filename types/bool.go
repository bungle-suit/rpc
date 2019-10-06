package types

import (
	"encoding/json"
	"reflect"

	"github.com/francoispqt/gojay"
)

type boolType struct{}

func (b boolType) Marshal(encoder *gojay.Encoder, v interface{}) error {
	val := v.(bool)
	if val {
		return encoder.Encode(1)
	} else {
		return encoder.Encode(0)
	}
}

func (b boolType) Unmarshal(decoder *json.Decoder, v reflect.Value) error {
	var i int
	if err := decoder.Decode(&i); err != nil {
		return err
	}

	bv := i != 0
	v.Elem().SetBool(bv)
	return nil
}
