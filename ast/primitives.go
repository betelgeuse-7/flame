package ast

type SignedIntegerLiteral interface{ SIL() }
type UnsignedIntegerLiteral interface{ USIL() }
type IFloatLiteral interface{ FL() }

type StringLiteral struct {
	Val string
}

func (s *StringLiteral) E()             {}
func (s *StringLiteral) Value() string  { return s.Val }
func (s *StringLiteral) String() string { return s.Val }

type BooleanLiteral struct {
	ValStr string
	Val    bool
}

func (b *BooleanLiteral) E()             {}
func (b *BooleanLiteral) Value() string  { return b.ValStr }
func (b *BooleanLiteral) String() string { return b.ValStr }

type IntLiteral struct {
	ValStr string
	Val    int64
}

func (i *IntLiteral) E()             {}
func (i *IntLiteral) Value() string  { return i.ValStr }
func (i *IntLiteral) SIL()           {}
func (i *IntLiteral) String() string { return i.ValStr }

type I32Literal struct {
	ValStr string
	Val    int32
}

func (i *I32Literal) SIL()           {}
func (i *I32Literal) E()             {}
func (i *I32Literal) Value() string  { return i.ValStr }
func (i *I32Literal) String() string { return i.ValStr }

type U32Literal struct {
	ValStr string
	Val    uint32
}

func (u *U32Literal) E()             {}
func (u *U32Literal) Value() string  { return u.ValStr }
func (u *U32Literal) USIL()          {}
func (u *U32Literal) String() string { return u.ValStr }

type UintLiteral struct {
	ValStr string
	Val    uint64
}

func (u *UintLiteral) E()             {}
func (u *UintLiteral) Value() string  { return u.ValStr }
func (u *UintLiteral) USIL()          {}
func (u *UintLiteral) String() string { return u.ValStr }

type FloatLiteral struct {
	ValStr string
	Val    float64
}

func (f *FloatLiteral) E()             {}
func (f *FloatLiteral) Value() string  { return f.ValStr }
func (f *FloatLiteral) FL()            {}
func (f *FloatLiteral) String() string { return f.ValStr }

type F32Literal struct {
	ValStr string
	Val    float32
}

func (f *F32Literal) E()             {}
func (f *F32Literal) Value() string  { return f.ValStr }
func (f *F32Literal) FL()            {}
func (f *F32Literal) String() string { return f.ValStr }
