package types_test

import (
	"testing"

	"github.com/bungle-suit/rpc/types"
	"github.com/stretchr/testify/assert"
)

func TestVoid(t *testing.T) {
	assertMarshal(t, "void", nil, "null")

	p := types.NewParser()
	p.DefinePrimitiveTypes()

	var back interface{}
	assert.NoError(t, types.Unmarshal(p, "void", []byte("null"), &back))
	assert.Nil(t, back)
}
