package extvals

import (
	"encoding/json"
	"fmt"

	"github.com/bungle-suit/rpc/extvals/decimal"
)

// Decimaller interface implemented by all DecimalX alias types
type Decimaller interface {
	// Decimal returns its decimal value.
	Decimal() decimal.Decimal

	// Scale returns expected scale, Decimal1 returns 1 for example
	Scale() uint8
}

// NullDecimaller interface implemented by all NullDecimalX alias types
type NullDecimaller interface {
	// NullDecimal returns its null decimal value.
	NullDecimal() decimal.NullDecimal

	// Scale returns expected scale, NullDecimal1 returns 1 for example
	Scale() uint8
}

func marshalJSON(d Decimaller) ([]byte, error) {
	return d.Decimal().MarshalJSON()
}

func marshalJSONN(d NullDecimaller) ([]byte, error) {
	return d.NullDecimal().MarshalJSON()
}

func unmarshalJSON(pv *decimal.Decimal, scale int, buf []byte) error {
	if err := pv.UnmarshalJSON(buf); err != nil {
		return err
	}

	if pv.Scale() != uint8(scale) {
		*pv = pv.Round(scale)
	}
	return nil
}

func unmarshalJSONN(pv *decimal.NullDecimal, scale int, buf []byte) error {
	if err := pv.UnmarshalJSON(buf); err != nil {
		return err
	}

	if pv.Valid && pv.Decimal.Scale() != uint8(scale) {
		pv.Decimal = pv.Decimal.Round(scale)
	}
	return nil
}

// Decimal0 alias for 0 scale, use alias to tell jsonrpc parser which decimal is.
type Decimal0 decimal.Decimal

// Decimal returns the wrapped decimal value
func (d Decimal0) Decimal() decimal.Decimal {
	return decimal.Decimal(d)
}

// Scale implement Decimaller interface
func (d Decimal0) Scale() uint8 {
	return 0
}

func (d Decimal0) String() string {
	return decimal.Decimal(d).String()
}

// MarshalJSON implement json.Marshaler interface
func (d Decimal0) MarshalJSON() ([]byte, error) {
	return marshalJSON(d)
}

// UnmarshalJSON impelmnet json.Unmarshaler interface
func (d *Decimal0) UnmarshalJSON(buf []byte) error {
	return unmarshalJSON((*decimal.Decimal)(d), 0, buf)
}

// UnmarshalJSON impelmnet json.Unmarshaler interface
func (d *Decimal1) UnmarshalJSON(buf []byte) error {
	return unmarshalJSON((*decimal.Decimal)(d), 1, buf)
}

// UnmarshalJSON impelmnet json.Unmarshaler interface
func (d *Decimal2) UnmarshalJSON(buf []byte) error {
	return unmarshalJSON((*decimal.Decimal)(d), 2, buf)
}

// UnmarshalJSON impelmnet json.Unmarshaler interface
func (d *Decimal3) UnmarshalJSON(buf []byte) error {
	return unmarshalJSON((*decimal.Decimal)(d), 3, buf)
}

// UnmarshalJSON impelmnet json.Unmarshaler interface
func (d *Decimal4) UnmarshalJSON(buf []byte) error {
	return unmarshalJSON((*decimal.Decimal)(d), 4, buf)
}

// UnmarshalJSON impelmnet json.Unmarshaler interface
func (d *Decimal5) UnmarshalJSON(buf []byte) error {
	return unmarshalJSON((*decimal.Decimal)(d), 5, buf)
}

// UnmarshalJSON impelmnet json.Unmarshaler interface
func (d *Decimal6) UnmarshalJSON(buf []byte) error {
	return unmarshalJSON((*decimal.Decimal)(d), 6, buf)
}

// UnmarshalJSON impelmnet json.Unmarshaler interface
func (d *Decimal7) UnmarshalJSON(buf []byte) error {
	return unmarshalJSON((*decimal.Decimal)(d), 7, buf)
}

// UnmarshalJSON impelmnet json.Unmarshaler interface
func (d *Decimal8) UnmarshalJSON(buf []byte) error {
	return unmarshalJSON((*decimal.Decimal)(d), 8, buf)
}

