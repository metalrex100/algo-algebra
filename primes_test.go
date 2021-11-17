package main

import (
	"fmt"
	"github.com/metalrex100/algo-tester"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetPrimesFullSearchTask(t *testing.T) {
	err := algo_tester.RunTests(t, getPrimesFullSearchTask(), fmt.Sprintf("%s/test-data/5.Primes", getPathToCurrentDir()))

	assert.NoError(t, err)
}

func TestGetPrimesSqrtSearchTask(t *testing.T) {
	err := algo_tester.RunTests(t, getPrimesSqrtSearchTask(), fmt.Sprintf("%s/test-data/5.Primes", getPathToCurrentDir()))

	assert.NoError(t, err)
}

func TestGetPrimesSieveOfEratosthenesTaskTask(t *testing.T) {
	err := algo_tester.RunTests(t, getPrimesSieveOfEratosthenesTask(), fmt.Sprintf("%s/test-data/5.Primes", getPathToCurrentDir()))

	assert.NoError(t, err)
}