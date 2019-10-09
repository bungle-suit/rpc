package types_test

import (
	"testing"

	"github.com/bungle-suit/rpc/extvals"
)

func TestNullBool(t *testing.T) {
	assertMarshal(t, "bool?", extvals.NullBool{}, "null")
	assertMarshal(t, "bool?", extvals.NullBool{V: true, Valid: true}, "1")
	assertMarshal(t, "bool?", extvals.NullBool{V: false, Valid: true}, "0")

	assertUnmarshal(t, "bool?", "null", extvals.NullBool{})
	assertUnmarshal(t, "bool?", "1", extvals.NullBool{V: true, Valid: true})
	assertUnmarshal(t, "bool?", "0", extvals.NullBool{V: false, Valid: true})
}

func TestNullInt(t *testing.T) {
	assertMarshalRoundTrip(t, "int?",
		extvals.NullInt32{},
		extvals.NullInt32{V: 33, Valid: true},
	)
}
