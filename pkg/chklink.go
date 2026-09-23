package pkg

import "log"

// 分割する四角形の中にL1点，L2点が含まれていないかチェックする
func ChkLinc(rect [][]float64, Lxy []float64) (chkL bool) {
	// L1点およびL2点を点eとし，四角形の4頂点を点a,b,c,dとする
	// 点aから点eへの外積を計算する，同様に点bから点dまで外積を求める
	// 全ての外積が＋であれば，点eは各辺の左側にあると言え四角形の中にある
	X := Lxy[0]
	Y := Lxy[1]
	log.Println("X=", X)
	log.Println("Y=", Y)
	// rectから点a～dのXY座標を得る
	var t [4]float64
	tTal := 0
	for i := 0; i < 4; i++ {
		t[i] = PosLine(rect[(i+1)%4][0], rect[i][0], rect[(i+1)%4][1], rect[i][1], Y, X)
		if t[i] > 0 {
			tTal++
		}
	}
	if tTal == 4 {
		chkL = false
	} else {
		chkL = true
	}
	return chkL
}
