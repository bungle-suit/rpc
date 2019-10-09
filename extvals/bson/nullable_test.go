package bson_test

import (
	"testing"
	"time"

	"github.com/bungle-suit/rpc/extvals"
	mybson "github.com/bungle-suit/rpc/extvals/bson"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestNullInt32(t *testing.T) { // nolint: dupl
	registry := mybson.Registry()
	tests := []extvals.NullInt32{
		{},
		{V: 32, Valid: true},
		{V: -32, Valid: true},
	}

	for _, item := range tests {
		rec := struct{ A extvals.NullInt32 }{item}
		buf, err := bson.MarshalWithRegistry(registry, rec)
		assert.NoError(t, err)

		var doc bson.M
		assert.NoError(t, bson.Unmarshal(buf, &doc))
		if !item.Valid {
			assert.Nil(t, doc["a"])
		} else {
			assert.Equal(t, item.V, doc["a"])
		}

		back := rec
		back.A = extvals.NullInt32{}
		assert.NoError(t, bson.UnmarshalWithRegistry(registry, buf, &back))
		assert.Equal(t, back, rec)
	}
}

func TestNullInt64(t *testing.T) { // nolint: dupl
	registry := mybson.Registry()
	tests := []extvals.NullInt64{
		{},
		{V: 32, Valid: true},
		{V: -32, Valid: true},
	}

	for _, item := range tests {
		rec := struct{ A extvals.NullInt64 }{item}
		buf, err := bson.MarshalWithRegistry(registry, rec)
		assert.NoError(t, err)

		var doc bson.M
		assert.NoError(t, bson.Unmarshal(buf, &doc))
		if !item.Valid {
			assert.Nil(t, doc["a"])
		} else {
			assert.Equal(t, item.V, doc["a"])
		}

		back := rec
		back.A = extvals.NullInt64{}
		assert.NoError(t, bson.UnmarshalWithRegistry(registry, buf, &back))
		assert.Equal(t, back, rec)
	}
}

func TestNullBool(t *testing.T) { // nolint: dupl
	registry := mybson.Registry()
	tests := []extvals.NullBool{
		{},
		{V: true, Valid: true},
		{V: false, Valid: true},
	}

	for _, item := range tests {
		rec := struct{ A extvals.NullBool }{item}
		buf, err := bson.MarshalWithRegistry(registry, rec)
		assert.NoError(t, err)

		var doc bson.M
		assert.NoError(t, bson.Unmarshal(buf, &doc))
		if !item.Valid {
			assert.Nil(t, doc["a"])
		} else {
			assert.Equal(t, item.V, doc["a"])
		}

		back := rec
		back.A = extvals.NullBool{}
		assert.NoError(t, bson.UnmarshalWithRegistry(registry, buf, &back))
		assert.Equal(t, back, rec)
	}
}

func TestNullTime(t *testing.T) {
	registry := mybson.Registry()
	tests := []extvals.NullTime{
		{},
		{V: time.Date(2019, 9, 28, 4, 5, 6, 0, time.UTC), Valid: true},
	}

	for _, item := range tests {
		rec := struct{ A extvals.NullTime }{item}
		buf, err := bson.MarshalWithRegistry(registry, rec)
		assert.NoError(t, err)

		var doc bson.M
		assert.NoError(t, bson.Unmarshal(buf, &doc))
		if !item.Valid {
			assert.Nil(t, doc["a"])
		} else {
			assert.Equal(t, item.V, doc["a"].(primitive.DateTime).Time().UTC())
		}

		back := rec
		back.A = extvals.NullTime{}
		assert.NoError(t, bson.UnmarshalWithRegistry(registry, buf, &back))
		assert.Equal(t, back, rec)
	}
}

func TestNullFloat64(t *testing.T) {
	registry := mybson.Registry()
	tests := []extvals.NullFloat64{
		{},
		{V: 33.34, Valid: true},
	}

	for _, item := range tests {
		rec := struct{ A extvals.NullFloat64 }{item}
		buf, err := bson.MarshalWithRegistry(registry, rec)
		assert.NoError(t, err)

		var doc bson.M
		assert.NoError(t, bson.Unmarshal(buf, &doc))
		if !item.Valid {
			assert.Nil(t, doc["a"])
		} else {
			assert.Equal(t, item.V, doc["a"].(float64))
		}

		back := rec
		back.A = extvals.NullFloat64{}
		assert.NoError(t, bson.UnmarshalWithRegistry(registry, buf, &back))
		assert.Equal(t, back, rec)
	}
}
