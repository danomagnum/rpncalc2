package rpncalc

import (
	"fmt"
	"strings"

	"github.com/shopspring/decimal"
)

type ReprType int

const (
	ReprAny ReprType = iota
	ReprDecimal
	ReprHex
	ReprBinary
)

type Value struct {
	DecimalValue decimal.Decimal
	Repr         ReprType
}

func (v *Value) Value() *Value {
	return v
}
func (v *Value) Variable() *Variable {
	return nil
}
func (v *Value) Operator() *Operator {
	return nil
}
func (v *Value) SubStack() *Interpreter {
	return nil
}
func (v *Value) String() string {
	val := v.DecimalValue
	switch v.Repr {
	case ReprAny:
		return val.String()
	case ReprBinary:
		numstr := fmt.Sprintf("%b", val.IntPart())
		return fmt.Sprintf("0b%s", format_number_with_spaces(numstr, 4))
	case ReprDecimal:
		numstr := fmt.Sprintf("%d", val.IntPart())
		return format_number_with_spaces(numstr, 3)
	case ReprHex:
		numstr := fmt.Sprintf("%X", val.IntPart())
		return fmt.Sprintf("0x%s", format_number_with_spaces(numstr, 4))
	default:
		return val.String()
	}
}

func NewValue(val decimal.Decimal) *Value {
	return &Value{DecimalValue: val}
}

func format_number_with_spaces(s string, cnt int) string {
	cnt0 := len(s)
	padding_needed := cnt - cnt0%cnt
	if padding_needed == cnt {
		padding_needed = 0
	}
	out := strings.Builder{}
	for i := 0; i < padding_needed; i++ {
		out.WriteByte('0')
	}
	for i := range s {
		totalpos := i + padding_needed
		if totalpos%cnt == 0 && totalpos != 0 {
			out.WriteByte('_')
		}
		out.WriteByte(s[i])
	}

	return out.String()
}
