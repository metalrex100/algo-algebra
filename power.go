package main

import (
	"github.com/metalrex100/algo-tester"
	"log"
	"strconv"
)

func getPowerIterationTask() algo_tester.Task {
	return func(data []string) string {
		number, err := strconv.ParseFloat(data[0], 64)
		if err != nil {
			log.Fatal(err)
		}
		pow, err := strconv.ParseUint(data[1], 10, 64)
		if err != nil {
			log.Fatal(err)
		}

		var result float64 = 1

		var i uint64
		for i = 1; i <= pow; i++ {
			result *= number
		}

		return strconv.FormatFloat(result, 'f', 11, 64)
	}
}

func getPowerMultiplyTask() algo_tester.Task {
	return func(data []string) string {
		number, err := strconv.ParseFloat(data[0], 64)
		if err != nil {
			log.Fatal(err)
		}
		pow, err := strconv.ParseUint(data[1], 10, 64)
		if err != nil {
			log.Fatal(err)
		}

		result := powerMultiply(number, pow)

		return strconv.FormatFloat(result, 'f', 11, 64)
	}
}

func getPowerDecompositionTask() algo_tester.Task {
	return func(data []string) string {
		number, err := strconv.ParseFloat(data[0], 64)
		if err != nil {
			log.Fatal(err)
		}
		pow, err := strconv.ParseUint(data[1], 10, 64)
		if err != nil {
			log.Fatal(err)
		}

		if pow == 0 {
			return strconv.FormatFloat(1.0, 'f', 11, 64)
		}

		var result float64 = 1
		powNumber := number
		for divisionResult := pow; divisionResult >= 1; divisionResult /= 2 {
			if divisionResult%2 == 1 {
				result *= powNumber
			}
			powNumber *= powNumber
		}

		return strconv.FormatFloat(result, 'f', 11, 64)
	}
}

func powerMultiply(number float64, pow uint64) float64 {
	const powCoef = 2

	var result float64 = 1
	if pow > 0 {
		result = number
	}

	var currentPow uint64 = 1
	for ; currentPow*powCoef <= pow; currentPow *= powCoef {
		result *= result
	}

	for ; currentPow < pow; currentPow++ {
		result *= number
	}

	return result
}
