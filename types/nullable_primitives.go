package types

import (
	"fmt"
	"time"

	"github.com/bungle-suit/json"
	"github.com/bungle-suit/rpc/extvals"
	"github.com/bungle-suit/rpc/extvals/decimal"
)

// isNullToken returns true if next token is json.Null,
// undo reader if not.
func isNullToken(r *json.Reader) bool {
	tt, _ := r.Next()
	if tt == json.Null {
		return true
	}

	r.Undo()
	return false
}

type nullBoolType struct{}

func (nullBoolType) Marshal(w *json.Writer, v interface{}) error {
	val := v.(extvals.NullBool)
	if !val.Valid {
		w.WriteNull()
		return nil
	}

	return boolType{}.Marshal(w, val.V)
}

func (nullBoolType) Unmarshal(r *json.Reader) (v interface{}, err error) {
	if isNullToken(r) {
		return extvals.NullBool{}, nil
	}

	bv, err := boolType{}.Unmarshal(r)
	if err != nil {
		return nil, err
	}
	return extvals.NullBool{V: bv.(bool), Valid: true}, nil
}

type nullIntType struct{}

func (n nullIntType) Marshal(w *json.Writer, v interface{}) error {
	val := v.(extvals.NullInt32)
	if !val.Valid {
		w.WriteNull()
		return nil
	}

	return intType{}.Marshal(w, val.V)
}

func (n nullIntType) Unmarshal(r *json.Reader) (v interface{}, err error) {
	if isNullToken(r) {
		return extvals.NullInt32{}, nil
	}

	bv, err := intType{}.Unmarshal(r)
	if err != nil {
		return nil, err
	}
	return extvals.NullInt32{V: bv.(int32), Valid: true}, nil
}

type nullLongType struct{}

func (n nullLongType) Marshal(w *json.Writer, v interface{}) error {
	val := v.(extvals.NullInt64)
	if !val.Valid {
		w.WriteNull()
		return nil
	}

	return longType{}.Marshal(w, val.V)
}

func (n nullLongType) Unmarshal(r *json.Reader) (v interface{}, err error) {
	if isNullToken(r) {
		return extvals.NullInt64{}, nil
	}

	bv, err := longType{}.Unmarshal(r)
	if err != nil {
		return nil, err
	}
	return extvals.NullInt64{V: bv.(int64), Valid: true}, nil
}

type nullFloatType struct{}

func (n nullFloatType) Marshal(w *json.Writer, v interface{}) error {
	val := v.(extvals.NullFloat64)
	if !val.Valid {
		w.WriteNull()
		return nil
	}

	return floatType{}.Marshal(w, val.V)
}

func (n nullFloatType) Unmarshal(r *json.Reader) (v interface{}, err error) {
	if isNullToken(r) {
		return extvals.NullFloat64{}, nil
	}

	bv, err := floatType{}.Unmarshal(r)
	if err != nil {
		return nil, err
	}
	return extvals.NullFloat64{V: bv.(float64), Valid: true}, nil
}

type nullDatetimeType struct{}

func (n nullDatetimeType) Marshal(w *json.Writer, v interface{}) error {
	val := v.(extvals.NullTime)
	if !val.Valid {
		w.WriteNull()
		return nil
	}

	return datetimeType{}.Marshal(w, val.V)
}

func (n nullDatetimeType) Unmarshal(r *json.Reader) (v interface{}, err error) {
	if isNullToken(r) {
		return extvals.NullTime{}, nil
	}

	bv, err := datetimeType{}.Unmarshal(r)
	if err != nil {
		return nil, err
	}
	return extvals.NullTime{V: bv.(time.Time), Valid: true}, nil
}

type nullDecimalType int

func (n nullDecimalType) Marshal(w *json.Writer, v interface{}) error {
	val := v.(decimal.NullDecimaller).NullDecimal()
	if !val.Valid {
		w.WriteNull()
		return nil
	}

	var dv decimal.Decimaller
	switch int(n) {
	case 0:
		dv = decimal.Decimal0(val.V)
	case 1:
		dv = decimal.Decimal1(val.V)
	case 2:
		dv = decimal.Decimal2(val.V)
	case 3:
		dv = decimal.Decimal3(val.V)
	case 4:
		dv = decimal.Decimal4(val.V)
	case 5:
		dv = decimal.Decimal5(val.V)
	case 6:
		dv = decimal.Decimal6(val.V)
	case 7:
		dv = decimal.Decimal7(val.V)
	case 8:
		dv = decimal.Decimal8(val.V)
	default:
		return fmt.Errorf("[%s] Unknown nullable decimal scale: %d", tag, int(n))
	}

	return decimalType(int(n)).Marshal(w, dv)
}

func (n nullDecimalType) Unmarshal(r *json.Reader) (v interface{}, err error) {
	if isNullToken(r) {
		switch int(n) {
		case 0:
			return decimal.NullDecimal0{}, nil
		case 1:
			return decimal.NullDecimal1{}, nil
		case 2:
			return decimal.NullDecimal2{}, nil
		case 3:
			return decimal.NullDecimal3{}, nil
		case 4:
			return decimal.NullDecimal4{}, nil
		case 5:
			return decimal.NullDecimal5{}, nil
		case 6:
			return decimal.NullDecimal6{}, nil
		case 7:
			return decimal.NullDecimal7{}, nil
		case 8:
			return decimal.NullDecimal8{}, nil
		default:
			return nil, fmt.Errorf("[%s] Unknown nullable decimal scale: %d", tag, int(n))
		}
	}

	bv, err := decimalType(int(n)).Unmarshal(r)
	if err != nil {
		return nil, err
	}
	dv := bv.(decimal.Decimaller).Decimal()
	switch int(n) {
	case 0:
		return decimal.NullDecimal0(
			decimal.NullDecimal{V: dv, Valid: true},
		), nil
	case 1:
		return decimal.NullDecimal1(
			decimal.NullDecimal{V: dv, Valid: true},
		), nil
	case 2:
		return decimal.NullDecimal2(
			decimal.NullDecimal{V: dv, Valid: true},
		), nil
	case 3:
		return decimal.NullDecimal3(
			decimal.NullDecimal{V: dv, Valid: true},
		), nil
	case 4:
		return decimal.NullDecimal4(
			decimal.NullDecimal{V: dv, Valid: true},
		), nil
	case 5:
		return decimal.NullDecimal5(
			decimal.NullDecimal{V: dv, Valid: true},
		), nil
	case 6:
		return decimal.NullDecimal6(
			decimal.NullDecimal{V: dv, Valid: true},
		), nil
	case 7:
		return decimal.NullDecimal7(
			decimal.NullDecimal{V: dv, Valid: true},
		), nil
	case 8:
		return decimal.NullDecimal8(
			decimal.NullDecimal{V: dv, Valid: true},
		), nil
	default:
		return nil, fmt.Errorf("[%s] Unknown nullable decimal scale: %d", tag, int(n))
	}
}
