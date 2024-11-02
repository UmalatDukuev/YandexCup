package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	a := make([]int, n, n)
	sc.Scan()
	elements := strings.Split(sc.Text(), " ")
	for i, value := range elements {
		a[i], _ = strconv.Atoi(value)
	}

}
