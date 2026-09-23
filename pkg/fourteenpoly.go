package pkg

import (
	"log"
	"strings"
)

// FortenPoly は14角形から四角形を分割して12角形以下の分割プログラムに渡す
func FortenPoly(lrPtn []string, lrIdx []int, deg2 []float64, cord2 [][]float64, order map[string]int) (slice1, slice2 [][]float64, result bool) {
	// 頂点並びからLRRLのパターンを捜す
	lrtxt := strings.Join(lrPtn, "")
	log.Println("lrtxt=", lrtxt)

	nod := len(cord2)
	var idx int
	// var slice1 [][]float64
	result = true

	if strings.Contains(lrtxt, "LRRRL") {
		log.Println("lrtxt include LRRRL")
		cnt := strings.Count(lrtxt, "LRRRL")
		log.Println("cnt=", cnt)

		var degsum1 float64
		var degsum2 float64

		if cnt > 1 {
			// "LRRRL"が２つ以上
			ptn1 := strings.Index(lrtxt, "LRRRL")
			log.Println("ptn1=", ptn1)
			idx1 := lrIdx[ptn1]
			log.Println("idx1=", idx1)
			for i := 0; i < 3; i++ {
				degsum1 = degsum1 + deg2[idx1+i]
			}
			// 最も後ろの"LRRRL"
			ptn2 := strings.LastIndex(lrtxt, "LRRRL")
			log.Println("ptn2=", ptn2)
			idx2 := lrIdx[ptn2]
			log.Println("idx2=", idx2)
			for i := 0; i < 3; i++ {
				degsum2 = degsum2 + deg2[idx2+i]
			}
			// 内角の和が大きい方を採用する
			if degsum1 > degsum2 {
				idx = idx1
			} else if degsum1 < degsum2 {
				idx = idx2
			}

		} else if cnt == 1 {
			// "LRRRL"が１つ
			ptn0 := strings.Index(lrtxt, "LRRRL")
			log.Println("ptn0=", ptn0)
			idx0 := lrIdx[ptn0]
			log.Println("idx0=", idx0)
			for i := 0; i < 3; i++ {
				degsum1 = degsum1 + deg2[idx0+i]
			}
			idx = idx0
			// 右端の"LRRR"と左端の"L"で"LRRRL"が１つ
			if strings.Contains(lrtxt, "LRRR") {
				ptn4 := strings.LastIndex(lrtxt, "LRRR")
				log.Println("ptn4=", ptn4)
				if ptn4 == (nod - 4) {
					log.Println("lrtxt include LRRR")
					idx4 := lrIdx[ptn4]
					log.Println("idx4=", idx4)
					// １つ目の"LRRRL"と内角の和を比較
					for i := 0; i < 3; i++ {
						degsum2 = degsum2 + deg2[(idx4+i)%nod]
					}
					if degsum1 > degsum2 {
						idx = idx0
					} else if degsum1 < degsum2 {
						idx = idx4
					}
				} else {
					// TODO:
				}
			}
		}
		log.Println("idx=", idx)
		// ５角形を分割
		if idx < (nod - 4) {
			slice1 = cord2[idx : idx+5]
			log.Println("slice1=", slice1)
		} else if idx > (nod - 5) {
			slice11 := cord2[idx:]
			slice12 := cord2[:idx-(nod-5)]
			slice1 = append(slice1, slice11...)
			slice1 = append(slice1, slice12...)
			log.Println("slice1=", slice1)
		}
		// 残りの多角形を分割
		if idx < (nod - 4) {
			slice21 := cord2[(idx-(nod-4)+nod)%nod:]
			slice22 := cord2[:idx+1]
			slice2 = append(slice2, slice21...)
			slice2 = append(slice2, slice22...)
			log.Println("slice2=", slice2)
		} else if idx > (nod - 5) {
			slice2 = cord2[idx-(nod-4) : idx]
			log.Println("slice2=", slice2)
		}

	} else if strings.Contains(lrtxt, "LRRR") {
		// 右端の"LRRR"と左端の"L"で"LRRRL"が１つ
		ptn3 := strings.LastIndex(lrtxt, "LRRR")
		log.Println("ptn3=", ptn3)
		if ptn3 == (nod - 4) {
			log.Println("lrtxt include LRRR")
			idx3 := lrIdx[ptn3]
			log.Println("idx3=", idx3)
			idx = idx3
		} else {
			// TODO:
			result = false
			return
		}

		log.Println("idx=", idx)
		// ５角形を分割
		if idx < (nod - 4) {
			slice1 = cord2[idx : idx+5]
			log.Println("slice1=", slice1)
		} else if idx > (nod - 5) {
			slice11 := cord2[idx:]
			slice12 := cord2[:idx-(nod-5)]
			slice1 = append(slice1, slice11...)
			slice1 = append(slice1, slice12...)
			log.Println("slice1=", slice1)
		}
		// 残りの多角形を分割
		if idx < (nod - 4) {
			slice21 := cord2[(idx-(nod-4)+nod)%nod:]
			slice22 := cord2[:idx+1]
			slice2 = append(slice2, slice21...)
			slice2 = append(slice2, slice22...)
			log.Println("slice2=", slice2)
		} else if idx > (nod - 5) {
			slice2 = cord2[idx-(nod-4) : idx+1]
			log.Println("slice2=", slice2)
		}

	} else if strings.Contains(lrtxt, "LRRL") {
		log.Println("lrtxt include LRRL")
		cnt := strings.Count(lrtxt, "LRRL")
		log.Println("cnt=", cnt)

		// LRRLパターンが複数ある場合，四角形分割による頂点削除後に内角が180度前後になる方を選ぶ
		if cnt > 1 {
			ptn1 := strings.Index(lrtxt, "LRRL")
			log.Println("ptn1=", ptn1)
			idx1 := lrIdx[ptn1]
			log.Println("idx1=", idx1)

			// 四角形分割による頂点削除後の内角
			x11 := cord2[idx1][0]
			y11 := cord2[idx1][1]
			x11f := cord2[(idx1+3)%14][0]
			y11f := cord2[(idx1+3)%14][1]
			x11b := cord2[(idx1-1+14)%14][0]
			y11b := cord2[(idx1-1+14)%14][1]

			deg1f := TriAngle(x11, y11, x11f, y11f, x11b, y11b)
			// log.Println("deg1f=", deg1f)

			x12 := cord2[(idx1+3)%14][0]
			y12 := cord2[(idx1+3)%14][1]
			x12f := cord2[(idx1+4)%14][0]
			y12f := cord2[(idx1+4)%14][1]
			x12b := cord2[(idx1)%14][0]
			y12b := cord2[(idx1)%14][1]

			deg1b := TriAngle(x12, y12, x12f, y12f, x12b, y12b)
			// log.Println("deg1b=", deg1b)

			deg1sum := deg1f + deg1b
			log.Println("deg1sum=", deg1sum)

			ptn2 := strings.LastIndex(lrtxt, "LRRL")
			log.Println("ptn2=", ptn2)
			idx2 := lrIdx[ptn2]
			log.Println("idx2=", idx2)

			// 四角形分割による頂点削除後の内角
			x21 := cord2[idx2][0]
			y21 := cord2[idx2][1]
			x21f := cord2[(idx2+3)%14][0]
			y21f := cord2[(idx2+3)%14][1]
			x21b := cord2[(idx2-1+14)%14][0]
			y21b := cord2[(idx2-1+14)%14][1]

			deg2f := TriAngle(x21, y21, x21f, y21f, x21b, y21b)
			// log.Println("deg2f=", deg2f)

			x22 := cord2[(idx2+3)%14][0]
			y22 := cord2[(idx2+3)%14][1]
			x22f := cord2[(idx2+4)%14][0]
			y22f := cord2[(idx2+4)%14][1]
			x22b := cord2[(idx2)%14][0]
			y22b := cord2[(idx2)%14][1]

			deg2b := TriAngle(x22, y22, x22f, y22f, x22b, y22b)
			// log.Println("deg2b=", deg2b)

			deg2sum := deg2f + deg2b
			log.Println("deg2sum=", deg2sum)

			if deg1sum > deg2sum {
				idx = idx1
			} else if deg1sum < deg2sum {
				idx = idx2
			}

		} else if cnt == 1 {
			ptn3 := strings.Index(lrtxt, "LRRL")
			log.Println("ptn3=", ptn3)
			idx3 := lrIdx[ptn3]
			log.Println("idx3=", idx3)

			// 四角形分割による頂点削除後の内角
			x31 := cord2[idx3][0]
			y31 := cord2[idx3][1]
			x31f := cord2[(idx3+3)%14][0]
			y31f := cord2[(idx3+3)%14][1]
			x31b := cord2[(idx3-1+14)%14][0]
			y31b := cord2[(idx3-1+14)%14][1]

			deg3f := TriAngle(x31, y31, x31f, y31f, x31b, y31b)
			// log.Println("deg3f=", deg3f)

			x32 := cord2[(idx3+3)%14][0]
			y32 := cord2[(idx3+3)%14][1]
			x32f := cord2[(idx3+4)%14][0]
			y32f := cord2[(idx3+4)%14][1]
			x32b := cord2[(idx3)%14][0]
			y32b := cord2[(idx3)%14][1]

			deg3b := TriAngle(x32, y32, x32f, y32f, x32b, y32b)
			// log.Println("deg3b=", deg3b)

			deg3sum := deg3f + deg3b
			log.Println("deg3sum=", deg3sum)

			if strings.Contains(lrtxt, "LRR") {
				ptn4 := strings.LastIndex(lrtxt, "LRR")
				log.Println("ptn4=", ptn4)
				if ptn4 == 11 {
					log.Println("lrtxt include LRR")
					idx4 := lrIdx[ptn4]
					log.Println("idx4=", idx4)

					// 四角形分割による頂点削除後の内角
					x41 := cord2[idx4][0]
					y41 := cord2[idx4][1]
					x41f := cord2[(idx4+3)%14][0]
					y41f := cord2[(idx4+3)%14][1]
					x41b := cord2[(idx4-1+14)%14][0]
					y41b := cord2[(idx4-1+14)%14][1]

					deg4f := TriAngle(x41, y41, x41f, y41f, x41b, y41b)
					// log.Println("deg4f=", deg4f)

					x42 := cord2[(idx4+3)%14][0]
					y42 := cord2[(idx4+3)%14][1]
					x42f := cord2[(idx4+4)%14][0]
					y42f := cord2[(idx4+4)%14][1]
					x42b := cord2[(idx4)%14][0]
					y42b := cord2[(idx4)%14][1]

					deg4b := TriAngle(x42, y42, x42f, y42f, x42b, y42b)
					// log.Println("deg4b=", deg4b)

					deg4sum := deg4f + deg4b
					log.Println("deg4sum=", deg4sum)

					if deg3sum > deg4sum {
						idx = idx3
					} else if deg3sum < deg4sum {
						idx = idx4
					}
				} else {
					idx = idx3
				}
			}
		}

		log.Println("idx=", idx)

		if idx < 11 {
			slice11 := cord2[idx+3 : idx+4]
			slice12 := cord2[idx : idx+3]
			slice1 = append(slice1, slice11...)
			slice1 = append(slice1, slice12...)
			// slice1 = cord2[idx : idx+4]
			// log.Println("slice1=", slice1)
		} else if idx > 10 {
			slice11 := cord2[(idx+3)%14 : (idx+4)%14]
			var slice12 [][]float64
			slice12t1 := cord2[idx:]
			slice12t2 := cord2[:(idx+3)%14]
			slice12 = append(slice12, slice12t1...)
			slice12 = append(slice12, slice12t2...)
			slice1 = append(slice1, slice11...)
			slice1 = append(slice1, slice12...)
			// log.Println("slice1=", slice1)
		}
		// slice1t1 := slice1[3:4]
		// slice1t2 := slice1[0:3]
		// slice1 = append(slice1, slice1t1...)
		// slice1 = append(slice1, slice1t2...)

		if idx < 11 {
			slice21 := cord2[idx+3:]
			slice22 := cord2[:idx+1]
			slice2 = append(slice2, slice21...)
			slice2 = append(slice2, slice22...)
			log.Println("slice2=", slice2)
		} else if idx > 10 {
			slice2 = cord2[(idx+3)%14 : idx+1]
			log.Println("slice2=", slice2)
		}

	} else if strings.Contains(lrtxt, "LRR") {
		// 右端の"LRR"と左端の"L"で"LRRL"が１つ
		ptn5 := strings.LastIndex(lrtxt, "LRR")
		log.Println("ptn5=", ptn5)
		if ptn5 == (nod - 3) {
			log.Println("lrtxt include LRRL")
			idx5 := lrIdx[ptn5]
			log.Println("idx3=", idx5)
			idx = idx5
		} else {
			// TODO:
			result = false
			return
		}

		log.Println("idx=", idx)

		if idx < 11 {
			slice11 := cord2[idx+3 : idx+4]
			slice12 := cord2[idx : idx+3]
			slice1 = append(slice1, slice11...)
			slice1 = append(slice1, slice12...)
		} else if idx > 10 {
			slice11 := cord2[(idx+3)%14 : (idx+4)%14]
			slice12 := cord2[idx : (idx+3)%14]
			slice1 = append(slice1, slice11...)
			slice1 = append(slice1, slice12...)
		}

		if idx < 11 {
			slice21 := cord2[idx+3:]
			slice22 := cord2[:idx+1]
			slice2 = append(slice2, slice21...)
			slice2 = append(slice2, slice22...)
			log.Println("slice2=", slice2)
		} else if idx > 10 {
			slice2 = cord2[(idx+3)%14 : idx+1]
			log.Println("slice2=", slice2)
		}

	} else {
		// TODO:
		result = false
	}

	extxhk := ExtChk(len(slice1), slice1)
	if !extxhk {
		result = false
	}
	log.Println("result=", result)
	return slice1, slice2, result
}
