package plugins

import (
	"fmt"
	"math"
	"rpncalc/rpncalc"

	"github.com/shopspring/decimal"
)

var Extended_Math_Ops = []rpncalc.Operator{}

func init() {

	no_op := []struct {
		name string
		f    float64
	}{
		{"ln10", math.Ln10},
		{"e", math.E},
		{"ln2", math.Ln2},
		{"log10e", math.Log10E},
		{"log2e", math.Log2E},
		{"phi", math.Phi},
		{"pi", math.Pi},
		{"sqrt2", math.Sqrt2},
		{"sqrte", math.SqrtE},
		{"sqrtphi", math.SqrtPhi},
		{"sqrtpi", math.SqrtPi},
	}

	for i := range no_op {
		str := no_op[i]
		var newop = rpncalc.Operator{
			Name: fmt.Sprintf("math.%s", str.name),
			Operation: func(i *rpncalc.Interpreter) error {
				result := str.f
				rval := rpncalc.NewValue(decimal.NewFromFloat(result))
				i.Push(rval)

				return nil
			},
		}

		Extended_Math_Ops = append(Extended_Math_Ops, newop)

	}

	single_op := []struct {
		name string
		f    func(float64) float64
	}{
		{"sin", math.Sin},
		{"cos", math.Cos},
		{"tan", math.Tan},
		{"asin", math.Asin},
		{"acos", math.Acos},
		{"atan", math.Atan},
		{"atanh", math.Atanh},
		{"asinh", math.Asinh},
		{"acosh", math.Acosh},
		{"abs", math.Abs},
		{"cbrt", math.Cbrt},
		{"ceil", math.Ceil},
		{"erf", math.Erf},
		{"erfc", math.Erfc},
		{"erfcinv", math.Erfcinv},
		{"erfinv", math.Erfinv},
		{"exp", math.Exp},
		{"exp2", math.Exp2},
		{"expm1", math.Expm1},
		{"floor", math.Floor},
		{"gamma", math.Gamma},
		{"j0", math.J0},
		{"j1", math.J1},
		{"log", math.Log},
		{"log10", math.Log10},
		{"log1p", math.Log1p},
		{"log2", math.Log2},
		{"logb", math.Logb},
		{"round", math.Round},
		{"roundtoeven", math.RoundToEven},
		{"sqrt", math.Sqrt},
		{"trunc", math.Trunc},
		{"y0", math.Y0},
		{"y1", math.Y1},
	}

	for i := range single_op {
		str := single_op[i]
		var newop = rpncalc.Operator{
			Name: fmt.Sprintf("math.%s", str.name),
			Operation: func(i *rpncalc.Interpreter) error {
				operand_0, err := i.Pop()
				if err != nil {
					return fmt.Errorf("not enough operands (require 1, got 0): %v", err)
				}
				value_0 := operand_0.Value()
				if value_0 == nil {
					return fmt.Errorf("first Operand was not a value")
				}

				result, _ := value_0.DecimalValue.Float64()
				result = str.f(result)
				rval := rpncalc.NewValue(decimal.NewFromFloat(result))
				i.Push(rval)

				return nil
			},
		}

		Extended_Math_Ops = append(Extended_Math_Ops, newop)

	}

	dual_op := []struct {
		name string
		f    func(float64, float64) float64
	}{
		{"atan2", math.Atan2},
		{"copysign", math.Copysign},
		{"dim", math.Dim},
		{"hypot", math.Hypot},
		{"max", math.Max},
		{"min", math.Min},
		{"mod", math.Mod},
		{"nextafter", math.Nextafter},
		{"pow", math.Pow},
		{"remainder", math.Remainder},
	}

	for i := range dual_op {
		str := dual_op[i]
		var newop = rpncalc.Operator{
			Name: fmt.Sprintf("math.%s", str.name),
			Operation: func(i *rpncalc.Interpreter) error {
				operand_0, err := i.Pop()
				if err != nil {
					return fmt.Errorf("not enough operands (require 2, got 0): %v", err)
				}
				value_0 := operand_0.Value()
				if value_0 == nil {
					return fmt.Errorf("first Operand was not a value")
				}

				operand_1, err := i.Pop()
				if err != nil {
					return fmt.Errorf("not enough operands (require 2, got 1): %v", err)
				}
				value_1 := operand_1.Value()
				if value_1 == nil {
					return fmt.Errorf("first Operand was not a value")
				}

				v0, _ := value_0.DecimalValue.Float64()
				v1, _ := value_1.DecimalValue.Float64()
				result := str.f(v0, v1)
				rval := rpncalc.NewValue(decimal.NewFromFloat(result))
				i.Push(rval)

				return nil
			},
		}

		Extended_Math_Ops = append(Extended_Math_Ops, newop)

	}

}
