package pkg

// PosLine2は直交条件を確認するために２点による線分に対してある点の位置を求める
// 外積のZ成分の計算，Z>0は線の左側，Z<0は線の右側，乗算の結果が負の場合は線分上で交差する
func PosLine2(orthoXY, faceXY [][]float64) (lnchk float64) {
	// 直交する辺の両端座標（一方がL点）
	x1 := orthoXY[0][0] // X座標
	y1 := orthoXY[0][1] // Y座標
	x2 := orthoXY[1][0] // X座標
	y2 := orthoXY[1][1] // Y座標
	// 対向する辺の両端座標
	x3 := faceXY[0][0] // X座標
	y3 := faceXY[0][1] // Y座標
	x4 := faceXY[1][0] // X座標
	y4 := faceXY[1][1] // Y座標

	t1 := (x2-x1)*(y3-y1) - (y2-y1)*(x3-x1)
	t2 := (x2-x1)*(y4-y1) - (y2-y1)*(x4-x1)
	lnchk = t1 * t2

	return lnchk
}
