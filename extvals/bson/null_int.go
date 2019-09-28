package bson

import (
	"errors"
	"reflect"

	"github.com/bungle-suit/rpc/extvals"
	"go.mongodb.org/mongo-driver/bson/bsoncodec"
	"go.mongodb.org/mongo-driver/bson/bsonrw"
	"go.mongodb.org/mongo-driver/bson/bsontype"
)

type nullInt32Codec struct{}

func (nullInt32Codec) DecodeValue(ctx bsoncodec.DecodeContext, r bsonrw.ValueReader, v reflect.Value) error {
	if r.Type() == bsontype.Null {
		v.Set(reflect.ValueOf(extvals.NullInt32{}))
		return r.ReadNull()
	}

	i, err := r.ReadInt32()
	v.Set(reflect.ValueOf(extvals.NullInt32{i, true}))
	return err
}

func (nullInt32Codec) EncodeValue(ctx bsoncodec.EncodeContext, w bsonrw.ValueWriter, v reflect.Value) error {
	i, ok := v.Interface().(extvals.NullInt32)
	if !ok {
		return errors.New("Not NullInt32")
	}

	if !i.Valid {
		return w.WriteNull()
	}
	return w.WriteInt32(i.Int32)
}

type nullInt64Codec struct{}

func (nullInt64Codec) DecodeValue(ctx bsoncodec.DecodeContext, r bsonrw.ValueReader, v reflect.Value) error {
	if r.Type() == bsontype.Null {
		v.Set(reflect.ValueOf(extvals.NullInt64{}))
		return r.ReadNull()
	}

	i, err := r.ReadInt64()
	v.Set(reflect.ValueOf(extvals.NullInt64{i, true}))
	return err
}

func (nullInt64Codec) EncodeValue(ctx bsoncodec.EncodeContext, w bsonrw.ValueWriter, v reflect.Value) error {
	i, ok := v.Interface().(extvals.NullInt64)
	if !ok {
		return errors.New("Not NullInt64")
	}

	if !i.Valid {
		return w.WriteNull()
	}
	return w.WriteInt64(i.Int64)
}
