package decimal

// Decimaller interface implemented by all DecimalX alias types
type Decimaller interface {
	// Decimal returns its decimal value.
	Decimal() Decimal

	// Scale returns expected scale, Decimal1 returns 1 for example
	Scale() uint8
}

// NullDecimaller interface implemented by all NullDecimalX alias types
type NullDecimaller interface {
	// NullDecimal returns its null decimal value.
	NullDecimal() NullDecimal

	// Scale returns expected scale, NullDecimal1 returns 1 for example
	Scale() uint8
}

func marshalJSON(d Decimaller) ([]byte, error) {
	return d.Decimal().MarshalJSON()
}

func marshalJSONN(d NullDecimaller) ([]byte, error) {
	return d.NullDecimal().MarshalJSON()
}

func unmarshalJSON(pv *Decimal, scale int, buf []byte) error {
	if err := pv.UnmarshalJSON(buf); err != nil {
		return err
	}

	if pv.Scale() != uint8(scale) {
		*pv = pv.Round(scale)
	}
	return nil
}

func unmarshalJSONN(pv *NullDecimal, scale int, buf []byte) error {
	if err := pv.UnmarshalJSON(buf); err != nil {
		return err
	}

	if pv.Valid && pv.V.Scale() != uint8(scale) {
		pv.V = pv.V.Round(scale)
	}
	return nil
}

// Decimal0 alias for 0 scale, use alias to tell jsonrpc parser which decimal is.
type Decimal0 Decimal

// Decimal returns the wrapped decimal value
func (d Decimal0) Decimal() Decimal {
	return Decimal(d)
}

// Scale implement Decimaller interface
func (d Decimal0) Scale() uint8 {
	return 0
}

func (d Decimal0) String() string {
	return Decimal(d).String()
}

// MarshalJSON implement json.Marshaler interface
func (d Decimal0) MarshalJSON() ([]byte, error) {
	return marshalJSON(d)
}

// UnmarshalJSON impelmnet json.Unmarshaler interface
func (d *Decimal0) UnmarshalJSON(buf []byte) error {
	return unmarshalJSON((*Decimal)(d), 0, buf)
}

// UnmarshalJSON impelmnet json.Unmarshaler interface
func (d *Decimal1) UnmarshalJSON(buf []byte) error {
	return unmarshalJSON((*Decimal)(d), 1, buf)
}

// UnmarshalJSON impelmnet json.Unmarshaler interface
func (d *Decimal2) UnmarshalJSON(buf []byte) error {
	return unmarshalJSON((*Decimal)(d), 2, buf)
}

// UnmarshalJSON impelmnet json.Unmarshaler interface
func (d *Decimal3) UnmarshalJSON(buf []byte) error {
	return unmarshalJSON((*Decimal)(d), 3, buf)
}

// UnmarshalJSON impelmnet json.Unmarshaler interface
func (d *Decimal4) UnmarshalJSON(buf []byte) error {
	return unmarshalJSON((*Decimal)(d), 4, buf)
}

// UnmarshalJSON impelmnet json.Unmarshaler interface
func (d *Decimal5) UnmarshalJSON(buf []byte) error {
	return unmarshalJSON((*Decimal)(d), 5, buf)
}

// UnmarshalJSON impelmnet json.Unmarshaler interface
func (d *Decimal6) UnmarshalJSON(buf []byte) error {
	return unmarshalJSON((*Decimal)(d), 6, buf)
}

// UnmarshalJSON impelmnet json.Unmarshaler interface
func (d *Decimal7) UnmarshalJSON(buf []byte) error {
	return unmarshalJSON((*Decimal)(d), 7, buf)
}

// UnmarshalJSON impelmnet json.Unmarshaler interface
func (d *Decimal8) UnmarshalJSON(buf []byte) error {
	return unmarshalJSON((*Decimal)(d), 8, buf)
}

// UnmarshalJSON impelmnet json.Unmarshaler interface
func (d *NullDecimal0) UnmarshalJSON(buf []byte) error {
	return unmarshalJSONN((*NullDecimal)(d), 0, buf)
}

// UnmarshalJSON impelmnet json.Unmarshaler interface
func (d *NullDecimal1) UnmarshalJSON(buf []byte) error {
	return unmarshalJSONN((*NullDecimal)(d), 1, buf)
}

// UnmarshalJSON impelmnet json.Unmarshaler interface
func (d *NullDecimal2) UnmarshalJSON(buf []byte) error {
	return unmarshalJSONN((*NullDecimal)(d), 2, buf)
}

// UnmarshalJSON impelmnet json.Unmarshaler interface
func (d *NullDecimal3) UnmarshalJSON(buf []byte) error {
	return unmarshalJSONN((*NullDecimal)(d), 3, buf)
}

