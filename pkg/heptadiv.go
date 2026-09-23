package pkg

import (
	"log"
	"math"
	"strings"
)

// HeptaDiv は７角形を３つに分割して片流れ屋根を掛ける
func HeptaDiv(lrPtn []string, deg2 []float64, cord2 [][]float64, order map[string]int) (slice1 [][]float64,
	slice2 [][]float64, slice3 [][]float64, type1L, type2L, type3L string, story []int, chk bool) {
	// L点が1つの場合と2つの場合で処理が異なる
	// L点を数える
	lrtxt := strings.Join(lrPtn, "")
	log.Println("lrtxt=", lrtxt)
	var roof5 string

	// ７角形を３つに分割
	// L点が連続する場合はL点の側に流れる片流れ屋根を掛ける
	if strings.Contains(lrtxt, "LL") {
		log.Println("lrtxt include LL")
		log.Println("order=", order)
		log.Println("order[\"L1\"]=", order["L1"])
		L1 := order["L1"]
		log.Println("order[\"L2\"]=", order["L2"])
		L2 := order["L2"]

		var slice1t [][]float64
		if L2 < 4 {
			slice1t = cord2[L2 : L2+4]
		} else if L2 > 3 {
			slice11 := cord2[L2:]
			slice12 := cord2[:(L2+4)%7]
			slice1t = append(slice1t, slice11...)
			slice1t = append(slice1t, slice12...)
		}
		slice1t1 := slice1t[2:]
		slice1t2 := slice1t[:2]
		slice1 = append(slice1, slice1t1...)
		slice1 = append(slice1, slice1t2...)
		log.Println("slice1=", slice1)

		if L1 < 3 {
			slice21 := cord2[L1+4:]
			slice22 := cord2[:L1+1]
			slice2 = append(slice2, slice21...)
			slice2 = append(slice2, slice22...)
		} else if L1 > 2 {
			slice21 := cord2[(L1+4)%7 : (L1+7)%7]
			slice22 := cord2[L1 : L1+1]
			slice2 = append(slice2, slice21...)
			slice2 = append(slice2, slice22...)
		}
		log.Println("slice2=", slice2)

		slice31 := cord2[(L2+3)%7 : (L2+3)%7+1]
		slice32 := cord2[L1 : L1+1]
		slice33 := cord2[L2 : L2+1]
		slice3 = append(slice3, slice31...)
		slice3 = append(slice3, slice32...)
		slice3 = append(slice3, slice33...)
		log.Println("slice3=", slice3)

		type1L = "kata1"
		type2L = "kata2"
		type3L = "heptri"
		for i := 0; i < 3; i++ {
			story = append(story, 2)
		}

	} else if strings.Contains(lrtxt, "L") {
		chk = true
		// L点が2つの場合はL点の間のR点の数に応じて屋根の掛け方を変える
		if strings.Count(lrtxt, "L") == 2 {
			log.Println("lrtxt include Lx2")
			log.Println("order=", order)
			log.Println("order[\"L1\"]=", order["L1"])
			L1 := order["L1"]
			log.Println("order[\"L2\"]=", order["L2"])
			L2 := order["L2"]

			// ７角形を３つに分割
			log.Println("deg2=", deg2)
			if L1 == 0 && L2 == 6 {
				// L点が連続する場合はL点の側に流れる片流れ屋根を掛ける
				slice11 := cord2[L1+2 : L1+4]
				slice12 := cord2[:L1+2]
				slice1 = append(slice1, slice11...)
				slice1 = append(slice1, slice12...)
				log.Println("slice1=", slice1)
				slice2 = cord2[L1+3:]
				log.Println("slice2=", slice2)

				slice31 := cord2[L1+3 : L1+4]
				slice32 := cord2[L2 : L2+1]
				slice33 := cord2[L1 : L1+1]
				slice3 = append(slice3, slice31...)
				slice3 = append(slice3, slice32...)
				slice3 = append(slice3, slice33...)
				log.Println("slice3=", slice3)

				type1L = "kata1"
				type2L = "kata2"
				type3L = "heptri"
				for i := 0; i < 3; i++ {
					story = append(story, 2)
				}

			} else if (L2-L1+7)%7 == 2 {
				// L点が２つ離れた王冠型−１
				log.Println("２つの三角形と１つの５角形に分割する-1")

				// 三角形・片流れ屋根
				slice11 := cord2[L1 : L1+1]
				var slice12 [][]float64
				if L1 < 1 {
					slice12 = cord2[L1+5:]
				} else if L1 == 1 {
					slice12t1 := cord2[L1+5:]
					slice12t2 := cord2[:(L1+7)%7]
					slice12 = append(slice12, slice12t1...)
					slice12 = append(slice12, slice12t2...)
				} else if L1 > 1 {
					slice12 = cord2[(L1+5)%7 : (L1+7)%7]
				}
				slice1 = append(slice1, slice11...)
				slice1 = append(slice1, slice12...)
				log.Println("slice1=", slice1)

				// 三角形・片流れ屋根
				var slice21 [][]float64
				if L2 < 5 {
					slice21 = cord2[L2+2 : L2+3]
				} else if L2 > 4 {
					slice21 = cord2[(L2+2)%7 : (L2+3)%7]
				}
				var slice22 [][]float64
				if L2 < 6 {
					slice22 = cord2[L2 : L2+2]
				} else if L2 == 6 {
					slice22t1 := cord2[L2:]
					slice22t2 := cord2[:(L2+2)%7]
					slice22 = append(slice22, slice22t1...)
					slice22 = append(slice22, slice22t2...)
				}
				slice2 = append(slice2, slice21...)
				slice2 = append(slice2, slice22...)
				log.Println("slice2=", slice2)

				// ５角形
				var slice31 [][]float64
				if L1 < 5 {
					slice31 = cord2[L1 : L1+3]
				} else if L1 == 5 {
					slice31t1 := cord2[L1 : L1+2]
					slice31t2 := cord2[(L1+2)%7]
					slice31 = append(slice31, slice31t1...)
					slice31 = append(slice31, slice31t2)
				} else if L1 > 5 {
					slice31t1 := cord2[L1]
					slice31t2 := cord2[:(L1+3)%7]
					slice31 = append(slice31, slice31t1)
					slice31 = append(slice31, slice31t2...)
				}
				var slice32 [][]float64
				if L1 < 2 {
					slice32 = cord2[(L1+4)%7 : (L1+6)%7]
				} else if L1 == 2 {
					slice32t1 := cord2[(L1+4)%7:]
					slice32t2 := cord2[:(L1+6)%7]
					slice32 = append(slice32, slice32t1...)
					slice32 = append(slice32, slice32t2...)
				} else if L1 > 2 {
					slice32 = cord2[(L1+4)%7 : (L1+6)%7]
				}
				slice3 = append(slice3, slice31...)
				slice3 = append(slice3, slice32...)
				log.Println("slice3=", slice3)

				deg5 := PentaDeg(slice3)
				slice3, roof5 = PentaNode(deg5, slice3)

				type1L = "tri"
				type2L = "tri"
				type3L = roof5
				story = append(story, 2)
				story = append(story, 2)
				story = append(story, 2)

			} else if (L2-L1+7)%7 == 5 {
				// L点が２つ離れた王冠型−2
				log.Println("２つの三角形と１つの５角形に分割する-2")

				// 三角形・片流れ屋根
				slice11 := cord2[L2 : L2+1]
				var slice12 [][]float64
				if L2 < 1 {
					slice12 = cord2[L2+5:]
				} else if L2 == 1 {
					slice12t1 := cord2[L2+5:]
					slice12t2 := cord2[:(L2+7)%7]
					slice12 = append(slice12, slice12t1...)
					slice12 = append(slice12, slice12t2...)
				} else if L2 > 1 {
					slice12 = cord2[(L2+5)%7 : (L2+7)%7]
				}
				slice1 = append(slice1, slice11...)
				slice1 = append(slice1, slice12...)
				log.Println("slice1=", slice1)

				// 三角形・片流れ屋根
				var slice21 [][]float64
				if L1 < 5 {
					slice21 = cord2[L1+2 : L1+3]
				} else if L1 > 4 {
					slice21 = cord2[(L1+2)%7 : (L1+3)%7]
				}
				var slice22 [][]float64
				if L1 < 6 {
					slice22 = cord2[L1 : L1+2]
				} else if L1 == 6 {
					slice22t1 := cord2[L1:]
					slice22t2 := cord2[:(L1+2)%7]
					slice22 = append(slice22, slice22t1...)
					slice22 = append(slice22, slice22t2...)
				}
				slice2 = append(slice2, slice21...)
				slice2 = append(slice2, slice22...)
				log.Println("slice2=", slice2)

				// ５角形
				var slice31 [][]float64
				if L2 < 5 {
					slice31 = cord2[L2 : L2+3]
				} else if L2 == 5 {
					slice31t1 := cord2[L2 : L2+2]
					slice31t2 := cord2[(L2+2)%7]
					slice31 = append(slice31, slice31t1...)
					slice31 = append(slice31, slice31t2)
				} else if L2 > 5 {
					slice31t1 := cord2[L2]
					slice31t2 := cord2[:(L2+3)%7]
					slice31 = append(slice31, slice31t1)
					slice31 = append(slice31, slice31t2...)
				}
				var slice32 [][]float64
				if L2 < 2 {
					slice32 = cord2[(L2+4)%7 : (L2+6)%7]
				} else if L2 == 2 {
					slice32t1 := cord2[(L2+4)%7:]
					slice32t2 := cord2[:(L2+6)%7]
					slice32 = append(slice32, slice32t1...)
					slice32 = append(slice32, slice32t2...)
				} else if L2 > 2 {
					slice32 = cord2[(L2+4)%7 : (L2+6)%7]
				}
				slice3 = append(slice3, slice31...)
				slice3 = append(slice3, slice32...)
				log.Println("slice3=", slice3)

				deg5 := PentaDeg(slice3)
				slice3, roof5 = PentaNode(deg5, slice3)

				type1L = "tri"
				type2L = "tri"
				type3L = roof5
				story = append(story, 2)
				story = append(story, 2)
				story = append(story, 2)

			} else if ((L1 < L2) && (L2-L1 == 3)) || ((L1 > L2) && (L1-L2 == 4)) {
				// L点が３つ離れたイカ型−１
				log.Println("１つの５角形と１つの四角形に分割する-1")

				// ５角形
				if L2 < 3 {
					slice1 = cord2[L2 : L2+5]
				} else if L2 > 2 {
					slice11 := cord2[L2:]
					slice12 := cord2[:L2-2]
					slice1 = append(slice1, slice11...)
					slice1 = append(slice1, slice12...)
				}
				log.Println("slice1=", slice1)

				deg5 := PentaDeg(slice1)
				slice1, roof5 = PentaNode(deg5, slice1)

				// 四角形・切妻屋根
				slice21 := cord2[L2 : L2+1]
				var slice22 [][]float64
				if L2 < 1 {
					slice22 = cord2[L2+4:]
				} else if L2 > 0 && L2 < 3 {
					slice22t1 := cord2[L2+4:]
					slice22t2 := cord2[:L2]
					slice22 = append(slice22, slice22t1...)
					slice22 = append(slice22, slice22t2...)
				} else if L2 > 2 {
					slice22 = cord2[L2-3 : L2]
				}
				slice2 = append(slice2, slice21...)
				slice2 = append(slice2, slice22...)
				log.Println("slice2=", slice2)

				// // 作成した５角形にL点が含まれるとうまく屋根が掛からない
				// // その場合は変形した四角形の切妻屋根とする
				// n := len(slice2)
				// ext, _, _ := TriVert(n, slice2)
				// newslice := make([][]float64, 0, 5)
				// for e, v := range ext {
				// 	if v <= 0.0 {
				// 		newslice = append(newslice, slice2[e])
				// 	}
				// }
				// m := len(newslice)
				// if m == 5 {
				// 	newpenta, yane := PentaNode(deg2, newslice)
				// 	slice2 = newpenta
				// 	log.Println("slice2=", slice2)
				// 	type2L = yane
				// } else if m == 4 {
				// 	slice2 = newslice
				// 	type2L = "kiri"
				// }
				// log.Println("slice2=", slice2)

				// ダミー三角形
				midcnt := (L2 + 2) % 7

				slice31 := cord2[L1 : L1+1]
				slice32 := cord2[midcnt : midcnt+1]
				slice33 := cord2[L2 : L2+1]
				slice3 = append(slice3, slice31...)
				slice3 = append(slice3, slice32...)
				slice3 = append(slice3, slice33...)
				log.Println("slice3=", slice3)

				type1L = roof5
				type2L = "kiri"
				type3L = "flat"
				story = append(story, 2)
				story = append(story, 2)
				story = append(story, 1)

			} else if ((L1 > L2) && (L1-L2 == 3)) || ((L1 < L2) && (L2-L1 == 4)) {
				// L点が３つ離れたイカ型−１
				log.Println("１つの５角形と１つの四角形に分割する-2")

				// ５角形
				if L1 < 3 {
					slice1 = cord2[L1 : L1+5]
				} else if L1 > 2 {
					slice11 := cord2[L1:]
					slice12 := cord2[:L1-2]
					slice1 = append(slice1, slice11...)
					slice1 = append(slice1, slice12...)
				}
				log.Println("slice1=", slice1)

				deg5 := PentaDeg(slice1)
				slice1, roof5 = PentaNode(deg5, slice1)

				// 四角形・切妻屋根
				slice21 := cord2[L1 : L1+1]
				var slice22 [][]float64
				if L1 < 1 {
					slice22 = cord2[L1+4:]
				} else if L1 > 0 && L1 < 3 {
					slice22t1 := cord2[L1+4:]
					slice22t2 := cord2[:L1]
					slice22 = append(slice22, slice22t1...)
					slice22 = append(slice22, slice22t2...)
				} else if L1 > 2 {
					slice22 = cord2[L1-3 : L1]
				}
				slice2 = append(slice2, slice21...)
				slice2 = append(slice2, slice22...)
				log.Println("slice2=", slice2)

				// // 作成した５角形にL点が含まれるとうまく屋根が掛からない
				// // その場合は変形した四角形の切妻屋根とする
				// n := len(slice2)
				// ext, _, _ := TriVert(n, slice2)
				// newslice := make([][]float64, 0, 5)
				// for e, v := range ext {
				// 	if v <= 0.0 {
				// 		newslice = append(newslice, slice2[e])
				// 	}
				// }
				// m := len(newslice)
				// if m == 5 {
				// 	newpenta, yane := PentaNode(deg2, newslice)
				// 	slice2 = newpenta
				// 	log.Println("slice2=", slice2)
				// 	type2L = yane
				// } else if m == 4 {
				// 	slice2 = newslice
				// 	type2L = "kiri"
				// }
				// log.Println("slice2=", slice2)

				// ダミー三角形
				midcnt := (L1 + 2) % 7

				slice31 := cord2[L1 : L1+1]
				slice32 := cord2[midcnt : midcnt+1]
				slice33 := cord2[L2 : L2+1]
				slice3 = append(slice3, slice31...)
				slice3 = append(slice3, slice32...)
				slice3 = append(slice3, slice33...)
				log.Println("slice3=", slice3)

				type1L = roof5
				type2L = "kiri"
				type3L = "flat"
				story = append(story, 2)
				story = append(story, 2)
				story = append(story, 1)
			}

		} else if strings.Count(lrtxt, "L") == 1 {
			// L1点から伸ばした線が対向する辺に直交する場合
			log.Println("直交する点で四角形と５角形に分割する")
			// L1点の頂点番号を確認する
			var num int
			for LRkey := range order {
				log.Println("LRkey=", LRkey) // Ctrl+/
				if LRkey == "L1" {
					num = order[LRkey]      // 頂点番号
					log.Println("num", num) // Ctrl+/
				}
			}
			// L1点のX座標
			p := cord2[num][0]
			// L1点のY座標
			q := cord2[num][1]

			// 対向する辺は，L点から２つ目と３つ目の点で結ばれる線分
			// 対向する辺２−３の座標ペア
			taikoCord1 := make([][]float64, 2)
			numP2 := (num + 2) % 7
			taikoCord1[0] = cord2[numP2]
			numP3 := (num + 3) % 7
			taikoCord1[1] = cord2[numP3]
			// 対向する辺の直線の方程式
			line1 := LineEquat(taikoCord1[0][0], taikoCord1[0][1], taikoCord1[1][0], taikoCord1[1][1])
			a1 := line1["m"]
			b1 := line1["n"]
			// L1点から対向する辺に下ろした垂線の交点の座標
			x1 := (a1*(q-b1) + p) / (math.Pow(a1, 2) + 1)
			y1 := (a1*(a1*(q-b1)+p))/(math.Pow(a1, 2)+1) + b1
			D1 := []float64{x1, y1}
			log.Println("D1=", D1)
			// 垂線の交点が辺の上にあるかどうかチェックする
			chk1 := PointonLine(taikoCord1[0][0], taikoCord1[0][1], taikoCord1[1][0], taikoCord1[1][1], x1, y1)
			// perpen1 := [][]float64{{p, q}, {x1, y1}}
			// chk1 := PosLine2(taikoCord1, perpen1)
			// chk1 := PosLine(taikoCord1[1][0], taikoCord1[0][0], taikoCord1[1][1], taikoCord1[0][1], y1, x1)
			log.Println("chk1=", chk1)

			// もう一方の対向する辺は，L点から４つ目と５つ目の点で結ばれる線分
			// 対向する辺4-5の座標ペア
			taikoCord2 := make([][]float64, 2)
			numN4 := (num + 4) % 7
			taikoCord2[0] = cord2[numN4]
			numN5 := (num + 5) % 7
			taikoCord2[1] = cord2[numN5]
			// 対向する辺の直線の方程式
			line2 := LineEquat(taikoCord2[0][0], taikoCord2[0][1], taikoCord2[1][0], taikoCord2[1][1])
			a2 := line2["m"]
			b2 := line2["n"]
			// L1点から対向する辺に下ろした垂線の交点の座標
			x2 := (a2*(q-b2) + p) / (math.Pow(a2, 2) + 1)
			y2 := (a2*(a2*(q-b2)+p))/(math.Pow(a2, 2)+1) + b2
			D2 := []float64{x2, y2}
			log.Println("D2=", D2)
			// 垂線の交点が辺の上にあるかどうかチェックする
			chk2 := PointonLine(taikoCord2[0][0], taikoCord2[0][1], taikoCord2[1][0], taikoCord2[1][1], x2, y2)
			// perpen2 := [][]float64{{p, q}, {x2, y2}}
			// chk2 := PosLine2(taikoCord2, perpen2)
			// chk2 := PosLine(taikoCord2[1][0], taikoCord2[0][0], taikoCord2[1][1], taikoCord2[0][1], y2, x2)
			log.Println("chk2=", chk2)

			// さらにもう一方つ対向する辺は，L点から３つ目と４つ目の点で結ばれる線分
			// 対向する辺3-4の座標ペア
			taikoCord3 := make([][]float64, 2)
			numC3 := (num + 3) % 7
			taikoCord3[0] = cord2[numC3]
			numC4 := (num + 4) % 7
			taikoCord3[1] = cord2[numC4]
			// 対向する辺の直線の方程式
			line3 := LineEquat(taikoCord3[0][0], taikoCord3[0][1], taikoCord3[1][0], taikoCord3[1][1])
			a3 := line3["m"]
			b3 := line3["n"]
			// L1点から対向する辺に下ろした垂線の交点の座標
			x3 := (a3*(q-b3) + p) / (math.Pow(a3, 2) + 1)
			y3 := (a3*(a3*(q-b3)+p))/(math.Pow(a3, 2)+1) + b3
			D3 := []float64{x3, y3}
			log.Println("D3=", D3)
			// 垂線の交点が辺の上にあるかどうかチェックする
			chk3 := PointonLine(taikoCord3[0][0], taikoCord3[0][1], taikoCord3[1][0], taikoCord3[1][1], x3, y3)
			// perpen3 := [][]float64{{p, q}, {x3, y3}}
			// chk3 := PosLine2(taikoCord3, perpen3)
			// chk2 := PosLine(taikoCord2[1][0], taikoCord2[0][0], taikoCord2[1][1], taikoCord2[0][1], y2, x2)
			log.Println("chk3=", chk3)

			// 四角形aを分割する
			if chk1 {
				log.Println("四角形１を分割する")
				slice1 = append(slice1, D1)
				var slice12 [][]float64
				if num < 5 {
					slice12 = cord2[num : num+3]
				} else if num > 4 {
					slice12t1 := cord2[num:]
					slice12t2 := cord2[:(num+3)%7]
					slice12 = append(slice12, slice12t1...)
					slice12 = append(slice12, slice12t2...)
				}
				slice1 = append(slice1, slice12...)
				type1L = "kiri"
				story = append(story, 2)

				if chk2 {
					log.Println("四角形２を分割する")
					slice21 := cord2[num : num+1]
					slice2 = append(slice2, slice21...)
					slice2 = append(slice2, D2)
					var slice22 [][]float64
					if num < 1 {
						slice22 = cord2[(num+5)%7 : num+7]
					} else if num == 1 {
						slice22t1 := cord2[(num+5)%7:]
						slice22t2 := cord2[:(num+7)%7]
						slice22 = append(slice22, slice22t1...)
						slice22 = append(slice22, slice22t2...)
					} else if num > 1 {
						slice22 = cord2[(num+5)%7 : num]
					}
					slice2 = append(slice2, slice22...)
					type2L = "kiri"
					story = append(story, 2)

					log.Println("5角形を分割する")
					slice31 := cord2[num : num+1]
					slice3 = append(slice3, slice31...)
					slice3 = append(slice3, D1)
					var slice32 [][]float64
					if num < 3 {
						slice32 = cord2[num+3 : num+5]
						slice3 = append(slice3, slice32...)
					} else if num == 3 {
						slice32t1 := cord2[num+3:]
						slice32t2 := cord2[:(num+5)%7]
						slice32 = append(slice32, slice32t1...)
						slice32 = append(slice32, slice32t2...)
						slice3 = append(slice3, slice32...)
					} else if num > 3 {
						slice32 = cord2[(num+3)%7 : (num+5)%7]
						slice3 = append(slice3, slice32...)
					}
					slice3 = append(slice3, D2)
					log.Println("slice3=", slice3)

					deg5 := PentaDeg(slice3)
					slice3, roof5 = PentaNode(deg5, slice3)
					type3L = roof5
					story = append(story, 2)

				} else if chk3 {
					log.Println("四角形３を分割する")
					slice21 := cord2[num : num+1]
					slice2 = append(slice2, slice21...)
					slice2 = append(slice2, D1)
					// slice22 := cord2[(num+3)%7 : (num+4)%7]
					var slice22 [][]float64
					if num < 4 {
						slice22 = cord2[num+3 : num+4]
						slice2 = append(slice2, slice22...)
					} else if num > 3 {
						slice22 = cord2[(num+3)%7 : (num+4)%7]
						slice2 = append(slice2, slice22...)
					}
					slice2 = append(slice2, D3)
					type2L = "kiri"
					story = append(story, 2)

					log.Println("５角形を分割する")
					slice31 := cord2[num : num+1]
					slice3 = append(slice3, slice31...)
					slice3 = append(slice3, D3)
					var slice32 [][]float64
					if num < 1 {
						slice32 = cord2[num+4 : num+7]
						slice3 = append(slice3, slice32...)
					} else if num > 0 && num < 3 {
						slice32t1 := cord2[num+4:]
						slice32t2 := cord2[:num]
						slice32 = append(slice32, slice32t1...)
						slice32 = append(slice32, slice32t2...)
						slice3 = append(slice3, slice32...)
					} else if num > 2 {
						slice32 = cord2[(num+4)%7 : num]
						slice3 = append(slice3, slice32...)
					}

					deg5 := PentaDeg(slice3)
					slice3, roof5 = PentaNode(deg5, slice3)
					type3L = roof5
					story = append(story, 2)
				} else {
					chk = false
				}

			} else if chk2 {
				log.Println("四角形２を分割する")
				slice11 := cord2[num : num+1]
				slice1 = append(slice1, slice11...)
				slice1 = append(slice1, D2)
				var slice12 [][]float64
				if num < 1 {
					slice12 = cord2[num+5 : num+7]
				} else if num == 1 {
					slice12t1 := cord2[num+5:]
					slice12t2 := cord2[:num]
					slice12 = append(slice12, slice12t1...)
					slice12 = append(slice12, slice12t2...)
				} else if num > 1 {
					slice12 = cord2[(num+5)%7 : num]
				}
				slice1 = append(slice1, slice12...)
				type1L = "kiri"
				story = append(story, 2)

				if chk3 {
					log.Println("5角形を分割する")
					slice21 := cord2[num : num+1]
					slice2 = append(slice2, slice21...)
					var slice22 [][]float64
					if num < 4 {
						slice22 = cord2[num+1 : num+4]
						slice2 = append(slice2, slice22...)
					} else if num > 3 && num < 6 {
						slice22t1 := cord2[num+1:]
						slice22t2 := cord2[:(num+4)%7]
						slice22 = append(slice22, slice22t1...)
						slice22 = append(slice22, slice22t2...)
						slice2 = append(slice2, slice22...)
					} else if num > 5 {
						slice22 = cord2[(num+1)%7 : (num+4)%7]
						slice2 = append(slice2, slice22...)
					}
					slice2 = append(slice2, D2)

					deg5 := PentaDeg(slice2)
					slice2, roof5 = PentaNode(deg5, slice2)
					type2L = roof5
					story = append(story, 2)

					log.Println("四角形４を分割する")
					slice31 := cord2[num : num+1]
					slice3 = append(slice3, slice31...)
					slice3 = append(slice3, D3)
					var slice32 [][]float64
					if num < 3 {
						slice32 = cord2[num+4 : num+5]
						slice3 = append(slice3, slice32...)
					} else if num > 2 {
						slice32 = cord2[(num+4)%7 : (num+5)%7]
						slice3 = append(slice3, slice32...)
					}
					slice3 = append(slice3, D2)
					type3L = "flat"
					story = append(story, 2)
				} else {
					chk = false
				}
			} else {
				chk = false
			}
		}

	} else if !strings.Contains(lrtxt, "L") {
		chk = false
	}

	log.Println("type1L=", type1L)
	log.Println("type2L=", type2L)
	log.Println("type3L=", type3L)

	return slice1, slice2, slice3, type1L, type2L, type3L, story, chk
}
