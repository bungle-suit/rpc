package types

import (
	"reflect"

	"github.com/bungle-suit/json"
)

type listType struct {
	inner Type
}

func (i listType) Marshal(w *json.Writer, val interface{}) error {
	v := reflect.ValueOf(val)
	w.BeginArray()
	for idx, l := 0, v.Len(); idx < l; idx++ {
		if err := i.inner.Marshal(w, v.Index(idx).Interface()); err != nil {
			return err
		}
	}
	w.EndArray()
	return nil
}

func (i listType) Unmarshal(r *json.Reader) (interface{}, error) {
	panic("not implemented")
	// if err := r.Expect(json.BeginArray); err != nil {
	// 	return nil, err
	// }

	// list := reflect.MakeSlice(i.Type(), 0, 0)
	// for t, err := r.Next(); t != json.EndArray; t, err = r.Next() {
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	r.Undo()
	// 	if v, err := i.inner.Unmarshal(r); err != nil {
	// 		return nil, err
	// 	} else {
	// 		list = reflect.Append(list, reflect.ValueOf(v))
	// 	}
	// }
	// return list.Interface(), nil
}
