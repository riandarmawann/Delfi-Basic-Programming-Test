package menu

import (
	"basic/usecase"
	"fmt"
)

func BilanganFibonacciMenu() {
	fibonacciAwal := 1
	fibonacciAkhir := 100
	fmt.Println("Deret bilangan Fibonacci dari", fibonacciAwal, "hingga", fibonacciAkhir)
	usecase.Bilanganfibonacci(fibonacciAwal, fibonacciAkhir)
}
