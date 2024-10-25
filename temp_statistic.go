package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func logic() {
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		input := sc.Text()
		args := strings.Fields(input)
		switch args[0] {
		case "+":
			operator := args[0]
			id, _ := strconv.Atoi(args[1])
			temperature, _ := strconv.ParseFloat(args[2], 64)
			fmt.Printf("Плюсик: %s, айди: %d, температура: %f\n", operator, id, temperature)

		case "-":
			operator := args[0]
			id, _ := strconv.Atoi(args[1])
			fmt.Printf("Минусик: %s, айди: %d\n", operator, id)

		case "~":
			operator := args[0]
			id, _ := strconv.Atoi(args[1])
			fmt.Printf("Тильда: %s, айди: %d\n", operator, id)

		case "?":
			fmt.Printf("Вопросик: %s\n", args[0])

		case "!":
			fmt.Printf("Восклицательный значок: %s\n", args[0])

		}
	}
}

func main() {
	logic()
}
