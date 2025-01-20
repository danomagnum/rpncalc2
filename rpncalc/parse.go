package rpncalc

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/shopspring/decimal"
)

func (i *Interpreter) Parse(s string) error {
	// replace all commas to allow numbers to be written more clearly?
	ss := i.SaveState()
	s = strings.ReplaceAll(s, ",", "")

	tokens := tokenize(s)

	for _, token := range tokens {
		if len(i.Stack) > 0 {
			item0 := i.Stack[len(i.Stack)-1].SubStack()
			if item0 != nil {
				if item0.Open {
					item0.Parse(token)
					continue
				}
			}
		}
		op, is_operator := i.operators[token]
		if is_operator {
			if i.Open {
				if token == OP_CloseSubStack.Name || token == OP_NewSubStack.Name {
					op.Operation(i)
					continue
				}
				op_ref := &op
				i.Push(op_ref)
			} else {
				err := op.Operation(i)
				if err != nil {
					i.RestoreState(ss)
					return fmt.Errorf("problem with token %s: %w", token, err)
				}
			}
			continue
		}

		if len(token) > 2 {
			if token[0] == '0' {
				// could be a fancy number
				if token[1] == 'x' || token[1] == 'X' {
					parsable := strings.ReplaceAll(token[2:], "_", "")
					val, err := strconv.ParseInt(parsable, 16, 64)
					if err != nil {
						return fmt.Errorf("could not parse '%s' as a hex number: %v", token, err)
					}
					val2 := NewValue(decimal.NewFromInt(val))
					val2.Repr = ReprHex
					i.Push(val2)
					continue
				}
				if token[1] == 'b' || token[1] == 'B' {
					parsable := strings.ReplaceAll(token[2:], "_", "")
					val, err := strconv.ParseInt(parsable, 2, 64)
					if err != nil {
						return fmt.Errorf("could not parse '%s' as a binary number: %v", token, err)
					}
					val2 := NewValue(decimal.NewFromInt(val))
					val2.Repr = ReprBinary
					i.Push(val2)
					continue
				}
			}
		}

		dv, err := decimal.NewFromString(token)
		if err == nil {
			i.Push(NewValue(dv))
			continue
		}

		//varname, is_var := i.Variables[token]
		varname := i.GetVariable(token)
		if varname != nil {
			i.Push(varname)
			continue
		}

		// assume we are defining a variable.
		varname = i.NewVariable(token)
		i.Push(varname)
	}

	return nil
}

func tokenize(s string) []string {
	result := make([]string, 0)
	comment := "#"
	s, _, _ = strings.Cut(s, comment)

	split_chars := " []+*/^%`?"
	ignore_chars := " "

	for len(s) > 0 {
		pos := strings.IndexAny(s, split_chars)
		if pos < 0 {
			result = append(result, s)
			break
		}
		before := s[:pos]
		splitter := string(s[pos])
		if pos < len(s) {
			s = s[pos+1:]
		} else {
			s = ""
		}
		if pos > 0 {
			result = append(result, before)
		}
		if !strings.Contains(ignore_chars, splitter) {
			result = append(result, splitter)
		}

	}

	return result

}
