package rpncalc

import "fmt"

var OP_Assign = Operator{
	Name:        ":=",
	Description: "Assign a variable on the end of the stack to the value of the next to last item on the stack.  Leave the newly assigned variable on the stack.",
	Operation: func(i *Interpreter) error {
		operand_0, err := i.Pop()
		if err != nil {
			return fmt.Errorf("not enough operands (require 2, got 0): %v", err)
		}

		x := operand_0.Variable()
		if x == nil {
			return fmt.Errorf("argument 0 must be a variable")
		}

		operand_1, err := i.Pop()
		if err != nil {
			return fmt.Errorf("not enough operands (require 2, got 1): %v", err)
		}

		y := operand_1.Variable()
		if y != nil {
			// this is a variable, so we will copy the value.
			i.Variables[x.Name] = i.Variables[y.Name]
			i.Push(x)
			return nil
		}

		y2 := operand_1.Value()
		if y2 != nil {
			//x.Contents = y2
			i.Variables[x.Name] = y2
			i.Push(x)
			return nil

		}

		i.Variables[x.Name] = operand_1
		//x.Contents = operand_1
		i.Push(x)
		return nil

	},
}

var OP_DeRef = Operator{
	Name:        "&",
	Description: "Get the value referenced by a variable (dereference)",
	Operation: func(i *Interpreter) error {
		operand_0, err := i.Pop()
		if err != nil {
			return fmt.Errorf("not enough operands (require 2, got 0): %v", err)
		}

		x := operand_0.Variable()
		if x == nil {
			return fmt.Errorf("argument 0 must be a variable")
		}

		y := x.Value()
		if y != nil {
			i.Push(y)
			return nil
		}

		y2 := x.SubStack()
		if y2 != nil {
			i.Push(y2)
			return nil
		}

		y3 := x.Operator()
		if y3 != nil {
			i.Push(y3)
			return nil
		}

		return nil

	},
}
