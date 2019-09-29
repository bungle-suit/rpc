package uid_test

import (
	"testing"

	"github.com/bungle-suit/rpc/extvals/uid"
	"github.com/stretchr/testify/assert"
)

type fakeService map[uid.High]int64

func (sv fakeService) MaxID(high uid.High) (int64, error) {
	return sv[high], nil
}

func TestNewID(t *testing.T) {
	uid.SetService(fakeService{
		3: int64(uid.New(3, 100)),
		5: int64(uid.New(5, 32)),
	})

	assert.Equal(t, uid.New(3, 101), uid.High(3).Next())
	assert.Equal(t, uid.New(3, 102), uid.High(3).Next())
	assert.Equal(t, uid.New(5, 33), uid.High(5).Next())
}
