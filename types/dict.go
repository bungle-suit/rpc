package types

import (
	"reflect"

	"github.com/bungle-suit/json"
)

type dictType struct {
	*Parser
	inner Type
	ts    string
}

func (d dictType) Marshal(w *json.Writer, val interface{}) error {
	m := reflect.ValueOf(val)

	w.BeginObject()
	for _, k := range m.MapKeys() {
		v := m.MapIndex(k)
		w.WriteName(k.String())
		if err := d.inner.Marshal(w, v.Interface()); err != nil {
			return err
		}
	}
	w.EndObject()
	return nil
}

func (d dictType) Unmarshal(r *json.Reader) (interface{}, error) {
	if err := r.Expect(json.BeginObject); err != nil {
		return nil, err
	}

	goType, err := d.ParseGoType(d.ts)
	if err != nil {
		return nil, err
	}
	dict := reflect.MakeMap(goType)
	for tt, err := r.Next(); tt != json.EndObject; tt, err = r.Next() {
		if err != nil {
			return nil, err
		}
		// spork/json ensure here must be property name,
		// and other Types are Implemented correctly.
		name := string(r.Str)
		val, err := d.inner.Unmarshal(r)
		if err != nil {
			return nil, err
		}

		dict.SetMapIndex(reflect.ValueOf(name), reflect.ValueOf(val))
	}
	return dict.Interface(), nil
}
