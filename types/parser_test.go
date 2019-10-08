package types_test

import (
	"reflect"
	"testing"

	"github.com/bungle-suit/json"
	"github.com/bungle-suit/rpc/types"
	"github.com/stretchr/testify/assert"
)

type fakeType string

func (fakeType) Marshal(w *json.Writer, v interface{}) error {
	panic("not implemented")
}

func (fakeType) Unmarshal(r *json.Reader, v reflect.Value) error {
	panic("not implemented")
}

func (fakeType) New() interface{} {
	panic("not implemented")
}

func TestParseKnownTypes(t *testing.T) {
	p := types.NewParser()
	p.Define("t1", fakeType("ty1"))
	p.Define("t2", fakeType("ty2"))

	ty, err := p.Parse("t1")
	assert.NoError(t, err)
	assert.Equal(t, "ty1", string(ty.(fakeType)))

	ty, err = p.Parse("t2")
	assert.NoError(t, err)
	assert.Equal(t, "ty2", string(ty.(fakeType)))

	_, err = p.Parse("unknown")
	assert.Error(t, err)
}

func TestParseList(t *testing.T) {
	p := types.NewParser()
	p.Define("int", fakeType("int"))

	t1, err := p.Parse("[int]")
	assert.NoError(t, err)
	t2, err := p.Parse("[int]")
	assert.NoError(t, err)
	assert.Equal(t, t1, t2)

	_, err = p.Parse("[[int]]")
	assert.NoError(t, err)
}

func TestParseDict(t *testing.T) {
	p := types.NewParser()
	p.Define("bool", fakeType("int"))

	_, err := p.Parse("{str:{str:bool}}")
	assert.NoError(t, err)
}

func TestNullable(t *testing.T) {
	p := types.NewParser()
	p.Define("int", fakeType("int"))
	p.Define("bool", fakeType("bool"))

	_, err := p.Parse("int?")
	assert.NoError(t, err)

	_, err = p.Parse("[bool?]")
	assert.NoError(t, err)
}
