package types

import (
	"encoding/json"
	"reflect"

	myjson "github.com/bungle-suit/json"
)

type nullType struct {
	inner Type
}

func (nullType) Marshal(w *myjson.Writer, v interface{}) error {
	panic("not implemented")
}

func (nullType) Unmarshal(decoder *json.Decoder, v reflect.Value) error {
	panic("not implemented")
}
