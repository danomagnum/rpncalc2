package rpncalc

type StackItem interface {
	Value() *Value
	Variable() *Variable
	Operator() *Operator
	SubStack() *Interpreter
	String() string
}

type SubStack struct {
}
