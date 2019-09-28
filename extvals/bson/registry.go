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

	types := []reflect.Type{
		reflect.TypeOf(decimal.Decimal0{}),
		reflect.TypeOf(decimal.Decimal1{}),
		reflect.TypeOf(decimal.Decimal2{}),
		reflect.TypeOf(decimal.Decimal3{}),
		reflect.TypeOf(decimal.Decimal4{}),
		reflect.TypeOf(decimal.Decimal5{}),
		reflect.TypeOf(decimal.Decimal6{}),
		reflect.TypeOf(decimal.Decimal7{}),
		reflect.TypeOf(decimal.Decimal8{}),
	}
	for _, t := range types {
		builder.RegisterCodec(t, decimalEncoderDecoderN{})
	}

	Registry = builder.Build()
}
