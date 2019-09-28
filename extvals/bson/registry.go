package bson

import (
	"reflect"

	"github.com/bungle-suit/rpc/extvals/decimal"
	"go.mongodb.org/mongo-driver/bson/bsoncodec"
)

var (
	// Registry is bson encoder/decoder for marshal value between mongo-db,
	// contains standard types and my extsion types.
	Registry *bsoncodec.Registry
)

func init() {
	builder := bsoncodec.NewRegistryBuilder()
	bsoncodec.DefaultValueEncoders{}.RegisterDefaultEncoders(builder)
	bsoncodec.DefaultValueDecoders{}.RegisterDefaultDecoders(builder)

	builder.
		RegisterCodec(reflect.TypeOf(decimal.Decimal{}), decimalEncoderDecoder{}).
		RegisterCodec(reflect.TypeOf(decimal.NullDecimal{}), nullDecimalEncoderDecoder{})
	Registry = builder.Build()
}
