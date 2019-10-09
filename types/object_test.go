package types_test

import (
	"testing"

	"github.com/bungle-suit/rpc/extvals"
	"github.com/bungle-suit/rpc/types"
)

func TestObject(t *testing.T) {
	p := types.NewParser()
	p.DefinePrimitiveTypes()

	assertMarshal(t, "object", extvals.Object{T: "int", V: int32(33)},
		`{"t":"int","v":33}`)
	assertUnmarshal(t, "object",
		`{"t":"int","v":33}`, extvals.Object{T: "int", V: int32(33)})
}
