package types_test

import "testing"

func TestString(t *testing.T) {
	assertMarshalRoundTrip(t, "str", "", `abc"foo"`)
}
