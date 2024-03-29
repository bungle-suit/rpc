package decimal

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// Decimal is a decimal implementation provides 18 max effective numbers.
// Decimal is immutable after creation, passed by value. Do not change its fields.
type Decimal struct {
	digits int64
	scale  uint8
}

// FromInt create Decimal from 64bit integer, set to scale to zero.
func FromInt(i int64) Decimal {
	return Decimal{i, 0}
}

// FromString create decimal from string, scale set from fragment part of number.
// Such as '3.00', scale is 2.
func FromString(s string) (Decimal, error) {
	return FromStringWithScale(s, -1)
}

// FromStringWithScale create decimal from string, with specific scale.
// Use number's actual scale if it larger than specific scale. Examples:
//
//  FromStringWithScale("3", 2) // 3.00
//  FromStringWithScale("3.33", 0) // 3.33
func FromStringWithScale(str string, scale int) (Decimal, error) {
	if err := checkScale(scale); err != nil {
		return Decimal{}, err
	}

	s := str
	actScale, dotIdx := 0, strings.IndexRune(s, '.')
	if dotIdx != -1 {
		actScale = len(s) - 1 - dotIdx
		s = s[0:dotIdx] + s[dotIdx+1:]
	}

	digits, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		e := err.(*strconv.NumError)
		switch e.Err {
		case strconv.ErrSyntax:
			return Decimal{}, fmt.Errorf("[%s] \"%s\" not a number", tag, str)
		case strconv.ErrRange:
			return Decimal{}, fmt.Errorf("[%s] \"%s\" effective number out of range", tag, str)
		default:
			return Decimal{}, err
		}
	}

	if err = checkScale(actScale); err != nil {
		return Decimal{}, err
	}

	d := Decimal{digits, uint8(actScale)}
	if scale != -1 && scale != actScale {
		return d.Round(scale), nil
	}
	return d, nil
}

// FromFloat convert float to decimal
func FromFloat(v float64, scale uint8) Decimal {
	return Decimal{
		digits: int64(v * math.Pow10(int(scale))),
		scale:  scale,
	}
}

// String implement fmt.Stringer interface, return decimal value in string format,
// appended 0 to scales. Such as 3.00, use ShortString() to get '3'.
func (d Decimal) String() string {
	if d.scale == 0 {
		return strconv.FormatInt(d.digits, 10)
	}

	var s string
	isNegative := d.digits < 0
	if isNegative {
		s = strconv.FormatInt(-d.digits, 10)
	} else {
		s = strconv.FormatInt(d.digits, 10)
	}

	dotIdx := len(s) - int(d.scale)
	switch {
	case dotIdx == 0:
		s = "0." + s
	case dotIdx < 0:
		s = "0." + strings.Repeat("0", -dotIdx) + s
	default:
		s = s[0:dotIdx] + "." + s[dotIdx:]
	}

	if isNegative {
		s = "-" + s
	}
	return s
}

// GoString implement fmt.GoStringer interface. Adding 'm' suffix to result of String().
func (d Decimal) GoString() string {
	return d.String() + "m"
}

// ShortString convert current value to string, removing ending 0s.
// Such as 3.00, returns 3.
func (d Decimal) ShortString() string {
	if d.IsZero() {
		return "0"
	}

	r := d.String()
	if d.scale == 0 {
		return r
	}

	if strings.ContainsRune(r, '.') {
		r = strings.TrimRight(r, "0")
		if r[len(r)-1] == '.' {
			return r[:len(r)-1]
		}
	}
	return r
}

// Scale return scale of this decimal value.
func (d Decimal) Scale() uint8 {
	return d.scale
}

// Int64 convert current value to int64, round tenth fragment.
func (d Decimal) Int64() int64 {
	if d.scale == 0 {
		return d.digits
	}

	r := d.digits
	r /= powerOf10(int(d.scale - 1))
	return roundLastDecimalBit(r)
}

