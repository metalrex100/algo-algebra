package main

import (
	"github.com/metalrex100/algo-tester"
	"log"
	"math"
	"strconv"
)

func getFibonacciIterationTask() algo_tester.Task {
	return func(data []string) string {
		number, err := strconv.ParseFloat(data[0], 64)
		if err != nil {
			log.Fatal(err)
		}

		var prev, last, cur float64 = 0, 1, 0
		var i float64
		for i = 2; i <= number; i++ {
			cur = last + prev
			prev, last = last, cur
		}

		if number == 1 {
			cur = 1
		}

		return strconv.FormatFloat(cur, 'f', -1, 64)
	}
}

func getFibonacciRecursionTask() algo_tester.Task {
	return func(data []string) string {
		number, err := strconv.Atoi(data[0])
		if err != nil {
			log.Fatal(err)
		}

		res := getFibonacciRec(number, nil)

		return strconv.FormatFloat(res, 'f', -1, 64)
	}
}

func getFibonacciGoldRatioTask() algo_tester.Task {
	return func(data []string) string {
		const phi = 1.618
		number, err := strconv.ParseUint(data[0], 10, 64)
		if err != nil {
			log.Fatal(err)
		}

		res := number
		if number > 1 {
			res = uint64((math.Pow(phi, float64(number)))/math.Sqrt(5) + 0.5)
		}

		return strconv.FormatUint(res, 10)
	}
}

func getFibonacciMatrixTask() algo_tester.Task {
	return func(data []string) string {
		number, err := strconv.ParseUint(data[0], 10, 64)
		if err != nil {
			log.Fatal(err)
		}

		if number < 2 {
			return strconv.FormatUint(number, 10)
		}

		transitionMatrix := NewSquareMatrix(1, 1, 1, 0)
		transitionMatrix = transitionMatrix.Power(number-1)

		return strconv.FormatUint(transitionMatrix.A11, 10)
	}
}

func getFibonacciRec(n int, fibo []float64) float64 {
	if n < 2 {
		return float64(n)
	}

	if fibo == nil {
		fibo = make([]float64, 0, n+1)
		fibo = append(fibo, 0, 1)
	}
	if len(fibo) == n+1 {
		return fibo[n]
	}

	fibo = append(fibo, fibo[len(fibo)-1] + fibo[len(fibo)-2])

	return getFibonacciRec(n, fibo)
}