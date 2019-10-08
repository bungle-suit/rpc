package types_test

import (
	"testing"

	"github.com/bungle-suit/rpc/types"
	"github.com/stretchr/testify/assert"
)

func TestInt(t *testing.T) {
	assertMarshalRoundTrip(t, "int", int32(0), int32(33), int32(-3124314))

	// should return error when parse number like string.
	p := types.NewParser()
	p.DefinePrimitiveTypes()
	_, err := types.Unmarshal(p, "int", []byte(`"1234"`))
	assert.Error(t, err)
}

func TestLong(t *testing.T) {
	assertMarshalRoundTrip(t, "long", int64(0), int64(33), int64(-3124314), int64(9000000000000000), int64(-9000000000000000))

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

	assertMarshalRoundTrip(t, "double", 0.0, 33.0, -312.4314, 9000000000000000.0, -9000000000000000.0)
}

func TestDecimal(t *testing.T) {
	p := types.NewParser()
	p.DefinePrimitiveTypes()

	var vals []interface{}
	for _, s := range []string{"0", "33", "-312.43", "9000000000000000", "-9000000000000000"} {
		vals = append(vals, parseDecimal2(s))
	}
	assertMarshalRoundTrip(t, "decimal(2)", vals...)
}
