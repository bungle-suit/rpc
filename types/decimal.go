package types

import (
	"encoding/json"
	"reflect"

	"github.com/bungle-suit/rpc/extvals/decimal"
	"github.com/francoispqt/gojay"
)

type decimalType int

func (d decimalType) Marshal(encoder *gojay.Encoder, v interface{}) error {
	val := v.(decimal.Decimaller).Decimal().Round(int(d))
	s := val.String()
	return encoder.Encode(s)
}

func (d decimalType) Unmarshal(decoder *json.Decoder, v reflect.Value) error {
	var s string
	if err := decoder.Decode(&s); err != nil {
		return err
	}

	dv, err := decimal.FromStringWithScale(s, int(d))
	if err != nil {
		return err
	}
	v.Elem().Set(reflect.ValueOf(dv).Convert(v.Type().Elem()))
	return nil
}