// UnmarshalJSON impelmnet json.Unmarshaler interface
func (d *NullDecimal0) UnmarshalJSON(buf []byte) error {
	return unmarshalJSONN((*decimal.NullDecimal)(d), 0, buf)
}

// UnmarshalJSON impelmnet json.Unmarshaler interface
func (d *NullDecimal1) UnmarshalJSON(buf []byte) error {
	return unmarshalJSONN((*decimal.NullDecimal)(d), 1, buf)
}

// UnmarshalJSON impelmnet json.Unmarshaler interface
func (d *NullDecimal2) UnmarshalJSON(buf []byte) error {
	return unmarshalJSONN((*decimal.NullDecimal)(d), 2, buf)
}

// UnmarshalJSON impelmnet json.Unmarshaler interface
func (d *NullDecimal3) UnmarshalJSON(buf []byte) error {
	return unmarshalJSONN((*decimal.NullDecimal)(d), 3, buf)
}

// UnmarshalJSON impelmnet json.Unmarshaler interface
func (d *NullDecimal4) UnmarshalJSON(buf []byte) error {
	return unmarshalJSONN((*decimal.NullDecimal)(d), 4, buf)
}

// UnmarshalJSON impelmnet json.Unmarshaler interface
func (d *NullDecimal5) UnmarshalJSON(buf []byte) error {
	return unmarshalJSONN((*decimal.NullDecimal)(d), 5, buf)
}

// UnmarshalJSON impelmnet json.Unmarshaler interface
func (d *NullDecimal6) UnmarshalJSON(buf []byte) error {
	return unmarshalJSONN((*decimal.NullDecimal)(d), 6, buf)
}

// UnmarshalJSON impelmnet json.Unmarshaler interface
func (d *NullDecimal7) UnmarshalJSON(buf []byte) error {
	return unmarshalJSONN((*decimal.NullDecimal)(d), 7, buf)
}

// UnmarshalJSON impelmnet json.Unmarshaler interface
func (d *NullDecimal8) UnmarshalJSON(buf []byte) error {
	return unmarshalJSONN((*decimal.NullDecimal)(d), 8, buf)
}

// NullDecimal0 alias for 0 scale, use alias to tell jsonrpc parser which decimal type is.
type NullDecimal0 decimal.NullDecimal

// NullDecimal returns wrapped value
func (d NullDecimal0) NullDecimal() decimal.NullDecimal {
	return decimal.NullDecimal(d)
}

// Scale implement Decimaller interface
func (d NullDecimal0) Scale() uint8 {
	return 0
}

func (d NullDecimal0) String() string {
	return decimal.NullDecimal(d).String()
}

// MarshalJSON implement json.Marshaler interface
func (d NullDecimal0) MarshalJSON() ([]byte, error) {
	return marshalJSONN(d)
}

// MarshalJSON implement json.Marshaler interface
func (d NullDecimal1) MarshalJSON() ([]byte, error) {
	return marshalJSONN(d)
}

// MarshalJSON implement json.Marshaler interface
func (d NullDecimal2) MarshalJSON() ([]byte, error) {
	return marshalJSONN(d)
}

// MarshalJSON implement json.Marshaler interface
func (d NullDecimal3) MarshalJSON() ([]byte, error) {
	return marshalJSONN(d)
}

// MarshalJSON implement json.Marshaler interface
func (d NullDecimal4) MarshalJSON() ([]byte, error) {
	return marshalJSONN(d)
}

// MarshalJSON implement json.Marshaler interface
func (d NullDecimal5) MarshalJSON() ([]byte, error) {
	return marshalJSONN(d)
}

// MarshalJSON implement json.Marshaler interface
func (d NullDecimal6) MarshalJSON() ([]byte, error) {
	return marshalJSONN(d)
}

// MarshalJSON implement json.Marshaler interface
func (d NullDecimal7) MarshalJSON() ([]byte, error) {
	return marshalJSONN(d)
}

// MarshalJSON implement json.Marshaler interface
func (d NullDecimal8) MarshalJSON() ([]byte, error) {
	return marshalJSONN(d)
}

