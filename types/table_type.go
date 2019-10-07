package types

import (
	"reflect"

	"github.com/bungle-suit/json"
)

type tableType struct{}

func (t tableType) Marshal(w *json.Writer, v interface{}) {
	panic("not implemented")
}

func (t tableType) Unmarshal(r *json.Reader, v reflect.Value) error {
	panic("not implemented")
}
