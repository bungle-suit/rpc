package hi

import (
	"fmt"
	"log"
	"sync"
)

var (
	sv   Service
	lock sync.Mutex

	ids map[High]int64
)

// Service for high id.
type Service interface {
	// MaxID returns current maxID for specific high
	//
	// Return empty string if no max id, such as empty database.
	MaxID(high High) (maxID string, err error)
}

// Setup service.
func SetService(svc Service) {
	lock.Lock()
	sv = svc
	ids = make(map[High]int64)
	lock.Unlock()
}

func nextID(high High) string {
	lock.Lock()
	defer lock.Unlock()

	if ids[high] == 0 {
		id, err := sv.MaxID(high)
		if err != nil {
			log.Panicf("[%s] failed get max id of %s", tag, high)
			return ""
		}

		low := int64(0)
		if id != "" {
			var hi High
			var err error
			hi, low, err = DecodeUID(id)
			if err != nil {
				log.Panicf("[%s] service returns invalid format maxID %s", tag, id)
			}
			if hi != high {
				log.Panicf("[%s] service returns mismatched high maxID exp: %s act: %s", tag, high, hi)
				return ""
			}
		}

		ids[high] = low
	}
	ids[high]++
	return fmt.Sprintf("%s:%012d", high, ids[high])
}
