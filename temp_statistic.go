package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func logic() {
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		input := sc.Text()
		symbols := strings.Fields(input)
		operator := symbols[0]
		switch operator {
		case "+":
			fmt.Printf("Плюсик: %s\n", operator)

		}
	}
}

func main() {
	logic()
}
