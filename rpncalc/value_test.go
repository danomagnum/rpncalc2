package rpncalc

import "testing"

func TestFormatNumbers(t *testing.T) {

	tests := []struct {
		name  string
		input string
		qty   int
		want  string
	}{
		{"hex", "F", 4, "000F"},
		{"hex", "EF", 4, "00EF"},
		{"hex", "EEF", 4, "0EEF"},
		{"hex", "BEEF", 4, "BEEF"},
		{"hex", "DBEEF", 4, "000D_BEEF"},
		{"hex", "ADBEEF", 4, "00AD_BEEF"},
		{"hex", "EADBEEF", 4, "0EAD_BEEF"},
		{"hex", "DEADBEEF", 4, "DEAD_BEEF"},
		{"dec", "1", 3, "001"},
		{"dec", "21", 3, "021"},
		{"dec", "321", 3, "321"},
		{"dec", "4321", 3, "004_321"},
		{"dec", "54321", 3, "054_321"},
		{"dec", "654321", 3, "654_321"},
	}

	for _, test := range tests {
		result := format_number_with_spaces(test.input, test.qty)
		if result != test.want {
			t.Errorf("%s: wanted %s got %s", test.name, test.want, result)
		}
	}

}
