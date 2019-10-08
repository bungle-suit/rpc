package types

import (
	"reflect"

	"github.com/bungle-suit/json"
)

type listType struct {
	inner Type
}

func (listType) Marshal(w *json.Writer, v interface{}) error {
	panic("not implemented")
}

func (listType) Unmarshal(r *json.Reader, v reflect.Value) error {
	panic("not implemented")
}

func (listType) New() interface{} {
	panic("not implemented")
}
