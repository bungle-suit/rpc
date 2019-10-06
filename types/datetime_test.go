package types_test

import (
	"testing"
	"time"
)

func TestDateTime(t *testing.T) {
	now := time.Now()
	nowSnapToSecond := time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second(), 0, time.Local)
	assertMarshalRoundTrip(t, "datetime", nowSnapToSecond)

	assertMarshal(t, "datetime", time.Unix(0, 0), "0")
	assertUnmarshal(t, "datetime", "0", time.Unix(0, 0))
}
