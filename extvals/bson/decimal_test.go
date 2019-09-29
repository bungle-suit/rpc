package bson_test

import (
	"testing"

	mybson "github.com/bungle-suit/rpc/extvals/bson"
	"github.com/bungle-suit/rpc/extvals/decimal"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestDecimal(t *testing.T) {
	registry := mybson.Registry()
	for _, s := range []string{"0", "-12345.7890", "1234567890123456"} {
		v := struct{ ID decimal.Decimal }{}
		back := v

		v.ID = parseDecimal(s)
		buf, err := bson.MarshalWithRegistry(registry, v)
		assert.NoError(t, err)

		assert.NoError(t, bson.UnmarshalWithRegistry(registry, buf, &back))
		assert.Equal(t, v.ID, back.ID)

		var doc bson.M
		assert.NoError(t, bson.Unmarshal(buf, &doc))
		assert.Equal(t, bson.M{"id": parseDecimal128(s)}, doc)
	}
}

func parseDecimal128(s string) primitive.Decimal128 {
	v, err := primitive.ParseDecimal128(s)
	if err != nil {
		panic(err)
	}
	return v
}

func parseDecimal(s string) decimal.Decimal {
	v, err := decimal.FromString(s)
	if err != nil {
		panic(err)
	}
	return v
}

func TestNullDecimal(t *testing.T) {
	registry := mybson.Registry()
	for _, s := range []string{"", "0", "-12345.7890", "1234567890123456"} {
		v := struct{ ID decimal.NullDecimal }{}
		back := v

		if s != "" {
			v.ID = decimal.NullDecimal{V: parseDecimal(s), Valid: true}
		}
		buf, err := bson.MarshalWithRegistry(registry, &v)
		assert.NoError(t, err)

		assert.NoError(t, bson.UnmarshalWithRegistry(registry, buf, &back))
		assert.Equal(t, v.ID, back.ID)

		var doc bson.M
		assert.NoError(t, bson.Unmarshal(buf, &doc))
		if s == "" {
			assert.Equal(t, bson.M{"id": nil}, doc)
		} else {
			assert.Equal(t, bson.M{"id": parseDecimal128(s)}, doc)
		}
	}
}

func TestDecimalN(t *testing.T) {
	registry := mybson.Registry()
	for _, s := range []string{"0", "-12345.7890", "1234567890123456"} {
		v := struct{ ID decimal.Decimal2 }{}
		back := v

		v.ID = decimal.Decimal2(parseDecimal(s))
		buf, err := bson.MarshalWithRegistry(registry, v)
		assert.NoError(t, err)

		assert.NoError(t, bson.UnmarshalWithRegistry(registry, buf, &back))
		assert.Equal(t, decimal.Decimal2(decimal.Decimal(v.ID).Round(2)), back.ID)
	}
}

func TestNullDecimalN(t *testing.T) {
	registry := mybson.Registry()
	for _, s := range []string{"", "0", "-12345.7890", "1234567890123456"} {
		v := struct{ ID decimal.NullDecimal2 }{}
		back := v

		if s != "" {
			v.ID = decimal.NullDecimal2(decimal.NullDecimal{V: parseDecimal(s), Valid: true})
		}
		buf, err := bson.MarshalWithRegistry(registry, v)
		assert.NoError(t, err)

		assert.NoError(t, bson.UnmarshalWithRegistry(registry, buf, &back))
		if s == "" {
			assert.Equal(t, decimal.NullDecimal2{}, back.ID)
		} else {
			assert.Equal(t,
				decimal.NullDecimal2(decimal.NullDecimal{V: v.ID.V.Round(2), Valid: true}),
				back.ID)
		}
	}
}
