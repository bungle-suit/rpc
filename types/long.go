package types

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/bungle-suit/json"
)

type intType struct{}

func (intType) Marshal(w *json.Writer, v interface{}) error {
	val := v.(int32)
	w.WriteNumber(float64(val))
	return nil
}

func (intType) Unmarshal(r *json.Reader, v reflect.Value) error {
	fv, err := r.ReadNumber()
	if err != nil {
		return err
	}

	// TODO: check fv do not have decimal part and in range.
	v.Elem().SetInt(int64(fv))
	return nil
}

func (intType) New() interface{} {
	return int32(0)
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

func (longType) Unmarshal(r *json.Reader, v reflect.Value) error {
	tt, err := r.Next()
	if err != nil {
		return err
	}

	var s string
	if tt == json.NUMBER {
		s = string(r.Buf[r.Start:r.End])
	} else if tt == json.STRING {
		s = string(r.Buf[r.Start+1 : r.End-1])
	} else {
		return fmt.Errorf("[%s] Unexpected long type", tag)
	}

	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return fmt.Errorf("[%s] Failed unmarshal '%s' to long: %w", tag, s, err)
	}
	v.Elem().SetInt(i)
	return nil
}

func (longType) New() interface{} {
	return int64(0)
}

type floatType struct{}

func (floatType) Marshal(w *json.Writer, v interface{}) error {
	val := v.(float64)
	w.WriteNumber(val)
	return nil
}

func (floatType) Unmarshal(r *json.Reader, v reflect.Value) error {
	fv, err := r.ReadNumber()
	if err != nil {
		return err
	}

	v.Elem().SetFloat(fv)
	return nil
}

func (floatType) New() interface{} {
	return float64(0)
}
