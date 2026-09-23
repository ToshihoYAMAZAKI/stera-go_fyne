package pkg

import "log"

// MakeOct1 は歯型1の８角形を３つの４角形に分割する
func MakeOct1(XY [][]float64, order map[string]int) (cord [][]float64,
	rect1List [][]float64, rect2List [][]float64, rect3List [][]float64,
	story []int, yane []string) {
	// octT := "歯型1"
	nodOct := len(XY)
	if nodOct != 8 {
		// TODO:8頂点でない多角形は，三角メッシュ分割
		return
	}

	num1 := order["L1"]
	// 直交する辺は．L1点と1つ次の点で結ばれる線分
	// 直交する辺の座標ペア
	chokuCord1 := make([][]float64, 2)
	num1N1 := (num1 + 1) % nodOct
	chokuCord1[0] = XY[num1]
	chokuCord1[1] = XY[num1N1]
	// 対向する辺は，L1点から５つ目と６つ目の点で結ばれる線分
	// 対向する辺の座標ペア
	taikoCord1 := make([][]float64, 2)
	num1N5 := (num1 + 5) % nodOct
	taikoCord1[0] = XY[num1N5]
	num1N6 := (num1 + 6) % nodOct
	taikoCord1[1] = XY[num1N6]
	// 直交する直線aと対向する辺との直交条件を確認する
	intX, intY, theta := OrthoAngle(chokuCord1, taikoCord1)
	int1stX := intX
	int1stY := intY
	// 交差角度が制限範囲内でない場合は処理を中断する
	if theta < 45 || theta > 135 {
		log.Println("theta=", theta)
		return
	}
	lnchk1 := PosLine2(chokuCord1, taikoCord1)
	if lnchk1 > 0 {
		log.Println("lnchk1=", lnchk1)
		return
	}

	num2 := order["L2"]
	// もう一本の直交する辺は．L2点と1つ前の点で結ばれる線分
	//  直交する辺の座標ペア
	chokuXY := make([][]float64, 2)
	num2P1 := (num2 - 1 + nodOct) % nodOct
	chokuXY[0] = XY[num2]
	chokuXY[1] = XY[num2P1]
	// 対向する辺は，L2点から２つ目と３つ目の点で結ばれる線分
	// 対向する辺の座標ペア
	taikoXY := make([][]float64, 2)
	num2P2 := (num2 + 2) % nodOct
	taikoXY[0] = XY[num2P2]
	num2P3 := (num2 + 3) % nodOct
	taikoXY[1] = XY[num2P3]
	// 直交する直線bと対向する辺との直交条件を確認する
	int2X, int2Y, theta2 := OrthoAngle(chokuXY, taikoXY)
	int2ndX := int2X
	int2ndY := int2Y
	// 交差角度が制限範囲内でない場合は処理を中断する
	if theta2 < 45 || theta2 > 135 {
		log.Println("theta=", theta)
		return
	}
	lnchk2 := PosLine2(chokuXY, taikoXY)
	if lnchk2 > 0 {
		log.Println("lnchk2=", lnchk2)
		return
	}

	// 分割点はD1点（交点１）
	d1 := []float64{int1stX, int1stY}
	log.Println("d1=", d1)
	// 座標値のリストにD1点の座標値を追加する
	XY = append(XY, d1)
	// 分割点はD2点（交点２）
	d2 := []float64{int2ndX, int2ndY}
	log.Println("d2=", d2)
	// 座標値のリストにD2点の座標値を追加する
	XY = append(XY, d2)
	// 頂点並びの辞書に分割点を追加する
	d1num := nodOct
	order["D1"] = d1num
	d2num := nodOct + 1
	order["D2"] = d2num

	// 四角形L1-D1-R5-R6
	// 切妻・片流れ／平屋
	rect1name := []string{"L1", "D1", "R5", "R6"}
	story = append(story, 1)
	yane = append(yane, "kata")
	// 四角形L2-R2-R3-D2
	// 切妻・片流れ／平屋
	rect2name := []string{"D2", "L2", "R2", "R3"}
	story = append(story, 1)
	yane = append(yane, "kata")
	// 四角形R1-D2-R4-D1
	// 寄棟屋根
	rect3name := []string{"R1", "D2", "R4", "D1"}
	story = append(story, 2)
	yane = append(yane, "yose")

	// 辞書の中身に従ってリストの座標データで四角形を作る
	rect1List = MakeRectList(XY, order, rect1name)
	log.Println("rect1List=", rect1List)
	rect2List = MakeRectList(XY, order, rect2name)
	log.Println("rect2List=", rect2List)
	rect3List = MakeRectList(XY, order, rect3name)
	log.Println("rect3List=", rect3List)

	// スライスをコピーし，コピーされた要素の個数を返す
	// n := copy(cord, XY)
	// log.Println(n)
	cord = XY
	// log.Println("cord=", cord)
	return cord, rect1List, rect2List, rect3List, story, yane
}
