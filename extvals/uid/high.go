package uid

// High is the high 16bit of ID, each corresponding to a entity/table
// use int16 instead of uint16, for easier to cooperate with other systems.
// Never use negative part, zero is reserved, accept range is [1..32768]
type High int16

// Next return next id for this high
func (h High) Next() UID {
	return nextID(h)
}
