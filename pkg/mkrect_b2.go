package pkg

// MkrectB2 はareaA ∉ areaB の四角形を作る．
func MkrectB2(cordz [][]float64, order2 map[string]int, keyList []string, nod,
	num0, d0Num int) (rect1name []string, rect4L [][]float64) {
	// 四角形を作る
	// D2点とL3/L4点の次のR点とその次のR点さらにその次のR点: d0Num, num0+1, num0+2，num0+3
	rect1name = append(rect1name, keyList[d0Num])
	num1 := (num0 + 1) % nod
	rect1name = append(rect1name, keyList[num1])
	num2 := (num0 + 2) % nod
	rect1name = append(rect1name, keyList[num2])
	num3 := (num0 + 3) % nod
	rect1name = append(rect1name, keyList[num3])
	for _, name := range rect1name {
		n := order2[name]
		rect4L = append(rect4L, cordz[n])
	}
	return rect1name, rect4L
}
