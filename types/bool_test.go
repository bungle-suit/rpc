package types_test

import (
	"testing"
)

func TestBool(t *testing.T) {
	assertMarshalRoundTrip(t, "bool", true, false)

	assertMarshal(t, "bool", true, "1")
	assertMarshal(t, "bool", false, "0")
}
