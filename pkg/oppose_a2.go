package pkg

// OpposeA2 は直交する辺aを分割線として対向する辺の頂点ペアを求める
func OpposeA2(n int, XY [][]float64, node int) (pairA [][]float64) {
	// pairA := make([][]float64, 2)
	pairA = append(pairA, XY[(n-3+node)%node])
	pairA = append(pairA, XY[(n-4+node)%node])

	return pairA
}
