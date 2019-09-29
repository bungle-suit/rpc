package types_test

import (
	"testing"

	"github.com/bungle-suit/rpc/types"
	"github.com/stretchr/testify/assert"
)

func TestInt(t *testing.T) {
	p := types.NewParser()
	p.DefinePrimitiveTypes()

	for _, v := range []int32{0, 33, -3124314} {
		buf, err := types.Marshal(p, "int", v)
		assert.NoError(t, err)

		var back int32
		assert.NoError(t, types.Unmarshal(p, "int", buf, &back))
		assert.Equal(t, v, back)
	}

	// should return error when parse number like string.
	var v int32
	assert.Error(t, types.Unmarshal(p, "int", []byte(`"1234"`), &v))
}
