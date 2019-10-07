package types

import (
	"encoding/json"
	"reflect"

	myjson "github.com/bungle-suit/json"
)

type boolType struct{}

func (b boolType) Marshal(w *myjson.Writer, v interface{}) error {
	panic("not implemented")
	// val := v.(bool)
	// if val {
	// 	return encoder.Encode(1)
	// } else {
	// 	return encoder.Encode(0)
	// }
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
