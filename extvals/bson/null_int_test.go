package bson_test

import (
	"testing"

	"github.com/bungle-suit/rpc/extvals"
	mybson "github.com/bungle-suit/rpc/extvals/bson"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
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
			assert.Equal(t, item.Int32, doc["a"])
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
			assert.Equal(t, item.Int64, doc["a"])
		}

		back := rec
		back.A = extvals.NullInt64{}
		assert.NoError(t, bson.UnmarshalWithRegistry(mybson.Registry, buf, &back))
		assert.Equal(t, back, rec)
	}
}
