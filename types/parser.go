package types

import (
	"fmt"
	"reflect"
	"sync"
	"time"

	"github.com/bungle-suit/rpc/ast"
	"github.com/bungle-suit/rpc/extvals"
	"github.com/bungle-suit/rpc/extvals/decimal"
	"github.com/bungle-suit/rpc/extvals/table"
	"github.com/pkg/errors"
)

// Parser to parse rpc type from type string.
type Parser struct {
	l   sync.Mutex
	lgo sync.Mutex

	cached   map[string]Type
	goCached map[string]reflect.Type
}

// NewParser create Parser.
//
// Parser is goroutine safe.
func NewParser() *Parser {
	return &Parser{
		cached:   make(map[string]Type),
		goCached: make(map[string]reflect.Type),
	}
}

// Parse type from type string.
func (p *Parser) Parse(ts string) (Type, error) {
	p.l.Lock()
	t, err := p.parse(ts)
	p.l.Unlock()
	return t, err
}

// parse without lock
func (p *Parser) parse(ts string) (Type, error) {
	if t, ok := p.cached[ts]; ok {
		return t, nil
	}

	t, err := p.parseComposite(ts)
	if err != nil {
		return nil, err
	}
	p.cached[ts] = t

	return t, nil
}

func (p *Parser) parseComposite(ts string) (Type, error) {
	n, err := ast.Parse(ts)
	if err != nil {
		return nil, err
	}

	if n.Nullable() {
		inner, err := p.parse(ts[0 : len(ts)-1])
		if err != nil {
			return nil, err
		}

		// TODO: remove generic nullType, nullable types should pre-registered
		return nullType{inner}, nil
	}

	switch n.Type() {
	case ast.List:
		inner, err := p.parse(n.(ast.ItemNode).Item)
		if err != nil {
			return nil, err
		}

		return listType{p, inner, ts}, nil

	case ast.Dict:
		innerTS := n.(ast.ItemNode).Item
		inner, err := p.parse(innerTS)
		if err != nil {
			return nil, err
		}

		return dictType{p, inner, innerTS}, nil
	}

	return nil, errors.Errorf("[%s] Failed parse type '%s'", tag, ts)
}

// Define type for specific type string.
//
// NOTE: define primitive types such as 'int', 'rpcObject', Parser parse
// compound types itself (including nullable types). Mainly used to define rpc
// objects.
func (p *Parser) Define(ts string, ty Type) {
	p.l.Lock()
	p.cached[ts] = ty
	p.l.Unlock()
}

// DefinePrimitiveTypes define builtin primitive types:
//
//   'int', 'void', 'long', 'bool', 'str', 'double', 'datetime',
//   'decimal', 'table', 'object'.
func (p *Parser) DefinePrimitiveTypes() {
	p.Define("int", intType{})
	p.DefineGoType("int", reflect.TypeOf(int32(0)))
	p.Define("double", floatType{})
	p.DefineGoType("double", reflect.TypeOf(float64(0)))
	p.Define("long", longType{})
	p.DefineGoType("long", reflect.TypeOf(int64(0)))
	p.Define("bool", boolType{})
	p.DefineGoType("bool", reflect.TypeOf(true))
	p.Define("str", stringType{})
	p.DefineGoType("str", reflect.TypeOf(""))
	p.Define("void", voidType{})
	p.DefineGoType("void", reflect.TypeOf(""))
	p.Define("datetime", datetimeType{})
	p.DefineGoType("datetime", reflect.TypeOf(time.Time{}))
	p.Define("table", tableType{p})
	p.DefineGoType("table", reflect.TypeOf((*table.Table)(nil)))
	p.Define("object", objectType{p})
	p.DefineGoType("object", reflect.TypeOf(extvals.Object{}))

	p.Define("decimal(0)", decimalType(0))
	p.DefineGoType("decimal(0)", reflect.TypeOf(decimal.Decimal0{}))
	p.Define("decimal(1)", decimalType(1))
	p.DefineGoType("decimal(1)", reflect.TypeOf(decimal.Decimal1{}))
	p.Define("decimal(2)", decimalType(2))
	p.DefineGoType("decimal(2)", reflect.TypeOf(decimal.Decimal2{}))
	p.Define("decimal(3)", decimalType(3))
	p.DefineGoType("decimal(3)", reflect.TypeOf(decimal.Decimal3{}))
	p.Define("decimal(4)", decimalType(4))
	p.DefineGoType("decimal(4)", reflect.TypeOf(decimal.Decimal4{}))
	p.Define("decimal(5)", decimalType(5))
	p.DefineGoType("decimal(5)", reflect.TypeOf(decimal.Decimal5{}))
	p.Define("decimal(6)", decimalType(6))
	p.DefineGoType("decimal(6)", reflect.TypeOf(decimal.Decimal6{}))
	p.Define("decimal(7)", decimalType(7))
	p.DefineGoType("decimal(7)", reflect.TypeOf(decimal.Decimal7{}))
	p.Define("decimal(8)", decimalType(8))
	p.DefineGoType("decimal(8)", reflect.TypeOf(decimal.Decimal8{}))

	p.defineNullTypes()
}

