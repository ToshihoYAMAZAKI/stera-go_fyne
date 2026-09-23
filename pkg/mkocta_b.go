package pkg

// MkoctaB はareaB を除いた8角形を作る．
func MkoctaB(cordz [][]float64, order2 map[string]int, keyList []string, nodDec,
	num0, d0Num int) (octa1name []string, octa1L [][]float64) {
	// 8角形を作る
	// keyListはインデックス0から順にＬ･Ｒ点が並んでいる．
	// num0とnum0-1，num0-2を削除する
	// d0Num, num0+1, num0+2, num0+3, num0+4, num0+5, num0+6, num0+7
	octa1name = append(octa1name, keyList[d0Num])
	for i := 1; i < 8; i++ {
		numi := (num0 + i) % nodDec
		octa1name = append(octa1name, keyList[numi])
	}
	for _, name := range octa1name {
		n := order2[name]
		octa1L = append(octa1L, cordz[n])
	}
	return octa1name, octa1L
}
