package rpncalc

import (
	"fmt"
)

var OP_ToHex = Operator{
	Name:        "hex",
	Description: "Set the last item on the stack to display as hex",
	Operation: func(i *Interpreter) error {
		ss := i.SaveState()
		operand_0, err := i.Pop()
		if err != nil {
			return fmt.Errorf("not enough operands (require 1, got 0): %v", err)
		}
		value_0 := operand_0.Value()
		if value_0 == nil {
			i.RestoreState(ss)
			return fmt.Errorf("first Operand was not a value")
		}

		value_0.Repr = ReprHex

		i.Push(operand_0)

		return nil
	},
}

var OP_ToBin = Operator{
	Name:        "bin",
	Description: "Set the last item on the stack to display as binary",
	Operation: func(i *Interpreter) error {
		ss := i.SaveState()
		operand_0, err := i.Pop()
		if err != nil {
			return fmt.Errorf("not enough operands (require 1, got 0): %v", err)
		}
		value_0 := operand_0.Value()
		if value_0 == nil {
			i.RestoreState(ss)
			return fmt.Errorf("first Operand was not a value")
		}

		value_0.Repr = ReprBinary

		i.Push(operand_0)

		return nil
	},
}

var OP_ToDec = Operator{
	Name:        "dec",
	Description: "Set the last item on the stack to display as decimal",
	Operation: func(i *Interpreter) error {
		ss := i.SaveState()
		operand_0, err := i.Pop()
		if err != nil {
			return fmt.Errorf("not enough operands (require 1, got 0): %v", err)
		}
		value_0 := operand_0.Value()
		if value_0 == nil {
			i.RestoreState(ss)
			return fmt.Errorf("first Operand was not a value")
		}

		value_0.Repr = ReprDecimal

		i.Push(operand_0)

		return nil
	},
}

var OP_ToAuto = Operator{
	Name:        "auto",
	Description: "Set the last item on the stack to display as an automatically formatted value",
	Operation: func(i *Interpreter) error {
		ss := i.SaveState()
		operand_0, err := i.Pop()
		if err != nil {
			return fmt.Errorf("not enough operands (require 1, got 0): %v", err)
		}
		value_0 := operand_0.Value()
		if value_0 == nil {
			i.RestoreState(ss)
			return fmt.Errorf("first Operand was not a value")
		}

		value_0.Repr = ReprAny

		i.Push(operand_0)

		return nil
	},
}
