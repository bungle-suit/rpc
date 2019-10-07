package types

import (
	"reflect"

	"github.com/bungle-suit/json"
)

type dictType struct {
	inner Type
}

func (dictType) Marshal(w *json.Writer, v interface{}) {
	panic("not implemented")
}

func (dictType) Unmarshal(r *json.Reader, v reflect.Value) error {
	panic("not implemented")
}
