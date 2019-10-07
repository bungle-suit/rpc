package types

import (
	"encoding/json"
	"reflect"

	myjson "github.com/bungle-suit/json"
	"github.com/bungle-suit/rpc/extvals/decimal"
)

type decimalType int

func (d decimalType) Marshal(w *myjson.Writer, v interface{}) {
	val := v.(decimal.Decimaller).Decimal().Round(int(d))
	w.WriteString(val.String())
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
