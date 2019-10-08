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

	back, err := types.Unmarshal(p, "void", []byte("null"))
	assert.NoError(t, err)
	assert.Nil(t, back)
}
