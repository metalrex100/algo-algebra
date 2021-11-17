package main

import (
	"github.com/metalrex100/algo-tester"
	"log"
	"math"
	"strconv"
)

func getPrimesFullSearchTask() algo_tester.Task {
	return func(data []string) string {
		number, err := strconv.ParseUint(data[0], 10, 64)
		if err != nil {
			log.Fatal(err)
		}

		var i, j, primesCount uint64

		for i = 2; i <= number; i++ {
			isPrime := true
			for j = 2; j < i; j++ {
				if i%j == 0 {
					isPrime = false
					break
				}
			}
			if isPrime {
				primesCount++
			}
		}

		return strconv.FormatUint(primesCount, 10)
	}
}

func getPrimesSqrtSearchTask() algo_tester.Task {
	return func(data []string) string {
		number, err := strconv.ParseUint(data[0], 10, 64)
		if err != nil {
			log.Fatal(err)
		}

		primes := make([]uint64, 0, number/2)
		if number > 1 {
			primes = append(primes, 2)
		}

		var i uint64 = 3
		for ; i <= number; i += 2 {
			if isPrimeSqrt(i, primes) {
				primes = append(primes, i)
			}
		}

		return strconv.Itoa(len(primes))
	}
}

func getPrimesSieveOfEratosthenesTask() algo_tester.Task {
	return func(data []string) string {
		number, err := strconv.ParseUint(data[0], 10, 64)
		if err != nil {
			log.Fatal(err)
		}

		sqrtNumber := uint64(math.Sqrt(float64(number)))

		primesCount := 0

		numbers := make([]bool, number+1)

		var i uint64 = 2
		for i = 2; i <= number; i++ {
			if numbers[i] == false {
				primesCount++
				if i > sqrtNumber {
					continue
				}
				for j := i * i; j <= number; j += i {
					numbers[j] = true
				}
			}
		}

		return strconv.Itoa(primesCount)
	}
}

func isPrimeSqrt(n uint64, foundPrimes []uint64) bool {
	if n == 2 {
		return true
	}
	if n%2 == 0 {
		return false
	}

	var i uint64
	sqrtN := uint64(math.Sqrt(float64(n)))
	for ; foundPrimes[i] <= sqrtN; i++ {
		if n%foundPrimes[i] == 0 {
			return false
		}
	}

	return true
}
