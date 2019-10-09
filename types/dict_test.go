package types_test

import "testing"

func TestDict(t *testing.T) {
	assertMarshal(t, "{str:int}", map[string]int32{}, "{}")
	assertMarshal(t, "{str:int}", map[string]int32{"a": 33}, `{"a":33}`)

	assertUnmarshal(t, "{str:int}", "{}", map[string]int32{})
	assertUnmarshal(t, "{str:int}", `{"a":33}`, map[string]int32{"a": 33})
}
