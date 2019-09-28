package ast

import (
	"fmt"
	"regexp"
	"strings"
)

var (
	basicTypes map[string]NodeType = map[string]NodeType{
		"void":     Void,
		"int":      Int32,
		"long":     Int64,
		"bool":     Bool,
		"str":      String,
		"double":   Float,
		"datetime": DateTime,
		"table":    Table,
		"object":   Object,
	}

	notAllowNullTypes map[NodeType]bool = map[NodeType]bool{
		Void: true, String: true, Table: true, List: true, Dict: true,
	}

	decimalPattern   *regexp.Regexp
	rpcObjectPattern *regexp.Regexp
)

func init() {
	var err error
	decimalPattern, err = regexp.Compile(`^decimal(\([0-8]\))$`)
	if err != nil {
		panic(err)
	}

	rpcObjectPattern, err = regexp.Compile(`^(\w+\.)+\w+$`)
	if err != nil {
		panic(err)
	}
}

// Parse rpc type string to Node.
func Parse(ts string) (Node, error) {
	nullable := strings.HasSuffix(ts, "?")
	absTS := ts // ts without '?'
	if nullable {
		absTS = ts[0 : len(ts)-1]
	}

	if t, ok := basicTypes[absTS]; ok {
		return checkNotAllowNull(BasicNode{t, nullable, ts})
	}

	if matches := decimalPattern.FindStringSubmatch(absTS); matches != nil {
		d := matches[1][0] - '0'
		return checkNotAllowNull(DecimalNode{
			BasicNode{Decimal, nullable, ts},
			int(d),
		})
	}

	if strings.HasPrefix(absTS, "[") && strings.HasSuffix(absTS, "]") {
		if itemNode, err := Parse(absTS[1 : len(absTS)-1]); err != nil {
			return nil, err
		} else {
			return checkNotAllowNull(ItemNode{
				BasicNode{List, nullable, ts},
				itemNode,
			})
		}
	}

	if strings.HasPrefix(absTS, "{str:") && strings.HasSuffix(absTS, "}") {
		if itemNode, err := Parse(absTS[5 : len(absTS)-1]); err != nil {
			return nil, err
		} else {
			return checkNotAllowNull(ItemNode{
				BasicNode{Dict, nullable, ts},
				itemNode,
			})
		}
	}

	if rpcObjectPattern.MatchString(absTS) {
		return checkNotAllowNull(BasicNode{
			RpcObject, nullable, ts,
		})
	}
	return nil, fmt.Errorf("[%s] Wrong type string: '%s'", tag, ts)
}

func checkNotAllowNull(n Node) (Node, error) {
	if n.Nullable() && notAllowNullTypes[n.Type()] {
		return nil, fmt.Errorf("[%s] '%s' not support nullable", tag, n)
	}
	return n, nil
}
