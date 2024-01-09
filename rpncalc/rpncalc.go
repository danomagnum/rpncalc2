package rpncalc

import (
	"fmt"
	"slices"

	"github.com/shopspring/decimal"
)

type Interpreter struct {
	Stack     []StackItem
	operators map[string]Operator
	Variables map[string]StackItem

	Open   bool
	Parent *Interpreter
}

func (i *Interpreter) GetVariable(name string) StackItem {
	local := i.Variables[name]
	if local == nil {
		if i.Parent != nil {
			local = i.Parent.GetVariable(name)
		}
	}
	if local == nil {
		return local
	}
	return &Variable{Name: name, i: i}
}
func (i *Interpreter) GetVariableValue(name string) StackItem {
	local := i.Variables[name]
	if local == nil {
		if i.Parent != nil {
			local = i.Parent.GetVariableValue(name)
		}
	}
	return local
}

func (i *Interpreter) Value() *Value {
	return nil
}
func (i *Interpreter) Variable() *Variable {
	return nil
}
func (i *Interpreter) Operator() *Operator {
	return nil
}
func (i *Interpreter) SubStack() *Interpreter {
	return i
}
func (i *Interpreter) String() string {
	return fmt.Sprintf("%v", i.Stack)
}

func NewInterpreter() *Interpreter {
	return newInterpreter(false, nil)
}

func newInterpreter(open bool, parent *Interpreter) *Interpreter {
	intp := &Interpreter{
		Stack:     make([]StackItem, 0),
		operators: make(map[string]Operator),
		Variables: make(map[string]StackItem),
		Open:      open,
		Parent:    parent,
	}
	if parent == nil {
		intp.AddOperators(BuiltinOperators)
	} else {
		intp.AddOperators(parent.Operators())
	}
	return intp
}

func (i *Interpreter) Operators() []Operator {
	o := make([]Operator, 0, len(i.operators))
	for k := range i.operators {
		o = append(o, i.operators[k])
	}
	return o
}

func (i *Interpreter) AddOperators(o []Operator) {
	for _, v := range o {
		i.AddOperator(v)
	}
}
func (i *Interpreter) AddOperator(o Operator) {
	i.operators[o.Name] = o
}

func (i *Interpreter) NewVariable(name string) *Variable {
	variable := &Variable{Name: name, i: i}
	i.Variables[name] = NewValue(decimal.NewFromInt(0))
	return variable
}

// Return a new snapshot of the interpreter state
func (i *Interpreter) SaveState() *Interpreter {
	//TODO: make sure we deep copy any slices or maps or whatnot
	newInt := &Interpreter{}
	newInt.Stack = slices.Clone(i.Stack)
	return newInt
}

func (i *Interpreter) RestoreState(newint *Interpreter) {
	//TODO: make sure we restore any slices or maps or whatnot.
	i.Stack = newint.Stack
}

type ERR_EmptyStack struct{}

func (e ERR_EmptyStack) Error() string {
	return "Stack empty"
}

// get the newest item from the stack
func (i *Interpreter) Pop() (StackItem, error) {
	if len(i.Stack) < 1 {
		return nil, ERR_EmptyStack{}
	}
	result := i.Stack[len(i.Stack)-1]
	i.Stack = i.Stack[:len(i.Stack)-1]
	return result, nil
}

// get the newest item from the stack
func (i *Interpreter) Push(si StackItem) {
	i.Stack = append(i.Stack, si)
}

func (i *Interpreter) ExecutSubStack(ss *Interpreter) error {
	for idx := range ss.Stack {
		item := ss.Stack[idx]
		op := item.Operator()
		if op == nil {
			i.Push(item)
			continue
		}
		err := op.Operation(i)
		if err != nil {
			return fmt.Errorf("problem with operation %s@%d: %v", op.Name, idx, err)
		}
	}
	return nil
}
