package hi_test

import (
	"testing"

	"github.com/bungle-suit/rpc/extvals/hi"
	"github.com/stretchr/testify/assert"
)

func TestDecodeUID(t *testing.T) {
	tests := []struct {
		uid string
		hi  hi.High
		low int64
	}{
		{"ABC:000000000001", "ABC", 1},
		{"CDE:100000000001", "CDE", 100000000001},
	}

	for _, rec := range tests {
		h, l, err := hi.DecodeUID(rec.uid)
		assert.NoError(t, err)
		assert.Equal(t, rec.hi, h)
		assert.Equal(t, rec.low, l)
	}
}

func TestErrDecodeUID(t *testing.T) {
	tests := []string{
		"aBC:000000000001", // high part should be three upper case letters
		"123:000000000001",
		"A_C:000000000001",
		"ABC:00000000001", // low part should be 12 positive numbers.
		"ABC:-00000000001",
		"ABC:00A000000001",
	}

	for _, s := range tests {
		_, _, err := hi.DecodeUID(s)
		assert.Error(t, err)
	}
}

type fakeService map[hi.High]string

func (sv fakeService) MaxID(high hi.High) (string, error) {
	return sv[high], nil
}

func TestNextID(t *testing.T) {
	hi.SetService(fakeService{
		"ABC": "ABC:000000000101",
		"CDE": "CDE:000000001103",
	})

	tests := []struct {
		hi  hi.High
		exp string
	}{
		{"ABC", "ABC:000000000102"},
		{"CDE", "CDE:000000001104"},
		{"ABC", "ABC:000000000103"},
		{"NEW", "NEW:000000000001"},
		{"CDE", "CDE:000000001105"},
	}

	for _, rec := range tests {
		assert.Equal(t, rec.exp, rec.hi.NextID())
	}
}
