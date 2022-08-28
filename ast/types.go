package ast

import (
	"fmt"
	"strings"
)

// variable / constant type
// string, int, uint, [string], [float], {string:uint}, or [CustomType], {CustomType:CustomType2}, ...
type Type interface {
	typ()
}

type StringType struct{ Val string }

func (StringType) typ() {}

type CharType struct{ Val rune }

func (CharType) typ() {}

type UintType struct{ Val uint }

func (UintType) typ() {}

type IntType struct{ Val int }

func (IntType) typ() {}

type BoolType struct{ Val bool }

func (BoolType) typ() {}

type FloatType struct{ Val float64 }

func (FloatType) typ() {}

type SliceType struct {
	Typ   Type
	Elems []Expr
}

func (SliceType) typ() {}
func (SliceType) E()   {}
func (s SliceType) String() string {
	builder := strings.Builder{}
	builder.WriteString("[")
	for i, v := range s.Elems {
		builder.WriteString(fmt.Sprintf("%s", v))
		if i != len(s.Elems)-1 {
			builder.WriteString(", ")
		}
	}
	builder.WriteString("]")
	return builder.String()
}

type MapType struct {
	KeyType, ValueType Type
	Key, Value         Expr
	Elems              []MapType
}

func (MapType) typ() {}
func (MapType) E()   {}
func (m MapType) String() string {
	b := strings.Builder{}
	if m.Key == nil && m.Value == nil {
		b.WriteString("{}")
		return b.String()
	}
	b.WriteString(fmt.Sprintf("{%s:%s", m.Key, m.Value))
	for _, v := range m.Elems {
		b.WriteString(fmt.Sprintf(", %s:%s", v.Key, v.Value))
	}
	b.WriteString("}")
	return b.String()
}
