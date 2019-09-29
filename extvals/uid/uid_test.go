package uid_test

import (
	"testing"

	"github.com/bungle-suit/rpc/extvals/uid"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	tests := []struct {
		h   uid.High
		l   int64
		exp int64
	}{
		{0, 0, 0},
		{0, 1, 1},
		{32767, 0, 0x7fff000000000000},
		{32767, 0x7fffffffffff, 0x7fff7fffffffffff},
	}

	for _, rec := range tests {
		id := uid.New(rec.h, rec.l)
		assert.Equal(t, rec.exp, int64(id))
		assert.Equal(t, rec.h, id.High())
		assert.Equal(t, rec.l, id.Low())
	}
}
