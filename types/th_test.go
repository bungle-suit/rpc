package types_test

import (
	"reflect"
	"strings"
	"testing"

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
