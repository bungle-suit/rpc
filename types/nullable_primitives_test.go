package types_test

import (
	"testing"
	"time"

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

func TestNullLong(t *testing.T) {
	assertMarshalRoundTrip(t, "long?",
		extvals.NullInt64{},
		extvals.NullInt64{V: 33, Valid: true},
		extvals.NullInt64{V: 9000000000000001, Valid: true},
	)
}

func TestNullFloat(t *testing.T) {
	assertMarshalRoundTrip(t, "double?",
		extvals.NullFloat64{},
		extvals.NullFloat64{V: 34.34, Valid: true},
	)
}

func TestNullDateTime(t *testing.T) {
	assertMarshalRoundTrip(t, "datetime?",
		extvals.NullTime{},
		extvals.NullTime{V: time.Date(2019, 10, 9, 8, 54, 34, 0, time.Local), Valid: true},
	)
}
