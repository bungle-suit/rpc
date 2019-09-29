package bson

import (
	"reflect"
	"time"

	"github.com/bungle-suit/rpc/extvals"
	"github.com/bungle-suit/rpc/extvals/decimal"
	"go.mongodb.org/mongo-driver/bson/bsoncodec"
)

// Registry creates bsoncodec registry to support rpc extension types.
func Registry() *bsoncodec.Registry {
	builder := bsoncodec.NewRegistryBuilder()
	bsoncodec.DefaultValueEncoders{}.RegisterDefaultEncoders(builder)
	bsoncodec.DefaultValueDecoders{}.RegisterDefaultDecoders(builder)

	builder.
		RegisterCodec(reflect.TypeOf(decimal.Decimal{}), decimalEncoderDecoder{})

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

	registryNotNull := builder.Build()

	builder.RegisterCodec(
		reflect.TypeOf(decimal.NullDecimal{}), newNullableCodec(registryNotNull, reflect.TypeOf(decimal.Decimal{})))

	types = []reflect.Type{
		reflect.TypeOf(decimal.NullDecimal0{}),
		reflect.TypeOf(decimal.NullDecimal1{}),
		reflect.TypeOf(decimal.NullDecimal2{}),
		reflect.TypeOf(decimal.NullDecimal3{}),
		reflect.TypeOf(decimal.NullDecimal4{}),
		reflect.TypeOf(decimal.NullDecimal5{}),
		reflect.TypeOf(decimal.NullDecimal6{}),
		reflect.TypeOf(decimal.NullDecimal7{}),
		reflect.TypeOf(decimal.NullDecimal8{}),
	}
	for _, t := range types {
		builder.RegisterCodec(t, nullDecimalEncoderDecoderN{})
	}

	builder.RegisterCodec(
		reflect.TypeOf(extvals.NullInt32{}), newNullableCodec(registryNotNull, reflect.TypeOf(int32(0))))
	builder.RegisterCodec(
		reflect.TypeOf(extvals.NullInt64{}), newNullableCodec(registryNotNull, reflect.TypeOf(int64(0))))
	builder.RegisterCodec(
		reflect.TypeOf(extvals.NullBool{}), newNullableCodec(registryNotNull, reflect.TypeOf(true)))
	builder.RegisterCodec(
		reflect.TypeOf(extvals.NullTime{}), newNullableCodec(registryNotNull, reflect.TypeOf(time.Time{})))

	return builder.Build()
}
