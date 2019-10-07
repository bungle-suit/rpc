package types

import (
	"reflect"

	"github.com/bungle-suit/json"
)

type nullType struct {
	inner Type
}

func (nullType) Marshal(w *json.Writer, v interface{}) {
	panic("not implemented")
}

func (nullType) Unmarshal(r *json.Reader, v reflect.Value) error {
	panic("not implemented")
}
