package types

import (
	"reflect"

	"github.com/bungle-suit/json"
)

type listType struct {
	inner Type
}

func (listType) Marshal(w *json.Writer, v interface{}) {
	panic("not implemented")
}

func (listType) Unmarshal(r *json.Reader, v reflect.Value) error {
	panic("not implemented")
}
