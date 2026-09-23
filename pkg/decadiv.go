package pkg

import (
	"log"
	"strings"
)

// DecaDiv は10角形を４つの四角形に分割する
func DecaDiv(cord2 [][]float64, order map[string]int, lrPtn []string,
	lrIdx []int) (rect1L [][]float64, rect2L [][]float64, rect3L [][]float64,
	rect4L [][]float64, story []int, yane []string) {
	var hex1L [][]float64
	var octa1name []string
	var octa1L [][]float64

	// 頂点データ数の確認
	nodDec := len(cord2)
	if nodDec != 10 {
		// TODO:関数から戻る
		return
	}
	d0Num := nodDec
	var num0 int

	// LR並びの確認  L点から始まっていなければエラー
	if lrPtn[0] != "L" {
		// TODO:関数から戻る
		return
	}
	// 検索用にLR並びから半角スペースを除く
	lrjoin := strings.Join(lrPtn, "")
	// log.Println("lrjoin=", lrjoin)

	// 分割線に対抗する辺の初期値（areaAとareaBが重複しない）
	var tp string

	switch lrjoin {
	// １－(1) areaA ∉ areaB
	case "LLRLRRRRRR", "LRLRRRRRRL", "LRRRRRRLLR":
		// L3点のインデックスがnum0
		if lrPtn[1] == "L" && lrPtn[3] == "L" {
			num0 = lrIdx[3]
		} else if lrPtn[2] == "L" && lrPtn[9] == "L" {
			num0 = lrIdx[2]
		} else if lrPtn[7] == "L" && lrPtn[8] == "L" {
			num0 = lrIdx[0]
		}
		// num0はL3点のインデックス番号
		// 10角形を１つの四角形と１つの8角形に分割する
		// 分割線bが対向する辺はL3点から３つ目と４つ目のR点が作る辺
		tp = "b2"
		cordz, order2, keyList, areaTag := DecaPrepro(num0, cord2, nodDec, order, tp)
		// 大きい耳となる方の四角形を分割し，残りで8角形を作る
		if areaTag == "areaA" {
			// areaAはareaBより必ず小さい
			// TODO:
		} else if areaTag == "areaB" {
			// 四角形を作る
			rect1name, _ := MkrectB2(cordz, order2, keyList, nodDec, num0, d0Num)
			// // rect1name = ['D2', 'R2', 'R3', 'R4']

			log.Println("rect1name", rect1name)
			rect4L = MakeRectList(cordz, order2, rect1name)
			log.Println("rect4L=", rect4L)
			// 8角形を作る
			octa1name, octa1L = MkoctaB2(cordz, order2, keyList, nodDec, num0, d0Num)
			// R2点，R3点, R4点を除外する
			log.Println("octa1name", octa1name)
			log.Println("octa1L=", octa1L)
		}

	// １－(2)
	case "LLRRLRRRRR", "LRRLRRRRRL", "LRRRRRLLRR":
		// L3点のインデックスがnum0
		if lrPtn[1] == "L" && lrPtn[4] == "L" {
			num0 = lrIdx[4]
		} else if lrPtn[3] == "L" && lrPtn[9] == "L" {
			num0 = lrIdx[3]
		} else if lrPtn[6] == "L" && lrPtn[7] == "L" {
			num0 = lrIdx[0]
		}
		// num0はL3点のインデックス番号
		// 10角形を１つの四角形と１つの8角形に分割する
		cordz, order2, keyList, areaTag := DecaPrepro(num0, cord2, nodDec, order, tp)
		// 大きい耳となる方の四角形を分割し，残りで8角形を作る
		if areaTag == "areaA" {
			// 四角形を作る
			rect1name, _ := MkrectA(cordz, order2, keyList, nodDec, num0, d0Num)
			// rect1name = ['D1', 'L3', 'R3', 'R4']

			log.Println("rect1name", rect1name)
			rect4L = MakeRectList(cordz, order2, rect1name)
			log.Println("rect4L=", rect4L)
			// 8角形を作る
			octa1name, octa1L = MkoctaA(cordz, order2, keyList, nodDec, num0, d0Num)
			// L3点，R3点，R4点を除外する
			log.Println("octa1name", octa1name)
			log.Println("octa1L=", octa1L)
		} else if areaTag == "areaB" {
			// 四角形を作る
			rect1name, _ := MkrectB(cordz, order2, keyList, nodDec, num0, d0Num)
			// rect1name = ['L3', 'D2', 'R1', 'R2']

			log.Println("rect1name", rect1name)
			rect4L = MakeRectList(cordz, order2, rect1name)
			log.Println("rect4L=", rect4L)
			// 8角形を作る
			octa1name, octa1L = MkoctaB(cordz, order2, keyList, nodDec, num0, d0Num)
			// L3点，R1点，R2点を除外する
			log.Println("octa1name", octa1name)
			log.Println("octa1L=", octa1L)
		}

	// １－(3)
	case "LLRRRLRRRR", "LRRRLRRRRL", "LRRRRLLRRR":
		if lrPtn[1] == "L" && lrPtn[5] == "L" {
			num0 = lrIdx[5]
		} else if lrPtn[4] == "L" && lrPtn[9] == "L" {
			num0 = lrIdx[4]
		} else if lrPtn[5] == "L" && lrPtn[6] == "L" {
			num0 = lrIdx[0]
		}

		// num0はL3点のインデックス番号
		// 10角形を１つの四角形と１つの8角形に分割する
		cordz, order2, keyList, areaTag := DecaPrepro(num0, cord2, nodDec, order, tp)
		// 大きい耳となる方の四角形を分割し，残りで8角形を作る
		if areaTag == "areaA" {
			// 四角形を作る
			rect1name, _ := MkrectA(cordz, order2, keyList, nodDec, num0, d0Num)
			// rect1name = ['D1', 'L3', 'R4', 'R5']

			log.Println("rect1name", rect1name)
			rect4L = MakeRectList(cordz, order2, rect1name)
			log.Println("rect4L=", rect4L)
			// 8角形を作る
			octa1name, octa1L = MkoctaA(cordz, order2, keyList, nodDec, num0, d0Num)
			// L3点，R4点，R5点を除外する
			log.Println("octa1name", octa1name)
			log.Println("octa1L=", octa1L)
		} else if areaTag == "areaB" {
			// 四角形を作る
			rect1name, _ := MkrectB(cordz, order2, keyList, nodDec, num0, d0Num)
			// rect1name = ['L3', 'D2', 'R2', 'R3']

			log.Println("rect1name", rect1name)
			rect4L = MakeRectList(cordz, order2, rect1name)
			log.Println("rect4L=", rect4L)
			// 8角形を作る
			octa1name, octa1L = MkoctaB(cordz, order2, keyList, nodDec, num0, d0Num)
			// L3点，R2点，R3点を除外する
			log.Println("octa1name", octa1name)
			log.Println("octa1L=", octa1L)
		}

	// １－(4)
	case "LLRRRRLRRR", "LRRRRLRRRL", "LRRRLLRRRR":
		if lrPtn[1] == "L" && lrPtn[6] == "L" {
			num0 = lrIdx[6]
		} else if lrPtn[5] == "L" && lrPtn[9] == "L" {
			num0 = lrIdx[5]
		} else if lrPtn[4] == "L" && lrPtn[5] == "L" {
			num0 = lrIdx[0]
		}
		log.Println("num0", num0)
		// num0はL3点のインデックス番号
		// 10角形を１つの四角形と１つの8角形に分割する
		cordz, order2, keyList, areaTag := DecaPrepro(num0, cord2, nodDec, order, tp)
		// 大きい耳となる方の四角形を分割し，残りで8角形を作る
		if areaTag == "areaA" {
			// 四角形を作る
			rect1name, _ := MkrectA(cordz, order2, keyList, nodDec, num0, d0Num)
			// rect1name = ['D1', 'L3', 'R5', 'R6']

			log.Println("rect1name", rect1name)
			rect4L = MakeRectList(cordz, order2, rect1name)
			log.Println("rect4L=", rect4L)
			// 8角形を作る
			octa1name, octa1L = MkoctaA(cordz, order2, keyList, nodDec, num0, d0Num)
			// L3点，R5点，R6点を除外する
			log.Println("octa1name", octa1name)
			log.Println("octa1L=", octa1L)
		} else if areaTag == "areaB" {
			// 四角形を作る
			rect1name, _ := MkrectB(cordz, order2, keyList, nodDec, num0, d0Num)
			// rect1name = ['L3', 'D2', 'R3', 'R4']

			log.Println("rect1name", rect1name)
			rect4L = MakeRectList(cordz, order2, rect1name)
			log.Println("rect4L=", rect4L)
			// 8角形を作る
			octa1name, octa1L = MkoctaB(cordz, order2, keyList, nodDec, num0, d0Num)
			// L3点，R3点，R4点を除外する
			log.Println("octa1name", octa1name)
			log.Println("octa1L=", octa1L)
		}

	// １－(5)
	case "LLRRRRRLRR", "LRRRRRLRRL", "LRRLLRRRRR":
		if lrPtn[1] == "L" && lrPtn[7] == "L" {
			num0 = lrIdx[7]
		} else if lrPtn[6] == "L" && lrPtn[9] == "L" {
			num0 = lrIdx[6]
		} else if lrPtn[3] == "L" && lrPtn[4] == "L" {
			num0 = lrIdx[0]
		}
		// num0はL3点のインデックス番号
		// 10角形を１つの四角形と１つの8角形に分割する
		cordz, order2, keyList, areaTag := DecaPrepro(num0, cord2, nodDec, order, tp)
		// 大きい耳となる方の四角形を分割し，残りで8角形を作る
		if areaTag == "areaA" {
			// 四角形を作る
			rect1name, _ := MkrectA(cordz, order2, keyList, nodDec, num0, d0Num)
			// rect1name = ['D1', 'L3', 'R6', 'R7']

			log.Println("rect1name", rect1name)
			rect4L = MakeRectList(cordz, order2, rect1name)
			log.Println("rect4L=", rect4L)
			// 8角形を作る
			octa1name, octa1L = MkoctaA(cordz, order2, keyList, nodDec, num0, d0Num)
			log.Println("octa1name", octa1name)
			log.Println("octa1L=", octa1L)
		} else if areaTag == "areaB" {
			// 四角形を作る
			rect1name, _ := MkrectB(cordz, order2, keyList, nodDec, num0, d0Num)
			// rect1name = ['L3', 'D2', 'R4', 'R5']

			log.Println("rect1name", rect1name)
			rect4L = MakeRectList(cordz, order2, rect1name)
			log.Println("rect4L=", rect4L)
			// 8角形を作る
			octa1name, octa1L = MkoctaB(cordz, order2, keyList, nodDec, num0, d0Num)
			// L3点，R4点，R5点を除外する
			log.Println("octa1name", octa1name)
			log.Println("octa1L=", octa1L)
		}

	// １－(6)
	case "LLRRRRRRLR", "LRRRRRRLRL", "LRLLRRRRRR":
		if lrPtn[1] == "L" && lrPtn[8] == "L" {
			num0 = lrIdx[8]
		} else if lrPtn[7] == "L" && lrPtn[9] == "L" {
			num0 = lrIdx[7]
		} else if lrPtn[2] == "L" && lrPtn[3] == "L" {
			num0 = lrIdx[0]
		}
		// num0はL3点のインデックス番号
		// 10角形を１つの四角形と１つの8角形に分割する
		// 分割線zが対向する辺はL3点から反対方向に３つ目と４つ目のR点が作る辺
		tp = "a2"
		cordz, order2, keyList, areaTag := DecaPrepro(num0, cord2, nodDec, order, tp)
		// 大きい耳となる方の四角形を分割し，残りで8角形を作る
		if areaTag == "areaA" {
			// 四角形を作る
			rect1name, _ := MkrectA2(cordz, order2, keyList, nodDec, num0, d0Num)
			// rect1name = ['R6', 'D1', 'R4', 'R5']

			log.Println("rect1name", rect1name)
			rect4L = MakeRectList(cordz, order2, rect1name)
			log.Println("rect4L=", rect4L)
			// 8角形を作る
			octa1name, octa1L = MkoctaA2(cordz, order2, keyList, nodDec, num0, d0Num)
			log.Println("octa1name", octa1name)
			log.Println("octa1L=", octa1L)
		}
		// areaBはareaAより必ず小さい

	// ２－(1)
	case "LRLRLRRRRR", "LRLRRRRRLR", "LRRRRRLRLR":
		if lrPtn[2] == "L" && lrPtn[4] == "L" {
			num0 = lrIdx[4]
		} else if lrPtn[2] == "L" && lrPtn[8] == "L" {
			num0 = lrIdx[2]
		} else if lrPtn[6] == "L" && lrPtn[8] == "L" {
			num0 = lrIdx[0]
		}
		// num0はL3点のインデックス番号
		// 10角形を１つの四角形と１つの8角形に分割する
		// 分割線bが対向する辺はL3点から３つ目と４つ目のR点が作る辺
		tp = "b2"
		cordz, order2, keyList, areaTag := DecaPrepro(num0, cord2, nodDec, order, tp)
		// 大きい耳となる方の四角形を分割し，残りで8角形を作る
		if areaTag == "areaA" {
			// areaAはareaBより必ず小さい
			// TODO:

		} else if areaTag == "areaB" {
			// 四角形を作る
			// L3点の次のR点とその次のR点とその次の々のR点とD2点: d0Num, num0+1, num0+2, num0+3
			rect1name, _ := MkrectB2(cordz, order2, keyList, nodDec, num0, d0Num)
			// rect1name = ['D2', 'R3', 'R4', 'R5']

			log.Println("rect1name", rect1name)
			rect4L = MakeRectList(cordz, order2, rect1name)
			log.Println("rect4L=", rect4L)
			// 8角形を作る
			octa1name, octa1L = MkoctaB2(cordz, order2, keyList, nodDec, num0, d0Num)
			// L3点，R3点，R4点を除外する
			// num0とnum0+1，num0+2を削除する
			// d0Num, num0+3, num0+4, num0+5, num0+6, num0+7, num0+8, num0+9
			log.Println("octa1name", octa1name)
			log.Println("octa1L=", octa1L)
		}

	// ２－(2)
	case "LRLRRLRRRR", "LRRLRRRRLR", "LRRRRLRLRR":
		if lrPtn[2] == "L" && lrPtn[5] == "L" {
			num0 = lrIdx[5]
		} else if lrPtn[3] == "L" && lrPtn[8] == "L" {
			num0 = lrIdx[3]
		} else if lrPtn[5] == "L" && lrPtn[7] == "L" {
			num0 = lrIdx[0]
		}
		// num0はL3点のインデックス番号
		// 10角形を１つの四角形と１つの8角形に分割する
		cordz, order2, keyList, areaTag := DecaPrepro(num0, cord2, nodDec, order, tp)
		// 大きい耳となる方の四角形を分割し，残りで8角形を作る
		if areaTag == "areaA" {
			// 四角形を作る
			// L3//L1点と次のR点とその次のR点とD点: num0, num0+1, num0+2, d0Num
			rect1name, _ := MkrectA(cordz, order2, keyList, nodDec, num0, d0Num)
			// rect1name = ['D1', 'L3', 'R4', 'R5']

			log.Println("rect1name", rect1name)
			rect4L = MakeRectList(cordz, order2, rect1name)
			log.Println("rect4L=", rect4L)
			// 8角形を作る
			// d0Num, num0+3, num0+4, num0+5, num0+6, num0+7, num0+8, num0+9
			octa1name, octa1L = MkoctaA(cordz, order2, keyList, nodDec, num0, d0Num)
			log.Println("octa1name", octa1name)
			log.Println("octa1L=", octa1L)
		} else if areaTag == "areaB" {
			// 四角形を作る
			// L3点とD点と前の々のR点と前のR点: num0, d0Num, num0+8, num0+9
			rect1name, _ := MkrectB(cordz, order2, keyList, nodDec, num0, d0Num)
			// rect1name = ['L3', 'D2', 'R2', 'R3']

			log.Println("rect1name", rect1name)
			rect4L = MakeRectList(cordz, order2, rect1name)
			log.Println("rect4L=", rect4L)
			// 8角形を作る
			// L3点，R2点，R3点を除外する
			// num0とnum0-1，num0-2を削除する
			// d0Num, num0+1, num0+2, num0+3, num0+4, num0+5, num0+6, num0+7
			octa1name, octa1L = MkoctaB(cordz, order2, keyList, nodDec, num0, d0Num)
			log.Println("octa1name", octa1name)
			log.Println("octa1L=", octa1L)
		}

	// ２－(3)
	case "LRLRRRLRRR", "LRRRLRRRLR", "LRRRLRLRRR":
		if lrPtn[2] == "L" && lrPtn[6] == "L" {
			num0 = lrIdx[6]
		} else if lrPtn[4] == "L" && lrPtn[8] == "L" {
			num0 = lrIdx[4]
		} else if lrPtn[4] == "L" && lrPtn[6] == "L" {
			num0 = lrIdx[0]
		}
		// num0はL3点のインデックス番号
		// 10角形を１つの四角形と１つの8角形に分割する
		cordz, order2, keyList, areaTag := DecaPrepro(num0, cord2, nodDec, order, tp)
		// 大きい耳となる方の四角形を分割し，残りで8角形を作る
		if areaTag == "areaA" {
			// 四角形を作る
			// L3//L1点と次のR点とその次のR点とD点: num0, num0+1, num0+2, d0Num
			rect1name, _ := MkrectA(cordz, order2, keyList, nodDec, num0, d0Num)
			// rect1name = ['D1', 'L3', 'R5', 'R6']

			log.Println("rect1name", rect1name)
			rect4L = MakeRectList(cordz, order2, rect1name)
			log.Println("rect4L=", rect4L)
			// 8角形を作る
			// d0Num, num0+3, num0+4, num0+5, num0+6, num0+7, num0+8, num0+9
			octa1name, octa1L = MkoctaA(cordz, order2, keyList, nodDec, num0, d0Num)
			log.Println("octa1name", octa1name)
			log.Println("octa1L=", octa1L)
		} else if areaTag == "areaB" {
			// 四角形を作る
			// L3点とD点と前の々のR点と前のR点: num0, d0Num, num0+8, num0+9
			rect1name, _ := MkrectB(cordz, order2, keyList, nodDec, num0, d0Num)
			// rect1name = ['L3', 'D2', 'R3', 'R4']

			log.Println("rect1name", rect1name)
			rect4L = MakeRectList(cordz, order2, rect1name)
			log.Println("rect4L=", rect4L)
			// 8角形を作る
			// L3点，R3点，R4点を除外する
			// num0とnum0-1，num0-2を削除する
			// d0Num, num0+1, num0+2, num0+3, num0+4, num0+5, num0+6, num0+7
			octa1name, octa1L = MkoctaB(cordz, order2, keyList, nodDec, num0, d0Num)
			log.Println("octa1name", octa1name)
			log.Println("octa1L=", octa1L)
		}

	// ２－(4)
	case "LRLRRRRLRR", "LRRRRLRRLR", "LRRLRLRRRR":
		if lrPtn[2] == "L" && lrPtn[7] == "L" {
			num0 = lrIdx[7]
		} else if lrPtn[5] == "L" && lrPtn[8] == "L" {
			num0 = lrIdx[5]
		} else if lrPtn[3] == "L" && lrPtn[5] == "L" {
			num0 = lrIdx[0]
		}
		// num0はL3点のインデックス番号
		// 10角形を１つの四角形と１つの8角形に分割する
		cordz, order2, keyList, areaTag := DecaPrepro(num0, cord2, nodDec, order, tp)
		// 大きい耳となる方の四角形を分割し，残りで8角形を作る
		if areaTag == "areaA" {
			// 四角形を作る
			// L3//L1点と次のR点とその次のR点とD点: num0, num0+1, num0+2, d0Num
			rect1name, _ := MkrectA(cordz, order2, keyList, nodDec, num0, d0Num)
			// rect1name = ['D1', 'L3', 'R6', 'R7']

			log.Println("rect1name", rect1name)
			rect4L = MakeRectList(cordz, order2, rect1name)
			log.Println("rect4L=", rect4L)
			// 8角形を作る
			// d0Num, num0+3, num0+4, num0+5, num0+6, num0+7, num0+8, num0+9
			octa1name, octa1L = MkoctaA(cordz, order2, keyList, nodDec, num0, d0Num)
			log.Println("octa1name", octa1name)
			log.Println("octa1L=", octa1L)
		} else if areaTag == "areaB" {
			// 四角形を作る
			// L3点とD点と前の々のR点と前のR点: num0, d0Num, num0+8, num0+9
			rect1name, _ := MkrectB(cordz, order2, keyList, nodDec, num0, d0Num)
			// rect1name = ['L3', 'D2', 'R4', 'R5']

			log.Println("rect1name", rect1name)
			rect4L = MakeRectList(cordz, order2, rect1name)
			log.Println("rect4L=", rect4L)
			// 8角形を作る
			// L3点，R2点，R3点を除外する
			// num0とnum0-1，num0-2を削除する
			// d0Num, num0+1, num0+2, num0+3, num0+4, num0+5, num0+6, num0+7
			octa1name, octa1L = MkoctaB(cordz, order2, keyList, nodDec, num0, d0Num)
			log.Println("octa1name", octa1name)
			log.Println("octa1L=", octa1L)
		}

	// ２－(5)
	case "LRRLRRLRRR", "LRRLRRRLRR", "LRRRLRRLRR":
		if lrPtn[3] == "L" && lrPtn[6] == "L" {
			num0 = lrIdx[6]
		} else if lrPtn[3] == "L" && lrPtn[7] == "L" {
			num0 = lrIdx[3]
		} else if lrPtn[4] == "L" && lrPtn[7] == "L" {
			num0 = lrIdx[0]
		}
		// num0はL3点のインデックス番号
		// 10角形を１つの四角形と１つの8角形に分割する
		cordz, order2, keyList, areaTag := DecaPrepro(num0, cord2, nodDec, order, tp)
		// 大きい耳となる方の四角形を分割し，残りで8角形を作る
		if areaTag == "areaA" {
			// 四角形を作る
			// L3//L1点と次のR点とその次のR点とD点: num0, num0+1, num0+2, d0Num
			rect1name, _ := MkrectA(cordz, order2, keyList, nodDec, num0, d0Num)
			// rect1name = ['L3', 'R5', 'R6', 'D1']

			log.Println("rect1name", rect1name)
			rect4L = MakeRectList(cordz, order2, rect1name)
			log.Println("rect4L=", rect4L)
			// 8角形を作る
			// d0Num, num0+3, num0+4, num0+5, num0+6, num0+7, num0+8, num0+9
			octa1name, octa1L = MkoctaA(cordz, order2, keyList, nodDec, num0, d0Num)
			log.Println("octa1name", octa1name)
			log.Println("octa1L=", octa1L)
		} else if areaTag == "areaB" {
			// 四角形を作る
			// L3点とD点と前の々のR点と前のR点: num0, d0Num, num0+8, num0+9
			rect1name, _ := MkrectB(cordz, order2, keyList, nodDec, num0, d0Num)
			// rect1name = ['L3', 'D2', 'R3', 'R4']

			log.Println("rect1name", rect1name)
			rect4L = MakeRectList(cordz, order2, rect1name)
			log.Println("rect4L=", rect4L)
			// 8角形を作る
			// L3点，R2点，R3点を除外する
			// num0とnum0-1，num0-2を削除する
			// d0Num, num0+1, num0+2, num0+3, num0+4, num0+5, num0+6, num0+7
			octa1name, octa1L = MkoctaB(cordz, order2, keyList, nodDec, num0, d0Num)
			log.Println("octa1name", octa1name)
			log.Println("octa1L=", octa1L)
		}

		// ３－(1)
	case "LLLRRRRRRR", "LLRRRRRRRL", "LRRRRRRRLL":
		if lrPtn[1] == "L" && lrPtn[2] == "L" {
			num0 = lrIdx[2]
		} else if lrPtn[1] == "L" && lrPtn[9] == "L" {
			num0 = lrIdx[1]
		} else if lrPtn[8] == "L" && lrPtn[9] == "L" {
			num0 = lrIdx[0]
		}
		// num0はL3点のインデックス番号
		// 10角形を１つの四角形と１つの8角形に分割する
		// 分割線bが対向する辺はL3点から３つ目と４つ目のR点が作る辺
		tp = "b2"
		cordz, order2, keyList, areaTag := DecaPrepro(num0, cord2, nodDec, order, tp)
		// 大きい耳となる方の四角形を分割し，残りで8角形を作る
		if areaTag == "areaA" {
			// areaAはareaBより必ず小さい
			// TODO:
		} else if areaTag == "areaB" {
			// 四角形を作る
			// L1点と次のR点とその次のR点とD点: num0, num0+1, num0+2, d0Num
			rect1name, _ := MkrectB2(cordz, order2, keyList, nodDec, num0, d0Num)
			// rect1name = ['D2', 'R1', 'R2', 'R3']

			log.Println("rect1name", rect1name)
			rect4L = MakeRectList(cordz, order2, rect1name)
			log.Println("rect4L=", rect4L)
			// 8角形を作る
			// num0, d0Num, num0+4, num0+5, num0+6, num0+7, num0+8, num0+9
			octa1name, octa1L = MkoctaB2(cordz, order2, keyList, nodDec, num0, d0Num)
			log.Println("octa1name", octa1name)
			log.Println("octa1L=", octa1L)
		}
	default:
		// TODO:
	}

	if octa1L != nil {
		log.Println("octa1L=", octa1L)
		// 頂点データ数の確認
		nod := len(octa1L)
		if nod != 8 {
			// TODO:
		}
		// 追加したD点のために外積の計算をやり直す必要がある
		extL, _, _ := TriVert(nod, octa1L)
		// L点，R点の辞書を作り直す
		_, _, orderN, lrPtn, lrIdx := Lexicogra(nod, octa1L, extL)
		log.Println("orderN=", orderN)
		// 8角形の四角形分割プログラムに渡す
		_, rect1L, rect2L, rect3L, story, yane = OctaDiv(octa1L, orderN, lrPtn, lrIdx)
		log.Println("rectO1L=", rect1L)
		log.Println("rectO2L=", rect2L)
		log.Println("rectO3L=", rect3L)
	}

	if hex1L != nil {
		// 頂点データ数の確認
		nod := len(hex1L)
		if nod != 6 {
			// TODO:
		}
		// 追加したD点のために外積の計算をやり直す必要がある
		extL, _, _ := TriVert(nod, hex1L)
		// L点，R点の辞書を作り直す
		_, _, orderN, _, _ := Lexicogra(nod, hex1L, extL)
		// log.Println("orderN=", orderN)
		// 6角形の四角形分割プログラムに渡す
		_, rect2L, rect3L = HexaDiv(hex1L, orderN)
		log.Println("rectH2L", rect2L)
		log.Println("rectH3L", rect3L)
	}
	return rect1L, rect2L, rect3L, rect4L, story, yane
}
