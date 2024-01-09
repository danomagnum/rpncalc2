package calctests

import (
	"rpncalc/rpncalc"
	"testing"

	"github.com/shopspring/decimal"
)

func TestInterp1(t *testing.T) {
	intp := rpncalc.NewInterpreter()
	intp.Push(rpncalc.NewValue(decimal.NewFromFloat(123.4)))
	intp.Push(rpncalc.NewValue(decimal.NewFromFloat(1.2)))
	err := rpncalc.OP_Add.Operation(intp)
	if err != nil {
		t.Errorf("problem adding: %v", err)
	}

	if len(intp.Stack) != 1 {
		t.Errorf("Expected 1 item on the stack but got %d", len(intp.Stack))
	}

	if !intp.Stack[0].Value().DecimalValue.Equal(decimal.NewFromFloat(124.6)) {
		t.Errorf("expected 124.6 but got %v", intp.Stack[0].Value().DecimalValue.String())

	}

}

func TestParseParts(t *testing.T) {
	intp := rpncalc.NewInterpreter()

	intp.Parse("123.4")
	intp.Parse("1.2")
	intp.Parse("+")

	if len(intp.Stack) != 1 {
		t.Errorf("Expected 1 item on the stack but got %d", len(intp.Stack))
	}

	if !intp.Stack[0].Value().DecimalValue.Equal(decimal.NewFromFloat(124.6)) {
		t.Errorf("expected 124.6 but got %v", intp.Stack[0].Value().DecimalValue.String())

	}

}

func TestParseArithmetic(t *testing.T) {

	tests := []struct {
		name string
		test string
		want float64
	}{
		{"add", "123.4 1.2 +", 124.6},
		{"sub", "123.4 1.2 -", 122.2},
		{"mul", "123.4 1.2 *", 148.08},
		{"div", "123.4 1.2 /", 123.4 / 1.2},
		{"div", "123 2 %", 123 % 2},
	}
	for _, test := range tests {
		intp := rpncalc.NewInterpreter()
		intp.Parse(test.test)
		if len(intp.Stack) != 1 {
			t.Errorf("%s: Expected 1 item on the stack but got %d", test.name, len(intp.Stack))
		}

		// rounding because the want is a floating point value.
		have := intp.Stack[0].Value().DecimalValue.Round(10)
		want := decimal.NewFromFloat(test.want).Round(10)

		if !have.Equal(want) {
			t.Errorf("%s: expected %v but got %v", test.name, want, have)
		}
	}

}

func TestComparisons(t *testing.T) {

	tests := []struct {
		name string
		test string
		want float64
	}{
		{"lt0", "123.4 1.2 <", 1},
		{"lt1", "123.4 123.4 <", 0},
		{"lt2", "123.4 123.5 <", 0},

		{"gt0", "123.4 1.2 >", 0},
		{"gt1", "123.4 123.4 >", 0},
		{"gt2", "123.4 123.5 >", 1},

		{"le0", "123.4 123.5 <=", 0},
		{"le1", "123.4 123.2 <=", 1},
		{"le2", "123.4 123.4 <=", 1},

		{"ge0", "123.4 1.2 >=", 0},
		{"ge1", "123.4 123.5 >=", 1},
		{"ge2", "123.4 123.4 >=", 1},

		{"eq0", "123.4 1.2 ==", 0},
		{"eq1", "123.4 123.4 ==", 1},
	}
	for _, test := range tests {
		intp := rpncalc.NewInterpreter()
		intp.Parse(test.test)
		if len(intp.Stack) != 1 {
			t.Errorf("%s: Expected 1 item on the stack but got %d", test.name, len(intp.Stack))
		}

		// rounding because the want is a floating point value.
		have := intp.Stack[0].Value().DecimalValue.Round(10)
		want := decimal.NewFromFloat(test.want).Round(10)

		if !have.Equal(want) {
			t.Errorf("%s: expected %v but got %v", test.name, want, have)
		}
	}

}

