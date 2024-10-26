package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func find_avg_temp(sc *bufio.Scanner) {
	temperatures := make(map[int]int)
	var sumTemp int
	count := 0
	for sc.Scan() {
		input := sc.Text()
		args := strings.Fields(input)
		switch args[0] {
		case "+":
			id, _ := strconv.Atoi(args[1])
			temperature, _ := strconv.Atoi(strings.ReplaceAll(args[2], ".", ""))
			temperatures[id] = temperature
			count++
			sumTemp += temperature

		case "-":
			id, _ := strconv.Atoi(args[1])
			sumTemp -= temperatures[id]
			count--
			delete(temperatures, id)

		case "~":
			id, _ := strconv.Atoi(args[1])
			temperature, _ := strconv.Atoi(strings.ReplaceAll(args[2], ".", ""))

			sumTemp = sumTemp - temperatures[id] + temperature
			temperatures[id] = temperature

		case "?":
			fmt.Printf("%.9f \n", float64(sumTemp)/float64(count)/10.0)

		case "!":
			return
		}
	}
}
