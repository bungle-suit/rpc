package types_test

import (
	"testing"

	"github.com/bungle-suit/rpc/extvals/decimal"
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

func TestFloat(t *testing.T) {
	p := types.NewParser()
	p.DefinePrimitiveTypes()

	for _, v := range []float64{0, 33, -312.4314, 9000000000000000, -9000000000000000} {
		buf, err := types.Marshal(p, "double", v)
		assert.NoError(t, err)

		var back float64
		assert.NoError(t, types.Unmarshal(p, "double", buf, &back))
		assert.Equal(t, v, back)
	}
}

func TestDecimal(t *testing.T) {
	p := types.NewParser()
	p.DefinePrimitiveTypes()

	for _, s := range []string{"0", "33", "-312.43", "9000000000000000", "-9000000000000000"} {
		v := parseDecimal2(s)
		buf, err := types.Marshal(p, "decimal(2)", v)
		assert.NoError(t, err)

		var back decimal.Decimal2
		assert.NoError(t, types.Unmarshal(p, "decimal(2)", buf, &back))
		assert.Equal(t, v, back)
	}
}
