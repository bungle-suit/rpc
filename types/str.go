package types

import (
	"encoding/json"
	"reflect"

	myjson "github.com/bungle-suit/json"
)

type stringType struct{}

func (stringType) Marshal(w *myjson.Writer, v interface{}) {
	w.WriteString(v.(string))
}

func (stringType) Unmarshal(decoder *json.Decoder, v reflect.Value) error {
	return decoder.Decode(v.Interface())
}
