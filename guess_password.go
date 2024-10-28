package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	"golang.org/x/tools/go/analysis/passes/ifaceassert"
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

	a := make([]int, n)
	for j := 0; j < m; j++{
		sum := 0

		for  i := 0; i < m; i++{
			for k := 0; k < 22; k++{
				if sum += x[j]*k{
					fmt.Println(1)
				}
			}
		}
		
	}
}
