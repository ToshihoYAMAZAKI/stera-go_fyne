package pkg

import (
	"log"
)

// PentaNode は５角形に屋根を掛けるために頂点座標の並びを整える
func PentaNode(deg []float64, cord [][]float64) (cord5 [][]float64, yane string) {
	// ５角形の各頂点の内角を確認する
	var wa []int
	for d := range deg {
		if deg[d] > 135 {
			wa = append(wa, d)
		}
	}
	log.Println("wa=", wa)

	maxcnt := 0
	// 内角が広角の頂点数に応じて屋根を掛ける
	if len(wa) == 1 {
		// ５角形に変形の寄棟屋根を掛ける
		degmax := 0.0
		for i := range deg {
			if deg[i] > degmax {
				degmax = deg[i]
				maxcnt = i
			}
		}
		yane = "penta"
	} else if len(wa) == 2 {
		if wa[1]-wa[0] == 1 {
			// 五角堂として屋根を掛けることを検討する
			// 広角に対向する辺の両端の角度が広角もしくは狭角の場合は五角堂屋根とする
			degC := deg[(wa[0]+2)%5]
			if degC > 80 && degC < 100 {
				maxcnt = wa[0]
				yane = "kiri5"
			} else {
				yane = "5kakudou"
			}
			// // 広角の２点がなす辺に対して、２番目の点から垂線が降ろせる場合は五角堂とする
			// // 対向するC点の頂点番号
			// num := (wa[0] + 2) & 7
			// // C点のX座標
			// pc := cord[num][0]
			// // C点のY座標
			// qc := cord[num][1]
			// // 広角の２点がなす辺の直線の方程式
			// line := LineEquat(cord[wa[0]][0], cord[wa[0]][1], cord[wa[1]][0], cord[wa[1]][1])
			// a1 := line["m"]
			// b1 := line["n"]
			// // C点から対向する辺に下ろした垂線の交点の座標
			// xc := (a1*(qc-b1) + pc) / (math.Pow(a1, 2) + 1)
			// yc := (a1*(a1*(qc-b1)+pc))/(math.Pow(a1, 2)+1) + b1
			// D1 := []float64{xc, yc}
			// log.Println("D1=", D1)
			// // 垂線の交点が辺の上にあるかどうかチェックする
			// chk := PointonLine(cord[wa[0]][0], cord[wa[0]][1], cord[wa[1]][0], cord[wa[1]][1], xc, yc)
			// log.Println("chk=", chk)

			// if chk {
			// 	yane = "5kakudou"
			// } else if !chk {
			// 	maxcnt = wa[0]
			// 	yane = "kiri5"
			// }

		} else if wa[1]-wa[0] == 2 {
			maxcnt = wa[0] + 1
			yane = "penta"
		}

	}
	log.Println("maxcnt=", maxcnt)

	if yane == "kiri5" {
		if maxcnt < 3 {
			slice1 := cord[maxcnt+2:]
			slice2 := cord[:maxcnt+2]
			cord5 = append(cord5, slice1...)
			cord5 = append(cord5, slice2...)
		} else if maxcnt == 3 {
			slice1 := cord[(maxcnt+2)%5:]
			cord5 = append(cord5, slice1...)
		} else if maxcnt > 3 {
			slice1 := cord[(maxcnt+2)%5:]
			slice2 := cord[:(maxcnt+2)%5]
			cord5 = append(cord5, slice1...)
			cord5 = append(cord5, slice2...)
		}
	} else if yane == "penta" {
		if maxcnt < 2 {
			slice1 := cord[maxcnt+3:]
			slice2 := cord[:maxcnt+3]
			cord5 = append(cord5, slice1...)
			cord5 = append(cord5, slice2...)
		} else if maxcnt == 2 {
			slice1 := cord[(maxcnt+3)%5:]
			cord5 = append(cord5, slice1...)
		} else if maxcnt > 2 {
			slice1 := cord[(maxcnt+3)%5:]
			slice2 := cord[:(maxcnt+3)%5]
			cord5 = append(cord5, slice1...)
			cord5 = append(cord5, slice2...)
		}
	} else {
		cord5 = cord
		yane = "flat"
	}

	// if maxcnt < 2 {
	// 	slice1 := cord[(maxcnt + 3):5]
	// 	slice2 := cord[0:(maxcnt + 3)]
	// 	cord5 = append(cord5, slice1...)
	// 	cord5 = append(cord5, slice2...)
	// } else if maxcnt > 2 {
	// 	slice1 := cord[(maxcnt - 2):5]
	// 	slice2 := cord[0:(maxcnt - 2)]
	// 	cord5 = append(cord5, slice1...)
	// 	cord5 = append(cord5, slice2...)
	// } else {
	// 	cord5 = cord
	// }

	log.Println("cord5=", cord5)

	return cord5, yane
}
