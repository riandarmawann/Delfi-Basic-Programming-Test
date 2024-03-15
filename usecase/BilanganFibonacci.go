package usecase

import (
	"basic/models"
	"fmt"
)

func Bilanganfibonacci(fibonacciAwal, fibonacciAkhir int) models.Fibonacci {
	var fibonacci models.Fibonacci
	a, b := 0, 1
	for b <= 100 {
		if b >= fibonacciAwal && b <= fibonacciAkhir {
			fmt.Println(b)
		}
		a, b = b, a+b
	}
	fibonacci.FibonacciAwal = fibonacciAwal
	fibonacci.FibonacciAkhir = fibonacciAkhir
	return fibonacci
}
