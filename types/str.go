package types

import (
	"github.com/bungle-suit/json"
)

type stringType struct{}

func (stringType) Marshal(w *json.Writer, v interface{}) error {
	w.WriteString(v.(string))
	return nil
}

func (stringType) Unmarshal(r *json.Reader) (interface{}, error) {
	return r.ReadString()
}
