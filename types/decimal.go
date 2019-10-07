package types

import (
	"reflect"

	"github.com/bungle-suit/json"
	"github.com/bungle-suit/rpc/extvals/decimal"
)

type decimalType int

func (d decimalType) Marshal(w *json.Writer, v interface{}) error {
	val := v.(decimal.Decimaller).Decimal().Round(int(d))
	w.WriteString(val.String())
	return nil
}

func (d decimalType) Unmarshal(r *json.Reader, v reflect.Value) error {
	s, err := r.ReadString()
	if err != nil {
		return err
	}

	dv, err := decimal.FromStringWithScale(s, int(d))
	if err != nil {
		return err
	}
	v.Elem().Set(reflect.ValueOf(dv).Convert(v.Type().Elem()))
	return nil
}
