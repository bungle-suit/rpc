package types

import (
	"github.com/bungle-suit/json"
	"github.com/bungle-suit/rpc/extvals"
)

type nullBoolType struct{}

func (nullBoolType) Marshal(w *json.Writer, v interface{}) error {
	val := v.(extvals.NullBool)
	if !val.Valid {
		w.WriteNull()
		return nil
	}

	return boolType{}.Marshal(w, val.V)
}

func (nullBoolType) Unmarshal(r *json.Reader) (v interface{}, err error) {
	if isNullToken(r) {
		return extvals.NullBool{}, nil
	}

	bv, err := boolType{}.Unmarshal(r)
	if err != nil {
		return nil, err
	}
	return extvals.NullBool{V: bv.(bool), Valid: true}, nil
}

// isNullToken returns true if next token is json.Null,
// undo reader if not.
func isNullToken(r *json.Reader) bool {
	tt, _ := r.Next()
	if tt == json.Null {
		return true
	}

	r.Undo()
	return false
}
