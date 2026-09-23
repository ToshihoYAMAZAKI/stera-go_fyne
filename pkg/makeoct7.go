package pkg

import (
	"log"
	"math"
)

// Pair は分割線の長さの構造体
type Pair struct {
	Key   string
	Value float64
}

// DivlineList は構造体のスライス
type DivlineList []Pair

// Sort関数のインターフェイス
func (p DivlineList) Len() int           { return len(p) }
func (p DivlineList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p DivlineList) Less(i, j int) bool { return p[i].Value < p[j].Value }

// 直交条件を確認するために２点による線分に対してある点の位置を求める
// 外積のZ成分の計算，Z>0は線の左側，Z<0は線の右側
// func posline(x1, x0, y1, y0, Y, X float64) (t float64) {
// 	t = (x1-x0)*(Y-y0) - (y1-y0)*(X-x0)
// 	return
// }

// MakeOct7 はＳ型の８角形を３つの４角形に分割する
func MakeOct7(XY [][]float64, order map[string]int) (cord [][]float64,
	rect1List [][]float64, rect2List [][]float64, rect3List [][]float64,
	story []int, yane []string) {
	// octT := "Ｓ型"
	nodOct := len(XY)
	if nodOct != 8 {
		// TODO:8頂点でない多角形は，三角メッシュ分割
		return
	}
	// L1点から頂点2-3，頂点3-4，頂点4-5，頂点5-6の辺との直交対向関係により型分類する
	var flagEdge2to3 bool // Line 1 A_S
	var flagEdge3to4 bool // Line 2 A_S
	var flagEdge4to5 bool // Line 3 B_S
	var flagEdge5to6 bool // Line 4 B_S

	// L1点からの分割線による対向する辺との交点までの距離を求める
	num1 := order["L1"]
	// 直交する辺は．L1点と1つ前の点で結ばれる線分
	// 直交する辺の座標ペア（分割線1a）
	chokuCord1ap := make([][]float64, 2)
	num1P1ap := (num1 - 1 + nodOct) % nodOct
	chokuCord1ap[0] = XY[num1]
	chokuCord1ap[1] = XY[num1P1ap]
	// L1点と1つ前の点の間の距離
	dl1p1a := DistVerts(XY[num1][0], XY[num1][1], XY[num1P1ap][0], XY[num1P1ap][1])

	// 頂点2-3の辺の座標ペア A_S
	// 対向する辺の座標ペア
	taikoCord1a1 := make([][]float64, 2)
	num1P2a1 := (num1 + 2) % nodOct
	taikoCord1a1[0] = XY[num1P2a1]
	num1P3a1 := (num1 + 3) % nodOct
	taikoCord1a1[1] = XY[num1P3a1]
	// 直交する直線1aと対向する辺との交点を求める
	int1a1X, int1a1Y, theta1a1 := OrthoAngle(chokuCord1ap, taikoCord1a1)

	var divLine1a1 float64
	// 交差角度が制限範囲内かどうか確認する
	if theta1a1 > 45 || theta1a1 < 135 {
		// 対向する辺の頂点が直行する直線に対して同じ側にある場合
		// 交点は対向する辺上にない．
		t1 := PosLine(chokuCord1ap[1][0], chokuCord1ap[0][0], chokuCord1ap[1][1],
			chokuCord1ap[0][1], taikoCord1a1[0][1], taikoCord1a1[0][0])
		t2 := PosLine(chokuCord1ap[1][0], chokuCord1ap[0][0], chokuCord1ap[1][1],
			chokuCord1ap[0][1], taikoCord1a1[1][1], taikoCord1a1[1][0])
		// 直交する直線に対し，対向する辺の2点のZ成分t1とt2の符号
		// 符号が異なっていればt1*t2は負となり直交関係にある
		if t1*t2 < 0 {
			// L1点から２つ目-３つ目の辺の直交対向条件を満たす
			flagEdge2to3 = true
			// 交点1a1までの距離
			divLine1a1 = DistVerts(XY[num1][0], XY[num1][1], int1a1X, int1a1Y)
		} else {
			// 対向する辺と直交する直線が交差しないのでＳ型回転体の可能性がある
			flagEdge2to3 = false
		}
	}

	// 頂点4-5の辺の座標ペア B_S
	// 対向する辺の座標ペア
	taikoCord1a3 := make([][]float64, 2)
	num1P4a3 := (num1 + 4) % nodOct
	taikoCord1a3[0] = XY[num1P4a3]
	num1P5a3 := (num1 + 5) % nodOct
	taikoCord1a3[1] = XY[num1P5a3]
	// 直交する直線1aと対向する辺との交点を求める
	int1a3X, int1a3Y, theta1a3 := OrthoAngle(chokuCord1ap, taikoCord1a3)

	var divLine1a3 float64
	// 交差角度が制限範囲内かどうか確認する
	if theta1a3 > 45 || theta1a3 < 135 {
		// 対向する辺の頂点が直行する直線に対して同じ側にある場合
		// 交点は対向する辺上にない．
		t1 := PosLine(chokuCord1ap[1][0], chokuCord1ap[0][0], chokuCord1ap[1][1],
			chokuCord1ap[0][1], taikoCord1a3[0][1], taikoCord1a3[0][0])
		t2 := PosLine(chokuCord1ap[1][0], chokuCord1ap[0][0], chokuCord1ap[1][1],
			chokuCord1ap[0][1], taikoCord1a3[1][1], taikoCord1a3[1][0])
		// 直交する直線に対し，対向する辺の2点のZ成分t1とt2の符号
		// 符号が異なっていればt1*t2は負となり直交関係にある
		if t1*t2 < 0 {
			// L1点から４つ目-５つ目の辺の直交対向条件を満たす
			flagEdge4to5 = true
			// 交点1a3までの距離
			divLine1a3 = DistVerts(XY[num1][0], XY[num1][1], int1a3X, int1a3Y)
		} else {
			// 対向する辺と直交する直線が交差しないのでＳ型回転体の可能性がある
			flagEdge4to5 = false
		}
	}

	// もう一方の直交する辺は．L点と次の点で結ばれる線分
	// 直交する辺の座標ペア（分割線1b）
	chokuCord1bn := make([][]float64, 2)
	num1N1bn := (num1 + 1) % nodOct
	chokuCord1bn[0] = XY[num1]
	chokuCord1bn[1] = XY[num1N1bn]
	// L1点と次の点の間の距離
	dl1n1b := DistVerts(XY[num1][0], XY[num1][1], XY[num1N1bn][0], XY[num1N1bn][1])

	// 頂点3-4の辺の座標ペア A_S
	// 対向する辺の座標ペア
	taikoCord1b2 := make([][]float64, 2)
	num1N3b1 := (num1 + 3) % nodOct
	taikoCord1b2[0] = XY[num1N3b1]
	num1N4b1 := (num1 + 4) % nodOct
	taikoCord1b2[1] = XY[num1N4b1]
	// 直交する直線1bと対向する辺との交点を求める
	int1b2X, int1b2Y, theta1b2 := OrthoAngle(chokuCord1bn, taikoCord1b2)

	var divLine1b2 float64
	// 交差角度が制限範囲内かどうか確認する
	if theta1b2 > 45 || theta1b2 < 135 {
		// 対向する辺の頂点が直行する直線に対して同じ側にある場合
		// 交点は対向する辺上にない．
		t1 := PosLine(chokuCord1bn[1][0], chokuCord1bn[0][0], chokuCord1bn[1][1],
			chokuCord1bn[0][1], taikoCord1b2[0][1], taikoCord1b2[0][0])
		t2 := PosLine(chokuCord1bn[1][0], chokuCord1bn[0][0], chokuCord1bn[1][1],
			chokuCord1bn[0][1], taikoCord1b2[1][1], taikoCord1b2[1][0])
		// 直交する直線に対し，対向する辺の2点のZ成分t1とt2の符号
		// 符号が異なっていればt1*t2は負となり直交関係にある
		if t1*t2 < 0 {
			// L1点から３つ目-４つ目の辺の直交対向条件を満たす
			flagEdge3to4 = true
			// 交点1b2までの距離
			divLine1b2 = DistVerts(XY[num1][0], XY[num1][1], int1b2X, int1b2Y)
		} else {
			// 対向する辺と直交する直線が交差しないのでＳ型回転体の可能性がある
			flagEdge3to4 = false
		}
	}

	// 頂点5-6の辺の座標ペア
	// 対向する辺の座標ペア
	taikoCord1b4 := make([][]float64, 2)
	num1N5b3 := (num1 + 5) % nodOct
	taikoCord1b4[0] = XY[num1N5b3]
	num1N6b3 := (num1 + 6) % nodOct
	taikoCord1b4[1] = XY[num1N6b3]
	// 直交する直線1bと対向する辺との交点を求める
	int1b4X, int1b4Y, theta1b4 := OrthoAngle(chokuCord1bn, taikoCord1b4)

	var divLine1b4 float64
	// 交差角度が制限範囲内かどうか確認する
	if theta1b4 > 45 || theta1b4 < 135 {
		// 対向する辺の頂点が直行する直線に対して同じ側にある場合
		// 交点は対向する辺上にない．
		t1 := PosLine(chokuCord1bn[1][0], chokuCord1bn[0][0], chokuCord1bn[1][1],
			chokuCord1bn[0][1], taikoCord1b4[0][1], taikoCord1b4[0][0])
		t2 := PosLine(chokuCord1bn[1][0], chokuCord1bn[0][0], chokuCord1bn[1][1],
			chokuCord1bn[0][1], taikoCord1b4[1][1], taikoCord1b4[1][0])
		// 直交する直線に対し，対向する辺の2点のZ成分t1とt2の符号
		// 符号が異なっていればt1*t2は負となり直交関係にある
		if t1*t2 < 0 {
			// L1点から５つ目-６つ目の辺の直交対向条件を満たす
			flagEdge5to6 = true
			// 交点1b4までの距離
			divLine1b4 = DistVerts(XY[num1][0], XY[num1][1], int1b4X, int1b4Y)
		} else {
			// 対向する辺と直交する直線が交差しないのでＳ型回転体の可能性がある
			flagEdge5to6 = false
		}
	}

	// 四角形の頂点のリストを３つ用意する．
	var rect1name []string
	var rect2name []string
	var rect3name []string

	// L2点からの分割線による対向する辺との交点までの距離を求める
	num2 := order["L2"]
	// 対抗する辺との直交条件から型分類を行う
	if flagEdge2to3 && flagEdge3to4 {
		octStype := "typeA_S"
		log.Println(octStype)
		// L2点と対向する辺との交点を求める
		// １つ目の直交する辺は．L2点と1つ前の点で結ばれる線分
		// 直交する辺の座標ペア（分割線2a）
		chokuXYap := make([][]float64, 2)
		num2P1ap := (num2 - 1 + nodOct) % nodOct
		chokuXYap[0] = XY[num2]
		chokuXYap[1] = XY[num2P1ap]
		// 頂点2-3の辺の座標ペア
		// 対向する辺の座標ペア
		taikoXYa1 := make([][]float64, 2)
		taikoXYa1[0] = XY[(num2+2)%nodOct]
		taikoXYa1[1] = XY[(num2+3)%nodOct]
		// 直交する直線2aと対向する辺との交点を求める
		int2a1X, int2a1Y, theta2a1 := OrthoAngle(chokuXYap, taikoXYa1)
		// 交差角度が制限範囲内かどうか確認する
		divLine2a1 := math.Inf(1)
		if theta2a1 > 45 || theta2a1 < 135 {
			// 対向する辺の頂点が直行する直線に対して同じ側にある場合
			// 交点は対向する辺上にない．
			t1 := PosLine(chokuXYap[1][0], chokuXYap[0][0], chokuXYap[1][1],
				chokuXYap[0][1], taikoXYa1[0][1], taikoXYa1[0][0])
			t2 := PosLine(chokuXYap[1][0], chokuXYap[0][0], chokuXYap[1][1],
				chokuXYap[0][1], taikoXYa1[1][1], taikoXYa1[1][0])
			// 直交する直線に対し，対向する辺の2点のZ成分t1とt2の符号
			// 符号が異なっていればt1*t2は負となり直交関係にある
			if t1*t2 < 0 {
				// 交点2a_2_3までの距離
				divLine2a1 = DistVerts(XY[num2][0], XY[num2][1], int2a1X, int2a1Y)
			} else {
				// L1点においてＳ型の条件を満たすがL2点において満たさない
				// 四角形分割せずに多角形でモデリングする
				divLine2a1 = 0.0
			}
		}
		// もう一方の直交する辺は．L2点と次の点で結ばれる線分
		// 直交する辺の座標ペア（分割線2b）
		chokuXYbn := make([][]float64, 2)
		num2N1bn := (num2 + 1) % nodOct
		chokuXYbn[0] = XY[num2]
		chokuXYbn[1] = XY[num2N1bn]
		// L2点と次の点の間の距離
		dl2n1b := DistVerts(XY[num2][0], XY[num2][1], XY[num2N1bn][0], XY[num2N1bn][1])
		// 頂点3-4の辺の座標ペア
		// 対向する辺の座標ペア
		taikoXYb2 := make([][]float64, 2)
		taikoXYb2[0] = XY[(num2+3)%nodOct]
		taikoXYb2[1] = XY[(num2+4)%nodOct]
		// 直交する直線2bと対向する辺との交点を求める
		int2b2X, int2b2Y, theta2b2 := OrthoAngle(chokuXYbn, taikoXYb2)
		// 交差角度が制限範囲内かどうか確認する
		divLine2b2 := math.Inf(1)
		if theta2b2 > 45 || theta2b2 < 135 {
			// 対向する辺の頂点が直行する直線に対して同じ側にある場合
			// 交点は対向する辺上にない．
			t1 := PosLine(chokuXYap[1][0], chokuXYap[0][0], chokuXYap[1][1],
				chokuXYap[0][1], taikoXYb2[0][1], taikoXYb2[0][0])
			t2 := PosLine(chokuXYap[1][0], chokuXYap[0][0], chokuXYap[1][1],
				chokuXYap[0][1], taikoXYb2[1][1], taikoXYb2[1][0])
			// 直交する直線に対し，対向する辺の2点のZ成分t1とt2の符号
			// 符号が異なっていればt1*t2は負となり直交関係にある
			if t1*t2 < 0 {
				// 交点2b_5_6までの距離
				divLine2b2 = DistVerts(XY[num2][0], XY[num2][1], int2b2X, int2b2Y)
			} else {
				// L1点においてＳ型の条件を満たすがL2点において満たさない
				// 四角形分割せずに多角形でモデリングする
				divLine2b2 = 0.0
			}
		}

		// 分割線1aと1bを比較する，分割線2aと2bを比較する
		// 距離の長い方の線分を分割線とする
		if divLine1a1 > divLine1b2 {
			// 分割点はD1a点（交点１）
			d1a := []float64{int1a1X, int1a1Y}
			// 座標値のリストにD1a点の座標値を追加する
			XY = append(XY, d1a)
			// 頂点並びの辞書に分割点を追加する
			d1anum := nodOct
			order["D1a"] = d1anum

			// 四角形D1a-L1-R1-R2
			rect1name = []string{"D1a", "L1", "R1", "R2"}
			story = append(story, 1)
			if dl1n1b > 1.8 {
				yane = append(yane, "kiri")
			} else {
				yane = append(yane, "kata")
			}

			// 距離の短い方の線分を分割線とする
			if divLine2a1 < divLine2b2 {
				// 分割点はD2a点（交点2）
				d2a := []float64{int2a1X, int2a1Y}
				// 座標値のリストにD2点の座標値を追加する
				XY = append(XY, d2a)
				// 頂点並びの辞書に分割点を追加する
				d2anum := nodOct + 1
				order["D2a"] = d2anum

				// 四角形R6-D1a-R3-D2a
				rect2name = []string{"R6", "D1a", "R3", "D2a"}
				story = append(story, 2)
				yane = append(yane, "kiri")
				// 四角形D2a-L2-R4-R5
				rect3name = []string{"D2a", "L2", "R4", "R5"}
				story = append(story, 1)
				if dl2n1b > 1.8 {
					yane = append(yane, "kiri")
				} else {
					yane = append(yane, "kata")
				}

			} else if divLine2a1 > divLine2b2 {
				// 分割点はD2b点（交点2）
				d2b := []float64{int2b2X, int2b2Y}
				// 座標値のリストにD2b点の座標値を追加する
				XY = append(XY, d2b)
				// 頂点並びの辞書に分割点を追加する
				d2bnum := nodOct + 1
				order["D2b"] = d2bnum

				// 四角形L2-D2b-D1a-R3
				rect2name = []string{"L2", "D2b", "D1a", "R3"}
				story = append(story, 2)
				yane = append(yane, "kiri")
				// 四角形D2b-R4-R5-R6
				rect3name = []string{"D2b", "R4", "R5", "R6"}
				story = append(story, 2)
				yane = append(yane, "kiri")
			}
		} else if divLine1a1 < divLine1b2 {
			// 分割点はD1b点（交点２）
			d1b := []float64{int1b2X, int1b2Y}
			// 座標値のリストにD1b点の座標値を追加する
			XY = append(XY, d1b)
			// 頂点並びの辞書に分割点を追加する
			d1bnum := nodOct
			order["D1b"] = d1bnum

			// 四角形D1b-R1-R2-R3
			rect1name = []string{"D1b", "R1", "R2", "R3"}
			story = append(story, 2)
			yane = append(yane, "kiri")

			// 距離の短い方の線分を分割線とする
			if divLine2a1 < divLine2b2 {
				// 分割点はD2a点（交点２）
				d2a := []float64{int2a1X, int2a1Y}
				// 座標値のリストにD2a点の座標値を追加する
				XY = append(XY, d2a)
				// 頂点並びの辞書に分割点を追加する
				d2anum := nodOct + 1
				order["D2a"] = d2anum

				// 四角形L1-D1b-D2a-R6
				rect2name = []string{"L1", "D1b", "D2a", "R6"}
				story = append(story, 2)
				yane = append(yane, "kiri")
				// 四角形L2-R4-R5-D2a
				rect3name = []string{"D2a", "L2", "R4", "R5"}
				story = append(story, 1)
				if dl2n1b > 1.8 {
					yane = append(yane, "kiri")
				} else {
					yane = append(yane, "kata")
				}
			} else if divLine2a1 > divLine2b2 {
				// 分割点はD2b点（交点２）
				d2b := []float64{int2b2X, int2b2Y}
				// 座標値のリストにD2b点の座標値を追加する
				XY = append(XY, d2b)
				// 頂点並びの辞書に分割点を追加する
				d2bnum := nodOct + 1
				order["D2b"] = d2bnum

				// L点と分割点が共に近接する場合，真ん中の四角形が非常に小さく細長くなる．
				// この場合は，2つのL点を結ぶ線を分割線とする．
				// L1点とD2b点の間の距離
				distL1D2b := DistVerts(XY[num1][0], XY[num1][1], int2b2X, int2b2Y)
				// L2点とD1b点の間の距離
				distL2D1b := DistVerts(XY[num2][0], XY[num2][1], int1b2X, int1b2Y)

				// L点と分割点の間の距離が共に短い場合は，L点を結ぶ線を分割線とする．
				if distL1D2b < 1.0 && distL2D1b < 1.0 {
					// ２つの四角形に分割する．なお，中間点となるL点は使用しない．
					// 四角形L1-D1b-L2-D2b
					rect2name = []string{"L1", "R4", "R5", "R6"}
					story = append(story, 2)
					yane = append(yane, "kiri")
					// 四角形D2b-R4-R5-R6
					rect3name = []string{"L1", "R3", "L2", "R6"}
					story = append(story, 1)
					yane = append(yane, "flat")
				} else {
					// 四角形L1-D1b-L2-D2b
					rect2name = []string{"L1", "D1b", "L2", "D2b"}
					story = append(story, 2)
					yane = append(yane, "kiri")
					// 四角形D2b-R4-R5-R6
					rect3name = []string{"D2b", "R4", "R5", "R6"}
					story = append(story, 2)
					yane = append(yane, "kiri")
				}
			}
		}

		// 辞書の中身に従ってリストの座標データで四角形を作る
		rect1List = MakeRectList(XY, order, rect1name)
		log.Println("rect1List=", rect1List)
		rect2List = MakeRectList(XY, order, rect2name)
		log.Println("rect2List=", rect2List)
		rect3List = MakeRectList(XY, order, rect3name)
		log.Println("rect3List=", rect3List)

	} else if flagEdge4to5 && flagEdge5to6 {
		octStype := "typeB_S"
		log.Println(octStype)
		// L2点と対向する辺との交点を求める
		// １つ目の直交する辺は．L2点と1つ前の点で結ばれる線分
		// 直交する辺の座標ペア（分割線2a）
		chokuXYap := make([][]float64, 2)
		num2P1ap := (num2 - 1 + nodOct) % nodOct
		chokuXYap[0] = XY[num2]
		chokuXYap[1] = XY[num2P1ap]
		// L2点と1つ前の点の間の距離
		dl2p1a := DistVerts(XY[num2][0], XY[num2][1], XY[num2P1ap][0], XY[num2P1ap][1])
		// 頂点4-5の辺の座標ペア
		// 対向する辺の座標ペア
		taikoXYa3 := make([][]float64, 2)
		taikoXYa3[0] = XY[(num2+4)%nodOct]
		taikoXYa3[1] = XY[(num2+5)%nodOct]
		// 直交する直線2aと対向する辺との交点を求める
		int2a3X, int2a3Y, theta2a3 := OrthoAngle(chokuXYap, taikoXYa3)
		// 交差角度が制限範囲内かどうか確認する
		divLine2a3 := math.Inf(1)
		if theta2a3 > 45 || theta2a3 < 135 {
			// 対向する辺の頂点が直行する直線に対して同じ側にある場合
			// 交点は対向する辺上にない．
			t1 := PosLine(chokuXYap[1][0], chokuXYap[0][0], chokuXYap[1][1],
				chokuXYap[0][1], taikoXYa3[0][1], taikoXYa3[0][0])
			t2 := PosLine(chokuXYap[1][0], chokuXYap[0][0], chokuXYap[1][1],
				chokuXYap[0][1], taikoXYa3[1][1], taikoXYa3[1][0])
			log.Println("t1=", t1)
			log.Println("t2=", t2)
			// 直交する直線に対し，対向する辺の2点のZ成分t1とt2の符号
			// 符号が異なっていればt1*t2は負となり直交関係にある
			if t1*t2 < 0 {
				// 交点2a_4_5までの距離
				divLine2a3 = DistVerts(XY[num2][0], XY[num2][1], int2a3X, int2a3Y)
				log.Println("divLine2a3=", divLine2a3)
			} else {
				// L1点においてＳ型の条件を満たすがL2点において満たさない
				// 四角形分割せずに多角形でモデリングする
				divLine2a3 = 0.0
			}
		}
		// もう一方の直交する辺は．L2点と次の点で結ばれる線分
		// 直交する辺の座標ペア（分割線2b）
		chokuXYbn := make([][]float64, 2)
		num2N1bn := (num2 + 1) % nodOct
		chokuXYbn[0] = XY[num2]
		chokuXYbn[1] = XY[num2N1bn]
		// 頂点5-6の辺の座標ペア
		// 対向する辺の座標ペア
		taikoXYb4 := make([][]float64, 2)
		taikoXYb4[0] = XY[(num2+5)%nodOct]
		taikoXYb4[1] = XY[(num2+6)%nodOct]
		// 直交する直線2bと対向する辺との交点を求める
		int2b4X, int2b4Y, theta2b4 := OrthoAngle(chokuXYbn, taikoXYb4)
		// 交差角度が制限範囲内かどうか確認する
		divLine2b4 := math.Inf(1)
		if theta2b4 > 45 || theta2b4 < 135 {
			// 対向する辺の頂点が直行する直線に対して同じ側にある場合
			// 交点は対向する辺上にない．
			t1 := PosLine(chokuXYap[1][0], chokuXYap[0][0], chokuXYap[1][1],
				chokuXYap[0][1], taikoXYb4[0][1], taikoXYb4[0][0])
			t2 := PosLine(chokuXYap[1][0], chokuXYap[0][0], chokuXYap[1][1],
				chokuXYap[0][1], taikoXYb4[1][1], taikoXYb4[1][0])
			log.Println("t1=", t1)
			log.Println("t2=", t2)
			// 直交する直線に対し，対向する辺の2点のZ成分t1とt2の符号
			// 符号が異なっていればt1*t2は負となり直交関係にある
			if t1*t2 < 0 {
				// 交点2b_5_6までの距離
				divLine2b4 = DistVerts(XY[num2][0], XY[num2][1], int2b4X, int2b4Y)
				log.Println("divLine2a3=", divLine2a3)
			} else {
				// L1点においてＳ型の条件を満たすがL2点において満たさない
				// 四角形分割せずに多角形でモデリングする
				divLine2b4 = 0.0
			}
		}

		// 分割線1aと1bを比較する，分割線2aと2bを比較する
		// 距離の長い方の線分を分割線とする
		if divLine1a3 > divLine1b4 {
			// 分割点はD1a点（交点１）
			d1a := []float64{int1a3X, int1a3Y}
			// 座標値のリストにD1a点の座標値を追加する
			XY = append(XY, d1a)
			// 頂点並びの辞書に分割点を追加する
			d1anum := nodOct
			order["D1a"] = d1anum

			// 四角形D1a-R4-R5-R6
			rect1name = []string{"D1a", "R4", "R5", "R6"}
			story = append(story, 2)
			yane = append(yane, "kiri")

			// 距離の短い方の線分を分割線とする
			if divLine2a3 < divLine2b4 {
				// log.Println("分割線はdivLine2a")
				// 分割点はD2a点（交点2）
				d2a := []float64{int2a3X, int2a3Y}
				// 座標値のリストにD2a点の座標値を追加する
				XY = append(XY, d2a)
				// 頂点並びの辞書に分割点を追加する
				d2anum := nodOct + 1
				order["D2a"] = d2anum

				// L点と分割点が共に近接する場合，真ん中の四角形が非常に小さく細長くなる．
				// この場合は，2つのL点を結ぶ線を分割線とする．
				// L1点とD2a点の間の距離
				distL1D2a := DistVerts(XY[num1][0], XY[num1][1], int2a3X, int2a3Y)
				// L2点とD2b点の間の距離
				distL2D1a := DistVerts(XY[num2][0], XY[num2][1], int1a3X, int1a3Y)

				// L点と分割点の間の距離が共に短い場合は，L点を結ぶ線を分割線とする．
				if distL1D2a < 1.0 && distL2D1a < 1.0 {
					// ２つの四角形に分割する．なお，中間点となるL点は使用しない．
					// 四角形D1a-L1-D2a-L2
					rect2name = []string{"R3", "L1", "R1", "R2"}
					story = append(story, 2)
					yane = append(yane, "kiri")
					// 四角形R3-D2a-R1-R2
					rect3name = []string{"L1", "R1", "L2", "R4"}
					story = append(story, 2)
					yane = append(yane, "flat")
				} else {
					// 四角形D1a-L1-D2a-L2
					rect2name = []string{"D1a", "L1", "D2a", "L2"}
					story = append(story, 1)
					yane = append(yane, "kiri")
					// 四角形R3-D2a-R1-R2
					rect3name = []string{"R3", "D2a", "R1", "R2"}
					story = append(story, 2)
					yane = append(yane, "flat")
				}
			} else if divLine2a3 > divLine2b4 {
				// log.Println("分割線はdivLine2b")
				// 分割点はD2b点（交点２）
				d2b := []float64{int2b4X, int2b4Y}
				// 座標値のリストにD2b点の座標値を追加する
				XY = append(XY, d2b)
				// 頂点並びの辞書に分割点を追加する
				d2bnum := nodOct + 1
				order["D2b"] = d2bnum

				// 四角形L2-D2b-R2-R3
				rect2name = []string{"L2", "D2b", "R2", "R3"}
				story = append(story, 1)
				if dl2p1a > 1.8 {
					yane = append(yane, "kiri")
				} else {
					yane = append(yane, "kata")
				}
				// 四角形D1a-L1-R1-D2b
				rect3name = []string{"D1a", "L1", "R1", "D2b"}
				story = append(story, 2)
				yane = append(yane, "kiri")
			}
		} else if divLine1a3 < divLine1b4 {
			// log.Println("分割線はdivLine1b")
			// 分割点はD1b点（交点２）
			d1b := []float64{int1b4X, int1b4Y}
			// 座標値のリストにD1b点の座標値を追加する
			XY = append(XY, d1b)
			// 頂点並びの辞書に分割点を追加する
			d1bnum := nodOct
			order["D1b"] = d1bnum

			// 四角形L1-D1b-R5-R6
			rect1name = []string{"L1", "D1b", "R5", "R6"}
			story = append(story, 1)
			if dl1p1a > 1.8 {
				yane = append(yane, "kiri")
			} else {
				yane = append(yane, "kata")
			}

			// 距離の短い方の線分を分割線とする
			if divLine2a3 < divLine2b4 {
				// 分割点はD2a点（交点２）
				d2a := []float64{int2a3X, int2a3Y}
				// 座標値のリストにD2a点の座標値を追加する
				XY = append(XY, d2a)
				// 頂点並びの辞書に分割点を追加する
				d2anum := nodOct + 1
				order["D2a"] = d2anum

				// 四角形D1b-D2a-L2-R4
				rect2name = []string{"D1b", "D2a", "L2", "R4"}
				story = append(story, 2)
				yane = append(yane, "kiri")
				// 四角形R3-D2a-R1-R2
				rect3name = []string{"R3", "D2a", "R1", "R2"}
				story = append(story, 2)
				yane = append(yane, "kiri")
			} else if divLine2a3 > divLine2b4 {
				// log.Println("分割線はdivLine2b")
				// 分割点はD2b点（交点２）
				d2b := []float64{int2b4X, int2b4Y}
				// 座標値のリストにD2b点の座標値を追加する
				XY = append(XY, d2b)
				// 頂点並びの辞書に分割点を追加する
				d2bnum := nodOct + 1
				order["D2b"] = d2bnum

				// 四角形D2b-R4-D1b-R1
				rect2name = []string{"D2b", "R4", "D1b", "R1"}
				story = append(story, 2)
				yane = append(yane, "kiri")
				// 四角形L2-D2b-R2-R3
				rect3name = []string{"L2", "D2b", "R2", "R3"}
				story = append(story, 1)
				if dl2p1a > 1.8 {
					yane = append(yane, "kiri")
				} else {
					yane = append(yane, "kata")
				}
			}
		}

		// 辞書の中身に従ってリストの座標データで四角形を作る
		rect1List = MakeRectList(XY, order, rect1name)
		log.Println("rect1List=", rect1List)
		rect2List = MakeRectList(XY, order, rect2name)
		log.Println("rect2List=", rect2List)
		rect3List = MakeRectList(XY, order, rect3name)
		log.Println("rect2List=", rect2List)

		// 対抗する辺との直交条件から型分類を行う
	} else if flagEdge2to3 && flagEdge5to6 {
		octStype := "Rotation_S"
		log.Println(octStype)
		// 2つの四角形に分割する
		// 分割点は2つのL点から伸ばした直交する線の交点

		// 1つ目の四角形の交点D1
		// L1点から直交する辺は．L1点と次の点で結ばれる線分
		// 直交する辺の座標ペア（分割線1b）
		chokuL1b := make([][]float64, 2)
		chokuL1b[0] = XY[num1]
		chokuL1b[1] = XY[(num1+1)%nodOct]
		// L1点から直交する辺は．L2点と1つ前の点で結ばれる線分
		// 直交する辺の座標ペア（分割線2a）
		chokuL2a := make([][]float64, 2)
		chokuL2a[0] = XY[num2]
		chokuL2a[1] = XY[(num2-1+nodOct)%nodOct]
		// 分割線1bと分割線2aの交点を求める
		intD1X, intD1Y, _ := OrthoAngle(chokuL1b, chokuL2a)

		// 分割点はD1点
		d1 := []float64{intD1X, intD1Y}
		// 座標値のリストにD2点の座標値を追加する
		XY = append(XY, d1)
		log.Println(XY)
		// 頂点並びの辞書に分割点を追加する
		d1Num := nodOct
		order["D1"] = d1Num
		log.Println("line1a", order)

		// 四角形 D1-R1-R2-R3
		rect1name = []string{"D1", "R1", "R2", "R3"}
		story = append(story, 2)
		yane = append(yane, "kiri")

		// 2つ目の四角形の交点D2
		// L1点から直交する辺は．L1点と1つ前の点で結ばれる線分
		// 直交する辺の座標ペア（分割線1a）
		chokuL1a := make([][]float64, 2)
		chokuL1a[0] = XY[num1]
		chokuL1a[1] = XY[(num1-1+nodOct)%nodOct]
		// L2点から直交する辺は．L2点と次の点で結ばれる線分
		// 直交する辺の座標ペア（分割線2b）
		chokuL2b := make([][]float64, 2)
		chokuL2b[0] = XY[num2]
		chokuL2b[1] = XY[(num2+1)%nodOct]
		// 分割線1aと分割線2bの交点を求める
		intD2X, intD2Y, _ := OrthoAngle(chokuL1a, chokuL2b)

		// 分割点はD2点
		d2 := []float64{intD2X, intD2Y}
		// 座標値のリストにD2点の座標値を追加する
		XY = append(XY, d2)
		log.Println(XY)
		// 頂点並びの辞書に分割点を追加する
		d2Num := nodOct + 1
		order["D2"] = d2Num
		log.Println("line1a", order)

		// 四角形 D2-R4-R5-R6
		rect2name = []string{"D2", "R4", "R5", "R6"}
		story = append(story, 2)
		yane = append(yane, "kiri")

		// 四角形 D1-L1-D2-L2
		rect3name = []string{"D1", "L1", "D2", "L2"}
		story = append(story, 2)
		yane = append(yane, "flat")

		// 辞書の中身に従ってリストの座標データで四角形を作る
		rect1List = MakeRectList(XY, order, rect1name)
		log.Println("rect1List=", rect1List)
		rect2List = MakeRectList(XY, order, rect2name)
		log.Println("rect2List=", rect2List)
		rect3List = MakeRectList(XY, order, rect3name)
		log.Println("rect3List=", rect3List)

	} else {
		octStype := "Others: 陸屋根"
		log.Println(octStype)

		return
	}

	cord = XY
	return cord, rect1List, rect2List, rect3List, story, yane
}
