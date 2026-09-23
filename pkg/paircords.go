package pkg

// PairCords は頂点ペア間の距離を求める
func PairCords(pair [][]float64) (dist float64) {
	x1 := pair[0][1]
	y1 := pair[0][0]
	x2 := pair[1][1]
	y2 := pair[1][0]
	dist = DistVerts(x1, y1, x2, y2)
	return dist
}
