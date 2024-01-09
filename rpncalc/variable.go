package rpncalc

import "fmt"

type Variable struct {
	Name string
	i    *Interpreter
}

func (v *Variable) Value() *Value {
	return v.i.GetVariableValue(v.Name).Value()
}
func (v *Variable) Variable() *Variable {
	return v
}
func (v *Variable) Operator() *Operator {
	return v.i.GetVariableValue(v.Name).Operator()
}
func (v *Variable) SubStack() *Interpreter {
	return v.i.GetVariableValue(v.Name).SubStack()
}
func (v *Variable) String() string {
	varval := v.i.GetVariableValue(v.Name)
	if varval == nil {
		return fmt.Sprintf("%s :=", v.Name)
	}
	if varval != nil {
		return fmt.Sprintf("%s %s :=", varval.String(), v.Name)
	}
	return fmt.Sprintf("Myster var: %s", v.Name)
}
