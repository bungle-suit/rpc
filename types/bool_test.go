package types_test

import (
	"testing"

	"github.com/bungle-suit/rpc/types"
	"github.com/stretchr/testify/assert"
)

func TestBool(t *testing.T) {
	p := types.NewParser()
	p.DefinePrimitiveTypes()

	for _, v := range []bool{true, false} {
		buf, err := types.Marshal(p, "bool", v)
		assert.NoError(t, err)

		var back bool
		assert.NoError(t, types.Unmarshal(p, "bool", buf, &back))
		assert.Equal(t, v, back)
	}

	assertMarshal(t, "bool", true, "1")
	assertMarshal(t, "bool", false, "0")
}
