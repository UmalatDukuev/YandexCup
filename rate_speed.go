package main

import (
	"fmt"
	"math/rand"
	"time"
)

func sumIntMap(m map[int]int) int {
	sum := 0
	for _, value := range m {
		sum += value
	}
	return sum
}

func sumFloatMap(m map[int]float64) float64 {
	sum := 0.0
	for _, value := range m {
		sum += value
	}
	return sum
}

func createIntMap(n int) map[int]int {
	m := make(map[int]int, n)
	for i := 0; i < n; i++ {
		m[i] = rand.Intn(100)
	}
	return m
}

func createFloatMap(n int) map[int]float64 {
	m := make(map[int]float64, n)
	for i := 0; i < n; i++ {
		m[i] = rand.Float64() * 100.0
	}
	return m
}

func rate_speed() {
	n := 500000000
	intMap := createIntMap(n)
	start := time.Now()
	sumInt := sumIntMap(intMap)
	duration := time.Since(start)
	fmt.Printf("Суммирование map[int]int с %d элементами заняло %v, результат: %d\n", n, duration, sumInt)

	floatMap := createFloatMap(n)
	start = time.Now()
	sumFloat := sumFloatMap(floatMap)
	duration = time.Since(start)
	fmt.Printf("Суммирование map[int]float64 с %d элементами заняло %v, результат: %f\n", n, duration, sumFloat)
}
