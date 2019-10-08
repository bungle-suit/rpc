package types_test

import (
	"testing"

	"github.com/bungle-suit/rpc/extvals/decimal"
	"github.com/bungle-suit/rpc/types"
	"github.com/stretchr/testify/assert"
)

func assertMarshalRoundTrip(t *testing.T, ts string, vals ...interface{}) {
	p := types.NewParser()
	p.DefinePrimitiveTypes()

	for _, v := range vals {
		buf, err := types.Marshal(p, ts, v)
		assert.NoError(t, err)

		back, err := types.Unmarshal(p, ts, buf)
		assert.NoError(t, err)
		assert.Equal(t, v, back)
	}
}

func assertMarshal(t *testing.T, ts string, v interface{}, exp string) {
	p := types.NewParser()
	p.DefinePrimitiveTypes()

	buf, err := types.Marshal(p, ts, v)
	assert.NoError(t, err)
	assert.JSONEq(t, exp, string(buf))
}

func assertUnmarshal(t *testing.T, ts, json string, exp interface{}) {
	p := types.NewParser()
	p.DefinePrimitiveTypes()

	back, err := types.Unmarshal(p, ts, []byte(json))
	assert.NoError(t, err)
	assert.Equal(t, exp, back)
}

func parseDecimal2(s string) decimal.Decimal2 {
	d, err := decimal.FromStringWithScale(s, 2)
	if err != nil {
		panic(err)
	}

	return decimal.Decimal2(d)
}
