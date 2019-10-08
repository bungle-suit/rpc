package types

import (
	"reflect"

	"github.com/bungle-suit/json"
)

type boolType struct{}

func (b boolType) Marshal(w *json.Writer, v interface{}) error {
	val := v.(bool)
	if val {
		w.WriteNumber(1)
	} else {
		w.WriteNumber(0)
	}
	return nil
}

func (b boolType) Unmarshal(r *json.Reader, v reflect.Value) error {
	fv, err := r.ReadNumber()
	if err != nil {
		return err
	}

	bv := fv != 0
	v.Elem().SetBool(bv)
	return nil
}
