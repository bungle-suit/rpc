package types

import (
	"fmt"
	"strconv"

	"github.com/bungle-suit/json"
)

type intType struct{}

func (intType) Marshal(w *json.Writer, v interface{}) error {
	val := v.(int32)
	w.WriteNumber(float64(val))
	return nil
}

func (intType) Unmarshal(r *json.Reader) (interface{}, error) {
	fv, err := r.ReadNumber()
	return int32(fv), err
}

type longType struct{}

const (
	maxSafeLong = int64(9000000000000000)
	minSafeLong = int64(-9000000000000000)
)

func (longType) Marshal(w *json.Writer, v interface{}) error {
	val := v.(int64)
	if val > maxSafeLong || val < minSafeLong {
		w.WriteString(strconv.FormatInt(val, 10))
	} else {
		w.WriteNumber(float64(val))
	}
	return nil
}

func (longType) Unmarshal(r *json.Reader) (interface{}, error) {
	tt, err := r.Next()
	if err != nil {
		return nil, err
	}

	var s string
	if tt == json.NUMBER {
		s = string(r.Buf[r.Start:r.End])
	} else if tt == json.STRING {
		s = string(r.Buf[r.Start+1 : r.End-1])
	} else {
		return nil, fmt.Errorf("[%s] Unexpected long type", tag)
	}

	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("[%s] Failed unmarshal '%s' to long: %w", tag, s, err)
	}
	return i, nil
}

type floatType struct{}

func (floatType) Marshal(w *json.Writer, v interface{}) error {
	val := v.(float64)
	w.WriteNumber(val)
	return nil
}

func (floatType) Unmarshal(r *json.Reader) (interface{}, error) {
	return r.ReadNumber()
}
