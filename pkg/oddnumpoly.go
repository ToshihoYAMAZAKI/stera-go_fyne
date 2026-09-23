package pkg

import (
	"log"
	"strings"
)

// OddPoly は13角形以下の奇数多角形から5角形を分割する
func OddPoly(lrPtn []string, lrIdx []int, deg2 []float64, cord2 [][]float64, order map[string]int) (slice1, slice2 [][]float64, result bool) {
	nod := len(cord2)

	// 頂点並びからLRRRLのパターンを捜す
	lrtxt := strings.Join(lrPtn, "")
	log.Println("lrtxt=", lrtxt)

	// var slice1 [][]float64
	// var slice2 [][]float64
	var idx int
	result = true

	// ９角形の場合は５角形"LRRRL"を分割することで６角形が残る
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
			for i := 1; i < 4; i++ {
				degsum1 = degsum1 + deg2[idx1+i]
			}
			// 最も後ろの"LRRRL"
			ptn2 := strings.LastIndex(lrtxt, "LRRRL")
			log.Println("ptn2=", ptn2)
			idx2 := lrIdx[ptn2]
			log.Println("idx2=", idx2)
			for i := 1; i < 4; i++ {
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
			for i := 1; i < 4; i++ {
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
					for i := 1; i < 4; i++ {
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
		// 11角形の場合は四角形"LRRRL"を分割することで９角形が残る
		log.Println("lrtxt include LRRL")
		cnt := strings.Count(lrtxt, "LRRL")
		log.Println("cnt=", cnt)

		var degsum1 float64
		var degsum2 float64

		if cnt > 1 {
			// "LRRL"が２つ以上
			ptn1 := strings.Index(lrtxt, "LRRL")
			log.Println("ptn1=", ptn1)
			idx1 := lrIdx[ptn1]
			log.Println("idx1=", idx1)
			for i := 1; i < 3; i++ {
				degsum1 = degsum1 + deg2[idx1+i]
			}
			// 最も後ろの"LRRL"
			ptn2 := strings.LastIndex(lrtxt, "LRRL")
			log.Println("ptn2=", ptn2)
			idx2 := lrIdx[ptn2]
			log.Println("idx2=", idx2)
			for i := 1; i < 3; i++ {
				degsum2 = degsum2 + deg2[idx2+i]
			}
			// 内角の和が大きい方を採用する
			if degsum1 > degsum2 {
				idx = idx1
			} else if degsum1 < degsum2 {
				idx = idx2
			}

		} else if cnt == 1 {
			// "LRRL"が１つ
			ptn0 := strings.Index(lrtxt, "LRRL")
			log.Println("ptn0=", ptn0)
			idx0 := lrIdx[ptn0]
			log.Println("idx0=", idx0)
			for i := 1; i < 3; i++ {
				degsum1 = degsum1 + deg2[(idx0+i)%nod]
			}
			idx = idx0
			// 右端の"LRRR"と左端の"L"で"LRRRL"が１つ
			if strings.Contains(lrtxt, "LRR") {
				ptn4 := strings.LastIndex(lrtxt, "LRR")
				log.Println("ptn4=", ptn4)
				if ptn4 == (nod - 3) {
					log.Println("lrtxt include LRRL")
					idx4 := lrIdx[ptn4]
					log.Println("idx4=", idx4)
					// １つ目の"LRRRL"と内角の和を比較
					for i := 1; i < 3; i++ {
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
		// 四角形を分割
		if idx < (nod - 3) {
			slice1 = cord2[idx : idx+4]
			log.Println("slice1=", slice1)
		} else if idx > (nod - 4) {
			slice11 := cord2[idx:]
			slice12 := cord2[:idx-(nod-4)]
			slice1 = append(slice1, slice11...)
			slice1 = append(slice1, slice12...)
			log.Println("slice1=", slice1)
		}
		// 残りの多角形を分割
		if idx < (nod - 3) {
			slice21 := cord2[idx+3:]
			slice22 := cord2[:idx+1]
			slice2 = append(slice2, slice21...)
			slice2 = append(slice2, slice22...)
			log.Println("slice2=", slice2)
		} else if idx > (nod - 4) {
			slice2 = cord2[idx-(nod-3) : idx+1]
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
