package pkg

import (
	"log"
	"sort"
)

// DecaPrepro は10角形から１つの四角形を分離する
func DecaPrepro(num int, XY [][]float64, nod int,
	order map[string]int, tp string) (cord [][]float64, ordering map[string]int,
	keyList []string, areaTag string) {
	log.Println("num", num)
	log.Println("nod", nod)
	// 直交する辺aは．L点と1つ前の点で結ばれる線分
	chokuPairA, distA := VertsA(num, nod, XY)
	log.Println("distA", distA)
	// もう一方の直交する辺bは．L点と次の点で結ばれる線分
	chokuPairB, distB := VertsB(num, nod, XY)
	log.Println("distB", distB)

	// 直交する辺aに対向する辺を求める
	var taikoPairA [][]float64
	if tp == "a2" {
		taikoPairA = OpposeA2(num, XY, nod)
	} else {
		taikoPairA = OpposeA(num, XY, nod)
	}
	log.Println("taikoPairA", taikoPairA)

	// 直交する辺bに対向する辺を求める
	var taikoPairB [][]float64
	if tp == "b2" {
		taikoPairB = OpposeB2(num, XY, nod)
	} else {
		taikoPairB = OpposeB(num, XY, nod)
	}
	log.Println("taikoPairB", taikoPairB)

	// 分割する四角形の中にL1点，L2点が含まれていないかチェックする
	L1num := order["L1"]
	L1xy := XY[L1num]
	L2num := order["L2"]
	L2xy := XY[L2num]

	// 直交する辺aによる交点aと四角形areaA
	// 直交する辺aに対向する辺との交点を求める
	intAX, intAY, thetaA := OrthoAngle(chokuPairA, taikoPairA)
	// 交差角度が制限範囲内かどうか確認する
	if thetaA < 75 || thetaA > 105 {
		// if thetaA < 60 || thetaA > 120 {
		// TODO:四角形分割は行わない
		// return
	}
	// 交点aが対向する辺上にあるかチェックする
	// 対向する辺の頂点が直行する直線に対して同じ側にある場合
	// 交点aは対向する辺上にない．
	A1 := PosLine(chokuPairA[1][0], chokuPairA[0][0], chokuPairA[1][1],
		chokuPairA[0][1], taikoPairA[0][1], taikoPairA[0][0])
	A2 := PosLine(chokuPairA[1][0], chokuPairA[0][0], chokuPairA[1][1],
		chokuPairA[0][1], taikoPairA[1][1], taikoPairA[1][0])
	var areaA float64
	if A1*A2 < 0 {
		// 四角形aの面積areaAを求める
		var Sa1 float64
		num2 := (num + 2) % nod
		Sa1 = TriArea(XY[num][0], XY[num][1], intAX, intAY, XY[num2][0], XY[num2][1])
		log.Println("Sa1", Sa1)
		var Sa2 float64
		num1 := (num + 1) % nod
		Sa2 = TriArea(XY[num][0], XY[num][1], XY[num1][0], XY[num1][1], XY[num2][0], XY[num2][1])
		log.Println("Sa2", Sa2)
		areaA = Sa1 + Sa2
	} else {
		areaA = 0.0
		// areaA = math.Inf(1)
	}
	log.Println("areaA=", areaA)

	// 四角形の中に他のL点が含まれていないかチェックする
	rectA := [][]float64{XY[num], XY[(num+1)%nod], XY[(num+2)%nod], XY[(num+3)%nod]}

	chkLa1 := ChkLinc(rectA, L1xy)
	if !chkLa1 {
		log.Println("点L1が分割した四角形の中に含まれる")
	}
	chkLa2 := ChkLinc(rectA, L2xy)
	if !chkLa2 {
		log.Println("点L2が分割した四角形の中に含まれる")
	}
	var areaAdis float64
	if chkLa1 && chkLa2 {
		// 交点aから隣り合うR点までの最短距離を求める
		// 交点aと１つ前のR点までの距離
		// R2点の座標　XY[(num+2)%6][0]，XY[(num+2)%6][1]
		DatoR2 := DistVerts(intAX, intAY, XY[(num+2)%nod][0], XY[(num+2)%6][1])
		// 交点aと１つ後ろのR点までの距離
		// R3点の座標　XY[(num+3)%6][0]，XY[(num+3)%6][1]
		DatoR3 := DistVerts(intAX, intAY, XY[(num+3)%nod][0], XY[(num+3)%6][1])

		if DatoR2 < DatoR3 {
			areaAdis = DatoR2
		} else {
			areaAdis = DatoR3
		}
	} else {
		areaAdis = 0.0
	}

	// 直交する辺bによる交点bと四角形areaB
	// 直交する辺bに対向する辺との交点を求める
	intBX, intBY, thetaB := OrthoAngle(chokuPairB, taikoPairB)
	log.Println("intBX=", intBX)
	log.Println("intBY=", intBY)
	log.Println("thetaB=", thetaB)
	// 交差角度が制限範囲内かどうか確認する
	if thetaB < 75 || thetaB > 105 {
		// if thetaB < 60 || thetaB > 120 {
		// TODO:四角形分割は行わない
		// return
	}
	// 交点bが対向する辺上にあるかチェックする
	// 対向する辺の頂点が直行する直線に対して同じ側にある場合
	// 交点bは対向する辺上にない．
	B1 := PosLine(chokuPairB[1][0], chokuPairB[0][0], chokuPairB[1][1],
		chokuPairB[0][1], taikoPairB[0][1], taikoPairB[0][0])
	B2 := PosLine(chokuPairB[1][0], chokuPairB[0][0], chokuPairB[1][1],
		chokuPairB[0][1], taikoPairB[1][1], taikoPairB[1][0])
	var areaB float64
	if B1*B2 < 0 {
		// 四角形bの面積areaBを求める
		var Sb1 float64
		numn2 := (num - 2 + nod) % nod
		Sb1 = TriArea(XY[num][0], XY[num][1], intBX, intBY, XY[numn2][0], XY[numn2][1])
		log.Println("Sb1", Sb1)
		var Sb2 float64
		numn1 := (num - 1 + nod) % nod
		Sb2 = TriArea(XY[num][0], XY[num][1], XY[numn1][0], XY[numn1][1], XY[numn2][0], XY[numn2][1])
		log.Println("Sb2", Sb2)
		areaB = Sb1 + Sb2
	} else {
		areaB = 0.0
		// areaB = math.Inf(1)
	}
	log.Println("areaB", areaB)

	// 四角形の中に他のL点が含まれていないかチェックする
	rectB := [][]float64{XY[num], XY[(num-3+nod)%nod], XY[(num-2+nod)%nod], XY[(num-1+nod)%nod]}

	chkLb1 := ChkLinc(rectB, L1xy)
	if !chkLb1 {
		log.Println("点L1が分割した四角形の中に含まれる")
	}
	chkLb2 := ChkLinc(rectB, L2xy)
	if !chkLb2 {
		log.Println("点L2が分割した四角形の中に含まれる")
	}
	var areaBdis float64
	if chkLb1 && chkLb2 {
		// 交点bから隣り合うR点までの最短距離を求める
		// 交点bと１つ前のR点までの距離
		// R5点の座標　XY[(num+3)%6][0]，XY[(num+3)%6][1]
		DbtoR5 := DistVerts(intBX, intBY, XY[(num-2+nod)%nod][0], XY[(num+3)%6][1])
		// 交点bと１つ後ろのR点までの距離
		// R6点の座標　XY[(num+4)%6][0]，XY[(num+4)%6][1]
		DbtoR6 := DistVerts(intBX, intBY, XY[(num-3+nod)%nod][0], XY[(num+4)%6][1])

		if DbtoR5 < DbtoR6 {
			areaBdis = DbtoR5
		} else {
			areaBdis = DbtoR6
		}
	} else {
		areaBdis = 0.0
	}

	// 分割点はD0点
	var d0 []float64
	// L1点・L2点を含む四角形は分割しない
	if areaAdis == 0.0 {
		d0 = []float64{intBX, intBY}
		areaTag = "areaB"
		log.Println("d0=", d0)
	} else if areaBdis == 0.0 {
		d0 = []float64{intAX, intAY}
		areaTag = "areaA"
		log.Println("d0=", d0)
	} else if areaAdis > 0.0 && areaBdis > 0.0 {
		// areaAとareaBを比較し大きい方を分割する
		if areaA > areaB {
			if areaAdis > areaBdis {
				d0 = []float64{intAX, intAY}
				areaTag = "areaA"
				log.Println("d0=", d0)
			} else if areaB > 0.0 {
				d0 = []float64{intBX, intBY}
				areaTag = "areaB"
				log.Println("d0=", d0)
			}
		} else if areaA < areaB {
			if areaAdis < areaBdis {
				d0 = []float64{intBX, intBY}
				areaTag = "areaB"
				log.Println("d0=", d0)
			} else if areaA > 0.0 {
				d0 = []float64{intAX, intAY}
				areaTag = "areaA"
				log.Println("d0=", d0)
			}
		}
	}

	// 座標値のリストにD0点の座標値を追加する
	cord = append(XY, d0)
	log.Println("cord", cord)
	// 頂点並びの辞書に分割点を追加する
	d0Num := nod
	order["D0"] = d0Num
	ordering = order
	log.Println("order", ordering)
	// 辞書からキーのリストを作る
	for key := range ordering {
		keyList = append(keyList, key)
	}
	// 頂点並びの順序で値（Ｌ･Ｒ点）のリストを作る
	sort.SliceStable(keyList, func(i int, j int) bool {
		return ordering[keyList[i]] < ordering[keyList[j]]
	})
	log.Println("keyList", keyList)

	return cord, ordering, keyList, areaTag
}