// Decimal1 alias for 1 scale, use alias to tell jsonrpc parser which decimal is.
type Decimal1 decimal.Decimal

// Decimal returns its decimal value.
func (d Decimal1) Decimal() decimal.Decimal {
	return decimal.Decimal(d)
}

// Scale implement Decimaller interface
func (d Decimal1) Scale() uint8 {
	return 1
}

func (d Decimal1) String() string {
	return decimal.Decimal(d).String()
}

// MarshalJSON implement json.Marshaler interface
func (d Decimal1) MarshalJSON() ([]byte, error) {
	return marshalJSON(d)
}

// MarshalJSON implement json.Marshaler interface
func (d Decimal2) MarshalJSON() ([]byte, error) {
	return marshalJSON(d)
}

// MarshalJSON implement json.Marshaler interface
func (d Decimal3) MarshalJSON() ([]byte, error) {
	return marshalJSON(d)
}

// MarshalJSON implement json.Marshaler interface
func (d Decimal4) MarshalJSON() ([]byte, error) {
	return marshalJSON(d)
}

// MarshalJSON implement json.Marshaler interface
func (d Decimal5) MarshalJSON() ([]byte, error) {
	return marshalJSON(d)
}

// MarshalJSON implement json.Marshaler interface
func (d Decimal6) MarshalJSON() ([]byte, error) {
	return marshalJSON(d)
}

// MarshalJSON implement json.Marshaler interface
func (d Decimal7) MarshalJSON() ([]byte, error) {
	return marshalJSON(d)
}

// MarshalJSON implement json.Marshaler interface
func (d Decimal8) MarshalJSON() ([]byte, error) {
	return marshalJSON(d)
}

// NullDecimal1 alias for 1 scale, use alias to tell jsonrpc parser which decimal type is.
type NullDecimal1 decimal.NullDecimal

// NullDecimal returns wrapped value
func (d NullDecimal1) NullDecimal() decimal.NullDecimal {
	return decimal.NullDecimal(d)
}

// Scale implement Decimaller interface
func (d NullDecimal1) Scale() uint8 {
	return 1
}

func (d NullDecimal1) String() string {
	return decimal.NullDecimal(d).String()
}

// Decimal2 alias for 2 scale, use alias to tell jsonrpc parser which decimal is.
type Decimal2 decimal.Decimal

// Decimal returns its decimal value.
func (d Decimal2) Decimal() decimal.Decimal {
	return decimal.Decimal(d)
}

// Scale implement Decimaller interface
func (d Decimal2) Scale() uint8 {
	return 2
}

func (d Decimal2) String() string {
	return decimal.Decimal(d).String()
}

// NullDecimal2 alias for 2 scale, use alias to tell jsonrpc parser which decimal type is.
type NullDecimal2 decimal.NullDecimal

// NullDecimal returns wrapped value
func (d NullDecimal2) NullDecimal() decimal.NullDecimal {
	return decimal.NullDecimal(d)
}

// Scale implement Decimaller interface
func (d NullDecimal2) Scale() uint8 {
	return 2
}

func (d NullDecimal2) String() string {
	return decimal.NullDecimal(d).String()
}

// Decimal3 alias for 3 scale, use alias to tell jsonrpc parser which decimal is.
type Decimal3 decimal.Decimal

// Decimal returns its decimal value.
func (d Decimal3) Decimal() decimal.Decimal {
	return decimal.Decimal(d)
}

// Scale implement Decimaller interface
func (d Decimal3) Scale() uint8 {
	return 3
}

func (d Decimal3) String() string {
	return decimal.Decimal(d).String()
}

// NullDecimal3 alias for 3 scale, use alias to tell jsonrpc parser which decimal type is.
type NullDecimal3 decimal.NullDecimal

// NullDecimal returns wrapped value
func (d NullDecimal3) NullDecimal() decimal.NullDecimal {
	return decimal.NullDecimal(d)
}

// Scale implement Decimaller interface
func (d NullDecimal3) Scale() uint8 {
	return 3
}

func (d NullDecimal3) String() string {
	return decimal.NullDecimal(d).String()
}

// Decimal4 alias for 4 scale, use alias to tell jsonrpc parser which decimal is.
type Decimal4 decimal.Decimal

