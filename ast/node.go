package ast

import "fmt"

// NodeType type of ast node
type NodeType int

const (
	Void NodeType = iota
	Int32
	Int64
	Bool
	String
	Float
	DateTime
	Decimal
	Table
	Object
	List
	Dict
	RPCObject
)

// Node ast node interface
type Node interface {
	// Type returns current node type
	Type() NodeType

	// Returns true if the type allows null value, which type string end with '?'
	//
	// 'void' type always to be null, its type string not end with '?', so Nullable()
	// returns false either.
	//
	// 'str', 'list', 'dict', 'table', use of empty values, these types are not
	// allowed to be null.
	//
	// Other types are allow to be nullable, includes 'object', 'rpcObject'.
	Nullable() bool

	fmt.Stringer
}

// BasicNode map to primitive types, includes 'rpcObject', 'object', 'table'.
type BasicNode struct {
	t        NodeType
	nullable bool
	ts       string
}

func (n BasicNode) Type() NodeType {
	return n.t
}

func (n BasicNode) Nullable() bool {
	return n.nullable
}

func (n BasicNode) String() string {
	return n.ts
}

// ItemNode is node with attached item Node, List, Dict is ItemNode.
type ItemNode struct {
	BasicNode

	Item string
}

// DecimalNode describe decimal node.
type DecimalNode struct {
	BasicNode

	// Scale of the decimal
	N int
}
