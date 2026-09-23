package pkg

// MkdecaA はareaA を除いた10角形を作る．
func MkdecaA(cordz [][]float64, order2 map[string]int, keyList []string, nod,
	num0, d0Num int) (deca1name []string, deca1L [][]float64) {
	// 10角形を作る
	// d0Num, num0+3, num0+4, num0+5, num0+6, num0+7, num0+8, num0+9, num0+10, num0+11
	deca1name = append(deca1name, keyList[d0Num])
	for i := 3; i < 12; i++ {
		numi := (num0 + i) % nod
		deca1name = append(deca1name, keyList[numi])
	}
	for _, name := range deca1name {
		n := order2[name]
		deca1L = append(deca1L, cordz[n])
	}
	return deca1name, deca1L
}
