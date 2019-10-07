package types

import (
	"encoding/json"
	"reflect"

	myjson "github.com/bungle-suit/json"
)

type listType struct {
	inner Type
}

func (listType) Marshal(w *myjson.Writer, v interface{}) error {
	panic("not implemented")
}

func (listType) Unmarshal(decoder *json.Decoder, v reflect.Value) error {
	panic("not implemented")
}
