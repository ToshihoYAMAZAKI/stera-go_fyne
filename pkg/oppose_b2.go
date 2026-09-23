package pkg

// OpposeB2 は直交する辺bを分割線として対向する辺の頂点ペアを求める
func OpposeB2(n int, XY [][]float64, node int) (pairB [][]float64) {
	// pairB := make([][]float64, 2)
	pairB = append(pairB, XY[(n+3+node)%node])
	pairB = append(pairB, XY[(n+4+node)%node])

	return pairB
}
