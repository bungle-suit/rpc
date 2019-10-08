package types

import (
	"github.com/bungle-suit/json"
)

type dictType struct {
	inner Type
}

func (dictType) Marshal(w *json.Writer, v interface{}) error {
	panic("not implemented")
}

func (dictType) Unmarshal(r *json.Reader) (interface{}, error) {
	panic("not implemented")
}