// Decimal returns its decimal value.
func (d Decimal4) Decimal() decimal.Decimal {
	return decimal.Decimal(d)
}

// Scale implement Decimaller interface
func (d Decimal4) Scale() uint8 {
	return 4
}

func (d Decimal4) String() string {
	return decimal.Decimal(d).String()
}

// NullDecimal4 alias for 4 scale, use alias to tell jsonrpc parser which decimal type is.
type NullDecimal4 decimal.NullDecimal

// NullDecimal returns wrapped value
func (d NullDecimal4) NullDecimal() decimal.NullDecimal {
	return decimal.NullDecimal(d)
}

// Scale implement Decimaller interface
func (d NullDecimal4) Scale() uint8 {
	return 4
}

func (d NullDecimal4) String() string {
	return decimal.NullDecimal(d).String()
}

// Decimal5 alias for 5 scale, use alias to tell jsonrpc parser which decimal is.
type Decimal5 decimal.Decimal

// Decimal returns its decimal value.
func (d Decimal5) Decimal() decimal.Decimal {
	return decimal.Decimal(d)
}

// Scale implement Decimaller interface
func (d Decimal5) Scale() uint8 {
	return 5
}

func (d Decimal5) String() string {
	return decimal.Decimal(d).String()
}

// NullDecimal5 alias for 5 scale, use alias to tell jsonrpc parser which decimal type is.
type NullDecimal5 decimal.NullDecimal

// NullDecimal returns wrapped value
func (d NullDecimal5) NullDecimal() decimal.NullDecimal {
	return decimal.NullDecimal(d)
}

// Scale implement Decimaller interface
func (d NullDecimal5) Scale() uint8 {
	return 5
}

func (d NullDecimal5) String() string {
	return decimal.NullDecimal(d).String()
}

// Decimal6 alias for 6 scale, use alias to tell jsonrpc parser which decimal is.
type Decimal6 decimal.Decimal

// Decimal returns its decimal value.
func (d Decimal6) Decimal() decimal.Decimal {
	return decimal.Decimal(d)
}

// Scale implement Decimaller interface
func (d Decimal6) Scale() uint8 {
	return 6
}

func (d Decimal6) String() string {
	return decimal.Decimal(d).String()
}

// NullDecimal6 alias for 6 scale, use alias to tell jsonrpc parser which decimal type is.
type NullDecimal6 decimal.NullDecimal

// NullDecimal returns wrapped value
func (d NullDecimal6) NullDecimal() decimal.NullDecimal {
	return decimal.NullDecimal(d)
}

// Scale implement Decimaller interface
func (d NullDecimal6) Scale() uint8 {
	return 6
}

func (d NullDecimal6) String() string {
	return decimal.NullDecimal(d).String()
}

// Decimal7 alias for 7 scale, use alias to tell jsonrpc parser which decimal is.
type Decimal7 decimal.Decimal

// Decimal returns its decimal value.
func (d Decimal7) Decimal() decimal.Decimal {
	return decimal.Decimal(d)
}

// Scale implement Decimaller interface
func (d Decimal7) Scale() uint8 {
	return 7
}

func (d Decimal7) String() string {
	return decimal.Decimal(d).String()
}

// NullDecimal7 alias for 7 scale, use alias to tell jsonrpc parser which decimal type is.
type NullDecimal7 decimal.NullDecimal

// NullDecimal returns wrapped value
func (d NullDecimal7) NullDecimal() decimal.NullDecimal {
	return decimal.NullDecimal(d)
}

// Scale implement Decimaller interface
func (d NullDecimal7) Scale() uint8 {
	return 7
}

func (d NullDecimal7) String() string {
	return decimal.NullDecimal(d).String()
}

// Decimal8 alias for 8 scale, use alias to tell jsonrpc parser which decimal is.
type Decimal8 decimal.Decimal

// Decimal returns wrapped value
func (d Decimal8) Decimal() decimal.Decimal {
	return decimal.Decimal(d)
}

// Scale implement Decimaller interface
func (d Decimal8) Scale() uint8 {
	return 8
}

