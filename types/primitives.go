package types

import (
	"fmt"
	"strconv"
	"time"

	"github.com/bungle-suit/json"
	"github.com/bungle-suit/rpc/extvals/decimal"
	"github.com/pkg/errors"
)

type boolType struct{}

func (b boolType) Marshal(w *json.Writer, v interface{}) error {
	val := v.(bool)
	if val {
		w.WriteNumber(1)
	} else {
		w.WriteNumber(0)
	}
	return nil
}

func (b boolType) Unmarshal(r *json.Reader) (interface{}, error) {
	fv, err := r.ReadNumber()
	if err != nil {
		return nil, err
	}

	return fv != 0, nil
}

type intType struct{}

func (intType) Marshal(w *json.Writer, v interface{}) error {
	val := v.(int32)
	w.WriteNumber(float64(val))
	return nil
}

func (intType) Unmarshal(r *json.Reader) (interface{}, error) {
	fv, err := r.ReadNumber()
	return int32(fv), err
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

func (longType) Unmarshal(r *json.Reader) (interface{}, error) {
	tt, err := r.Next()
	if err != nil {
		return nil, err
	}

	var s string
	switch tt {
	case json.Number:
		s = string(r.Buf[r.Start:r.End])
	case json.String:
		s = string(r.Buf[r.Start+1 : r.End-1])
	default:
		return nil, fmt.Errorf("[%s] Unexpected long type", tag)
	}

	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("[%s] Failed unmarshal '%s' to long: %w", tag, s, err)
	}
	return i, nil
}

type floatType struct{}

func (floatType) Marshal(w *json.Writer, v interface{}) error {
	val := v.(float64)
	w.WriteNumber(val)
	return nil
}

func (floatType) Unmarshal(r *json.Reader) (interface{}, error) {
	return r.ReadNumber()
}

type stringType struct{}

func (stringType) Marshal(w *json.Writer, v interface{}) error {
	w.WriteString(v.(string))
	return nil
}

func (stringType) Unmarshal(r *json.Reader) (interface{}, error) {
	return r.ReadString()
}

type datetimeType struct{}

func (datetimeType) Marshal(w *json.Writer, v interface{}) error {
	val := v.(time.Time)
	secs := val.Unix()
	w.WriteNumber(float64(secs))
	return nil
}

func (datetimeType) Unmarshal(r *json.Reader) (interface{}, error) {
	fv, err := r.ReadNumber()
	if err != nil {
		return nil, err
	}

	return time.Unix(int64(fv), 0), nil
}

type voidType struct{}

func (voidType) Marshal(w *json.Writer, v interface{}) error {
	w.WriteNull()
	return nil
}

func (voidType) Unmarshal(r *json.Reader) (interface{}, error) {
	return nil, r.Expect(json.Null)
}

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
