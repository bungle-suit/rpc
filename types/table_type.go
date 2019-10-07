package types

import (
	"encoding/json"
	"reflect"

	myjson "github.com/bungle-suit/json"
)

type tableType struct{}

func (t tableType) Marshal(w *myjson.Writer, v interface{}) {
	panic("not implemented")
}

func (t tableType) Unmarshal(decoder *json.Decoder, v reflect.Value) error {
	panic("not implemented")
}
