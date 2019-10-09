package types

import (
	"time"

	"github.com/bungle-suit/json"
	"github.com/bungle-suit/rpc/extvals"
)

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

type nullIntType struct{}

func (n nullIntType) Marshal(w *json.Writer, v interface{}) error {
	val := v.(extvals.NullInt32)
	if !val.Valid {
		w.WriteNull()
		return nil
	}

	return intType{}.Marshal(w, val.V)
}

func (n nullIntType) Unmarshal(r *json.Reader) (v interface{}, err error) {
	if isNullToken(r) {
		return extvals.NullInt32{}, nil
	}

	bv, err := intType{}.Unmarshal(r)
	if err != nil {
		return nil, err
	}
	return extvals.NullInt32{V: bv.(int32), Valid: true}, nil
}

type nullLongType struct{}

func (n nullLongType) Marshal(w *json.Writer, v interface{}) error {
	val := v.(extvals.NullInt64)
	if !val.Valid {
		w.WriteNull()
		return nil
	}

	return longType{}.Marshal(w, val.V)
}

func (n nullLongType) Unmarshal(r *json.Reader) (v interface{}, err error) {
	if isNullToken(r) {
		return extvals.NullInt64{}, nil
	}

	bv, err := longType{}.Unmarshal(r)
	if err != nil {
		return nil, err
	}
	return extvals.NullInt64{V: bv.(int64), Valid: true}, nil
}

type nullFloatType struct{}

func (n nullFloatType) Marshal(w *json.Writer, v interface{}) error {
	val := v.(extvals.NullFloat64)
	if !val.Valid {
		w.WriteNull()
		return nil
	}

	return floatType{}.Marshal(w, val.V)
}

func (n nullFloatType) Unmarshal(r *json.Reader) (v interface{}, err error) {
	if isNullToken(r) {
		return extvals.NullFloat64{}, nil
	}

	bv, err := floatType{}.Unmarshal(r)
	if err != nil {
		return nil, err
	}
	return extvals.NullFloat64{V: bv.(float64), Valid: true}, nil
}

type nullDatetimeType struct{}

func (n nullDatetimeType) Marshal(w *json.Writer, v interface{}) error {
	val := v.(extvals.NullTime)
	if !val.Valid {
		w.WriteNull()
		return nil
	}

	return datetimeType{}.Marshal(w, val.V)
}

func (n nullDatetimeType) Unmarshal(r *json.Reader) (v interface{}, err error) {
	if isNullToken(r) {
		return extvals.NullTime{}, nil
	}

	bv, err := datetimeType{}.Unmarshal(r)
	if err != nil {
		return nil, err
	}
	return extvals.NullTime{V: bv.(time.Time), Valid: true}, nil
}