// Float64 convert current value to float.
func (d Decimal) Float64() float64 {
	if d.scale == 0 {
		return float64(d.digits)
	}
	return float64(d.digits) / math.Pow10(int(d.scale))
}

// Round decimal to specific scale.
func (d Decimal) Round(scale int) Decimal {
	if err := checkScale(scale); err != nil {
		panic(err.Error())
	}

	diff := scale - int(d.scale)
	digits := d.digits
	switch {
	case diff == 0:
		return d
	case diff > 0:
		digits *= powerOf10(diff)
	default:
		diff = -diff
		digits /= powerOf10(diff - 1)
		digits = roundLastDecimalBit(digits)
	}

	return Decimal{digits, uint8(scale)}
}

// Sign returns 1 if current value greater than 0, -1 if less than 0, 0 if equals to 0.
func (d Decimal) Sign() int {
	switch {
	case d.digits > 0:
		return 1
	case d.digits < 0:
		return -1
	default:
		return 0
	}
}

// Neg returns negative value
func (d Decimal) Neg() Decimal {
	return Decimal{-d.digits, d.scale}
}

// Add this value with other value, use two values' highest scale as result scale, such as
// 3.45 + 1 = 4.45.
func (d Decimal) Add(other Decimal) Decimal {
	va, vb, scale := d.digits, other.digits, d.scale
	diff := int(d.scale) - int(other.scale)
	switch {
	case diff > 0:
		vb *= powerOf10(diff)
	case diff < 0:
		scale = other.scale
		va *= powerOf10(-diff)
	}
	return Decimal{va + vb, scale}
}

// AddToScale this value with other value round to specific scale.
func (d Decimal) AddToScale(other Decimal, scale int) Decimal {
	return d.Add(other).Round(scale)
}

// Sub subtract the other value.
func (d Decimal) Sub(other Decimal) Decimal {
	return d.Add(other.Neg())
}

// SubToScale subtract the other value to specific scale.
func (d Decimal) SubToScale(other Decimal, scale int) Decimal {
	return d.AddToScale(other.Neg(), scale)
}

// Mul multiply the other value.
func (d Decimal) Mul(other Decimal) Decimal {
	return d.MulToScale(other, max(d.scale, other.scale))
}

// MulToScale multiply the other value and round to specific scale
func (d Decimal) MulToScale(other Decimal, scale int) Decimal {
	digits := d.digits * other.digits
	scaleDiff := int(d.scale) + int(other.scale) - scale
	switch {
	case scaleDiff > 0:
		digits /= powerOf10(scaleDiff - 1)
		digits = roundLastDecimalBit(digits)
	case scaleDiff < 0:
		digits *= powerOf10(-scaleDiff)
	}
	return Decimal{digits, uint8(scale)}
}

// Div the other value, scale use max scale of current and other decimal.
func (d Decimal) Div(other Decimal) Decimal {
	return d.DivToScale(other, max(d.scale, other.scale))
}

// DivToScale the other value and round result to specific scale.
func (d Decimal) DivToScale(other Decimal, scale int) Decimal {
	scaleDiff := scale - (int(d.scale) - int(other.scale))
	digits := d.digits
	if scaleDiff > 0 {
		digits *= powerOf10(scaleDiff + 1)
		digits /= other.digits
		digits = roundLastDecimalBit(digits)
	} else {
		digits /= other.digits
		if scaleDiff < 0 {
			scaleDiff = -scaleDiff
			digits /= powerOf10(scaleDiff - 1)
			digits = roundLastDecimalBit(digits)
		}
	}
	return Decimal{digits, uint8(scale)}
}

// Cmp the other value return -1 if < other, 1 if > other, 0 if equal.
// Cmp ignore scale, so 0.00 equals to 0
func (d Decimal) Cmp(other Decimal) int {
	sub := d.Sub(other)
	switch {
	case sub.digits > 0:
		return 1
	case sub.digits < 0:
		return -1
	default:
		return 0
	}
}

