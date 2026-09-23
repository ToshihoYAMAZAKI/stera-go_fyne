package pkg

// PosLineは直交条件を確認するために２点による線分に対してある点の位置を求める
// 外積のZ成分の計算，Z>0は線の左側，Z<0は線の右側
func PosLine(x1, x0, y1, y0, Y, X float64) (t float64) {
	t = (x1-x0)*(Y-y0) - (y1-y0)*(X-x0)
	return
}
