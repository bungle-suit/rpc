package types

import (
	"github.com/bungle-suit/json"
)

type voidType struct{}

func (voidType) Marshal(w *json.Writer, v interface{}) error {
	w.WriteNull()
	return nil
}

func (voidType) Unmarshal(r *json.Reader) (interface{}, error) {
	return nil, r.Expect(json.NULL)
}
