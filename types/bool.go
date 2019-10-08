package types

import (
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

func (b boolType) Unmarshal(r *json.Reader) (interface{}, error) {
	fv, err := r.ReadNumber()
	if err != nil {
		return nil, err
	}

	return fv != 0, nil
}
