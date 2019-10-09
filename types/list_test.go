package types_test

import "testing"

func _TestNullList(t *testing.T) {
	assertMarshal(t, "[int]", []int32{}, "[]")
	assertMarshal(t, "[int]", []int32{3, 4}, "[3,4]")

	assertUnmarshal(t, "[int]", "[]", []int32{})
}
