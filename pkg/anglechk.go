package pkg

// AnglChk は内閣を計算して傾斜屋根を掛けるかどうか判断する
func AnglChk(list [][]float64) (chk bool) {
	// 超点数を求める
	nod := len(list)
	// 内角を求める
	var ang float64
	for i := 0; i < nod; i++ {
		if i == nod-1 {
			x1 := list[i][0]
			y1 := list[i][1]
			x2 := list[0][0]
			y2 := list[0][1]
			ang = CrossAngl(x1, y1, x2, y2)
		} else {
			x1 := list[i][0]
			y1 := list[i][1]
			x2 := list[i+1][0]
			y2 := list[i+1][1]
			ang = CrossAngl(x1, y1, x2, y2)
		}
	}
	if ang < 45 || ang > 135 {
		chk = false
	} else {
		chk = true
	}
	return chk
}
