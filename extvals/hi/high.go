package hi

import (
	"fmt"
	"regexp"
	"strconv"
)

// High is a three letters prefix.
type High string

// NextID returns next id for the high entity.
func (h High) NextID() string {
	return nextID(h)
}

var uidPattern = regexp.MustCompile(`^[A-Z]{3}:[0-9]{12}$`)

// DecodeUID decode uid to high/low parts.
func DecodeUID(uid string) (high High, low int64, err error) {
	if !uidPattern.MatchString(uid) {
		err = fmt.Errorf("[%s] Invalid UID: %s", tag, uid)
		return
	}

	high = High(uid[0:3])
	low, err = strconv.ParseInt(uid[4:], 10, 64)
	return
}
