package main

import (
	"fmt"
	"github.com/metalrex100/algo-tester"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetPowerIterationTask(t *testing.T) {
	err := algo_tester.RunTests(t, getPowerIterationTask(), fmt.Sprintf("%s/test-data/3.Power", getPathToCurrentDir()))

	assert.NoError(t, err)
}

func TestGetPowerMultiplyTask(t *testing.T) {
	err := algo_tester.RunTests(t, getPowerMultiplyTask(), fmt.Sprintf("%s/test-data/3.Power", getPathToCurrentDir()))

	assert.NoError(t, err)
}

func TestGetPowerDecompositionTask(t *testing.T) {
	err := algo_tester.RunTests(t, getPowerDecompositionTask(), fmt.Sprintf("%s/test-data/3.Power", getPathToCurrentDir()))

	assert.NoError(t, err)
}
