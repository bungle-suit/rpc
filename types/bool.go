package types

import (
	"encoding/json"
	"reflect"

	myjson "github.com/bungle-suit/json"
)

type boolType struct{}

func (b boolType) Marshal(w *myjson.Writer, v interface{}) {
	val := v.(bool)
	if val {
		w.WriteNumber(1)
	} else {
		w.WriteNumber(0)
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
