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

func TestLong(t *testing.T) {
	p := types.NewParser()
	p.DefinePrimitiveTypes()

	for _, v := range []int64{0, 33, -3124314, 9000000000000000, -9000000000000000} {
		buf, err := types.Marshal(p, "long", v)
		assert.NoError(t, err)

		var back int64
		assert.NoError(t, types.Unmarshal(p, "long", buf, &back))
		assert.Equal(t, v, back)
	}

	// marshal as string if abs() > 9000000000000000
	assertMarshal(t, "long", int64(9000000000000000), "9000000000000000")
	assertMarshal(t, "long", int64(-9000000000000000), "-9000000000000000")
	assertMarshal(t, "long", int64(9000000000000001), `"9000000000000001"`)
	assertMarshal(t, "long", int64(-9000000000000001), `"-9000000000000001"`)

	assertUnmarshal(t, "long", `"9000000000000001"`, int64(9000000000000001))
	assertUnmarshal(t, "long", "-9000000000000000", int64(-9000000000000000))
	assertUnmarshal(t, "long", `"9000000000000001"`, int64(9000000000000001))
	assertUnmarshal(t, "long", `"-9000000000000001"`, int64(-9000000000000001))
}
