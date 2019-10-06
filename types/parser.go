package types

import (
	"sync"

	"github.com/bungle-suit/rpc/ast"
	"github.com/pkg/errors"
)

// Parser to parse rpc type from type string.
type Parser struct {
	l sync.Mutex

	cached map[string]Type
}

// NewParser create Parser.
//
// Parser is goroutine safe.
func NewParser() *Parser {
	return &Parser{
		cached: make(map[string]Type),
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

		return nullType{inner}, nil
	}

	switch n.Type() {
	case ast.List:
		inner, err := p.parse(n.(ast.ItemNode).Item)
		if err != nil {
			return nil, err
		}

		return listType{inner}, nil

	case ast.Dict:
		inner, err := p.parse(n.(ast.ItemNode).Item)
		if err != nil {
			return nil, err
		}

		return dictType{inner}, nil
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
	p.Define("int", directType{})
	p.Define("double", directType{})
	p.Define("long", longType{})
	p.Define("bool", boolType{})
	p.Define("str", directType{})
	p.Define("void", directType{})
	p.Define("datetime", datetimeType{})

	p.Define("decimal(0)", decimalType(0))
	p.Define("decimal(1)", decimalType(1))
	p.Define("decimal(2)", decimalType(2))
	p.Define("decimal(3)", decimalType(3))
	p.Define("decimal(4)", decimalType(4))
	p.Define("decimal(5)", decimalType(5))
	p.Define("decimal(6)", decimalType(6))
	p.Define("decimal(7)", decimalType(7))
	p.Define("decimal(8)", decimalType(8))
}
