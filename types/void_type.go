package types

import (
	"encoding/json"
	"reflect"

	myjson "github.com/bungle-suit/json"
)

type voidType struct{}

func (voidType) Marshal(w *myjson.Writer, v interface{}) error {
	panic("not implemented")
}

func (voidType) Unmarshal(decoder *json.Decoder, v reflect.Value) error {
	panic("not implemented")
}
