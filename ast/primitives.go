package ast

import "fmt"

type SignedIntegerLiteral interface {
	Expr
	SIL()
}
type UnsignedIntegerLiteral interface {
	Expr
	USIL()
}
type IFloatLiteral interface {
	Expr
	FL()
}

type StringLiteral struct {
	Val string
}

func (s *StringLiteral) E()                    {}
func (s *StringLiteral) Value() string         { return s.Val }
func (s *StringLiteral) String() string        { return fmt.Sprintf("\"%s\"", s.Val) }
func (s *StringLiteral) IsANumericValue() bool { return false }

type BooleanLiteral struct {
	ValStr string
	Val    bool
}

func (b *BooleanLiteral) E()                    {}
func (b *BooleanLiteral) Value() string         { return b.ValStr }
func (b *BooleanLiteral) String() string        { return b.ValStr }
func (b *BooleanLiteral) IsANumericValue() bool { return false }

type IntLiteral struct {
	ValStr string
	Val    int64
}

func (i *IntLiteral) E()                    {}
func (i *IntLiteral) Value() string         { return i.ValStr }
func (i *IntLiteral) SIL()                  {}
func (i *IntLiteral) String() string        { return i.ValStr }
func (i *IntLiteral) IsANumericValue() bool { return true }

type UintLiteral struct {
	ValStr string
	Val    uint64
}

func (u *UintLiteral) E()                    {}
func (u *UintLiteral) Value() string         { return u.ValStr }
func (u *UintLiteral) USIL()                 {}
func (u *UintLiteral) String() string        { return u.ValStr }
func (u *UintLiteral) IsANumericValue() bool { return true }

type FloatLiteral struct {
	ValStr string
	Val    float64
}

func (f *FloatLiteral) E()                    {}
func (f *FloatLiteral) Value() string         { return f.ValStr }
func (f *FloatLiteral) FL()                   {}
func (f *FloatLiteral) String() string        { return f.ValStr }
func (f *FloatLiteral) IsANumericValue() bool { return true }
