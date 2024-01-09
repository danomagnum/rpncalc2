package rpncalc

import (
	"fmt"

	"github.com/shopspring/decimal"
)

var OP_Drop = Operator{
	Name:        "`",
	Description: "Discard the last item on the stack (shorthand)",
	Operation: func(i *Interpreter) error {
		_, err := i.Pop()
		if err != nil {
			return fmt.Errorf("not enough operands (require 1, got 0): %v", err)
		}
		return nil
	},
}
var OP_Drop2 = Operator{
	Name:        "drop",
	Description: "Discard the last item on the stack",
	Operation: func(i *Interpreter) error {
		_, err := i.Pop()
		if err != nil {
			return fmt.Errorf("not enough operands (require 1, got 0): %v", err)
		}
		return nil
	},
}

var OP_Swap = Operator{
	Name:        "swap",
	Description: "Swap positions of the last item and the second to last item on the stack",
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
		i.Push(operand_0)
		i.Push(operand_1)
		return nil
	},
}

var OP_Dup = Operator{
	Name:        "dup",
	Description: "Duplicate last item on the stack",
	Operation: func(i *Interpreter) error {
		operand_0, err := i.Pop()
		if err != nil {
			return fmt.Errorf("not enough operands (require 2, got 0): %v", err)
		}

		i.Push(operand_0)
		i.Push(operand_0)
		return nil
	},
}

var OP_Rot = Operator{
	Name:        "rot",
	Description: "Rotate the last 3 items on the stack. 3 -> 2, 2 -> 1, 1 -> 3",
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
			return fmt.Errorf("not enough operands (require 3, got 2): %v", err)
		}

		i.Push(operand_1)
		i.Push(operand_0)
		i.Push(operand_2)
		return nil
	},
}

var OP_Over = Operator{
	Name:        "over",
	Description: "Duplicate the next to last item in the stack and push it back on the end",
	Operation: func(i *Interpreter) error {
		operand_0, err := i.Pop()
		if err != nil {
			return fmt.Errorf("not enough operands (require 2, got 0): %v", err)
		}
		operand_1, err := i.Pop()
		if err != nil {
			return fmt.Errorf("not enough operands (require 2, got 1): %v", err)
		}

		i.Push(operand_1)
		i.Push(operand_0)
		i.Push(operand_1)
		return nil
	},
}

var OP_Tuck = Operator{
	Name:        "tuck",
	Description: "Duplicate the last item in the stack and insert it above the next to last item.  2,1,0 -> 2,0,1,0",
	Operation: func(i *Interpreter) error {
		operand_0, err := i.Pop()
		if err != nil {
			return fmt.Errorf("not enough operands (require 2, got 0): %v", err)
		}
		operand_1, err := i.Pop()
		if err != nil {
			return fmt.Errorf("not enough operands (require 2, got 1): %v", err)
		}

		i.Push(operand_0)
		i.Push(operand_1)
		i.Push(operand_0)
		return nil
	},
}

var OP_Pick = Operator{
	Name:        "pick",
	Description: "Pop the last item from the stack.  Interpret it as a number, and use it as a pointer up the stack.  Duplicate that item and push it onto the end of the stack",
	Operation: func(i *Interpreter) error {
		operand_0, err := i.Pop()
		if err != nil {
			return fmt.Errorf("not enough operands (require 2, got 0): %v", err)
		}

		x := operand_0.Value()
		if x == nil {
			return fmt.Errorf("argument 0 must be a number")
		}

		if !x.DecimalValue.IsInteger() {
			return fmt.Errorf("argument 0 must be a whole number. got %s", x.String())
		}

		index := int(x.DecimalValue.IntPart())

		if index >= len(i.Stack) || index < 0 {
			return fmt.Errorf("referenced index out of bound [0,%d): %d", len(i.Stack), index)
		}

		relative_index := len(i.Stack) - index - 1

		i.Push(i.Stack[relative_index])

		return nil
	},
}

var OP_Roll = Operator{
	Name: "roll",
	Description: "Pop the last item from the stack.  Interpret it as a number, and use it as a pointer up the stack." +
		" Move that item to the end of the stack",
	Operation: func(i *Interpreter) error {
		operand_0, err := i.Pop()
		if err != nil {
			return fmt.Errorf("not enough operands (require 2, got 0): %v", err)
		}

		x := operand_0.Value()
		if x == nil {
			return fmt.Errorf("argument 0 must be a number")
		}

		if !x.DecimalValue.IsInteger() {
			return fmt.Errorf("argument 0 must be a whole number. got %s", x.String())
		}

		index := int(x.DecimalValue.IntPart())

		if index >= len(i.Stack) || index < 0 {
			return fmt.Errorf("referenced index out of bound [0,%d): %d", len(i.Stack), index)
		}

		relative_index := len(i.Stack) - index - 1
		item := i.Stack[relative_index]
		pre := i.Stack[:relative_index]
		post := i.Stack[relative_index+1:]
		i.Stack = append(pre, post...)
		//slices.Delete(i.Stack, relative_index-1, relative_index)
		i.Push(item)

		return nil
	},
}

var OP_Size = Operator{
	Name:        "size",
	Description: "adds the total current count of items on the stack to the stack",
	Operation: func(i *Interpreter) error {
		item := NewValue(decimal.NewFromInt(int64(len(i.Stack))))
		i.Push(item)

		return nil
	},
}
