package main

import (
	"fmt"
	algo_tester "github.com/metalrex100/algo-tester"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetFibonacciIterationTask(t *testing.T) {
	err := algo_tester.RunTests(t, getFibonacciIterationTask(), fmt.Sprintf("%s/test-data/4.Fibo", getPathToCurrentDir()))

	assert.NoError(t, err)
}

func TestGetFibonacciRecursionTask(t *testing.T) {
	err := algo_tester.RunTests(t, getFibonacciRecursionTask(), fmt.Sprintf("%s/test-data/4.Fibo", getPathToCurrentDir()))

	assert.NoError(t, err)
}

func TestGetFibonacciRatioTask(t *testing.T) {
	err := algo_tester.RunTests(t, getFibonacciGoldRatioTask(), fmt.Sprintf("%s/test-data/4.Fibo", getPathToCurrentDir()))

	assert.NoError(t, err)
}

func TestGetFibonacciMatrixTask(t *testing.T) {
	err := algo_tester.RunTests(t, getFibonacciMatrixTask(), fmt.Sprintf("%s/test-data/4.Fibo", getPathToCurrentDir()))

	assert.NoError(t, err)
}