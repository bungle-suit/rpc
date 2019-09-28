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

func TestNullInt32(t *testing.T) {
	tests := []extvals.NullInt32{
		{},
		{32, true},
		{-32, true},
	}

	for _, item := range tests {
		rec := struct{ A extvals.NullInt32 }{item}
		buf, err := bson.MarshalWithRegistry(mybson.Registry, rec)
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
		assert.NoError(t, bson.UnmarshalWithRegistry(mybson.Registry, buf, &back))
		assert.Equal(t, back, rec)
	}
}

func TestNullInt64(t *testing.T) {
	tests := []extvals.NullInt64{
		{},
		{32, true},
		{-32, true},
	}

	for _, item := range tests {
		rec := struct{ A extvals.NullInt64 }{item}
		buf, err := bson.MarshalWithRegistry(mybson.Registry, rec)
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
		assert.NoError(t, bson.UnmarshalWithRegistry(mybson.Registry, buf, &back))
		assert.Equal(t, back, rec)
	}
}

func TestNullBool(t *testing.T) {
	tests := []extvals.NullBool{
		{},
		{true, true},
		{false, true},
	}

	for _, item := range tests {
		rec := struct{ A extvals.NullBool }{item}
		buf, err := bson.MarshalWithRegistry(mybson.Registry, rec)
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
		assert.NoError(t, bson.UnmarshalWithRegistry(mybson.Registry, buf, &back))
		assert.Equal(t, back, rec)
	}
}

func _TestNullTime(t *testing.T) {
	tests := []extvals.NullTime{
		{},
		{time.Date(2019, 9, 28, 4, 5, 6, 0, time.Local), true},
	}

	for _, item := range tests {
		rec := struct{ A extvals.NullTime }{item}
		buf, err := bson.MarshalWithRegistry(mybson.Registry, rec)
		assert.NoError(t, err)

		var doc bson.M
		assert.NoError(t, bson.Unmarshal(buf, &doc))
		if !item.Valid {
			assert.Nil(t, doc["a"])
		} else {
			assert.Equal(t, item.V, doc["a"].(primitive.DateTime).Time().Local())
		}

		back := rec
		back.A = extvals.NullTime{}
		assert.NoError(t, bson.UnmarshalWithRegistry(mybson.Registry, buf, &back))
		assert.Equal(t, back, rec)
	}
}
