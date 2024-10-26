package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func guess_password(sc *bufio.Scanner) {
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	fmt.Println(n)
	sc.Scan()
	m, _ := strconv.Atoi(sc.Text())

	sc.Scan()
	xStr := strings.Fields(sc.Text())
	x := make([]int, m)
	for i := 0; i < m; i++ {
		x[i], _ = strconv.Atoi(xStr[i])
	}

	sc.Scan()
	bStr := strings.Fields(sc.Text())
	b := make([]int, m)
	for i := 0; i < m; i++ {
		b[i], _ = strconv.Atoi(bStr[i])
	}
}
