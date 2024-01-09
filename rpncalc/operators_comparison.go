package rpncalc

import (
	"fmt"

	"github.com/shopspring/decimal"
)

var OP_GT = Operator{
	Name: ">",
	Operation: func(i *Interpreter) error {
		operand_0, err := i.Pop()
		if err != nil {
			return fmt.Errorf("not enough operands (require 2, got 0): %v", err)
		}
		operand_1, err := i.Pop()
		if err != nil {
			return fmt.Errorf("not enough operands (require 2, got 1): %v", err)
		}
		value_0 := operand_0.Value()
		if value_0 == nil {
			return fmt.Errorf("first Operand was not a value")
		}
		value_1 := operand_1.Value()
		if value_1 == nil {
			return fmt.Errorf("second Operand was not a value")
		}

		result := value_0.DecimalValue.GreaterThan(value_1.DecimalValue)
		if result {
			rval := NewValue(decimal.NewFromInt(1))
			i.Push(rval)
		} else {
			rval := NewValue(decimal.NewFromInt(0))
			i.Push(rval)
		}

		return nil
	},
}
var OP_GTE = Operator{
	Name: ">=",
	Operation: func(i *Interpreter) error {
		operand_0, err := i.Pop()
		if err != nil {
			return fmt.Errorf("not enough operands (require 2, got 0): %v", err)
		}
		operand_1, err := i.Pop()
		if err != nil {
			return fmt.Errorf("not enough operands (require 2, got 1): %v", err)
		}
		value_0 := operand_0.Value()
		if value_0 == nil {
			return fmt.Errorf("first Operand was not a value")
		}
		value_1 := operand_1.Value()
		if value_1 == nil {
			return fmt.Errorf("second Operand was not a value")
		}

		result := value_0.DecimalValue.GreaterThanOrEqual(value_1.DecimalValue)
		if result {
			rval := NewValue(decimal.NewFromInt(1))
			i.Push(rval)
		} else {
			rval := NewValue(decimal.NewFromInt(0))
			i.Push(rval)
		}

		return nil
	},
}

var OP_LT = Operator{
	Name: "<",
	Operation: func(i *Interpreter) error {
		operand_0, err := i.Pop()
		if err != nil {
			return fmt.Errorf("not enough operands (require 2, got 0): %v", err)
		}
		operand_1, err := i.Pop()
		if err != nil {
			return fmt.Errorf("not enough operands (require 2, got 1): %v", err)
		}
		value_0 := operand_0.Value()
		if value_0 == nil {
			return fmt.Errorf("first Operand was not a value")
		}
		value_1 := operand_1.Value()
		if value_1 == nil {
			return fmt.Errorf("second Operand was not a value")
		}

		result := value_0.DecimalValue.LessThan(value_1.DecimalValue)
		if result {
			rval := NewValue(decimal.NewFromInt(1))
			i.Push(rval)
		} else {
			rval := NewValue(decimal.NewFromInt(0))
			i.Push(rval)
		}

		return nil
	},
}
var OP_LTE = Operator{
	Name: "<=",
	Operation: func(i *Interpreter) error {
		operand_0, err := i.Pop()
		if err != nil {
			return fmt.Errorf("not enough operands (require 2, got 0): %v", err)
		}
		operand_1, err := i.Pop()
		if err != nil {
			return fmt.Errorf("not enough operands (require 2, got 1): %v", err)
		}
		value_0 := operand_0.Value()
		if value_0 == nil {
			return fmt.Errorf("first Operand was not a value")
		}
		value_1 := operand_1.Value()
		if value_1 == nil {
			return fmt.Errorf("second Operand was not a value")
		}

		result := value_0.DecimalValue.LessThanOrEqual(value_1.DecimalValue)
		if result {
			rval := NewValue(decimal.NewFromInt(1))
			i.Push(rval)
		} else {
			rval := NewValue(decimal.NewFromInt(0))
			i.Push(rval)
		}

		return nil
	},
}
var OP_EQ = Operator{
	Name: "==",
	Operation: func(i *Interpreter) error {
		operand_0, err := i.Pop()
		if err != nil {
			return fmt.Errorf("not enough operands (require 2, got 0): %v", err)
		}
		operand_1, err := i.Pop()
		if err != nil {
			return fmt.Errorf("not enough operands (require 2, got 1): %v", err)
		}
		value_0 := operand_0.Value()
		if value_0 == nil {
			return fmt.Errorf("first Operand was not a value")
		}
		value_1 := operand_1.Value()
		if value_1 == nil {
			return fmt.Errorf("second Operand was not a value")
		}

		result := value_0.DecimalValue.Equal(value_1.DecimalValue)
		if result {
			rval := NewValue(decimal.NewFromInt(1))
			i.Push(rval)
		} else {
			rval := NewValue(decimal.NewFromInt(0))
			i.Push(rval)
		}

		return nil
	},
}

var OP_If = Operator{
	Name:        "if",
	Description: "Pops the last item from the stack.  if it is truthy (not 0) we execute (!) the next item on the stack.  Otherwise we drop it.",
	Operation: func(i *Interpreter) error {
		operand_0, err := i.Pop()
		if err != nil {
			return fmt.Errorf("not enough operands (require 2, got 0): %v", err)
		}
		operand_1, err := i.Pop()
		if err != nil {
			return fmt.Errorf("not enough operands (require 2, got 1): %v", err)
		}
		value_0 := operand_0.Value()
		if value_0 == nil {
			return fmt.Errorf("first Operand was not a value")
		}
		value_1 := operand_1.SubStack()
		if value_1 == nil {
			return fmt.Errorf("second Operand was not a substack")
		}

		is_false := value_0.DecimalValue.Equal(decimal.NewFromInt(0))
		if !is_false {
			return i.ExecutSubStack(value_1)
		}

		return nil
	},
}

var OP_IfElse = Operator{
	Name: "ifelse",
	Description: "Pops the last 3 items from the stack.  if the last item is truthy (not 0) we execute (!) the next item on the stack." +
		"If it is false, we execute the item after that. with the result of the executed item pushed back on",
	Operation: func(i *Interpreter) error {
		operand_0, err := i.Pop()
		if err != nil {
			return fmt.Errorf("not enough operands (require 3, got 0): %v", err)
		}
		operand_1, err := i.Pop()
		if err != nil {
			return fmt.Errorf("not enough operands (require 3, got 1): %v", err)
		}
		operand_2, err := i.Pop()
		if err != nil {
			return fmt.Errorf("not enough operands (require 3, got 1): %v", err)
		}
		value_0 := operand_0.Value()
		if value_0 == nil {
			return fmt.Errorf("first Operand was not a value")
		}
		value_1 := operand_1.SubStack()
		if value_1 == nil {
			return fmt.Errorf("second Operand was not a substack")
		}
		value_2 := operand_2.SubStack()
		if value_1 == nil {
			return fmt.Errorf("third Operand was not a substack")
		}

		is_false := value_0.DecimalValue.Equal(decimal.NewFromInt(0))
		if is_false {
			return i.ExecutSubStack(value_2)
		} else {
			return i.ExecutSubStack(value_1)
		}

	},
}