// LT returns true if current value less than other
func (d Decimal) LT(other Decimal) bool {
	return d.Cmp(other) < 0
}

// GT returns true if current value greater than other.
func (d Decimal) GT(other Decimal) bool {
	return d.Cmp(other) > 0
}

// LTE returns true if current value less or equal to other.
func (d Decimal) LTE(other Decimal) bool {
	return d.Cmp(other) <= 0
}

// GTE returns true if current value greater or equal to other.
func (d Decimal) GTE(other Decimal) bool {
	return d.Cmp(other) >= 0
}

// EQ returns true if current value equals to other.
func (d Decimal) EQ(other Decimal) bool {
	return d.Cmp(other) == 0
}

// NE returns true if current value not equals to other.
func (d Decimal) NE(other Decimal) bool {
	return d.Cmp(other) != 0
}

// IsZero returns true if the value is 0, no matter what scale is.
func (d Decimal) IsZero() bool {
	return d.digits == 0
}

// MarshalJSON implement json.Marshaler interface
func (d Decimal) MarshalJSON() ([]byte, error) {
	return []byte(d.String()), nil
}

// UnmarshalJSON implement json.Unmarshaler interface
func (d *Decimal) UnmarshalJSON(buf []byte) error {
	var err error
	*d, err = FromString(string(buf))
	return err
}

// ToDecimal128 convert to IEEE 754 decimal128
func (d Decimal) ToDecimal128() (low, high uint64) {
	sign := d.Sign()
	if sign < 0 {
		low = uint64(-d.digits)
	} else {
		low = uint64(d.digits)
	}

	high = 6176 - uint64(d.scale)
	high <<= 49
	if sign < 0 {
		high |= 0x8000000000000000
	}
	return
}

const (
	maxVal int64 = 1<<63 - 1
)

// FromDecimal128 convert IEEE 754 decimal128 to Decimal. Decimal128 has greater range than Decimal,
// FromDecimal128 expect the argument must in range of Decimal.
func FromDecimal128(low, high uint64) Decimal {
	if high&0x3fffffffffff != 0 {
		panic("FromDecimal128 value too big 2")
	}

	high >>= 46
	neg := (high & 0x20000) != 0
	scale := high & 0x1ffff
	is11 := (high & 0x18000) == 0x18000
	if !is11 {
		scale >>= 3
	} else {
		panic("FromDecimal128 not support scale start with 11")
	}

	scale = -(scale - 6176)
	if scale > 8 {
		panic("FromDecimal128 scale out of range")
	}
	if low > uint64(maxVal) {
		panic("FromDecimal128 value too big")
	}

	digits := int64(low)
	if neg {
		digits = -digits
	}
	return Decimal{
		digits: digits,
		scale:  uint8(scale),
	}
}

// Zero returns zero decimal value with specific scale.
func Zero(scale int) Decimal {
	if err := checkScale(scale); err != nil {
		panic(err.Error())
	}

	return Decimal{0, uint8(scale)}
}

func max(a, b uint8) int {
	if a > b {
		return int(a)
	}
	return int(b)
}

var tenth = [18]int64{1, 10, 100, 1000, 10000, 100000, 1000000, 10000000, 100000000, 1000000000,
	10000000000, 100000000000, 1000000000000, 10000000000000, 100000000000000, 1000000000000000,
	10000000000000000, 100000000000000000}

func powerOf10(n int) int64 {
	return tenth[n]
}

// checkScale checks scale, return non-nil error if out of range
func checkScale(scale int) error {
	if scale >= 10 || scale < -1 {
		return fmt.Errorf("[%s] scale %d out of range", tag, scale)
	}
	return nil
}

// roundLastDecimalBit round last decimal bit and divide by 10.
func roundLastDecimalBit(v int64) int64 {
	mod, v := v%10, v/10
	switch {
	case mod >= 5:
		v++
	case mod <= -5:
		v--
	}
	return v
}
