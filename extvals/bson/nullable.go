package bson

import (
	"errors"
	"reflect"

	"github.com/bungle-suit/rpc/extvals"
	"go.mongodb.org/mongo-driver/bson/bsoncodec"
	"go.mongodb.org/mongo-driver/bson/bsonrw"
	"go.mongodb.org/mongo-driver/bson/bsontype"
)

var (
	trueVal = reflect.ValueOf(true)
)

type nullableCodec struct {
	inner     bsoncodec.ValueCodec
	innerType reflect.Type
}

type combinedCodec struct {
	encoder bsoncodec.ValueEncoder
	decoder bsoncodec.ValueDecoder
}

func (c combinedCodec) EncodeValue(ctx bsoncodec.EncodeContext, w bsonrw.ValueWriter, v reflect.Value) error {
	return c.encoder.EncodeValue(ctx, w, v)
}

func (c combinedCodec) DecodeValue(ctx bsoncodec.DecodeContext, r bsonrw.ValueReader, v reflect.Value) error {
	return c.decoder.DecodeValue(ctx, r, v)
}

func newNullableCodec(notNullRegistry *bsoncodec.Registry, innerType reflect.Type) bsoncodec.ValueCodec {
	decoder, err := notNullRegistry.LookupDecoder(innerType)
	if err != nil {
		panic(err)
	}

	encoder, err := notNullRegistry.LookupEncoder(innerType)
	if err != nil {
		panic(err)
	}

	return nullableCodec{
		inner: combinedCodec{
			encoder, decoder,
		},
		innerType: innerType,
	}
}

func (nc nullableCodec) DecodeValue(ctx bsoncodec.DecodeContext, r bsonrw.ValueReader, v reflect.Value) error {
	if r.Type() == bsontype.Null {
		v.Set(reflect.Zero(v.Type()))
		return r.ReadNull()
	}

	if err := nc.inner.DecodeValue(ctx, r, v.Field(0)); err != nil {
		return err
	}
	v.Field(1).Set(trueVal)
	return nil
}

func (nc nullableCodec) EncodeValue(ctx bsoncodec.EncodeContext, w bsonrw.ValueWriter, v reflect.Value) error {
	n, ok := v.Interface().(extvals.Nullable)
	if !ok {
		return errors.New("not Nullable")
	}

	if n.IsNull() {
		return w.WriteNull()
	}
	return nc.inner.EncodeValue(ctx, w, v.Field(0))
}
