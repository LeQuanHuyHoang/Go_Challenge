package main

import (
	"fmt"
	m "math"
)

func main() {
	A := []int{100, 50, 100}
	B := []int{98, 150, 100}
	Solution(A, B, 100, 100)
}

func Solution(A []int, B []int, X int, Y int) {
	for i := 0; i < len(A); i++ {
		r1 := float64(A[i]) - float64(X)
		r2 := float64(B[i]) - float64(Y)
		rs := m.Sqrt((r1 * r1) + (r2 * r2))
		if rs <= 0 {
			fmt.Printf("The tapped point (%v,%v) is also the center of icon (%v,%v)\n", X, Y, A[i], B[i])
		} else if 0 < rs && rs <= 20 {
			fmt.Printf("The tapped point (%v,%v) is in range of icon (%v,%v)\n", X, Y, A[i], B[i])
		} else if rs > 20 {
			fmt.Printf("The tapped point (%v,%v) is out range of icon (%v,%v)\n", X, Y, A[i], B[i])
		}
	}
}
