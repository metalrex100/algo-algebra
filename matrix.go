package main

func NewSquareMatrix(a11, a12, a21, a22 uint64) SquareMatrix {
	return SquareMatrix{A11: a11, A12: a12, A21: a21, A22: a22}
}

type SquareMatrix struct {
	A11, A12 uint64
	A21, A22 uint64
}

func (m SquareMatrix) Power(pow uint64) SquareMatrix {
	if pow < 2 {
		return m
	}

	matrixesToMultiply := make([]SquareMatrix, 0)

	powMatrix := m
	for divisionResult := pow; divisionResult >= 1; divisionResult /= 2 {
		if divisionResult%2 == 1 {
			matrixesToMultiply = append(matrixesToMultiply, powMatrix)
		}
		powMatrix = powMatrix.SelfMultiply()
	}

	resultMatrix := matrixesToMultiply[0]
	for i := 1; i < len(matrixesToMultiply); i++ {
		resultMatrix = resultMatrix.MultiplyByMatrix(matrixesToMultiply[i])
	}

	return resultMatrix
}

func (m SquareMatrix) PowerMultiply(pow uint64) SquareMatrix {
	const powCoef = 2

	copyM := m

	var currentPow uint64 = 1
	for ; currentPow*powCoef <= pow; currentPow *= powCoef {
		copyM = copyM.SelfMultiply()
	}

	for ; currentPow < pow; currentPow++ {
		copyM = copyM.MultiplyByMatrix(m)
	}

	return copyM
}

func (m SquareMatrix) MultiplyByMatrix(sm SquareMatrix) SquareMatrix {
	return SquareMatrix{
		A11: m.A11*sm.A11 + m.A12*sm.A21,
		A12: m.A11*sm.A12 + m.A12*sm.A22,
		A21: m.A21*sm.A11 + m.A22*sm.A21,
		A22: m.A21*sm.A12 + m.A22*sm.A22,
	}
}

func (m SquareMatrix) SelfMultiply() SquareMatrix {
	return m.MultiplyByMatrix(m)
}
