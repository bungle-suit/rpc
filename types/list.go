package types

import (
	"github.com/bungle-suit/json"
)

type listType struct {
	inner Type
}

func (listType) Marshal(w *json.Writer, v interface{}) error {
	panic("not implemented")
}

func (listType) Unmarshal(r *json.Reader) (interface{}, error) {
	panic("not implemented")
}
