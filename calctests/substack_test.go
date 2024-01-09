package calctests

import (
	"rpncalc/rpncalc"
	"testing"

	"github.com/shopspring/decimal"
)

func TestSubStack(t *testing.T) {

	tests := []struct {
		name string
		test string
		want []float64
	}{
		{"simple substack1", "1 2 [ 3 ] !", []float64{1, 2, 3}},
		{"simple operating substack", "1 2 [ 3 4 + ] !", []float64{1, 2, 7}},
		{"simple operating substack", "1 2 [ 3 4 + ] a := ! a !", []float64{1, 2, 7, 7}},
	}
	for _, test := range tests {
		intp := rpncalc.NewInterpreter()
		intp.Parse(test.test)

		if len(intp.Stack) != len(test.want) {
			t.Errorf("%s: expected %d items on stack but got %d", test.name, len(test.want), len(intp.Stack))
		}
		for i, want := range test.want {
			have := intp.Stack[i]
			if !have.Value().DecimalValue.Equal(decimal.NewFromFloat(want)) {
				t.Errorf("%s: expected position %d to be %v but got %s", test.name, i, test.want[i], intp.Stack[i].String())
			}

		}

	}
}

func TestIfConditional(t *testing.T) {

	tests := []struct {
		name string
		test string
		want []float64
	}{
		{"if true", "1 2 [ 3 ] 4 if", []float64{1, 2, 3}},
		{"if false", "1 2 [ 3 ] 0 if", []float64{1, 2}},
		{"if true 2", "1 2 [ 3 4 + ] 1 if", []float64{1, 2, 7}},
		{"if false 2", "1 2 [ 3 4 + ] 0 if", []float64{1, 2}},
		{"if true 3", "1 2 [ 3 ] 1 0 < if", []float64{1, 2, 3}},
		{"if false 3", "1 2 [ 3 ] 1 0 > if", []float64{1, 2}},
	}
	for _, test := range tests {
		intp := rpncalc.NewInterpreter()
		intp.Parse(test.test)

		if len(intp.Stack) != len(test.want) {
			t.Errorf("%s: expected %d items on stack but got %d", test.name, len(test.want), len(intp.Stack))
		}
		for i, want := range test.want {
			have := intp.Stack[i]
			if !have.Value().DecimalValue.Equal(decimal.NewFromFloat(want)) {
				t.Errorf("%s: expected position %d to be %v but got %s", test.name, i, test.want[i], intp.Stack[i].String())
			}

		}

	}
}

func TestIfElseConditional(t *testing.T) {

	tests := []struct {
		name string
		test string
		want []float64
	}{
		{"if true", "1 2 [5] [ 3 ] 4 ifelse", []float64{1, 2, 3}},
		{"if false", "1 2 [5] [ 3 ] 0 ifelse", []float64{1, 2, 5}},
		{"if true 2", "1 2 [5] [ 3 4 + ] 1 ifelse", []float64{1, 2, 7}},
		{"if false 2", "1 2 [5] [ 3 4 + ] 0 ifelse", []float64{1, 2, 5}},
		{"if true 3", "1 2 [5] [ 3 ] 1 0 < ifelse", []float64{1, 2, 3}},
		{"if false 3", "1 2 [5] [ 3 ] 1 0 > ifelse", []float64{1, 2, 5}},
	}
	for _, test := range tests {
		intp := rpncalc.NewInterpreter()
		intp.Parse(test.test)

		if len(intp.Stack) != len(test.want) {
			t.Errorf("%s: expected %d items on stack but got %d", test.name, len(test.want), len(intp.Stack))
		}
		for i, want := range test.want {
			have := intp.Stack[i]
			if !have.Value().DecimalValue.Equal(decimal.NewFromFloat(want)) {
				t.Errorf("%s: expected position %d to be %v but got %s", test.name, i, test.want[i], intp.Stack[i].String())
			}

		}

	}
}

func TestWhile(t *testing.T) {

	tests := []struct {
		name string
		test string
		want []float64
	}{
		{"while countdown with var", "5 a := drop  [a &  a 1 - a := ] while", []float64{5, 4, 3, 2, 1}},
		{"while countup with var", "0 a := drop  [a & a 1 + a := 5 > ] while", []float64{0, 1, 2, 3, 4}},

		{"while countdown novar", "5 [ dup 1 - dup ] while", []float64{5, 4, 3, 2, 1, 0}},
		{"while countup novar", "0 [dup  1 + dup 5 >] while", []float64{0, 1, 2, 3, 4, 5}},
	}
	for _, test := range tests {
		intp := rpncalc.NewInterpreter()
		intp.Parse(test.test)

		if len(intp.Stack) != len(test.want) {
			t.Errorf("%s: expected %d items on stack but got %d", test.name, len(test.want), len(intp.Stack))
			continue
		}
		for i, want := range test.want {
			have := intp.Stack[i]
			if !have.Value().DecimalValue.Equal(decimal.NewFromFloat(want)) {
				t.Errorf("%s: expected position %d to be %v but got %s", test.name, i, test.want[i], intp.Stack[i].String())
			}

		}

	}
}

func TestNestedSubStack(t *testing.T) {

	tests := []struct {
		name string
		test string
		want []float64
	}{
		{"simple nested", "1 2 [ 3 [4] ] ! ! ", []float64{1, 2, 3, 4}},
	}
	for _, test := range tests {
		intp := rpncalc.NewInterpreter()
		intp.Parse(test.test)

		if len(intp.Stack) != len(test.want) {
			t.Errorf("%s: expected %d items on stack but got %d", test.name, len(test.want), len(intp.Stack))
		}
		for i, want := range test.want {
			have := intp.Stack[i]
			if !have.Value().DecimalValue.Equal(decimal.NewFromFloat(want)) {
				t.Errorf("%s: expected position %d to be %v but got %d", test.name, i, len(test.want), len(intp.Stack))
			}

		}

	}
}
