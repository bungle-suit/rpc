package decimal_test

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"strconv"
	"testing"

	"github.com/bungle-suit/rpc/extvals/decimal"
	"github.com/stretchr/testify/assert"
)

func TestFromString(t *testing.T) {
	for _, s := range []string{
		"0", "0.00", "123456789012345678", "0.123456789", "-1.30",
	} {
		d, err := decimal.FromString(s)
		assert.NoError(t, err)
		assert.Equal(t, s, d.String())
	}
}

func TestFromStringError(t *testing.T) {
	tests := []struct {
		s     string
		scale int
		msg   string
	}{
		{"", 0, `[decimal] "" not a number`},
		{"abc", 0, `[decimal] "abc" not a number`},
		{"1.2.3", 0, `[decimal] "1.2.3" not a number`},
		{"12345678901234567890", 0, `[decimal] "12345678901234567890" effective number out of range`},
		{"0.1234567890", 1, `[decimal] scale 10 out of range`},
		{"0.0", 10, `[decimal] scale 10 out of range`},
		{"0.0", -2, `[decimal] scale -2 out of range`},
	}

	for _, rec := range tests {
		_, err := decimal.FromStringWithScale(rec.s, rec.scale)
		assert.Error(t, err)
		assert.Equal(t, rec.msg, err.Error())
	}
}

func TestFromStringWithScale(t *testing.T) {
	tests := []struct {
		s     string
		scale uint8
		exp   string
	}{
		{"0", 0, "0"},
		{"0.00", 2, "0.00"},
		{"3.3", 3, "3.300"},
		{"3.333", 1, "3.3"},
		{"3.353", 1, "3.4"},
	}

	for _, rec := range tests {
		d, err := decimal.FromStringWithScale(rec.s, int(rec.scale))
		assert.NoError(t, err)
		assert.Equal(t, rec.scale, d.Scale(), rec.s)
		assert.Equal(t, rec.exp, d.String())
	}
}

func TestFromFloat(t *testing.T) {
	tests := []struct {
		v     float64
		scale uint8
		exp   string
	}{
		{0.0, 0, "0"},
		{1.0, 0, "1"},
		{1.33333, 2, "1.33"},
		{1.66666, 2, "1.66"},
		{-1.44444, 2, "-1.44"},
		{1.0, 2, "1.00"},
	}

	for _, rec := range tests {
		d := decimal.FromFloat(rec.v, rec.scale)
		assert.Equal(t, rec.scale, d.Scale())
		assert.Equal(t, rec.exp, d.String())
	}
}

func TestFromInt(t *testing.T) {
	for _, i := range []int64{0, -1, 123456789012345678, -123456789012345678} {
		d := decimal.FromInt(i)
		assert.Equal(t, uint8(0), d.Scale())
		assert.Equal(t, fmt.Sprint(i), d.String())
	}
}

func TestGoStringer(t *testing.T) {
	d, err := decimal.FromString("3.30")
	assert.NoError(t, err)
	assert.Equal(t, d.GoString(), "3.30m")
}

func TestShortString(t *testing.T) {
	tests := []struct {
		s   string
		exp string
	}{
		{"3.4", "3.4"},
		{"300", "300"},
		{"3.00", "3"},
		{"0.00", "0"},
		{"10.00", "10"},
		{"10", "10"},
		{"3.30", "3.3"},
	}

	for _, rec := range tests {
		d, err := decimal.FromString(rec.s)
		assert.NoError(t, err)
		assert.Equal(t, rec.exp, d.ShortString())
	}
}

func TestToInt(t *testing.T) {
	tests := []struct {
		s   string
		exp int64
	}{
		{"345", 345},
		{"-345", -345},
		{"123456789012345678", 123456789012345678},
		{"3.345", 3},
		{"-3.345", -3},
		{"3.5", 4},
		{"-3.5", -4},
	}

	for _, rec := range tests {
		d, err := decimal.FromString(rec.s)
		assert.NoError(t, err)
		assert.Equal(t, rec.exp, d.Int64())
	}
}

