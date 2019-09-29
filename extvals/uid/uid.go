package uid

import (
	"fmt"
	"strconv"
)

// ID is alias of int64, use Next() method to generate next unique ID.
// Use int64 instead of uint64 for better cooperate with other systems,
// such as database and javascript clients.
// Never use negative part.
type UID int64

// NewID Create a Id from low and high part.
// Use cases of this function are rare, this is low level function, no argument
// check is done for performance reason, use it with care.
func New(high High, low int64) UID {
	return UID(low + int64(high)<<48)
}

// High return high part of the id
func (l UID) High() High {
	return High(int16(l >> 48))
}

// Low return low part of the id
func (l UID) Low() int64 {
	return int64(l) & 0xffffffffffff
}

// String implement fmt.Stringer interface
func (l UID) String() string {
	return strconv.FormatInt(int64(l.Low()), 10)
}

// GoString implement fmt.GoStringer interface
func (l UID) GoString() string {
	return fmt.Sprintf("UID(%d, %d)", l.High(), l.Low())
}
