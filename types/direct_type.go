package types

import (
	"encoding/json"
	"reflect"

	myjson "github.com/bungle-suit/json"
)

type directType struct{}

func (directType) Marshal(w *myjson.Writer, v interface{}) error {
	panic("not implemented")
	// return encoder.Encode(v)
}

func (directType) Unmarshal(decoder *json.Decoder, v reflect.Value) error {
	return decoder.Decode(v.Interface())
}
