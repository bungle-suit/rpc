package types

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"

	myjson "github.com/bungle-suit/json"
)

type longType struct{}

const (
	maxSafeLong = int64(9000000000000000)
	minSafeLong = int64(-9000000000000000)
)

func (longType) Marshal(w *myjson.Writer, v interface{}) error {
	panic("not implemented")
	// val := v.(int64)
	// if val > maxSafeLong || val < minSafeLong {
	// 	return encoder.Encode(strconv.FormatInt(val, 10))
	// }
	// return encoder.Encode(v)
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