func TestToFloat64(t *testing.T) {
	tests := []struct {
		s   string
		exp float64
	}{
		{"0", 0.0},
		{"1000", 1000.0},
		{"-1.456", -1.456},
	}

	for _, rec := range tests {
		d, err := decimal.FromString(rec.s)
		assert.NoError(t, err)
		assert.Equal(t, rec.exp, d.Float64())
	}
}

func TestNegate(t *testing.T) {
	tests := []struct {
		s   string
		exp string
	}{
		{"0", "0"},
		{"0.00", "0.00"},
		{"3.456", "-3.456"},
		{"-3.4456", "3.4456"},
	}

	for _, rec := range tests {
		d, err := decimal.FromString(rec.s)
		assert.NoError(t, err)
		assert.Equal(t, rec.exp, d.Neg().String())
	}
}

type BinTestCase struct {
	a   string
	b   string
	exp string
}

func assertBinOp(t *testing.T, tests []BinTestCase, op func(x, y decimal.Decimal) decimal.Decimal) {
	for _, rec := range tests {
		a, err := decimal.FromString(rec.a)
		assert.NoError(t, err)

		b, err := decimal.FromString(rec.b)
		assert.NoError(t, err)

		act := op(a, b)
		assert.Equal(t, rec.exp, act.String())
	}
}

type BinTestCaseWithOption struct {
	a   string
	b   string
	exp string
	c   int
}

func assertBinOpWithOption(t *testing.T, tests []BinTestCaseWithOption, op func(x, y decimal.Decimal, c int) decimal.Decimal) {
	for _, rec := range tests {
		a, err := decimal.FromString(rec.a)
		assert.NoError(t, err)

		b, err := decimal.FromString(rec.b)
		assert.NoError(t, err)

		act := op(a, b, rec.c)
		assert.Equal(t, rec.exp, act.String())
	}
}

func TestAdd(t *testing.T) {
	tests := []BinTestCase{
		{"3", "4", "7"},
		{"3", "-4", "-1"},
		{"1.1", "2.2", "3.3"},
		{"1", "4.5", "5.5"},
		{"1.2", "2.8", "4.0"},
		{"0.0003", "0.0007", "0.0010"},
		{"0.3", "300", "300.3"},
	}

	assertBinOp(t, tests, func(a, b decimal.Decimal) decimal.Decimal {
		return a.Add(b)
	})
}

func TestAddToScale(t *testing.T) {
	tests := []BinTestCaseWithOption{
		{"3.10", "4.01", "7.11", 2},
		{"3.10", "1", "4.100", 3},
		{"3.1", "0.05", "3.2", 1},
		{"3.1", "0.04", "3.1", 1},
	}

	assertBinOpWithOption(t, tests, func(a, b decimal.Decimal, c int) decimal.Decimal {
		return a.AddToScale(b, c)
	})
}

func TestSubtract(t *testing.T) {
	tests := []BinTestCase{
		{"4", "3", "1"},
		{"-4", "3", "-7"},
		{"2.2", "1.1", "1.1"},
		{"4.5", "1", "3.5"},
		{"2.8", "1.2", "1.6"},
		{"0.0007", "0.0003", "0.0004"},
		{"300", "0.3", "299.7"},
	}

	assertBinOp(t, tests, func(a, b decimal.Decimal) decimal.Decimal {
		return a.Sub(b)
	})
}

func TestSubtractToScale(t *testing.T) {
	tests := []BinTestCaseWithOption{
		{"1.2", "0.2", "1.0", 1},
		{"1.2", "0.2", "1.00", 2},
		{"1", "0.4", "1", 0},
		{"1", "0.6", "0", 0},
	}

	assertBinOpWithOption(t, tests, func(a, b decimal.Decimal, c int) decimal.Decimal {
		return a.SubToScale(b, c)
	})
}