// UnmarshalJSON impelmnet json.Unmarshaler interface
func (d *NullDecimal4) UnmarshalJSON(buf []byte) error {
	return unmarshalJSONN((*NullDecimal)(d), 4, buf)
}

// UnmarshalJSON impelmnet json.Unmarshaler interface
func (d *NullDecimal5) UnmarshalJSON(buf []byte) error {
	return unmarshalJSONN((*NullDecimal)(d), 5, buf)
}

// UnmarshalJSON impelmnet json.Unmarshaler interface
func (d *NullDecimal6) UnmarshalJSON(buf []byte) error {
	return unmarshalJSONN((*NullDecimal)(d), 6, buf)
}

// UnmarshalJSON impelmnet json.Unmarshaler interface
func (d *NullDecimal7) UnmarshalJSON(buf []byte) error {
	return unmarshalJSONN((*NullDecimal)(d), 7, buf)
}

// UnmarshalJSON impelmnet json.Unmarshaler interface
func (d *NullDecimal8) UnmarshalJSON(buf []byte) error {
	return unmarshalJSONN((*NullDecimal)(d), 8, buf)
}

// NullDecimal0 alias for 0 scale, use alias to tell jsonrpc parser which decimal type is.
type NullDecimal0 NullDecimal

// NullDecimal returns wrapped value
func (d NullDecimal0) NullDecimal() NullDecimal {
	return NullDecimal(d)
}

// Scale implement Decimaller interface
func (d NullDecimal0) Scale() uint8 {
	return 0
}

