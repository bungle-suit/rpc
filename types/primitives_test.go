package types_test

import (
	"testing"
	"time"

	"github.com/bungle-suit/rpc/types"
	"github.com/stretchr/testify/assert"
)

func TestBool(t *testing.T) {
	assertMarshalRoundTrip(t, "bool", true, false)

	assertMarshal(t, "bool", true, "1")
	assertMarshal(t, "bool", false, "0")
}

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

func TestDateTime(t *testing.T) {
	now := time.Now()
	nowSnapToSecond := time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), 0, time.Local)
	assertMarshalRoundTrip(t, "datetime", nowSnapToSecond)

	assertMarshal(t, "datetime", time.Unix(0, 0), "0")
	assertUnmarshal(t, "datetime", "0", time.Unix(0, 0))
}

func TestString(t *testing.T) {
	assertMarshalRoundTrip(t, "str", "", `abc"foo"`)
}

func TestVoid(t *testing.T) {
	assertMarshal(t, "void", nil, "null")

	p := types.NewParser()
	p.DefinePrimitiveTypes()

	back, err := types.Unmarshal(p, "void", []byte("null"))
	assert.NoError(t, err)
	assert.Nil(t, back)
}