func (p *Parser) defineNullTypes() {
	p.Define("bool?", nullBoolType{})
	p.DefineGoType("bool?", reflect.TypeOf(extvals.NullBool{}))
	p.Define("int?", nullIntType{})
	p.DefineGoType("int?", reflect.TypeOf(extvals.NullInt32{}))
	p.Define("long?", nullLongType{})
	p.DefineGoType("long?", reflect.TypeOf(extvals.NullInt64{}))
	p.Define("double?", nullFloatType{})
	p.DefineGoType("double?", reflect.TypeOf(extvals.NullFloat64{}))
	p.Define("datetime?", nullDatetimeType{})
	p.DefineGoType("datetime?", reflect.TypeOf(extvals.NullTime{}))

	p.Define("decimal(0)?", nullDecimalType(0))
	p.DefineGoType("decimal(0)?", reflect.TypeOf(decimal.NullDecimal0{}))
	p.Define("decimal(1)?", nullDecimalType(1))
	p.DefineGoType("decimal(1)?", reflect.TypeOf(decimal.NullDecimal1{}))
	p.Define("decimal(2)?", nullDecimalType(2))
	p.DefineGoType("decimal(2)?", reflect.TypeOf(decimal.NullDecimal2{}))
	p.Define("decimal(3)?", nullDecimalType(3))
	p.DefineGoType("decimal(3)?", reflect.TypeOf(decimal.NullDecimal3{}))
	p.Define("decimal(4)?", nullDecimalType(4))
	p.DefineGoType("decimal(4)?", reflect.TypeOf(decimal.NullDecimal4{}))
	p.Define("decimal(5)?", nullDecimalType(5))
	p.DefineGoType("decimal(5)?", reflect.TypeOf(decimal.NullDecimal5{}))
	p.Define("decimal(6)?", nullDecimalType(6))
	p.DefineGoType("decimal(6)?", reflect.TypeOf(decimal.NullDecimal6{}))
	p.Define("decimal(7)?", nullDecimalType(7))
	p.DefineGoType("decimal(7)?", reflect.TypeOf(decimal.NullDecimal7{}))
	p.Define("decimal(8)?", nullDecimalType(8))
	p.DefineGoType("decimal(8)?", reflect.TypeOf(decimal.NullDecimal8{}))
}

// ParseGoType returns golang type of corresponding type string.
func (p *Parser) ParseGoType(ts string) (reflect.Type, error) {
	p.lgo.Lock()
	t, ok := p.goCached[ts]
	p.lgo.Unlock()

	if ok {
		return t, nil
	}

	ty, err := p.parseCompositeGoType(ts)
	if err != nil {
		return nil, err
	}

	p.DefineGoType(ts, ty)
	return ty, nil
}

func (p *Parser) DefineGoType(ts string, ty reflect.Type) {
	p.lgo.Lock()
	p.goCached[ts] = ty
	p.lgo.Unlock()
}

func (p *Parser) parseCompositeGoType(ts string) (reflect.Type, error) {
	n, err := ast.Parse(ts)
	if err != nil {
		return nil, err
	}

	switch n.Type() {
	case ast.List:
		inner, err := p.ParseGoType(n.(ast.ItemNode).Item)
		if err != nil {
			return nil, err
		}
		return reflect.SliceOf(inner), nil
	}

	return nil, fmt.Errorf("[%s] Failed find golang type for '%s'", tag, ts)
}
