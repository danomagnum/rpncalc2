package main

import (
	"bufio"
	"fmt"
	"os"
	"rpncalc/plugins"
	"rpncalc/rpncalc"
)

func main() {

	intp := rpncalc.NewInterpreter()
	intp.AddOperators(plugins.Extended_Math_Ops)
	intp.AddOperators(plugins.Conversion_Ops)

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Printf("Current Stack: \n")
		c := len(intp.Stack)
		for i, v := range intp.Stack {
			fmt.Printf("%2d: %s\n", c-i-1, v.String())
		}
		fmt.Print(">")
		scanner.Scan()
		user_input := scanner.Text()
		err := intp.Parse(user_input)
		if err != nil {
			fmt.Printf("Problem parsing input '%s'\n\n%v\n\n", user_input, err)
		}

	}

}
