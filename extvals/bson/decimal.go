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

func readNullDecimal(r bsonrw.ValueReader) (d decimal.NullDecimal, err error) {
	if r.Type() == bsontype.Null {
		err = r.ReadNull()
		return
	}

	v, err := readDecimal(r)
	d = decimal.NullDecimal{v, true}
	return
}

func writeNullDecimal(w bsonrw.ValueWriter, d decimal.NullDecimal) error {
	if !d.Valid {
		return w.WriteNull()
	}

	return writeDecimal(w, d.Decimal)
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
	d, err := readNullDecimal(r)
	if err != nil {
		return err
	}

	v.Set(reflect.ValueOf(d))
	return nil
}

func (nullDecimalEncoderDecoder) EncodeValue(ctx bsoncodec.EncodeContext, w bsonrw.ValueWriter, v reflect.Value) error {
	d, ok := v.Interface().(decimal.NullDecimal)
	if !ok {
		return errors.New("Not null decimal value")
	}

	return writeNullDecimal(w, d)
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

type nullDecimalEncoderDecoderN struct{}

func (nullDecimalEncoderDecoderN) DecodeValue(ctx bsoncodec.DecodeContext, r bsonrw.ValueReader, v reflect.Value) error {
	nd, ok := v.Interface().(decimal.NullDecimaller)
	if !ok {
		return errors.New("Not decimal value 5")
	}

	d, err := readNullDecimal(r)
	if err != nil {
		return err
	}
	if d.Valid && d.Decimal.Scale() != nd.Scale() {
		d = decimal.NullDecimal{
			d.Decimal.Round(int(nd.Scale())),
			d.Valid,
		}
	}

	val := reflect.ValueOf(d).Convert(v.Type())
	v.Set(val)
	return nil
}

func (nullDecimalEncoderDecoderN) EncodeValue(ctx bsoncodec.EncodeContext, w bsonrw.ValueWriter, v reflect.Value) error {
	nd, ok := v.Interface().(decimal.NullDecimaller)
	if !ok {
		return errors.New("Not decimal value 6")
	}

	d := nd.NullDecimal()
	if d.Valid && d.Decimal.Scale() != nd.Scale() {
		d = decimal.NullDecimal{
			d.Decimal.Round(int(nd.Scale())),
			d.Valid,
		}
	}

	return writeNullDecimal(w, d)
}
