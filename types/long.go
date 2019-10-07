package types

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"

	myjson "github.com/bungle-suit/json"
)

type intType struct{}

func (intType) Marshal(w *myjson.Writer, v interface{}) {
	val := v.(int32)
	w.WriteNumber(float64(val))
}

func (intType) Unmarshal(decoder *json.Decoder, v reflect.Value) error {
	return decoder.Decode(v.Interface())
}

type longType struct{}

const (
	maxSafeLong = int64(9000000000000000)
	minSafeLong = int64(-9000000000000000)
)

func (longType) Marshal(w *myjson.Writer, v interface{}) {
	val := v.(int64)
	if val > maxSafeLong || val < minSafeLong {
		w.WriteNumber(float64(val))
	} else {
		w.WriteString(strconv.FormatInt(val, 10))
	}
}

func (longType) Unmarshal(decoder *json.Decoder, v reflect.Value) error {
	tok, err := decoder.Token()
	if err != nil {
		return err
	}

	switch val := tok.(type) {
	case string:
		i, err := strconv.ParseInt(val, 10, 64)
		if err != nil {
			return fmt.Errorf("[%s] Failed unmarshal '%s' to long: %w", tag, tok, err)
		}
		v.Elem().SetInt(i)

	case float64:
		v.Elem().SetInt(int64(val))

	default:
		return fmt.Errorf("[%s] Failed unmarshal '%v' to long", tag, tok)
	}
	return nil
}

type floatType struct{}

func (floatType) Marshal(w *myjson.Writer, v interface{}) {
	val := v.(float64)
	w.WriteNumber(val)
}

func (floatType) Unmarshal(decoder *json.Decoder, v reflect.Value) error {
	return decoder.Decode(v.Interface())
}
