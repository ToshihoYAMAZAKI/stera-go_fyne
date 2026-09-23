package pkg

// PentaDeg は５角形の内角の角度を計算する
func PentaDeg(cord [][]float64) (deg5 []float64) {
	for i := 0; i < 4; i++ {
		deg := CrossAngl(cord[i][0], cord[i][1], cord[i+1][0], cord[i+1][1])
		deg5 = append(deg5, deg)
	}
	deg := CrossAngl(cord[4][0], cord[4][1], cord[0][0], cord[0][1])
	deg5 = append(deg5, deg)

	return deg5
}
