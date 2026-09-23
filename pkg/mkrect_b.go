package pkg

// MkrectB はareaB の四角形を作る．
func MkrectB(cordz [][]float64, order2 map[string]int, keyList []string, nod,
	num0, d0Num int) (rect1name []string, rect4L [][]float64) {
	// 四角形を作る
	// L3/L4点とD2点と２つ前のR点と前のR点: num0, d0Num, num0-2, num0-1
	rect1name = append(rect1name, keyList[num0])
	rect1name = append(rect1name, keyList[d0Num])
	num_2 := (num0 - 2 + nod) % nod
	rect1name = append(rect1name, keyList[num_2])
	num_1 := (num0 - 1 + nod) % nod
	rect1name = append(rect1name, keyList[num_1])
	for _, name := range rect1name {
		n := order2[name]
		rect4L = append(rect4L, cordz[n])
	}
	return rect1name, rect4L
}
