package types_test

import (
	"reflect"
	"strings"
	"testing"

	"github.com/bungle-suit/rpc/extvals/decimal"
	"github.com/bungle-suit/rpc/types"
	"github.com/stretchr/testify/assert"
)

func assertMarshal(t *testing.T, ts string, v interface{}, exp string) {
	p := types.NewParser()
	p.DefinePrimitiveTypes()

	buf, err := types.Marshal(p, ts, v)
	assert.NoError(t, err)
	assert.Equal(t, exp, strings.TrimSpace(string(buf)))
}

func assertUnmarshal(t *testing.T, ts, json string, exp interface{}) {
	p := types.NewParser()
	p.DefinePrimitiveTypes()

	back := reflect.New(reflect.TypeOf(exp))

	assert.NoError(t, types.Unmarshal(p, ts, []byte(json), back.Interface()))
	assert.Equal(t, exp, back.Elem().Interface())
}

func parseDecimal2(s string) decimal.Decimal2 {
	d, err := decimal.FromStringWithScale(s, 2)
	if err != nil {
		panic(err)
	}

	return decimal.Decimal2(d)
}