func TestMultiply(t *testing.T) {
	tests := []BinTestCase{
		{"3", "0", "0"},
		{"3", "4", "12"},
		{"3", "0.4", "1.2"},
		{"0.3", "0.5", "0.2"},
		{"-0.3", "0.5", "-0.2"},
		{"0.3", "0.4", "0.1"},
		{"-100", "0.01", "-1.00"},
		{"-0.001", "-0.01", "0.000"},
		{"40", "50", "2000"},
	}

	assertBinOp(t, tests, func(a, b decimal.Decimal) decimal.Decimal {
		return a.Mul(b)
	})
}

func TestMultiplyToScale(t *testing.T) {
	tests := []BinTestCaseWithOption{
		{"1.2", "0.6", "0.7", 1},
		{"1.2", "0.6", "0.72", 2},
		{"1.5555", "1", "1.6", 1},
		{"1.455", "1", "1", 0},
		{"1.445", "0.1", "0.14", 2},
		{"1.0", "2", "2.000", 3},
	}

	assertBinOpWithOption(t, tests, func(a, b decimal.Decimal, c int) decimal.Decimal {
		return a.MulToScale(b, c)
	})
}

func TestDiv(t *testing.T) {
	tests := []BinTestCase{
		{"9", "3", "3"},
		{"1", "4", "0"},
		{"1", "4.0", "0.3"},
		{"1", "3.0", "0.3"},
		{"1", "1.00001", "0.99999"},
		{"100000", "3", "33333"},
		{"5", "6000", "0"},
		{"5.0", "6000", "0.0"},
		{"5.0000", "6000", "0.0008"},
		{"-1", "3", "0"},
		{"-1", "6.0", "-0.2"},
		{"1", "0.25", "4.00"},
	}

	assertBinOp(t, tests, func(a, b decimal.Decimal) decimal.Decimal {
		return a.Div(b)
	})
}

func TestDivToScale(t *testing.T) {
	tests := []BinTestCaseWithOption{
		{"9", "3", "3", 0},
		{"10", "4", "2.50", 2},
		{"1.454", "1", "1.5", 1},
		{"1.454", "1", "1.45", 2},
	}

	assertBinOpWithOption(t, tests, func(a, b decimal.Decimal, c int) decimal.Decimal {
		return a.DivToScale(b, c)
	})
}

func TestDivByZero(t *testing.T) {
	assert.Panics(t, func() {
		decimal.FromInt(1).Div(decimal.FromInt(0))
	})
}

func TestRound(t *testing.T) {
	tests := []struct {
		s   string
		n   int
		exp string
	}{
		{"3.4", 1, "3.4"},
		{"3.4", 3, "3.400"},
		{"3.45", 1, "3.5"},
		{"3.44", 1, "3.4"},
		{"-3.45", 1, "-3.5"},
		{"-3.44", 1, "-3.4"},
	}

	for _, rec := range tests {
		d, err := decimal.FromString(rec.s)
		assert.NoError(t, err)

		assert.Equal(t, rec.exp, d.Round(rec.n).String())
	}
}

func TestZero(t *testing.T) {
	for i := 0; i < 9; i++ {
		d := decimal.Zero(i)
		assert.True(t, decimal.FromInt(0).EQ(d))
		assert.Equal(t, uint8(i), d.Scale())
	}
}

func TestIsZero(t *testing.T) {
	tests := []struct {
		s   string
		exp bool
	}{
		{"0.00", true},
		{"0.01", false},
	}

	for _, rec := range tests {
		d, err := decimal.FromString(rec.s)
		assert.NoError(t, err)

		assert.Equal(t, rec.exp, d.IsZero())
	}
}

