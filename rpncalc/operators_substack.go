package rpncalc

import (
	"fmt"

	"github.com/shopspring/decimal"
)

var OP_NewSubStack = Operator{
	Name:        "[",
	Description: "Start a new paused sub-interpreter stack",
	Operation: func(i *Interpreter) error {
		newi := newInterpreter(true, i)
		i.Push(newi)
		return nil
	},
}

var OP_CloseSubStack = Operator{
	Name:        "]",
	Description: "Start a new paused sub-interpreter stack",
	Operation: func(i *Interpreter) error {
		if !i.Open {
			return fmt.Errorf("substack is not open")
		}
		i.Open = false

		return nil
	},
}

var OP_RunSubStack = Operator{
	Name:        "!",
	Description: "Execute a substack",
	Operation: func(i *Interpreter) error {
		operand_0, err := i.Pop()
		if err != nil {
			return fmt.Errorf("not enough operands (require 1, got 0): %v", err)
		}

		ss := operand_0.SubStack()
		if ss == nil {
			return fmt.Errorf("operand 0 is not a substack")
		}

		return i.ExecutSubStack(ss)
	},
}

var OP_While = Operator{
	Name: "while",
	Description: "Pops 2 substacks.  The first one is the condition and the second is the function to execute." +
		"  First the condition is executed.  The last item left on the stack is evaluated for truthiness.  If it is truthy," +
		" The second item is executed.  Then the first condition is executed and the result checked again in a loop until the" +
		" value is falsey at which point execution stops.",
	Operation: func(i *Interpreter) error {
		operand_0, err := i.Pop()
		if err != nil {
			return fmt.Errorf("not enough operands (require 2, got 0): %v", err)
		}

		ss_func := operand_0.SubStack()
		if ss_func == nil {
			return fmt.Errorf("operand 0 is not a substack")
		}

		iters := 0
		for {
			err := i.ExecutSubStack(ss_func)
			if err != nil {
				return fmt.Errorf("problem running conditional on loop %d: %w", iters, err)
			}
			result, err := i.Pop()
			if err != nil {
				return fmt.Errorf("problem getting result of condition on loop %d: %v", iters, err)
			}

			resval := result.Value()
			if resval == nil {
				return fmt.Errorf("bad result value on loop %d: %v", iters, resval)
			}

			if resval.DecimalValue.Equal(decimal.NewFromInt(0)) {
				break
			}

			iters++
		}

		return nil
	},
}
