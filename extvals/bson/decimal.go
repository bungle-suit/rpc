package bson

import (
	"errors"
	"reflect"

	"github.com/bungle-suit/rpc/extvals/decimal"
	"go.mongodb.org/mongo-driver/bson/bsoncodec"
	"go.mongodb.org/mongo-driver/bson/bsonrw"
	"go.mongodb.org/mongo-driver/bson/bsontype"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// decimalEncoderDecoder implement bsoncodec.ValueDecode/Encoder to
// convert bosn value to/from DecimalX.
type decimalEncoderDecoder struct{}

func readDecimal(r bsonrw.ValueReader) (d decimal.Decimal, err error) {
	dbVal, err := r.ReadDecimal128()
	if err != nil {
		return
	}

	h, l := dbVal.GetBytes()
	return decimal.FromDecimal128(l, h), nil
}

func writeDecimal(w bsonrw.ValueWriter, d decimal.Decimal) error {
	l, h := d.ToDecimal128()
	return w.WriteDecimal128(primitive.NewDecimal128(h, l))
}

func (decimalEncoderDecoder) DecodeValue(ctx bsoncodec.DecodeContext, r bsonrw.ValueReader, v reflect.Value) error {
	d, err := readDecimal(r)
	if err != nil {
		return err
	}

	v.Set(reflect.ValueOf(d))
	return nil
}

func (decimalEncoderDecoder) EncodeValue(ctx bsoncodec.EncodeContext, w bsonrw.ValueWriter, v reflect.Value) error {
	d, ok := v.Interface().(decimal.Decimal)
	if !ok {
		return errors.New("Not decimal value")
	}

	return writeDecimal(w, d)
}

type nullDecimalEncoderDecoder struct{}

func (nullDecimalEncoderDecoder) DecodeValue(ctx bsoncodec.DecodeContext, r bsonrw.ValueReader, v reflect.Value) error {
	if r.Type() == bsontype.Null {
		r.ReadNull()
		v.Set(reflect.ValueOf(decimal.NullDecimal{}))
		return nil
	}

	d, err := readDecimal(r)
	if err != nil {
		return err
	}

	v.Set(reflect.ValueOf(decimal.NullDecimal{d, true}))
	return nil
}

func (nullDecimalEncoderDecoder) EncodeValue(ctx bsoncodec.EncodeContext, w bsonrw.ValueWriter, v reflect.Value) error {
	d, ok := v.Interface().(decimal.NullDecimal)
	if !ok {
		return errors.New("Not null decimal value")
	}

	if !d.Valid {
		return w.WriteNull()
	}

	return writeDecimal(w, d.Decimal)
}

type decimalEncoderDecoderN struct{}

func (decimalEncoderDecoderN) DecodeValue(ctx bsoncodec.DecodeContext, r bsonrw.ValueReader, v reflect.Value) error {
	nd, ok := v.Interface().(decimal.Decimaller)
	if !ok {
		return errors.New("Not decimal value 4")
	}

	d, err := readDecimal(r)
	if err != nil {
		return err
	}
	d = d.Round(int(nd.Scale()))

	val := reflect.ValueOf(d).Convert(v.Type())
	v.Set(val)

	return nil
}

func (decimalEncoderDecoderN) EncodeValue(ctx bsoncodec.EncodeContext, w bsonrw.ValueWriter, v reflect.Value) error {
	d, ok := v.Interface().(decimal.Decimaller)
	if !ok {
		return errors.New("Not decimal value 3")
	}

	return writeDecimal(w, d.Decimal().Round(int(d.Scale())))
}
