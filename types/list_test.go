package types_test

import "testing"

func TestList(t *testing.T) {
	assertMarshal(t, "[int]", []int32{}, "[]")
	assertMarshal(t, "[int]", []int32{3, 4}, "[3,4]")

	assertUnmarshal(t, "[int]", "[]", []int32{})
	assertUnmarshal(t, "[int]", "[3,4]", []int32{3, 4})

	assertMarshal(t, "[[int]]", [][]int32{}, "[]")
	assertMarshal(t, "[[int]]", [][]int32{[]int32{3}, []int32{}}, "[[3],[]]")

	assertUnmarshal(t, "[[int]]", "[]", [][]int32{})
	assertUnmarshal(t, "[[int]]", "[[3],[]]", [][]int32{[]int32{3}, []int32{}})
}
