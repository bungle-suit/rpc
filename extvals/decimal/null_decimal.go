package decimal

import "bytes"

// NullDecimal nullable decimal value
type NullDecimal struct {
	V     Decimal
	Valid bool // Valid is true if current value is not NULL
}

func (d NullDecimal) String() string {
	if d.Valid {
		return d.V.String()
	}
	return ""
}

// MarshalJSON implement json.Marshaler interface
func (d NullDecimal) MarshalJSON() ([]byte, error) {
	if !d.Valid {
		return []byte("null"), nil
	}

	return d.V.MarshalJSON()
}

// UnmarshalJSON implement json.Unmarshaler interface
func (d *NullDecimal) UnmarshalJSON(buf []byte) error {
	if bytes.Equal(buf, []byte("null")) {
		*d = NullDecimal{}
		return nil
	}

	var v Decimal
	if err := v.UnmarshalJSON(buf); err != nil {
		return err
	}

	*d = NullDecimal{v, true}
	return nil
}

func (i NullDecimal) IsNull() bool {
	return !i.Valid
}

func (i NullDecimal) Val() interface{} {
	return i.V
}
