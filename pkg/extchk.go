package pkg

// ExtChk は外積を計算して凹角の有無を確認する
func ExtChk(n int, XY [][]float64) (extchk bool) {

	// 各頂点の外積を計算する
	for i := 0; i < n; i++ {
		xs := XY[i][0]
		ys := XY[i][1]
		xp := XY[(i-1+n)%n][0]
		yp := XY[(i-1+n)%n][1]
		xn := XY[(i+1)%n][0]
		yn := XY[(i+1)%n][1]

		// 外積を計算する
		// a x b = |a||b|sinθ = ax * by - ay * bx
		s := (xp-xs)*(yn-ys) - (xn-xs)*(yp-ys)
		// 外積の計算結から凹角かどうか判断する
		if s > 0 {
			extchk = false
		}
	}
	return extchk
}
