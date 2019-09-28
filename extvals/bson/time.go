package bson

import (
	"reflect"

	"go.mongodb.org/mongo-driver/bson/bsoncodec"
	"go.mongodb.org/mongo-driver/bson/bsonrw"
)

// bsoncodec package support marshal time.Time, but can not maintains time zone,
// when restore from bson, always set timezone to UTC.
//
// timeCodec set restored timezone to local timezone.
type timeCodec struct{}

func (timeCodec) EncodeValue(ctx bsoncodec.EncodeContext, w bsonrw.ValueWriter, v reflect.Value) error {
	return bsoncodec.DefaultValueEncoders{}.TimeEncodeValue(ctx, w, v)
}

func (timeCodec) DecodeValue(ctx bsoncodec.DecodeContext, r bsonrw.ValueReader, v reflect.Value) error {
	return bsoncodec.DefaultValueDecoders{}.TimeDecodeValue(ctx, r, v)
}
