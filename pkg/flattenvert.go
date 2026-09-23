package pkg

import "log"

// FlatVert は内角が約180°の頂点を削除する
func FlatVert(num int, XY [][]float64, ext []float64, deg []float64) (nod2 int,
	cord2 [][]float64, ext2 []float64, deg2 []float64) {

	var fltLst []int
	for i := 0; i < num; i++ {
		// if (175.0 < deg[i]) && (deg[i] < 185.0) {
		if deg[i] > 175.0 {
			fltLst = append(fltLst, i)
		}
	}
	log.Println("fltLst=", fltLst)

	// log.Println("平坦な頂点の番号", fltLst)
	delCnt := len(fltLst)
	if delCnt != 0 {
		inCnt := 0
		for i := 0; i < delCnt; i++ {
			// log.Println("180Lst", fltLst[i])
			XY = append(XY[:fltLst[i]-inCnt], XY[fltLst[i]+1-inCnt:]...)
			ext = append(ext[:fltLst[i]-inCnt], ext[fltLst[i]+1-inCnt:]...)
			deg = append(deg[:fltLst[i]-inCnt], deg[fltLst[i]+1-inCnt:]...)
			inCnt++
		}
	}
	nod2 = num - delCnt
	cord2 = XY
	ext2 = ext
	deg2 = deg

	return nod2, cord2, ext2, deg2
}
