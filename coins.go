package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func coins() {
	file, _ := os.Open("coins/coins.in")
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	fmt.Println(n)

}
