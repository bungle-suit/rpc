package extvals

import "time"

// NullInt32 represents a int32 value that may be null. Similar to
// database/sql.NullInt64.
type NullInt32 struct {
	Int32 int32
	Valid bool // Valid is true if Int32 is not NULL
}

// NullInt64 represents a int32 value that may be null. Similar to
// database/sql.NullInt64.
type NullInt64 struct {
	Int64 int64
	Valid bool
}

// NullBool represents a bool value that may be null. Similar to
// database/sql.NullInt64.
type NullBool struct {
	Bool  bool
	Valid bool
}

// NullTime represents a time value that may be null. Similar to
// database/sql.NullInt64.
type NullTime struct {
	Time  time.Time
	Valid bool
}

func (i NullInt32) IsNull() bool {
	return i.Valid
}

func (i NullInt32) Val() interface{} {
	return i.Int32
}

func (i NullInt64) IsNull() bool {
	return i.Valid
}

func (i NullInt64) Val() interface{} {
	return i.Int64
}

func (i NullBool) IsNull() bool {
	return i.Valid
}

func (i NullBool) Val() interface{} {
	return i.Bool
}

func (i NullTime) IsNull() bool {
	return i.Valid
}

func (i NullTime) Val() interface{} {
	return i.Time
}
