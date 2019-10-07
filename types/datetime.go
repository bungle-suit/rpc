package types

import (
	"encoding/json"
	"reflect"
	"time"

	myjson "github.com/bungle-suit/json"
)

type datetimeType struct{}

func (datetimeType) Marshal(w *myjson.Writer, v interface{}) error {
	panic("not implemented")
	// val := v.(time.Time)
	// secs := val.Unix()
	// return encoder.Encode(secs)
}

func (datetimeType) Unmarshal(decoder *json.Decoder, v reflect.Value) error {
	var val int64
	if err := decoder.Decode(&val); err != nil {
		return err
	}

	t := time.Unix(val, 0)
	v.Elem().Set(reflect.ValueOf(t))
	return nil
}