func (d NullDecimal0) String() string {
	return NullDecimal(d).String()
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
type Decimal1 Decimal

// Decimal returns its decimal value.
func (d Decimal1) Decimal() Decimal {
	return Decimal(d)
}

// Scale implement Decimaller interface
func (d Decimal1) Scale() uint8 {
	return 1
}

func (d Decimal1) String() string {
	return Decimal(d).String()
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
type NullDecimal1 NullDecimal

// NullDecimal returns wrapped value
func (d NullDecimal1) NullDecimal() NullDecimal {
	return NullDecimal(d)
}

// Scale implement Decimaller interface
func (d NullDecimal1) Scale() uint8 {
	return 1
}

func (d NullDecimal1) String() string {
	return NullDecimal(d).String()
}

// Decimal2 alias for 2 scale, use alias to tell jsonrpc parser which decimal is.
type Decimal2 Decimal

// Decimal returns its decimal value.
func (d Decimal2) Decimal() Decimal {
	return Decimal(d)
}

// Scale implement Decimaller interface
func (d Decimal2) Scale() uint8 {
	return 2
}

func (d Decimal2) String() string {
	return Decimal(d).String()
}

// NullDecimal2 alias for 2 scale, use alias to tell jsonrpc parser which decimal type is.
type NullDecimal2 NullDecimal

// NullDecimal returns wrapped value
func (d NullDecimal2) NullDecimal() NullDecimal {
	return NullDecimal(d)
}

// Scale implement Decimaller interface
func (d NullDecimal2) Scale() uint8 {
	return 2
}

func (d NullDecimal2) String() string {
	return NullDecimal(d).String()
}

// Decimal3 alias for 3 scale, use alias to tell jsonrpc parser which decimal is.
type Decimal3 Decimal

// Decimal returns its decimal value.
func (d Decimal3) Decimal() Decimal {
	return Decimal(d)
}

// Scale implement Decimaller interface
func (d Decimal3) Scale() uint8 {
	return 3
}

func (d Decimal3) String() string {
	return Decimal(d).String()
}

// NullDecimal3 alias for 3 scale, use alias to tell jsonrpc parser which decimal type is.
type NullDecimal3 NullDecimal

// NullDecimal returns wrapped value
func (d NullDecimal3) NullDecimal() NullDecimal {
	return NullDecimal(d)
}

// Scale implement Decimaller interface
func (d NullDecimal3) Scale() uint8 {
	return 3
}

func (d NullDecimal3) String() string {
	return NullDecimal(d).String()
}

// Decimal4 alias for 4 scale, use alias to tell jsonrpc parser which decimal is.
type Decimal4 Decimal

// Decimal returns its decimal value.
func (d Decimal4) Decimal() Decimal {
	return Decimal(d)
}

// Scale implement Decimaller interface
func (d Decimal4) Scale() uint8 {
	return 4
}

func (d Decimal4) String() string {
	return Decimal(d).String()
}

// NullDecimal4 alias for 4 scale, use alias to tell jsonrpc parser which decimal type is.
type NullDecimal4 NullDecimal

// NullDecimal returns wrapped value
func (d NullDecimal4) NullDecimal() NullDecimal {
	return NullDecimal(d)
}

// Scale implement Decimaller interface
func (d NullDecimal4) Scale() uint8 {
	return 4
}

func (d NullDecimal4) String() string {
	return NullDecimal(d).String()
}

// Decimal5 alias for 5 scale, use alias to tell jsonrpc parser which decimal is.
type Decimal5 Decimal

// Decimal returns its decimal value.
func (d Decimal5) Decimal() Decimal {
	return Decimal(d)
}

// Scale implement Decimaller interface
func (d Decimal5) Scale() uint8 {
	return 5
}

func (d Decimal5) String() string {
	return Decimal(d).String()
}

// NullDecimal5 alias for 5 scale, use alias to tell jsonrpc parser which decimal type is.
type NullDecimal5 NullDecimal

// NullDecimal returns wrapped value
func (d NullDecimal5) NullDecimal() NullDecimal {
	return NullDecimal(d)
}

// Scale implement Decimaller interface
func (d NullDecimal5) Scale() uint8 {
	return 5
}

func (d NullDecimal5) String() string {
	return NullDecimal(d).String()
}

// Decimal6 alias for 6 scale, use alias to tell jsonrpc parser which decimal is.
type Decimal6 Decimal

// Decimal returns its decimal value.
func (d Decimal6) Decimal() Decimal {
	return Decimal(d)
}

// Scale implement Decimaller interface
func (d Decimal6) Scale() uint8 {
	return 6
}

func (d Decimal6) String() string {
	return Decimal(d).String()
}

// NullDecimal6 alias for 6 scale, use alias to tell jsonrpc parser which decimal type is.
type NullDecimal6 NullDecimal

// NullDecimal returns wrapped value
func (d NullDecimal6) NullDecimal() NullDecimal {
	return NullDecimal(d)
}

// Scale implement Decimaller interface
func (d NullDecimal6) Scale() uint8 {
	return 6
}

func (d NullDecimal6) String() string {
	return NullDecimal(d).String()
}

// Decimal7 alias for 7 scale, use alias to tell jsonrpc parser which decimal is.
type Decimal7 Decimal

// Decimal returns its decimal value.
func (d Decimal7) Decimal() Decimal {
	return Decimal(d)
}

// Scale implement Decimaller interface
func (d Decimal7) Scale() uint8 {
	return 7
}

func (d Decimal7) String() string {
	return Decimal(d).String()
}

// NullDecimal7 alias for 7 scale, use alias to tell jsonrpc parser which decimal type is.
type NullDecimal7 NullDecimal

// NullDecimal returns wrapped value
func (d NullDecimal7) NullDecimal() NullDecimal {
	return NullDecimal(d)
}

// Scale implement Decimaller interface
func (d NullDecimal7) Scale() uint8 {
	return 7
}

func (d NullDecimal7) String() string {
	return NullDecimal(d).String()
}

// Decimal8 alias for 8 scale, use alias to tell jsonrpc parser which decimal is.
type Decimal8 Decimal

// Decimal returns wrapped value
func (d Decimal8) Decimal() Decimal {
	return Decimal(d)
}

// Scale implement Decimaller interface
func (d Decimal8) Scale() uint8 {
	return 8
}

func (d Decimal8) String() string {
	return Decimal(d).String()
}

// NullDecimal8 alias for 8 scale, use alias to tell jsonrpc parser which decimal type is.
type NullDecimal8 NullDecimal

// NullDecimal returns wrapped value
func (d NullDecimal8) NullDecimal() NullDecimal {
	return NullDecimal(d)
}

// Scale implement Decimaller interface
func (d NullDecimal8) Scale() uint8 {
	return 8
}

func (d NullDecimal8) String() string {
	return NullDecimal(d).String()
}

func (d NullDecimal0) IsNull() bool {
	return d.Valid
}

func (d NullDecimal0) Val() interface{} {
	return d.V
}

func (d NullDecimal1) IsNull() bool {
	return d.Valid
}

func (d NullDecimal1) Val() interface{} {
	return d.V
}

func (d NullDecimal2) IsNull() bool {
	return d.Valid
}

func (d NullDecimal2) Val() interface{} {
	return d.V
}

func (d NullDecimal3) IsNull() bool {
	return d.Valid
}

func (d NullDecimal3) Val() interface{} {
	return d.V
}

func (d NullDecimal4) IsNull() bool {
	return d.Valid
}

func (d NullDecimal4) Val() interface{} {
	return d.V
}

func (d NullDecimal5) IsNull() bool {
	return d.Valid
}

func (d NullDecimal5) Val() interface{} {
	return d.V
}

func (d NullDecimal6) IsNull() bool {
	return d.Valid
}

func (d NullDecimal6) Val() interface{} {
	return d.V
}

func (d NullDecimal7) IsNull() bool {
	return d.Valid
}

func (d NullDecimal7) Val() interface{} {
	return d.V
}

func (d NullDecimal8) IsNull() bool {
	return d.Valid
}

func (d NullDecimal8) Val() interface{} {
	return d.V
}
