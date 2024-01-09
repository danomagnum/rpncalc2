package rpncalc

import "fmt"

var OP_Add = Operator{
	Name:        "+",
	Description: "Add.  Second to last item plus the last item. Result pushed onto the stack.",
	Operation: func(i *Interpreter) error {
		ss := i.SaveState()
		operand_0, err := i.Pop()
		if err != nil {
			return fmt.Errorf("not enough operands (require 2, got 0): %v", err)
		}
		operand_1, err := i.Pop()
		if err != nil {
			i.RestoreState(ss)
			return fmt.Errorf("not enough operands (require 2, got 1): %v", err)
		}
		value_0 := operand_0.Value()
		if value_0 == nil {
			i.RestoreState(ss)
			return fmt.Errorf("first Operand was not a value")
		}
		value_1 := operand_1.Value()
		if value_1 == nil {
			i.RestoreState(ss)
			return fmt.Errorf("second Operand was not a value")
		}

		result := value_0.DecimalValue.Add(value_1.DecimalValue)
		rval := NewValue(result)
		i.Push(rval)

		return nil
	},
}

var OP_Sub = Operator{
	Name:        "-",
	Description: "Subtract.  Second to last item minus the last item. Result pushed onto the stack.",
	Operation: func(i *Interpreter) error {
		ss := i.SaveState()
		operand_0, err := i.Pop()
		if err != nil {
			return fmt.Errorf("not enough operands (require 2, got 0): %v", err)
		}
		operand_1, err := i.Pop()
		if err != nil {
			i.RestoreState(ss)
			return fmt.Errorf("not enough operands (require 2, got 1): %v", err)
		}
		value_0 := operand_0.Value()
		if value_0 == nil {
			i.RestoreState(ss)
			return fmt.Errorf("first Operand was not a value")
		}
		value_1 := operand_1.Value()
		if value_1 == nil {
			i.RestoreState(ss)
			return fmt.Errorf("second Operand was not a value")
		}

		result := value_1.DecimalValue.Sub(value_0.DecimalValue)
		rval := NewValue(result)
		i.Push(rval)

		return nil
	},
}
var OP_Mul = Operator{
	Name:        "*",
	Description: "Multiply.  Second to last item multipiled by the last item. Result pushed onto the stack.",
	Operation: func(i *Interpreter) error {
		ss := i.SaveState()
		operand_0, err := i.Pop()
		if err != nil {
			return fmt.Errorf("not enough operands (require 2, got 0): %v", err)
		}
		operand_1, err := i.Pop()
		if err != nil {
			i.RestoreState(ss)
			return fmt.Errorf("not enough operands (require 2, got 1): %v", err)
		}
		value_0 := operand_0.Value()
		if value_0 == nil {
			i.RestoreState(ss)
			return fmt.Errorf("first Operand was not a value")
		}
		value_1 := operand_1.Value()
		if value_1 == nil {
			i.RestoreState(ss)
			return fmt.Errorf("second Operand was not a value")
		}

		result := value_1.DecimalValue.Mul(value_0.DecimalValue)
		rval := NewValue(result)
		i.Push(rval)

		return nil
	},
}
var OP_Div = Operator{
	Name:        "/",
	Description: "Division.  Second to last item divided by the last item. Result pushed onto the stack.",
	Operation: func(i *Interpreter) error {
		ss := i.SaveState()
		operand_0, err := i.Pop()
		if err != nil {
			return fmt.Errorf("not enough operands (require 2, got 0): %v", err)
		}
		operand_1, err := i.Pop()
		if err != nil {
			i.RestoreState(ss)
			return fmt.Errorf("not enough operands (require 2, got 1): %v", err)
		}
		value_0 := operand_0.Value()
		if value_0 == nil {
			i.RestoreState(ss)
			return fmt.Errorf("first Operand was not a value")
		}
		value_1 := operand_1.Value()
		if value_1 == nil {
			i.RestoreState(ss)
			return fmt.Errorf("second Operand was not a value")
		}

		result := value_1.DecimalValue.Div(value_0.DecimalValue)
		rval := NewValue(result)
		i.Push(rval)

		return nil
	},
}

var OP_Mod = Operator{
	Name:        "%",
	Description: "Remainder / Modulo Operation second to last item mod last item. Result pushed onto the stack.",
	Operation: func(i *Interpreter) error {
		ss := i.SaveState()
		operand_0, err := i.Pop()
		if err != nil {
			return fmt.Errorf("not enough operands (require 2, got 0): %v", err)
		}
		operand_1, err := i.Pop()
		if err != nil {
			i.RestoreState(ss)
			return fmt.Errorf("not enough operands (require 2, got 1): %v", err)
		}
		value_0 := operand_0.Value()
		if value_0 == nil {
			i.RestoreState(ss)
			return fmt.Errorf("first Operand was not a value")
		}
		value_1 := operand_1.Value()
		if value_1 == nil {
			i.RestoreState(ss)
			return fmt.Errorf("second Operand was not a value")
		}

		result := value_1.DecimalValue.Mod(value_0.DecimalValue)
		rval := NewValue(result)
		i.Push(rval)

		return nil
	},
}
