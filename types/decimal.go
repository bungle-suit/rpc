package types

import (
	"github.com/bungle-suit/json"
	"github.com/bungle-suit/rpc/extvals/decimal"
	"github.com/pkg/errors"
)

type decimalType int

func (d decimalType) Marshal(w *json.Writer, v interface{}) error {
	val := v.(decimal.Decimaller).Decimal().Round(int(d))
	w.WriteString(val.String())
	return nil
}

func (d decimalType) Unmarshal(r *json.Reader) (interface{}, error) {
	s, err := r.ReadString()
	if err != nil {
		return nil, err
	}

	dv, err := decimal.FromStringWithScale(s, int(d))
	if err != nil {
		return nil, err
	}

	switch d {
	case 0:
		return decimal.Decimal0(dv), nil
	case 1:
		return decimal.Decimal1(dv), nil
	case 2:
		return decimal.Decimal2(dv), nil
	case 3:
		return decimal.Decimal3(dv), nil
	case 4:
		return decimal.Decimal4(dv), nil
	case 5:
		return decimal.Decimal5(dv), nil
	case 6:
		return decimal.Decimal6(dv), nil
	case 7:
		return decimal.Decimal7(dv), nil
	case 8:
		return decimal.Decimal8(dv), nil
	default:
		return nil, errors.New("Unknown decimal type")
	}
}
