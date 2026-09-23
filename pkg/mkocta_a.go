package pkg

// MkoctaA はareaA を除いた8角形を作る．
func MkoctaA(cordz [][]float64, order2 map[string]int, keyList []string, nodDec,
	num0, d0Num int) (octa1name []string, octa1L [][]float64) {
	// 8角形を作る
	// d0Num, num0+3, num0+4, num0+5, num0+6, num0+7, num0+8, num0+9
	octa1name = append(octa1name, keyList[d0Num])
	for i := 3; i < 10; i++ {
		numi := (num0 + i) % nodDec
		octa1name = append(octa1name, keyList[numi])
	}
	for _, name := range octa1name {
		n := order2[name]
		octa1L = append(octa1L, cordz[n])
	}
	return octa1name, octa1L
}
