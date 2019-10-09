package types_test

import (
	"testing"

	"github.com/bungle-suit/rpc/extvals"
)

func TestNullableBool(t *testing.T) {
	assertMarshal(t, "bool?", extvals.NullBool{}, "null")
	assertMarshal(t, "bool?", extvals.NullBool{V: true, Valid: true}, "1")
	assertMarshal(t, "bool?", extvals.NullBool{V: false, Valid: true}, "0")

	assertUnmarshal(t, "bool?", "null", extvals.NullBool{})
	assertUnmarshal(t, "bool?", "1", extvals.NullBool{V: true, Valid: true})
	assertUnmarshal(t, "bool?", "0", extvals.NullBool{V: false, Valid: true})
}
