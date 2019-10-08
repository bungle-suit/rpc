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

func (d decimalType) New() interface{} {
	switch int(d) {
	case 0:
		return decimal.Decimal0{}
	case 1:
		return decimal.Decimal1{}
	case 2:
		return decimal.Decimal2{}
	case 3:
		return decimal.Decimal3{}
	case 4:
		return decimal.Decimal4{}
	case 5:
		return decimal.Decimal5{}
	case 6:
		return decimal.Decimal6{}
	case 7:
		return decimal.Decimal7{}
	case 8:
		return decimal.Decimal8{}
	default:
		panic("Unknown decimal type")
	}
}
