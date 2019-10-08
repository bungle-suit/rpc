package types

import (
	"reflect"

	"github.com/bungle-suit/json"
)

type voidType struct{}

func (voidType) Marshal(w *json.Writer, v interface{}) error {
	w.WriteNull()
	return nil
}

func (voidType) Unmarshal(r *json.Reader, v reflect.Value) error {
	return r.Expect(json.NULL)
}
