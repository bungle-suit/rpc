package types

import (
	"encoding/json"
	"reflect"
	"time"

	"github.com/francoispqt/gojay"
)

type datetimeType struct{}

func (datetimeType) Marshal(encoder *gojay.Encoder, v interface{}) error {
	val := v.(time.Time)
	secs := val.Unix()
	return encoder.Encode(secs)
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
