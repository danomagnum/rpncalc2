package calctests

import (
	"rpncalc/rpncalc"
	"testing"

	"github.com/shopspring/decimal"
)

func TestAssign(t *testing.T) {

	tests := []struct {
		name string
		test string
		want []float64
	}{
		{"assign val", "123.4 1.2 a :=", []float64{123.4, 1.2}},
		{"lookup val", "123.4 1.2 a := a", []float64{123.4, 1.2, 1.2}},
		{"assign var", "123.4 1.2 a := b :=", []float64{123.4, 1.2}},
		{"assign var", "123.4 1.2 a := b := 1 + b := a", []float64{123.4, 2.2, 1.2}},
	}
	for _, test := range tests {
		intp := rpncalc.NewInterpreter()
		intp.Parse(test.test)

		if len(intp.Stack) != len(test.want) {
			t.Errorf("%s: expected %d items on stack but got %d", test.name, len(test.want), len(intp.Stack))
			for i, want := range test.want {
				have := intp.Stack[i]
				if !have.Value().DecimalValue.Equal(decimal.NewFromFloat(want)) {
					t.Errorf("%s: expected position %d to be %v but got %d", test.name, i, len(test.want), len(intp.Stack))
				}

			}
		}

	}

}
