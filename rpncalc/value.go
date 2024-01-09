package rpncalc

import (
	"fmt"

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
		return fmt.Sprintf("0b%b", val.IntPart())
	case ReprDecimal:
		return fmt.Sprintf("%d", val.IntPart())
	case ReprHex:
		return fmt.Sprintf("0x%X", val.IntPart())
	default:
		return val.String()
	}
}

func NewValue(val decimal.Decimal) *Value {
	return &Value{DecimalValue: val}
}