func (d Decimal8) String() string {
	return decimal.Decimal(d).String()
}

// NullDecimal8 alias for 8 scale, use alias to tell jsonrpc parser which decimal type is.
type NullDecimal8 decimal.NullDecimal

// NullDecimal returns wrapped value
func (d NullDecimal8) NullDecimal() decimal.NullDecimal {
	return decimal.NullDecimal(d)
}

// Scale implement Decimaller interface
func (d NullDecimal8) Scale() uint8 {
	return 8
}

func (d NullDecimal8) String() string {
	return decimal.NullDecimal(d).String()
}

// CheckScale checks DecimalN value's scale matches to its type
// causedBy argument specify CausedBy of returned error.
func CheckScale(d Decimaller) error {
	if d.Scale() != d.Decimal().Scale() {
		return fmt.Errorf("Mismatched decimaller scale, %d %d", d.Scale(), d.Decimal().Scale())
	}
	return nil
}

// CheckScaleNull checks NullDecimalN value's scale matches to its type
// causedBy argument specify CausedBy of returned error.
func CheckScaleNull(d NullDecimaller) error {
	if !d.NullDecimal().Valid {
		return nil
	}

	if d.Scale() != d.NullDecimal().Decimal.Scale() {
		return fmt.Errorf("Mismatched null decimaller scale, %d %d", d.Scale(), d.NullDecimal().Decimal.Scale())
	}
	return nil
}

/*
// decimalEncoderDecoder implement bsoncodec.ValueDecode/Encoder to
// convert bosn value to/from DecimalX.
type decimalEncoderDecoder struct {
}

func (decimalEncoderDecoder) DecodeValue(ctx bsoncodec.DecodeContext, r bsonrw.ValueReader, v reflect.Value) error {
	return nil
}

func (decimalEncoderDecoder) EncodeValue(ctx bsoncodec.EncodeContext, w bsonrw.ValueWriter, v reflect.Value) error {
	if v.Kind() != reflect.Struct || !v.CanInterface() {
		return errors.New("Not decimal value")
	}

	d, ok := v.Interface().(Decimaller)
	if !ok {
		return errors.New("Not decimal value 2")
	}

	return nil
}
*/

var (
	_ json.Marshaler = Decimal0{}
	_ json.Marshaler = Decimal1{}
	_ json.Marshaler = Decimal2{}
	_ json.Marshaler = Decimal3{}
	_ json.Marshaler = Decimal4{}
	_ json.Marshaler = Decimal5{}
	_ json.Marshaler = Decimal6{}
	_ json.Marshaler = Decimal7{}
	_ json.Marshaler = Decimal8{}

	_ json.Unmarshaler = &Decimal0{}
	_ json.Unmarshaler = &Decimal1{}
	_ json.Unmarshaler = &Decimal2{}
	_ json.Unmarshaler = &Decimal3{}
	_ json.Unmarshaler = &Decimal4{}
	_ json.Unmarshaler = &Decimal5{}
	_ json.Unmarshaler = &Decimal6{}
	_ json.Unmarshaler = &Decimal7{}
	_ json.Unmarshaler = &Decimal8{}

	_ json.Marshaler = NullDecimal0{}
	_ json.Marshaler = NullDecimal1{}
	_ json.Marshaler = NullDecimal2{}
	_ json.Marshaler = NullDecimal3{}
	_ json.Marshaler = NullDecimal4{}
	_ json.Marshaler = NullDecimal5{}
	_ json.Marshaler = NullDecimal6{}
	_ json.Marshaler = NullDecimal7{}
	_ json.Marshaler = NullDecimal8{}

	_ json.Unmarshaler = &NullDecimal0{}
	_ json.Unmarshaler = &NullDecimal1{}
	_ json.Unmarshaler = &NullDecimal2{}
	_ json.Unmarshaler = &NullDecimal3{}
	_ json.Unmarshaler = &NullDecimal4{}
	_ json.Unmarshaler = &NullDecimal5{}
	_ json.Unmarshaler = &NullDecimal6{}
	_ json.Unmarshaler = &NullDecimal7{}
	_ json.Unmarshaler = &NullDecimal8{}
)