func TestSign(t *testing.T) {
	tests := []struct {
		s   string
		exp int
	}{
		{"0", 0},
		{"1.333", 1},
		{"-2.3", -1},
	}

	for _, rec := range tests {
		d, err := decimal.FromString(rec.s)
		assert.NoError(t, err)

		assert.Equal(t, rec.exp, d.Sign())
	}
}

func TestCmp(t *testing.T) {
	tests := []struct {
		a, b string
		exp  int
	}{
		{"0", "0", 0},
		{"1.00", "1.000", 0},
		{"1", "9", -1},
		{"2.1", "2", 1},
	}

	for _, rec := range tests {
		a, err := decimal.FromString(rec.a)
		assert.NoError(t, err)

		b, err := decimal.FromString(rec.b)
		assert.NoError(t, err)

		assert.Equal(t, rec.exp, a.Cmp(b))
	}
}

func TestLTOrGTE(t *testing.T) {
	tests := []struct {
		a, b string
		exp  bool
	}{
		{"0.0", "0.00", false},
		{"0.0", "1.00", true},
		{"0.01", "0.00", false},
	}

	for _, rec := range tests {
		a, err := decimal.FromString(rec.a)
		assert.NoError(t, err)

		b, err := decimal.FromString(rec.b)
		assert.NoError(t, err)

		assert.Equal(t, rec.exp, a.LT(b))
		assert.Equal(t, !rec.exp, a.GTE(b))
	}
}

func TestGTOrLTE(t *testing.T) {
	tests := []struct {
		a, b string
		exp  bool
	}{
		{"0.0", "0.00", false},
		{"0.0", "1.00", false},
		{"0.01", "0.00", true},
	}

	for _, rec := range tests {
		a, err := decimal.FromString(rec.a)
		assert.NoError(t, err)

		b, err := decimal.FromString(rec.b)
		assert.NoError(t, err)

		assert.Equal(t, rec.exp, a.GT(b))
		assert.Equal(t, !rec.exp, a.LTE(b))
	}
}

func TestEQNotEQ(t *testing.T) {
	tests := []struct {
		a, b string
		exp  bool
	}{
		{"0.0", "0.00", true},
		{"0.0", "1.00", false},
		{"0.01", "0.00", false},
	}

	for _, rec := range tests {
		a, err := decimal.FromString(rec.a)
		assert.NoError(t, err)

		b, err := decimal.FromString(rec.b)
		assert.NoError(t, err)

		assert.Equal(t, rec.exp, a.EQ(b))
		assert.Equal(t, !rec.exp, a.NE(b))
	}
}

func TestToFromDecimal128(t *testing.T) {
	tests := []struct {
		s         string
		low, high uint64
	}{
		{"0", 0, 0x3040000000000000},
		{"0.00", 0, 0x303c000000000000},
		{"1", 1, 0x3040000000000000},
		{"-1", 1, 0xb040000000000000},
	}

	for _, rec := range tests {
		d, err := decimal.FromString(rec.s)
		assert.NoError(t, err)

		low, high := d.ToDecimal128()
		assert.Equal(t, rec.low, low)
		assert.Equal(t, rec.high, high)

		back := decimal.FromDecimal128(low, high)
		assert.Equal(t, d, back)
	}
}

func TestRandomToFromDecimal128(t *testing.T) {
	for i := 0; i < 200; i++ {
		scale := rand.Intn(9)
		d, err := decimal.FromStringWithScale(
			strconv.FormatFloat(rand.Float64()*float64(rand.Int31()), 'f', scale, 64),
			scale)
		assert.NoError(t, err)
		low, high := d.ToDecimal128()
		back := decimal.FromDecimal128(low, high)
		assert.Equal(t, d, back)
	}
}

func TestJson(t *testing.T) {
	d, err := decimal.FromString("3.33")
	assert.NoError(t, err)

	buf, err := json.Marshal(d)
	assert.NoError(t, err)

	var back decimal.Decimal
	assert.NoError(t, json.Unmarshal(buf, &back))
	assert.True(t, d.EQ(back))
}
