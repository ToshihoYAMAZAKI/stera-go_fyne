package pkg

// MkrectA3 はareaB ∉ areaA の四角形を作る．
func MkrectA3(cordz [][]float64, order2 map[string]int, keyList []string, nod,
	num0, d0Num int) (rect1name []string, rect4L [][]float64) {
	// 四角形を作る
	// D1点とL3/L4点の３つ前のR点と２つ前のR点と前のR点: d0Num, num0-3, num0-2, num0-1
	num_1 := (num0 - 1 + nod) % nod
	rect1name = append(rect1name, keyList[num_1])
	rect1name = append(rect1name, keyList[d0Num])
	num_3 := (num0 - 3 + nod) % nod
	rect1name = append(rect1name, keyList[num_3])
	num_2 := (num0 - 2 + nod) % nod
	rect1name = append(rect1name, keyList[num_2])

	for _, name := range rect1name {
		n := order2[name]
		rect4L = append(rect4L, cordz[n])
	}
	return rect1name, rect4L
}