func TestDrop(t *testing.T) {
	intp := rpncalc.NewInterpreter()

	err := intp.Parse("123.4")
	if err != nil {
		t.Errorf("problem parsing first value: %v", err)
	}
	err = intp.Parse("1.2")
	if err != nil {
		t.Errorf("problem parsing second value: %v", err)
	}
	err = intp.Parse("`")
	if err != nil {
		t.Errorf("problem parsing drop cmd: %v", err)
	}

	if len(intp.Stack) != 1 {
		t.Errorf("Expected 1 items on the stack but got %d", len(intp.Stack))
	}

	if !intp.Stack[0].Value().DecimalValue.Equal(decimal.NewFromFloat(123.4)) {
		t.Errorf("expected 123.4 but got %v", intp.Stack[0].Value().DecimalValue.String())
	}

}

func TestSwap(t *testing.T) {
	intp := rpncalc.NewInterpreter()

	err := intp.Parse("123.4")
	if err != nil {
		t.Errorf("problem parsing first value: %v", err)
	}
	err = intp.Parse("1.2")
	if err != nil {
		t.Errorf("problem parsing second value: %v", err)
	}
	err = intp.Parse("swap")
	if err != nil {
		t.Errorf("problem parsing drop cmd: %v", err)
	}

	if len(intp.Stack) != 2 {
		t.Errorf("Expected 2 items on the stack but got %d", len(intp.Stack))
	}

	if !intp.Stack[1].Value().DecimalValue.Equal(decimal.NewFromFloat(123.4)) {
		t.Errorf("expected 123.4 but got %v", intp.Stack[1].Value().DecimalValue.String())
	}
	if !intp.Stack[0].Value().DecimalValue.Equal(decimal.NewFromFloat(1.2)) {
		t.Errorf("expected 1.2 but got %v", intp.Stack[0].Value().DecimalValue.String())
	}
}

func TestMovement(t *testing.T) {

	tests := []struct {
		name string
		test string
		want []float64
	}{
		{"dup", "123.4 1.2 dup", []float64{123.4, 1.2, 1.2}},

		{"drop1", "123.4 1.2 `", []float64{123.4}},
		{"drop2", "123.4 1.2 drop", []float64{123.4}},

		{"dupdrop", "123.4 1.2 dup drop", []float64{123.4, 1.2}},

		{"rot1", "1 2 3 rot", []float64{3, 1, 2}},
		{"rot2", "1 2 3 rot rot", []float64{2, 3, 1}},
		{"rot3", "1 2 3 rot rot rot", []float64{1, 2, 3}},

		{"over", "1 2 3 over", []float64{1, 2, 3, 2}},
		{"tuck", "1 2 3 tuck", []float64{1, 3, 2, 3}},

		{"pick1", "1 2 3 0 pick", []float64{1, 2, 3, 3}},
		{"pick2", "1 2 3 1 pick", []float64{1, 2, 3, 2}},
		{"pick3", "1 2 3 2 pick", []float64{1, 2, 3, 1}},

		{"roll1", "1 2 3 0 roll", []float64{1, 2, 3}},
		{"roll2", "1 2 3 1 roll", []float64{1, 3, 2}},
		{"roll3", "1 2 3 2 roll", []float64{2, 3, 1}},
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

func TestMovementFail(t *testing.T) {

	tests := []struct {
		name  string
		test  string
		test2 string
		want  []float64
	}{
		{"dup", "", "dup", []float64{}},

		{"drop1", "", "`", []float64{}},
		{"drop2", "", "drop", []float64{}},

		{"dupdrop", "", "dup drop", []float64{}},

		{"rot1", "1 2", "rot", []float64{1, 2}},

		{"over", "1", "over", []float64{1}},
		{"tuck", "1", "tuck", []float64{1}},

		{"pick1", "1 2 3 -1", "pick", []float64{1, 2, 3, -1}},
		{"pick1", "1 2 3 -1.0000000000001", "pick", []float64{1, 2, 3, -1}},
		{"pick2", "1 2 3 3", "pick", []float64{1, 2, 3, 3}},

		{"roll1", "1 2 3 -1", "roll", []float64{1, 2, 3, -1}},
		{"roll1", "1 2 3 -1.0000000000001", "roll", []float64{1, 2, 3, -1}},
		{"roll1", "1 2 3 3", "roll", []float64{1, 2, 3, -1}},
	}
	for _, test := range tests {
		intp := rpncalc.NewInterpreter()
		err := intp.Parse(test.test)
		if err != nil {
			t.Errorf("%s: failed to parse initial part: %v", test.name, err)
		}
		err = intp.Parse(test.test2)
		if err == nil {
			t.Errorf("%s: succeeded to parse part that should have failed: %v", test.name, err)
		}

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
