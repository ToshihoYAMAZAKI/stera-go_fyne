package pkg

// MkrectA はareaA の四角形を作る．
func MkrectA(cordz [][]float64, order2 map[string]int, keyList []string, nod,
	num0, d0Num int) (rect1name []string, rect4L [][]float64) {
	// 四角形を作る
	// D1点とL3/L4点と次のR点とその次のR点: d0Num, num0, num0+1, num0+2
	rect1name = append(rect1name, keyList[d0Num])
	rect1name = append(rect1name, keyList[num0])
	num1 := (num0 + 1) % nod
	rect1name = append(rect1name, keyList[num1])
	num2 := (num0 + 2) % nod
	rect1name = append(rect1name, keyList[num2])
	for _, name := range rect1name {
		n := order2[name]
		rect4L = append(rect4L, cordz[n])
	}
	return rect1name, rect4L
}
