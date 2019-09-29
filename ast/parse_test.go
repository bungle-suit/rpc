package ast_test

import (
	"fmt"
	"testing"

	"github.com/bungle-suit/rpc/ast"
	"github.com/stretchr/testify/assert"
)

func TestBasicTypes(t *testing.T) {
	tests := []struct {
		ts        string
		ty        ast.NodeType
		allowNull bool
	}{
		{"void", ast.Void, false},
		{"int", ast.Int32, false},
		{"long", ast.Int64, false},
		{"bool", ast.Bool, false},
		{"str", ast.String, false},
		{"double", ast.Float, false},
		{"datetime", ast.DateTime, false},
		{"table", ast.Table, false},
		{"object", ast.Object, false},
		{"refID", ast.RefID, false},
		{"id", ast.ID, false},
		{"int?", ast.Int32, true},
		{"long?", ast.Int64, true},
		{"bool?", ast.Bool, true},
		{"double?", ast.Float, true},
		{"datetime?", ast.DateTime, true},
		{"object?", ast.Object, true},
	}

	for _, rec := range tests {
		n, err := ast.Parse(rec.ts)
		assert.NoError(t, err)
		assert.Equal(t, rec.ts, n.String())
		assert.Equal(t, rec.ty, n.Type())
		assert.Equal(t, rec.allowNull, n.Nullable())
	}
}

func TestDecimal(t *testing.T) {
	for i := 0; i < 9; i++ {
		ts := fmt.Sprintf("decimal(%d)", i)
		n, err := ast.Parse(ts)
		assert.NoError(t, err)
		dn := n.(ast.DecimalNode)
		assert.Equal(t, ts, dn.String())
		assert.False(t, dn.Nullable())
		assert.Equal(t, ast.Decimal, dn.Type())

		ts += "?"
		n, err = ast.Parse(ts)
		assert.NoError(t, err)
		dn = n.(ast.DecimalNode)
		assert.Equal(t, ts, dn.String())
		assert.True(t, dn.Nullable())
		assert.Equal(t, ast.Decimal, dn.Type())
	}
}

func TestList(t *testing.T) {
	tests := []string{"int", "[int?]", "[str]"}
	for _, itemTS := range tests {
		ts := "[" + itemTS + "]"
		n, err := ast.Parse(ts)
		assert.NoError(t, err)
		assert.Equal(t, ast.List, n.Type())
		assert.False(t, n.Nullable())
		assert.Equal(t, ts, n.String())

		itemNode := n.(ast.ItemNode)
		assert.Equal(t, itemTS, itemNode.Item.String())
	}
}

func TestDict(t *testing.T) {
	tests := []string{"int", "[int?]", "{str:[str]}"}
	for _, itemTS := range tests {
		ts := "{str:" + itemTS + "}"
		n, err := ast.Parse(ts)
		assert.NoError(t, err)
		assert.Equal(t, ast.Dict, n.Type())
		assert.False(t, n.Nullable())
		assert.Equal(t, ts, n.String())

		itemNode := n.(ast.ItemNode)
		assert.Equal(t, itemTS, itemNode.Item.String())
	}
}

func TestRpcObject(t *testing.T) {
	tests := []string{"sys.order.Order"}
	for _, ts := range tests {
		n, err := ast.Parse(ts)
		assert.NoError(t, err)
		assert.Equal(t, ast.RPCObject, n.Type())
		assert.False(t, n.Nullable())
		assert.Equal(t, ts, n.String())

		n, err = ast.Parse(ts + "?")
		assert.NoError(t, err)
		assert.Equal(t, ast.RPCObject, n.Type())
		assert.True(t, n.Nullable())
		assert.Equal(t, ts+"?", n.String())
	}

	_, err := ast.Parse("foo")
	assert.EqualError(t, err, "[rpc/ast] Wrong type string: 'foo'")
}

func TestNotAllowedNullTypes(t *testing.T) {
	tests := []string{
		"str?", "void?", "table?", "[int]?", "[str?]",
		"{str:int}?", "refID?", "id?",
	}
	for _, ts := range tests {
		_, err := ast.Parse(ts)
		assert.Errorf(t, err, "[rpc/ast] '%s' not support nullable", ts)
	}
}
