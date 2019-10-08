package types

import (
	"github.com/bungle-suit/json"
)

type nullType struct {
	inner Type
}

func (nullType) Marshal(w *json.Writer, v interface{}) error {
	panic("not implemented")
}

func (nullType) Unmarshal(r *json.Reader) (interface{}, error) {
	panic("not implemented")
}
