package types

import (
	"reflect"

	"github.com/bungle-suit/json"
)

type stringType struct{}

func (stringType) Marshal(w *json.Writer, v interface{}) {
	w.WriteString(v.(string))
}

func (stringType) Unmarshal(r *json.Reader, v reflect.Value) error {
	s, err := r.ReadString()
	v.Elem().SetString(s)
	return err
}
