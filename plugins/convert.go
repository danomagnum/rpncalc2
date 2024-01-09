package plugins

import (
	"fmt"
	"rpncalc/rpncalc"

	"github.com/shopspring/decimal"
)

var Conversion_Ops = []rpncalc.Operator{}

func init() {

	no_op := []struct {
		name string
		f    float64
	}{
		{"mm_to_in", 1.0 / 25.4},
		{"m_to_in", 1000.0 / 25.4},
		{"m_to_ft", 1000.0 / 25.4 / 12.0},
		{"mm_to_ft", 1.0 / 25.4 / 12.0},
		{"in_to_mm", 25.4},
		{"in_to_m", 25.4 / 1000},
		{"ft_to_m", 12.0 * 25.4 / 1000},
		{"ft_to_mm", 25.4 / 12.0},
		{"N_to_lb", 0.22480894244319},
		{"lb_to_N", 4.448221628250858},
		{"kN_to_ton", 0.22480894244319 / 2.0},
		{"ton_to_kN", 4.448221628250858 * 2.0},
	}

	for i := range no_op {
		str := no_op[i]
		var newop = rpncalc.Operator{
			Name: fmt.Sprintf("conv.%s", str.name),
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
				result *= str.f
				rval := rpncalc.NewValue(decimal.NewFromFloat(result))
				i.Push(rval)

				return nil
			},
		}

		Conversion_Ops = append(Conversion_Ops, newop)

	}

}
