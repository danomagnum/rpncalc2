package rpncalc

type Operator struct {
	Operation   func(*Interpreter) error
	Name        string
	Description string
}

func (o *Operator) Value() *Value {
	return nil
}
func (o *Operator) Variable() *Variable {
	return nil
}
func (o *Operator) Operator() *Operator {
	return o
}
func (o *Operator) SubStack() *Interpreter {
	return nil
}
func (o *Operator) String() string {
	return o.Name
}

var BuiltinOperators []Operator

func init() {

	BuiltinOperators = []Operator{
		OP_Add, OP_Sub, OP_Mul, OP_Div, OP_Mod, OP_Pow, // Arithmetic Operators
		OP_GT, OP_LT, OP_GTE, OP_LTE, OP_EQ, OP_If, OP_IfElse, // Comparison Operators
		OP_Drop, OP_Drop2, OP_Swap, OP_Dup, OP_Rot, OP_Over, OP_Tuck, OP_Pick, OP_Roll, // Movement Operators
		OP_Assign, OP_DeRef, // Variable Operators
		OP_ToHex, OP_ToDec, OP_ToBin, OP_ToAuto, // View operators
		OP_NewSubStack, OP_CloseSubStack, OP_RunSubStack, OP_While, // SubStack Operators
	}
}
