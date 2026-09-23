package pkg

import (
	"log"
)

// HexaDiv は６角形を２つの４角形に分割する
func HexaDiv(XY [][]float64, order map[string]int) (cord [][]float64,
	rect1List [][]float64, rect2List [][]float64) {
	// 頂点データ数の確認
	nodHex := len(XY)
	log.Println("len(XY)=", nodHex) // Ctrl+/
	if nodHex != 6 {
		return
	}
	var num int

	// L点の直交条件．対向する辺との交点の角度制限を確認する．
	// var LRkey []string
	var int1stX float64
	var int1stY float64
	var int2ndX float64
	var int2ndY float64

	// L1点があるか確認する．ない場合は六角堂
	if val, ok := order["L1"]; ok {
		log.Println(ok)
		log.Println(val)
		// 交点が対向する辺の上にあるか確認する
		var lnincl1 bool
		var lnincl2 bool

		// L!点の頂点番号を確認する
		for LRkey := range order {
			log.Println("LRkey=", LRkey) // Ctrl+/
			if LRkey == "L1" {
				num = order[LRkey]      // 頂点番号
				log.Println("num", num) // Ctrl+/

				// 直交する辺は．L点と1つ前の点で結ばれる線分
				// 直交する辺の座標ペア
				chokuCord1 := make([][]float64, 2)
				numP1 := (num - 1 + nodHex) % nodHex
				chokuCord1[0] = XY[num]
				chokuCord1[1] = XY[numP1]
				// 対向する辺は，L点から２つ目と３つ目の点で結ばれる線分
				// 対向する辺の座標ペア
				taikoCord1 := make([][]float64, 2)
				numP2 := (num + 2) % nodHex
				taikoCord1[0] = XY[numP2]
				numP3 := (num + 3) % nodHex
				taikoCord1[1] = XY[numP3]
				// 直交する直線aと対向する辺との直交条件を確認する
				intX, intY, theta := OrthoAngle(chokuCord1, taikoCord1) // OrthoAngleの戻り値X･Yが逆転
				int1stX = intX
				log.Println("int1stX=", int1stX) // Ctrl+/
				int1stY = intY
				log.Println("int1stY=", int1stY) // Ctrl+/
				// 交差角度が制限範囲内でない場合は処理を中断する
				log.Println("theta=", theta) // Ctrl+/
				if theta < 45 || theta > 135 {
					// TODO:折れ曲がりの切妻屋根
					return
				}
				lnchk1 := PosLine2(chokuCord1, taikoCord1)
				if lnchk1 < 0 {
					lnincl1 = true
				}

				// もう一方の直交する辺は．L点と1つ次の点で結ばれる線分
				// 直交する辺の座標ペア
				chokuCord2 := make([][]float64, 2)
				numN1 := (num + 1) % nodHex
				chokuCord2[0] = XY[num]
				chokuCord2[1] = XY[numN1]
				// もう一方の対向する辺は，L点から３つ目と４つ目の点で結ばれる線分
				// 対向する辺の座標ペア
				taikoCord2 := make([][]float64, 2)
				numN3 := (num + 3) % nodHex
				taikoCord2[0] = XY[numN3]
				numN4 := (num + 4) % nodHex
				taikoCord2[1] = XY[numN4]
				// 直交する直線bと対向する辺との直交条件を確認する
				int2X, int2Y, theta2 := OrthoAngle(chokuCord2, taikoCord2) // OrthoAngleの戻り値X･Yが逆転
				int2ndX = int2X
				log.Println("int2ndX=", int2ndX) // Ctrl+/
				int2ndY = int2Y
				log.Println("int2ndY=", int2ndY) // Ctrl+/
				// 交差角度が制限範囲内でない場合は処理を中断する
				log.Println("theta2=", theta2) // Ctrl+/
				if theta < 45 || theta > 135 {
					// TODO:折れ曲がりの切妻屋根
					return
				}
				lnchk2 := PosLine2(chokuCord2, taikoCord2)
				if lnchk2 < 0 {
					lnincl2 = true
				}

			} else {
				// L1点が見つかるまで処理を繰り返す
				continue
			}
			log.Println("normal termination") // Ctrl+/
			continue
			// TODO:L点がない６角形は六角堂，三角メッシュ分割
		}
		log.Println("All finished") // Ctrl+/

		// L点から対向する二辺までの距離を比較する
		// L点の座標
		log.Println("X座標", XY[num][0]) // Ctrl+/
		log.Println("Y座標", XY[num][1]) // Ctrl+/
		// 交点１までの距離
		divLa := DistVerts(XY[num][0], XY[num][1], int1stX, int1stY)
		log.Println("divLa=", divLa)
		// 交点２までの距離
		divLb := DistVerts(XY[num][0], XY[num][1], int2ndX, int2ndY)
		log.Println("divLb=", divLb)
		var D1splt float64
		var D2splt float64

		if lnincl1 {
			// 交点１から隣り合うR点までの最短距離を求める
			// D1点（交点１）と１つ前のR2点までの距離
			D1toR2 := DistVerts(int1stX, int1stY, XY[(num+2)%6][0], XY[(num+2)%6][1])
			// D1点（交点１）と１つ後ろのR3点までの距離
			D1toR3 := DistVerts(int1stX, int1stY, XY[(num+3)%6][0], XY[(num+3)%6][1])
			if D1toR2 < D1toR3 {
				D1splt = D1toR2
			} else {
				D1splt = D1toR3
			}
		}
		if lnincl2 {
			// 交点２から隣り合うR点までの最短距離を求める
			// D2点（交点２）と１つ前のR3点までの距離
			D2toR3 := DistVerts(int2ndX, int2ndY, XY[(num+3)%6][0], XY[(num+3)%6][1])
			// D2点（交点２）と１つ後ろのR4点までの距離
			D2toR4 := DistVerts(int2ndX, int2ndY, XY[(num+4)%6][0], XY[(num+4)%6][1])
			if D2toR3 < D2toR4 {
				D2splt = D2toR3
			} else {
				D2splt = D2toR4
			}
		}

		// 四角形のスライスを用意する
		var rect1name []string
		var rect2name []string

		// 交点からR点までの最短距離を比較する
		if (D1splt > D2splt) && D1splt >= 1.0 {
			if divLa >= 1.0 {
				// 分割点はD1点（交点１）
				d1 := []float64{int1stX, int1stY}
				log.Println("d1=", d1)
				// 座標値のリストにD1点の座標値を追加する
				XY = append(XY, d1)
				log.Println(XY) // Ctrl+/
				// 頂点並びの辞書に分割点を追加する
				d1Num := nodHex
				order["D1"] = d1Num
				log.Println("line_a", order) // Ctrl+/

				// 四角形D1-L1-R1-R2(1-2頂点，3-4頂点が妻面)
				rect1name = []string{"D1", "L1", "R1", "R2"}
				// 四角形R5-D1-R3-R4(1-2頂点，3-4頂点が妻面)
				rect2name = []string{"R5", "D1", "R3", "R4"}
			}
		} else if (D1splt < D2splt) && D2splt >= 1.0 {
			if divLb >= 1.0 {
				// 分割点はD2点（交点２）
				d2 := []float64{int2ndX, int2ndY}
				log.Println("d2=", d2)
				// 座標値のリストにD2点の座標値を追加する
				XY = append(XY, d2)
				log.Println(XY) // Ctrl+/
				// 頂点並びの辞書に分割点を追加する
				d2Num := nodHex
				order["D2"] = d2Num
				log.Println("line_b", order) // Ctrl+/

				// 四角形D2-R1-R2-R3(1-2頂点，3-4頂点が妻面)
				rect1name = []string{"D2", "R1", "R2", "R3"}
				// 四角形L1-D2-R4-R5(1-2頂点，3-4頂点が妻面)
				rect2name = []string{"L1", "D2", "R4", "R5"}
			}
		} else {
			// 四角形分割しない
			return
		}

		// 辞書の中身に従ってリストの座標データで四角形を作る
		// var rect1List [][]float64
		for _, v := range rect1name {
			// log.Println(i, v)
			n := order[v]
			rect1List = append(rect1List, XY[n])
		}
		log.Println(rect1List) // Ctrl+/

		// 辞書の中身に従ってリストの座標データで四角形を作る
		// var rect2List [][]float64
		for _, v := range rect2name {
			// log.Println(i, v)
			n := order[v]
			rect2List = append(rect2List, XY[n])
		}
		log.Println(rect2List) // Ctrl+/
	}

	cord = XY
	log.Println("cord=", cord) // Ctrl+/
	return cord, rect1List, rect2List
}
