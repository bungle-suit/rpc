package uid

import (
	"log"
	"sync"
)

// Service for uid.
type Service interface {
	// MaxID returns current maxID for specific high
	MaxID(high High) (maxID int64, err error)
}

var (
	sv   Service
	lock sync.Mutex

	ids []UID
)

// Setup service.
func SetService(svc Service) {
	lock.Lock()
	sv = svc
	ids = nil
	lock.Unlock()
}

func nextID(high High) UID {
	lock.Lock()
	defer lock.Unlock()

	if len(ids) < int(high) {
		newIds := make([]UID, high+16)
		copy(newIds, ids)
		ids = newIds
	}

	if ids[high] == 0 {
		id, err := sv.MaxID(high)
		if err != nil {
			log.Panicf("[%s] failed get max id of %d", tag, high)
			return 0
		}

		if UID(id).High() != high {
			log.Panicf("[%s] service returns wrong maxID high %d %d", tag, int16(high), id)
			return 0
		}
		ids[high] = UID(id)
	}
	ids[high]++
	return ids[high]
}
