package types

import (
	"encoding/json"
	"reflect"

	myjson "github.com/bungle-suit/json"
)

type dictType struct {
	inner Type
}

func (dictType) Marshal(w *myjson.Writer, v interface{}) {
	panic("not implemented")
}

func (dictType) Unmarshal(decoder *json.Decoder, v reflect.Value) error {
	panic("not implemented")
}
