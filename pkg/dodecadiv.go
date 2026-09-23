package pkg

import (
	"log"
	"strings"
)

// DodecaDiv は12角形を５つの四角形に分割する
func DodecaDiv(cord2 [][]float64, order map[string]int, lrPtn []string,
	lrIdx []int) (rect1L [][]float64, rect2L [][]float64, rect3L [][]float64,
	rect4L [][]float64, rect5L [][]float64, story []int, yane []string) {
	var hex1L [][]float64
	var octa1L [][]float64
	var deca1name []string
	var deca1L [][]float64

	// 頂点データ数の確認
	nodDode := len(cord2)
	if nodDode != 12 {
		// TODO:関数から戻る
		return
	}
	d0Num := nodDode
	var num0 int

	// LR並びの確認　L点から始まっていなければエラー
	if lrPtn[0] != "L" {
		// TODO:関数から戻る
		return
	}
	// 検索用にLR並びから半角スペースを除く
	lrjoin := strings.Join(lrPtn, "")
	log.Println("lrjoin=", lrjoin)

	// 分割線に対抗する辺の初期値（areaAとareaBが重複しない）
	var tp string

	switch lrjoin {
	// １－(1)
	case "LRLRLRLRRRRR", "LRLRLRRRRRLR", "LRLRRRRRLRLR", "LRRRRRLRLRLR":
		if lrPtn[2] == "L" && lrPtn[4] == "L" && lrPtn[6] == "L" {
			num0 = lrIdx[6]
		} else if lrPtn[2] == "L" && lrPtn[4] == "L" && lrPtn[10] == "L" {
			num0 = lrIdx[4]
		} else if lrPtn[2] == "L" && lrPtn[8] == "L" && lrPtn[10] == "L" {
			num0 = lrIdx[2]
		} else if lrPtn[6] == "L" && lrPtn[8] == "L" && lrPtn[10] == "L" {
			num0 = lrIdx[0]
		}
		log.Println("１－(1)", lrjoin)

		// num0はL4点のインデックス番号
		// 12角形を１つの四角形と１つの10角形に分割する
		// 分割線bが対向する辺はL4点から３つ目と４つ目のR点が作る辺
		tp = "b2"
		cordz, order2, keyList, areaTag := DodePrepro(num0, cord2, nodDode, order, tp)
		// 大きい耳となる方の四角形を分割し，残りで10角形を作る
		if areaTag == "areaA" {
			// areaAはareaBより必ず小さい
			// TODO:
		} else if areaTag == "areaB" {
			// 四角形を作る
			rect5name, _ := MkrectB2(cordz, order2, keyList, nodDode, num0, d0Num)
			// rect5name = ['D2', 'R4', 'R5', 'R6']

			log.Println("rect5name", rect5name)
			rect5L = MakeRectList(cordz, order2, rect5name)
			log.Println("rect5L=", rect5L)
			// 10角形を作る
			deca1name, deca1L = MkdecaB2(cordz, order2, keyList, nodDode, num0, d0Num)
			// R4点，R5点，R6点を除外する
			log.Println("deca1name", deca1name)
			log.Println("deca1L=", deca1L)
		}

	// １－(2)
	case "LRLRLRRLRRRR", "LRLRRLRRRRLR", "LRRLRRRRLRLR", "LRRRRLRLRLRR":
		if lrPtn[2] == "L" && lrPtn[4] == "L" && lrPtn[7] == "L" {
			num0 = lrIdx[7]
		} else if lrPtn[2] == "L" && lrPtn[5] == "L" && lrPtn[10] == "L" {
			num0 = lrIdx[5]
		} else if lrPtn[3] == "L" && lrPtn[8] == "L" && lrPtn[10] == "L" {
			num0 = lrIdx[3]
		} else if lrPtn[5] == "L" && lrPtn[7] == "L" && lrPtn[9] == "L" {
			num0 = lrIdx[0]
		}
		log.Println("１－(2)", lrjoin)

		// num0はL4点のインデックス番号
		// 12角形を１つの四角形と１つの10角形に分割する
		cordz, order2, keyList, areaTag := DodePrepro(num0, cord2, nodDode, order, tp)
		// 大きい耳となる方の四角形を分割し，残りで10角形を作る
		if areaTag == "areaA" {
			// 四角形を作る
			rect5name, _ := MkrectA(cordz, order2, keyList, nodDode, num0, d0Num)
			// rect5name = ['D1', 'L4', 'R5', 'R6']

			log.Println("rect5name", rect5name)
			rect5L = MakeRectList(cordz, order2, rect5name)
			log.Println("rect5L=", rect5L)
			// 10角形を作る
			deca1name, deca1L = MkdecaA(cordz, order2, keyList, nodDode, num0, d0Num)
			// L4点，R5点，R6点を除外する
			log.Println("deca1name", deca1name)
			log.Println("deca1L=", deca1L)
		} else if areaTag == "areaB" {
			// 四角形を作る
			rect5name, _ := MkrectB(cordz, order2, keyList, nodDode, num0, d0Num)
			// rect5name = ['L4', 'D2', 'R3', 'R4']

			log.Println("rect5name", rect5name)
			rect5L = MakeRectList(cordz, order2, rect5name)
			log.Println("rect5L=", rect5L)
			// 10角形を作る
			deca1name, deca1L = MkdecaB(cordz, order2, keyList, nodDode, num0, d0Num)
			// L4点，R3点，R4点を除外する
			log.Println("deca1name", deca1name)
			log.Println("deca1L=", deca1L)
		}

	// １－(3)
	case "LRLRLRRRLRRR", "LRLRRRLRRRLR", "LRRRLRRRLRLR", "LRRRLRLRLRRR":
		if lrPtn[2] == "L" && lrPtn[4] == "L" && lrPtn[8] == "L" {
			num0 = lrIdx[8]
		} else if lrPtn[2] == "L" && lrPtn[6] == "L" && lrPtn[10] == "L" {
			num0 = lrIdx[6]
		} else if lrPtn[4] == "L" && lrPtn[8] == "L" && lrPtn[10] == "L" {
			num0 = lrIdx[4]
		} else if lrPtn[4] == "L" && lrPtn[6] == "L" && lrPtn[8] == "L" {
			num0 = lrIdx[0]
		}
		log.Println("１－(3)", lrjoin)

		// num0はL4点のインデックス番号
		// 12角形を１つの四角形と１つの10角形に分割する
		cordz, order2, keyList, areaTag := DodePrepro(num0, cord2, nodDode, order, tp)
		// 大きい耳となる方の四角形を分割し，残りで10角形を作る
		if areaTag == "areaA" {
			// 四角形を作る
			rect5name, _ := MkrectA(cordz, order2, keyList, nodDode, num0, d0Num)
			// rect5name = ['D1', 'L4', 'R6', 'R7']

			log.Println("rect5name", rect5name)
			rect5L = MakeRectList(cordz, order2, rect5name)
			log.Println("rect5L=", rect5L)
			// 10角形を作る
			deca1name, deca1L = MkdecaA(cordz, order2, keyList, nodDode, num0, d0Num)
			// L4点，R6点，R7点を除外する
			log.Println("deca1name", deca1name)
			log.Println("deca1L=", deca1L)
		} else if areaTag == "areaB" {
			// 四角形を作る
			rect5name, _ := MkrectB(cordz, order2, keyList, nodDode, num0, d0Num)
			// rect5name = ['L4', 'D2', 'R4', 'R5']

			log.Println("rect5name", rect5name)
			rect5L = MakeRectList(cordz, order2, rect5name)
			log.Println("rect5L=", rect5L)
			// 10角形を作る
			deca1name, deca1L = MkdecaB(cordz, order2, keyList, nodDode, num0, d0Num)
			// L4点，R4点，R5点を除外する
			log.Println("deca1name", deca1name)
			log.Println("deca1L=", deca1L)
		}

	// １－(4)
	case "LRLRLRRRRLRR", "LRLRRRRLRRLR", "LRRRRLRRLRLR", "LRRLRLRLRRRR":
		if lrPtn[2] == "L" && lrPtn[4] == "L" && lrPtn[9] == "L" {
			num0 = lrIdx[9]
		} else if lrPtn[2] == "L" && lrPtn[7] == "L" && lrPtn[10] == "L" {
			num0 = lrIdx[7]
		} else if lrPtn[5] == "L" && lrPtn[8] == "L" && lrPtn[10] == "L" {
			num0 = lrIdx[5]
		} else if lrPtn[3] == "L" && lrPtn[5] == "L" && lrPtn[7] == "L" {
			num0 = lrIdx[0]
		}
		log.Println("１－(4)", lrjoin)

		// num0はL4点のインデックス番号
		// 12角形を１つの四角形と１つの10角形に分割する
		cordz, order2, keyList, areaTag := DodePrepro(num0, cord2, nodDode, order, tp)
		// 大きい耳となる方の四角形を分割し，残りで10角形を作る
		if areaTag == "areaA" {
			// 四角形を作る
			rect5name, _ := MkrectA(cordz, order2, keyList, nodDode, num0, d0Num)
			// rect5name = ['D1', 'L4', 'R7', 'R8']

			log.Println("rect5name", rect5name)
			rect5L = MakeRectList(cordz, order2, rect5name)
			log.Println("rect5L=", rect5L)
			// 10角形を作る
			deca1name, deca1L = MkdecaA(cordz, order2, keyList, nodDode, num0, d0Num)
			// R4点，R5点，R6点を除外する
			log.Println("deca1name", deca1name)
			log.Println("deca1L=", deca1L)
		} else if areaTag == "areaB" {
			// 四角形を作る
			rect5name, _ := MkrectB(cordz, order2, keyList, nodDode, num0, d0Num)
			// rect5name = ['L4', 'D2', 'R5', 'R6']

			log.Println("rect5name", rect5name)
			rect5L = MakeRectList(cordz, order2, rect5name)
			log.Println("rect5L=", rect5L)
			// 10角形を作る
			deca1name, deca1L = MkdecaB(cordz, order2, keyList, nodDode, num0, d0Num)
			// L4点，R5点，R6点を除外する
			log.Println("deca1name", deca1name)
			log.Println("deca1L=", deca1L)
		}

	// １－(5)
	case "LRLRRLRLRRRR", "LRRLRLRRRRLR", "LRLRRRRLRLRR", "LRRRRLRLRRLR":
		if lrPtn[2] == "L" && lrPtn[5] == "L" && lrPtn[7] == "L" {
			num0 = lrIdx[7]
		} else if lrPtn[3] == "L" && lrPtn[5] == "L" && lrPtn[10] == "L" {
			num0 = lrIdx[5]
		} else if lrPtn[2] == "L" && lrPtn[7] == "L" && lrPtn[9] == "L" {
			num0 = lrIdx[2]
		} else if lrPtn[5] == "L" && lrPtn[7] == "L" && lrPtn[10] == "L" {
			num0 = lrIdx[0]
		}
		log.Println("１－(5)", lrjoin)

		// num0はL4点のインデックス番号
		// 12角形を１つの四角形と１つの10角形に分割する
		// 分割線bが対向する辺はL4点から３つ目と４つ目のR点が作る辺
		tp = "b2"
		cordz, order2, keyList, areaTag := DodePrepro(num0, cord2, nodDode, order, tp)
		// 大きい耳となる方の四角形を分割し，残りで10角形を作る
		if areaTag == "areaA" {
			// 四角形を作る
			rect5name, _ := MkrectA(cordz, order2, keyList, nodDode, num0, d0Num)
			// rect5name = ['D1', 'L4', 'R5', 'R6']

			log.Println("rect5name", rect5name)
			rect5L = MakeRectList(cordz, order2, rect5name)
			log.Println("rect5L=", rect5L)
			// 10角形を作る
			deca1name, deca1L = MkdecaA(cordz, order2, keyList, nodDode, num0, d0Num)
			// L4点，R5点，R6点を除外する
			log.Println("deca1name", deca1name)
			log.Println("deca1L=", deca1L)
		} else if areaTag == "areaB" {
			// areaBはareaAより必ず大きい
			// しかし，分割線bはエラーになる可能性がある
			// 四角形を作る
			rect5name, _ := MkrectB2(cordz, order2, keyList, nodDode, num0, d0Num)
			// rect5name = ['D2', 'R5', 'R6', 'R7']

			log.Println("rect5name", rect5name)
			rect5L = MakeRectList(cordz, order2, rect5name)
			log.Println("rect5L=", rect5L)
			// 10角形を作る
			deca1name, deca1L = MkdecaB2(cordz, order2, keyList, nodDode, num0, d0Num)
			// R5点，R6点，R7点を除外する
			log.Println("deca1name", deca1name)
			log.Println("deca1L=", deca1L)
		}

	// １－(6)
	case "LRLRRLRRLRRR", "LRRLRRLRRRLR", "LRRLRRRLRLRR", "LRRRLRLRRLRR":
		if lrPtn[2] == "L" && lrPtn[5] == "L" && lrPtn[8] == "L" {
			num0 = lrIdx[8]
		} else if lrPtn[3] == "L" && lrPtn[6] == "L" && lrPtn[10] == "L" {
			num0 = lrIdx[6]
		} else if lrPtn[3] == "L" && lrPtn[7] == "L" && lrPtn[9] == "L" {
			num0 = lrIdx[3]
		} else if lrPtn[4] == "L" && lrPtn[6] == "L" && lrPtn[9] == "L" {
			num0 = lrIdx[0]
		}
		log.Println("１－(6)", lrjoin)

		// num0はL4点のインデックス番号
		// 12角形を１つの四角形と１つの10角形に分割する
		cordz, order2, keyList, areaTag := DodePrepro(num0, cord2, nodDode, order, tp)
		// 大きい耳となる方の四角形を分割し，残りで10角形を作る
		if areaTag == "areaA" {
			// 四角形を作る
			rect5name, _ := MkrectA(cordz, order2, keyList, nodDode, num0, d0Num)
			// rect5name = ['D1', 'L4', 'R6', 'R7']

			log.Println("rect5name", rect5name)
			rect5L = MakeRectList(cordz, order2, rect5name)
			log.Println("rect5L=", rect5L)
			// 10角形を作る
			deca1name, deca1L = MkdecaA(cordz, order2, keyList, nodDode, num0, d0Num)
			// L4点，R6点，R7点を除外する
			log.Println("deca1name", deca1name)
			log.Println("deca1L=", deca1L)
		} else if areaTag == "areaB" {
			// 四角形を作る
			rect5name, _ := MkrectB(cordz, order2, keyList, nodDode, num0, d0Num)
			// rect5name = ['L4', 'D2', 'R4', 'R5']

			log.Println("rect5name", rect5name)
			rect5L = MakeRectList(cordz, order2, rect5name)
			log.Println("rect5L=", rect5L)
			// 10角形を作る
			deca1name, deca1L = MkdecaB(cordz, order2, keyList, nodDode, num0, d0Num)
			// L4点，R4点，R5点を除外する
			log.Println("deca1name", deca1name)
			log.Println("deca1L=", deca1L)
		}

	// １－(7)
	case "LRLRRLRRRLRR", "LRRLRRRLRRLR", "LRRRLRRLRLRR", "LRRLRLRRLRRR":
		if lrPtn[2] == "L" && lrPtn[5] == "L" && lrPtn[9] == "L" {
			num0 = lrIdx[9]
		} else if lrPtn[3] == "L" && lrPtn[7] == "L" && lrPtn[10] == "L" {
			num0 = lrIdx[7]
		} else if lrPtn[4] == "L" && lrPtn[7] == "L" && lrPtn[9] == "L" {
			num0 = lrIdx[4]
		} else if lrPtn[3] == "L" && lrPtn[5] == "L" && lrPtn[8] == "L" {
			num0 = lrIdx[0]
		}
		log.Println("１－(7)", lrjoin)

		// num0はL4点のインデックス番号
		// 12角形を１つの四角形と１つの10角形に分割する
		cordz, order2, keyList, areaTag := DodePrepro(num0, cord2, nodDode, order, tp)
		// 大きい耳となる方の四角形を分割し，残りで10角形を作る
		if areaTag == "areaA" {
			// 四角形を作る
			rect5name, _ := MkrectA(cordz, order2, keyList, nodDode, num0, d0Num)
			// rect5name = ['D1', 'L4', 'R7', 'R8']

			log.Println("rect5name", rect5name)
			rect5L = MakeRectList(cordz, order2, rect5name)
			log.Println("rect5L=", rect5L)
			// 10角形を作る
			deca1name, deca1L = MkdecaA(cordz, order2, keyList, nodDode, num0, d0Num)
			// R4点，R5点，R6点を除外する
			log.Println("deca1name", deca1name)
			log.Println("deca1L=", deca1L)
		} else if areaTag == "areaB" {
			// 四角形を作る
			rect5name, _ := MkrectB(cordz, order2, keyList, nodDode, num0, d0Num)
			// rect5name = ['L4', 'D2', 'R5', 'R6']

			log.Println("rect5name", rect5name)
			rect5L = MakeRectList(cordz, order2, rect5name)
			log.Println("rect5L=", rect5L)
			// 10角形を作る
			deca1name, deca1L = MkdecaB(cordz, order2, keyList, nodDode, num0, d0Num)
			// L4点，R5点，R6点を除外する
			log.Println("deca1name", deca1name)
			log.Println("deca1L=", deca1L)
		}

	// １－(8)
	case "LRLRRRLRLRRR", "LRRRLRLRRRLR":
		if lrPtn[2] == "L" && lrPtn[6] == "L" && lrPtn[8] == "L" {
			num0 = lrIdx[8]
		} else if lrPtn[4] == "L" && lrPtn[6] == "L" && lrPtn[10] == "L" {
			num0 = lrIdx[4]
		}
		log.Println("１－(8)", lrjoin)

		// num0はL4点のインデックス番号
		// 12角形を１つの四角形と１つの10角形に分割する
		// 分割線bが対向する辺はL4点から３つ目と４つ目のR点が作る辺
		tp = "b2"
		cordz, order2, keyList, areaTag := DodePrepro(num0, cord2, nodDode, order, tp)
		// 大きい耳となる方の四角形を分割し，残りで10角形を作る
		if areaTag == "areaA" {
			// 四角形を作る
			rect5name, _ := MkrectA(cordz, order2, keyList, nodDode, num0, d0Num)
			// rect5name = ['D1', 'L4', 'R6', 'R7']

			log.Println("rect5name", rect5name)
			rect5L = MakeRectList(cordz, order2, rect5name)
			log.Println("rect5L=", rect5L)
			// 10角形を作る
			deca1name, deca1L = MkdecaA(cordz, order2, keyList, nodDode, num0, d0Num)
			// L4点，R6点，R7点を除外する
			log.Println("deca1name", deca1name)
			log.Println("deca1L=", deca1L)
		} else if areaTag == "areaB" {
			// areaBはareaAより必ず大きい
			// しかし，分割線bはエラーになる可能性がある
			// 四角形を作る
			rect5name, _ := MkrectB2(cordz, order2, keyList, nodDode, num0, d0Num)
			// rect5name = ['D2', 'R6', 'R7', 'R8']

			log.Println("rect5name", rect5name)
			rect5L = MakeRectList(cordz, order2, rect5name)
			log.Println("rect5L=", rect5L)
			// 10角形を作る
			deca1name, deca1L = MkdecaB2(cordz, order2, keyList, nodDode, num0, d0Num)
			// R6点，R7点，R8点を除外する
			log.Println("deca1name", deca1name)
			log.Println("deca1L=", deca1L)
		}

	// １－(9)
	case "LRLRRRLRRLRR", "LRRRLRRLRRLR", "LRRLRRLRLRRR", "LRRLRLRRRLRR":
		if lrPtn[2] == "L" && lrPtn[6] == "L" && lrPtn[9] == "L" {
			num0 = lrIdx[9]
		} else if lrPtn[4] == "L" && lrPtn[7] == "L" && lrPtn[10] == "L" {
			num0 = lrIdx[7]
		} else if lrPtn[3] == "L" && lrPtn[6] == "L" && lrPtn[8] == "L" {
			num0 = lrIdx[3]
		} else if lrPtn[3] == "L" && lrPtn[5] == "L" && lrPtn[9] == "L" {
			num0 = lrIdx[0]
		}
		log.Println("１－(9)", lrjoin)

		// num0はL4点のインデックス番号
		// 12角形を１つの四角形と１つの10角形に分割する
		cordz, order2, keyList, areaTag := DodePrepro(num0, cord2, nodDode, order, tp)
		// 大きい耳となる方の四角形を分割し，残りで10角形を作る
		if areaTag == "areaA" {
			// 四角形を作る
			rect5name, _ := MkrectA(cordz, order2, keyList, nodDode, num0, d0Num)
			// rect5name = ['D1', 'L4', 'R7', 'R8']

			log.Println("rect5name", rect5name)
			rect5L = MakeRectList(cordz, order2, rect5name)
			log.Println("rect5L=", rect5L)
			// 10角形を作る
			deca1name, deca1L = MkdecaA(cordz, order2, keyList, nodDode, num0, d0Num)
			// L4点，R7点，R8点を除外する
			log.Println("deca1name", deca1name)
			log.Println("deca1L=", deca1L)
		} else if areaTag == "areaB" {
			// 四角形を作る
			rect5name, _ := MkrectB(cordz, order2, keyList, nodDode, num0, d0Num)
			// rect5name = ['L4', 'D2', 'R5', 'R6']

			log.Println("rect5name", rect5name)
			rect5L = MakeRectList(cordz, order2, rect5name)
			log.Println("rect5L=", rect5L)
			// 10角形を作る
			deca1name, deca1L = MkdecaB(cordz, order2, keyList, nodDode, num0, d0Num)
			// L4点，R5点，R6点を除外する
			log.Println("deca1name", deca1name)
			log.Println("deca1L=", deca1L)
		}

	// １－(10)
	case "LRRLRRLRRLRR":
		num0 = lrIdx[9]
		log.Println("１－(10)", lrjoin)

		// num0はL4点のインデックス番号
		// 12角形を１つの四角形と１つの10角形に分割する
		cordz, order2, keyList, areaTag := DodePrepro(num0, cord2, nodDode, order, tp)
		// 大きい耳となる方の四角形を分割し，残りで10角形を作る
		if areaTag == "areaA" {
			// 四角形を作る
			rect5name, _ := MkrectA(cordz, order2, keyList, nodDode, num0, d0Num)
			// rect5name = ['D1', 'L4', 'R7', 'R8']

			log.Println("rect5name", rect5name)
			rect5L = MakeRectList(cordz, order2, rect5name)
			log.Println("rect5L=", rect5L)
			// 10角形を作る
			deca1name, deca1L = MkdecaA(cordz, order2, keyList, nodDode, num0, d0Num)
			// L4点，R7点，R8点を除外する
			log.Println("deca1name", deca1name)
			log.Println("deca1L=", deca1L)
		} else if areaTag == "areaB" {
			// 四角形を作る
			rect5name, _ := MkrectB(cordz, order2, keyList, nodDode, num0, d0Num)
			// rect5name = ['L4', 'D2', 'R5', 'R6']

			log.Println("rect5name", rect5name)
			rect5L = MakeRectList(cordz, order2, rect5name)
			log.Println("rect5L=", rect5L)
			// 10角形を作る
			deca1name, deca1L = MkdecaB(cordz, order2, keyList, nodDode, num0, d0Num)
			// L4点，R5点，R6点を除外する
			log.Println("deca1name", deca1name)
			log.Println("deca1L=", deca1L)
		}

	// ２－(1)
	case "LLRLLRRRRRRR", "LRLLRRRRRRRL", "LLRRRRRRRLLR", "LRRRRRRRLLRL":
		if lrPtn[1] == "L" && lrPtn[3] == "L" && lrPtn[4] == "L" {
			num0 = lrIdx[4]
		} else if lrPtn[2] == "L" && lrPtn[3] == "L" && lrPtn[11] == "L" {
			num0 = lrIdx[3]
		} else if lrPtn[1] == "L" && lrPtn[9] == "L" && lrPtn[10] == "L" {
			num0 = lrIdx[1]
		} else if lrPtn[8] == "L" && lrPtn[9] == "L" && lrPtn[11] == "L" {
			num0 = lrIdx[0]
		}
		log.Println("２－(1)", lrjoin)

		// num0はL4点のインデックス番号
		// 12角形を１つの四角形と１つの10角形に分割する
		// 分割線bが対向する辺はL4点から３つ目と４つ目のR点が作る辺
		tp = "b2"
		cordz, order2, keyList, areaTag := DodePrepro(num0, cord2, nodDode, order, tp)
		// 大きい耳となる方の四角形を分割し，残りで10角形を作る
		if areaTag == "areaA" {
			// areaAはareaBより必ず小さい
			// TODO:
		} else if areaTag == "areaB" {
			// 四角形を作る
			rect5name, _ := MkrectB2(cordz, order2, keyList, nodDode, num0, d0Num)
			// rect5name = ['D2', 'R2', 'R3', 'R4']

			log.Println("rect5name", rect5name)
			rect5L = MakeRectList(cordz, order2, rect5name)
			log.Println("rect5L=", rect5L)
			// 10角形を作る
			deca1name, deca1L = MkdecaB2(cordz, order2, keyList, nodDode, num0, d0Num)
			// R2点，R3点，R4点を除外する
			log.Println("deca1name", deca1name)
			log.Println("deca1L=", deca1L)
		}

	// ２－(2)
	case "LLRRLLRRRRRR", "LRRLLRRRRRRL", "LLRRRRRRLLRR", "LRRRRRRLLRRL":
		if lrPtn[1] == "L" && lrPtn[4] == "L" && lrPtn[5] == "L" {
			num0 = lrIdx[5]
		} else if lrPtn[3] == "L" && lrPtn[4] == "L" && lrPtn[11] == "L" {
			num0 = lrIdx[4]
		} else if lrPtn[1] == "L" && lrPtn[8] == "L" && lrPtn[9] == "L" {
			num0 = lrIdx[1]
		} else if lrPtn[7] == "L" && lrPtn[8] == "L" && lrPtn[11] == "L" {
			num0 = lrIdx[0]
		}
		log.Println("２－(2)", lrjoin)

		// num0はL4点のインデックス番号
		// 12角形を１つの四角形と１つの10角形に分割する
		// 分割線bが対向する辺はL4点から３つ目と４つ目のR点が作る辺
		tp = "b2"
		cordz, order2, keyList, areaTag := DodePrepro(num0, cord2, nodDode, order, tp)
		// 大きい耳となる方の四角形を分割し，残りで10角形を作る
		if areaTag == "areaA" {
			// areaAはareaBより必ず小さい
			// TODO:
		} else if areaTag == "areaB" {
			// 四角形を作る
			rect5name, _ := MkrectB2(cordz, order2, keyList, nodDode, num0, d0Num)
			// rect5name = ['D2', 'R3', 'R4', 'R5']

			log.Println("rect5name", rect5name)
			rect5L = MakeRectList(cordz, order2, rect5name)
			log.Println("rect5L=", rect5L)
			// 10角形を作る
			deca1name, deca1L = MkdecaB2(cordz, order2, keyList, nodDode, num0, d0Num)
			// R3点，R4点，R5点を除外する
			log.Println("deca1name", deca1name)
			log.Println("deca1L=", deca1L)
		}

		// ２－(3)
	case "LLRRRLLRRRRR", "LRRRLLRRRRRL", "LLRRRRRLLRRR", "LRRRRRLLRRRL":
		if lrPtn[1] == "L" && lrPtn[5] == "L" && lrPtn[6] == "L" {
			num0 = lrIdx[6]
		} else if lrPtn[4] == "L" && lrPtn[5] == "L" && lrPtn[11] == "L" {
			num0 = lrIdx[5]
		} else if lrPtn[1] == "L" && lrPtn[7] == "L" && lrPtn[8] == "L" {
			num0 = lrIdx[1]
		} else if lrPtn[6] == "L" && lrPtn[7] == "L" && lrPtn[11] == "L" {
			num0 = lrIdx[0]
		}
		log.Println("２－(3)", lrjoin)

		// num0はL4点のインデックス番号
		// 12角形を１つの四角形と１つの10角形に分割する
		// 分割線bが対向する辺はL4点から３つ目と４つ目のR点が作る辺
		tp = "b2"
		cordz, order2, keyList, areaTag := DodePrepro(num0, cord2, nodDode, order, tp)
		// 大きい耳となる方の四角形を分割し，残りで10角形を作る
		if areaTag == "areaA" {
			// 四角形を作る
			rect5name, _ := MkrectA(cordz, order2, keyList, nodDode, num0, d0Num)
			// rect5name = ['D1', 'L4', 'R4', 'R5']

			log.Println("rect5name", rect5name)
			rect5L = MakeRectList(cordz, order2, rect5name)
			log.Println("rect5L=", rect5L)
			// 10角形を作る
			deca1name, deca1L = MkdecaA(cordz, order2, keyList, nodDode, num0, d0Num)
			// L4点，R4点，R5点を除外する
			log.Println("deca1name", deca1name)
			log.Println("deca1L=", deca1L)
		} else if areaTag == "areaB" {
			// areaBはareaAより必ず大きい
			// しかし，分割線bはエラーになる可能性がある
			// 四角形を作る
			rect5name, _ := MkrectB2(cordz, order2, keyList, nodDode, num0, d0Num)
			// rect5name = ['D2', 'R4', 'R5', 'R6']

			log.Println("rect5name", rect5name)
			rect5L = MakeRectList(cordz, order2, rect5name)
			log.Println("rect5L=", rect5L)
			// 10角形を作る
			deca1name, deca1L = MkdecaB2(cordz, order2, keyList, nodDode, num0, d0Num)
			// R4点，R5点，R6点を除外する
			log.Println("deca1name", deca1name)
			log.Println("deca1L=", deca1L)
		}

		// ２－(4)
	case "LLRRRRLLRRRR", "LRRRRLLRRRRL":
		if lrPtn[1] == "L" && lrPtn[6] == "L" && lrPtn[7] == "L" {
			num0 = lrIdx[7]
		} else if lrPtn[5] == "L" && lrPtn[6] == "L" && lrPtn[11] == "L" {
			num0 = lrIdx[5]
		}
		log.Println("２－(4)", lrjoin)

		// num0はL4点のインデックス番号
		// 12角形を１つの四角形と１つの10角形に分割する
		// 分割線bが対向する辺はL4点から３つ目と４つ目のR点が作る辺
		tp = "b2"
		cordz, order2, keyList, areaTag := DodePrepro(num0, cord2, nodDode, order, tp)
		// 大きい耳となる方の四角形を分割し，残りで10角形を作る
		if areaTag == "areaA" {
			// 四角形を作る
			rect5name, _ := MkrectA(cordz, order2, keyList, nodDode, num0, d0Num)
			// rect5name = ['D1', 'L4', 'R5', 'R6']

			log.Println("rect5name", rect5name)
			rect5L = MakeRectList(cordz, order2, rect5name)
			log.Println("rect5L=", rect5L)
			// 10角形を作る
			deca1name, deca1L = MkdecaA(cordz, order2, keyList, nodDode, num0, d0Num)
			// L4点，R5点，R6点を除外する
			log.Println("deca1name", deca1name)
			log.Println("deca1L=", deca1L)
		} else if areaTag == "areaB" {
			// areaBはareaAより必ず大きい
			// しかし，分割線bはエラーになる可能性がある
			// 四角形を作る
			rect5name, _ := MkrectB2(cordz, order2, keyList, nodDode, num0, d0Num)
			// rect5name = ['D2', 'R5', 'R6', 'R7']

			log.Println("rect5name", rect5name)
			rect5L = MakeRectList(cordz, order2, rect5name)
			log.Println("rect5L=", rect5L)
			// 10角形を作る
			deca1name, deca1L = MkdecaB2(cordz, order2, keyList, nodDode, num0, d0Num)
			// R5点，R6点，R7点を除外する
			log.Println("deca1name", deca1name)
			log.Println("deca1L=", deca1L)
		}

	// ２－(5)
	case "LLRLRLRRRRRR", "LRLRLRRRRRRL", "LRLRRRRRRLLR", "LRRRRRRLLRLR":
		if lrPtn[1] == "L" && lrPtn[3] == "L" && lrPtn[5] == "L" {
			num0 = lrIdx[5]
		} else if lrPtn[2] == "L" && lrPtn[4] == "L" && lrPtn[11] == "L" {
			num0 = lrIdx[4]
		} else if lrPtn[2] == "L" && lrPtn[9] == "L" && lrPtn[10] == "L" {
			num0 = lrIdx[2]
		} else if lrPtn[7] == "L" && lrPtn[8] == "L" && lrPtn[10] == "L" {
			num0 = lrIdx[0]
		}
		log.Println("２－(5)", lrjoin)

		// num0はL4点のインデックス番号
		// 12角形を１つの四角形と１つの10角形に分割する
		// 分割線bが対向する辺はL4点から３つ目と４つ目のR点が作る辺
		tp = "b2"
		cordz, order2, keyList, areaTag := DodePrepro(num0, cord2, nodDode, order, tp)
		// 大きい耳となる方の四角形を分割し，残りで10角形を作る
		if areaTag == "areaA" {
			// areaAはareaBより必ず小さい
			// TODO:
		} else if areaTag == "areaB" {
			// 四角形を作る
			rect5name, _ := MkrectB2(cordz, order2, keyList, nodDode, num0, d0Num)
			// rect5name = ['D2', 'R3', 'R4', 'R5']

			log.Println("rect5name", rect5name)
			rect5L = MakeRectList(cordz, order2, rect5name)
			log.Println("rect5L=", rect5L)
			// 10角形を作る
			deca1name, deca1L = MkdecaB2(cordz, order2, keyList, nodDode, num0, d0Num)
			// R3点，R4点，R5点を除外する
			log.Println("deca1name", deca1name)
			log.Println("deca1L=", deca1L)
		}

	// ２－(6)
	case "LLRLRRLRRRRR", "LRLRRLRRRRRL", "LRRLRRRRRLLR", "LRRRRRLLRLRR":
		if lrPtn[1] == "L" && lrPtn[3] == "L" && lrPtn[6] == "L" {
			num0 = lrIdx[6]
		} else if lrPtn[2] == "L" && lrPtn[5] == "L" && lrPtn[11] == "L" {
			num0 = lrIdx[5]
		} else if lrPtn[3] == "L" && lrPtn[9] == "L" && lrPtn[10] == "L" {
			num0 = lrIdx[3]
		} else if lrPtn[6] == "L" && lrPtn[7] == "L" && lrPtn[9] == "L" {
			num0 = lrIdx[0]
		}
		log.Println("２－(6)", lrjoin)

		// num0はL4点のインデックス番号
		// 12角形を１つの四角形と１つの10角形に分割する
		cordz, order2, keyList, areaTag := DodePrepro(num0, cord2, nodDode, order, tp)
		// 大きい耳となる方の四角形を分割し，残りで10角形を作る
		if areaTag == "areaA" {
			// 四角形を作る
			rect5name, _ := MkrectA(cordz, order2, keyList, nodDode, num0, d0Num)
			// rect5name = ['D1', 'L4', 'R4', 'R5']

			log.Println("rect5name", rect5name)
			rect5L = MakeRectList(cordz, order2, rect5name)
			log.Println("rect5L=", rect5L)
			// 10角形を作る
			deca1name, deca1L = MkdecaA(cordz, order2, keyList, nodDode, num0, d0Num)
			// L4点，R4点，R5点を除外する
			log.Println("deca1name", deca1name)
			log.Println("deca1L=", deca1L)
		} else if areaTag == "areaB" {
			// 四角形を作る
			rect5name, _ := MkrectB(cordz, order2, keyList, nodDode, num0, d0Num)
			// rect5name = ['L4', 'D2', 'R2', 'R3']

			log.Println("rect5name", rect5name)
			rect5L = MakeRectList(cordz, order2, rect5name)
			log.Println("rect5L=", rect5L)
			// 10角形を作る
			deca1name, deca1L = MkdecaB(cordz, order2, keyList, nodDode, num0, d0Num)
			// L4点，R2点，R3点を除外する
			log.Println("deca1name", deca1name)
			log.Println("deca1L=", deca1L)
		}

	// ２－(7)
	case "LLRLRRRLRRRR", "LRLRRRLRRRRL", "LRRRLRRRRLLR", "LRRRRLLRLRRR":
		if lrPtn[1] == "L" && lrPtn[3] == "L" && lrPtn[7] == "L" {
			num0 = lrIdx[7]
		} else if lrPtn[2] == "L" && lrPtn[6] == "L" && lrPtn[11] == "L" {
			num0 = lrIdx[6]
		} else if lrPtn[4] == "L" && lrPtn[9] == "L" && lrPtn[10] == "L" {
			num0 = lrIdx[4]
		} else if lrPtn[5] == "L" && lrPtn[6] == "L" && lrPtn[8] == "L" {
			num0 = lrIdx[0]
		}
		log.Println("２－(7)", lrjoin)

		// num0はL4点のインデックス番号
		// 12角形を１つの四角形と１つの10角形に分割する
		cordz, order2, keyList, areaTag := DodePrepro(num0, cord2, nodDode, order, tp)
		// 大きい耳となる方の四角形を分割し，残りで10角形を作る
		if areaTag == "areaA" {
			// 四角形を作る
			rect5name, _ := MkrectA(cordz, order2, keyList, nodDode, num0, d0Num)
			// rect5name = ['D1', 'L4', 'R5', 'R6']

			log.Println("rect5name", rect5name)
			rect5L = MakeRectList(cordz, order2, rect5name)
			log.Println("rect5L=", rect5L)
			// 10角形を作る
			deca1name, deca1L = MkdecaA(cordz, order2, keyList, nodDode, num0, d0Num)
			// L4点，R5点，R6点を除外する
			log.Println("deca1name", deca1name)
			log.Println("deca1L=", deca1L)
		} else if areaTag == "areaB" {
			// 四角形を作る
			rect5name, _ := MkrectB(cordz, order2, keyList, nodDode, num0, d0Num)
			// rect5name = ['L4', 'D2', 'R3', 'R4']

			log.Println("rect5name", rect5name)
			rect5L = MakeRectList(cordz, order2, rect5name)
			log.Println("rect5L=", rect5L)
			// 10角形を作る
			deca1name, deca1L = MkdecaB(cordz, order2, keyList, nodDode, num0, d0Num)
			// L4点，R3点，R4点を除外する
			log.Println("deca1name", deca1name)
			log.Println("deca1L=", deca1L)
		}

	// ２－(8)
	case "LLRLRRRRLRRR", "LRLRRRRLRRRL", "LRRRRLRRRLLR", "LRRRLLRLRRRR":
		if lrPtn[1] == "L" && lrPtn[3] == "L" && lrPtn[8] == "L" {
			num0 = lrIdx[8]
		} else if lrPtn[2] == "L" && lrPtn[7] == "L" && lrPtn[11] == "L" {
			num0 = lrIdx[7]
		} else if lrPtn[5] == "L" && lrPtn[9] == "L" && lrPtn[10] == "L" {
			num0 = lrIdx[5]
		} else if lrPtn[4] == "L" && lrPtn[5] == "L" && lrPtn[7] == "L" {
			num0 = lrIdx[0]
		}
		log.Println("２－(8)", lrjoin)

		// num0はL4点のインデックス番号
		// 12角形を１つの四角形と１つの10角形に分割する
		cordz, order2, keyList, areaTag := DodePrepro(num0, cord2, nodDode, order, tp)
		// 大きい耳となる方の四角形を分割し，残りで10角形を作る
		if areaTag == "areaA" {
			// 四角形を作る
			rect5name, _ := MkrectA(cordz, order2, keyList, nodDode, num0, d0Num)
			// rect5name = ['D1', 'L4', 'R6', 'R7']

			log.Println("rect5name", rect5name)
			rect5L = MakeRectList(cordz, order2, rect5name)
			log.Println("rect5L=", rect5L)
			// 10角形を作る
			deca1name, deca1L = MkdecaA(cordz, order2, keyList, nodDode, num0, d0Num)
			// L4点，R6点，R7点を除外する
			log.Println("deca1name", deca1name)
			log.Println("deca1L=", deca1L)
		} else if areaTag == "areaB" {
			// 四角形を作る
			rect5name, _ := MkrectB(cordz, order2, keyList, nodDode, num0, d0Num)
			// rect5name = ['L4', 'D2', 'R4', 'R5']

			log.Println("rect5name", rect5name)
			rect5L = MakeRectList(cordz, order2, rect5name)
			log.Println("rect5L=", rect5L)
			// 10角形を作る
			deca1name, deca1L = MkdecaB(cordz, order2, keyList, nodDode, num0, d0Num)
			// L4点，R4点，R5点を除外する
			log.Println("deca1name", deca1name)
			log.Println("deca1L=", deca1L)
		}

	// ２－(9)
	case "LLRLRRRRRLRR", "LRLRRRRRLRRL", "LRRRRRLRRLLR", "LRRLLRLRRRRR":
		if lrPtn[1] == "L" && lrPtn[3] == "L" && lrPtn[9] == "L" {
			num0 = lrIdx[9]
		} else if lrPtn[2] == "L" && lrPtn[8] == "L" && lrPtn[11] == "L" {
			num0 = lrIdx[8]
		} else if lrPtn[6] == "L" && lrPtn[9] == "L" && lrPtn[10] == "L" {
			num0 = lrIdx[6]
		} else if lrPtn[3] == "L" && lrPtn[4] == "L" && lrPtn[6] == "L" {
			num0 = lrIdx[0]
		}
		log.Println("２－(9)", lrjoin)

		// num0はL4点のインデックス番号
		// 12角形を１つの四角形と１つの10角形に分割する
		cordz, order2, keyList, areaTag := DodePrepro(num0, cord2, nodDode, order, tp)
		// 大きい耳となる方の四角形を分割し，残りで10角形を作る
		if areaTag == "areaA" {
			// 四角形を作る
			rect5name, _ := MkrectA(cordz, order2, keyList, nodDode, num0, d0Num)
			// rect5name = ['D1', 'L4', 'R7', 'R8']

			log.Println("rect5name", rect5name)
			rect5L = MakeRectList(cordz, order2, rect5name)
			log.Println("rect5L=", rect5L)
			// 10角形を作る
			deca1name, deca1L = MkdecaA(cordz, order2, keyList, nodDode, num0, d0Num)
			// L4点，R7点，R8点を除外する
			log.Println("deca1name", deca1name)
			log.Println("deca1L=", deca1L)
		} else if areaTag == "areaB" {
			// 四角形を作る
			rect5name, _ := MkrectB(cordz, order2, keyList, nodDode, num0, d0Num)
			// rect5name = ['L4', 'D2', 'R5', 'R6']

			log.Println("rect5name", rect5name)
			rect5L = MakeRectList(cordz, order2, rect5name)
			log.Println("rect5L=", rect5L)
			// 10角形を作る
			deca1name, deca1L = MkdecaB(cordz, order2, keyList, nodDode, num0, d0Num)
			// L4点，R5点，R6点を除外する
			log.Println("deca1name", deca1name)
			log.Println("deca1L=", deca1L)
		}

	// ２－(10)
	case "LLRLRRRRRRLR", "LRLRRRRRRLRL", "LRRRRRRLRLLR", "LRLLRLRRRRRR":
		if lrPtn[1] == "L" && lrPtn[3] == "L" && lrPtn[10] == "L" {
			num0 = lrIdx[10]
		} else if lrPtn[2] == "L" && lrPtn[9] == "L" && lrPtn[11] == "L" {
			num0 = lrIdx[9]
		} else if lrPtn[7] == "L" && lrPtn[9] == "L" && lrPtn[10] == "L" {
			num0 = lrIdx[7]
		} else if lrPtn[2] == "L" && lrPtn[3] == "L" && lrPtn[5] == "L" {
			num0 = lrIdx[0]
		}
		log.Println("２－(10)", lrjoin)

		// num0はL4点のインデックス番号
		// 12角形を１つの四角形と１つの10角形に分割する
		// 分割線aが対向する辺はL4点から３つ前と４つ前のR点が作る辺
		tp = "a2"
		cordz, order2, keyList, areaTag := DodePrepro(num0, cord2, nodDode, order, tp)
		// 大きい耳となる方の四角形を分割し，残りで10角形を作る
		if areaTag == "areaA" {
			// 四角形を作る
			rect5name, _ := MkrectA3(cordz, order2, keyList, nodDode, num0, d0Num)
			// rect5name = ['R7', 'D1', 'R5', 'R6']

			log.Println("rect5name", rect5name)
			rect5L = MakeRectList(cordz, order2, rect5name)
			log.Println("rect5L=", rect5L)
			// 10角形を作る
			deca1name, deca1L = MkdecaA2(cordz, order2, keyList, nodDode, num0, d0Num)
			// R5点，R6点，R7点を除外する
			log.Println("deca1name", deca1name)
			log.Println("deca1L=", deca1L)
		}
		// areaBはareaAより必ず小さい

	// ２－(11)
	case "LLRRLRLRRRRR", "LRRLRLRRRRRL", "LRLRRRRRLLRR", "LRRRRRLLRRLR":
		if lrPtn[1] == "L" && lrPtn[4] == "L" && lrPtn[6] == "L" {
			num0 = lrIdx[6]
		} else if lrPtn[3] == "L" && lrPtn[5] == "L" && lrPtn[11] == "L" {
			num0 = lrIdx[5]
		} else if lrPtn[2] == "L" && lrPtn[8] == "L" && lrPtn[9] == "L" {
			num0 = lrIdx[2]
		} else if lrPtn[6] == "L" && lrPtn[7] == "L" && lrPtn[10] == "L" {
			num0 = lrIdx[0]
		}
		log.Println("２－(11)", lrjoin)

		// num0はL4点のインデックス番号
		// 12角形を１つの四角形と１つの10角形に分割する
		// 分割線bが対向する辺はL4点から３つ目と４つ目のR点が作る辺
		tp = "b2"
		cordz, order2, keyList, areaTag := DodePrepro(num0, cord2, nodDode, order, tp)
		// 大きい耳となる方の四角形を分割し，残りで10角形を作る
		if areaTag == "areaA" {
			// areaBはareaAより必ず大きい
			// しかし，分割線bはエラーになる可能性がある
			// 四角形を作る
			rect5name, _ := MkrectA(cordz, order2, keyList, nodDode, num0, d0Num)
			// rect5name = ['D1', 'L4', 'R4', 'R5']

			log.Println("rect5name", rect5name)
			rect5L = MakeRectList(cordz, order2, rect5name)
			log.Println("rect5L=", rect5L)
			// 10角形を作る
			deca1name, deca1L = MkdecaA(cordz, order2, keyList, nodDode, num0, d0Num)
			// L4点，R4点，R5点を除外する
			log.Println("deca1name", deca1name)
			log.Println("deca1L=", deca1L)
		} else if areaTag == "areaB" {
			// 四角形を作る
			rect5name, _ := MkrectB2(cordz, order2, keyList, nodDode, num0, d0Num)
			// rect5name = ['D2', 'R4', 'R5', 'R6']

			log.Println("rect5name", rect5name)
			rect5L = MakeRectList(cordz, order2, rect5name)
			log.Println("rect5L=", rect5L)
			// 10角形を作る
			deca1name, deca1L = MkdecaB2(cordz, order2, keyList, nodDode, num0, d0Num)
			// R4点，R5点，R6点を除外する
			log.Println("deca1name", deca1name)
			log.Println("deca1L=", deca1L)
		}

	// ２－(12)
	case "LLRRLRRLRRRR", "LRRLRRLRRRRL", "LRRLRRRRLLRR", "LRRRRLLRRLRR":
		if lrPtn[1] == "L" && lrPtn[4] == "L" && lrPtn[7] == "L" {
			num0 = lrIdx[7]
		} else if lrPtn[3] == "L" && lrPtn[6] == "L" && lrPtn[11] == "L" {
			num0 = lrIdx[6]
		} else if lrPtn[3] == "L" && lrPtn[8] == "L" && lrPtn[9] == "L" {
			num0 = lrIdx[3]
		} else if lrPtn[5] == "L" && lrPtn[6] == "L" && lrPtn[9] == "L" {
			num0 = lrIdx[0]
		}
		log.Println("２－(12)", lrjoin)

		// num0はL4点のインデックス番号
		// 12角形を１つの四角形と１つの10角形に分割する
		cordz, order2, keyList, areaTag := DodePrepro(num0, cord2, nodDode, order, tp)
		// 大きい耳となる方の四角形を分割し，残りで10角形を作る
		if areaTag == "areaA" {
			// 四角形を作る
			rect5name, _ := MkrectA(cordz, order2, keyList, nodDode, num0, d0Num)
			// rect5name = ['D1', 'L4', 'R5', 'R6']

			log.Println("rect5name", rect5name)
			rect5L = MakeRectList(cordz, order2, rect5name)
			log.Println("rect5L=", rect5L)
			// 10角形を作る
			deca1name, deca1L = MkdecaA(cordz, order2, keyList, nodDode, num0, d0Num)
			// L4点，R5点，R6点を除外する
			log.Println("deca1name", deca1name)
			log.Println("deca1L=", deca1L)
		} else if areaTag == "areaB" {
			// 四角形を作る
			rect5name, _ := MkrectB(cordz, order2, keyList, nodDode, num0, d0Num)
			// rect5name = ['L4', 'D2', 'R3', 'R4']

			log.Println("rect5name", rect5name)
			rect5L = MakeRectList(cordz, order2, rect5name)
			log.Println("rect5L=", rect5L)
			// 10角形を作る
			deca1name, deca1L = MkdecaB(cordz, order2, keyList, nodDode, num0, d0Num)
			// L4点，R3点，R4点を除外する
			log.Println("deca1name", deca1name)
			log.Println("deca1L=", deca1L)
		}

	// ２－(13)
	case "LLRRLRRRLRRR", "LRRLRRRLRRRL", "LRRRLRRRLLRR", "LRRRLLRRLRRR":
		if lrPtn[1] == "L" && lrPtn[4] == "L" && lrPtn[8] == "L" {
			num0 = lrIdx[8]
		} else if lrPtn[3] == "L" && lrPtn[7] == "L" && lrPtn[11] == "L" {
			num0 = lrIdx[7]
		} else if lrPtn[4] == "L" && lrPtn[8] == "L" && lrPtn[9] == "L" {
			num0 = lrIdx[4]
		} else if lrPtn[4] == "L" && lrPtn[5] == "L" && lrPtn[8] == "L" {
			num0 = lrIdx[0]
		}
		log.Println("２－(13)", lrjoin)

		// num0はL4点のインデックス番号
		// 12角形を１つの四角形と１つの10角形に分割する
		cordz, order2, keyList, areaTag := DodePrepro(num0, cord2, nodDode, order, tp)
		// 大きい耳となる方の四角形を分割し，残りで10角形を作る
		if areaTag == "areaA" {
			// 四角形を作る
			rect5name, _ := MkrectA(cordz, order2, keyList, nodDode, num0, d0Num)
			// rect5name = ['D1', 'L4', 'R6', 'R7']

			log.Println("rect5name", rect5name)
			rect5L = MakeRectList(cordz, order2, rect5name)
			log.Println("rect5L=", rect5L)
			// 10角形を作る
			deca1name, deca1L = MkdecaA(cordz, order2, keyList, nodDode, num0, d0Num)
			// L4点，R6点，R7点を除外する
			log.Println("deca1name", deca1name)
			log.Println("deca1L=", deca1L)
		} else if areaTag == "areaB" {
			// 四角形を作る
			rect5name, _ := MkrectB(cordz, order2, keyList, nodDode, num0, d0Num)
			// rect5name = ['L4', 'D2', 'R4', 'R5']

			log.Println("rect5name", rect5name)
			rect5L = MakeRectList(cordz, order2, rect5name)
			log.Println("rect5L=", rect5L)
			// 10角形を作る
			deca1name, deca1L = MkdecaB(cordz, order2, keyList, nodDode, num0, d0Num)
			// L4点，R4点，R5点を除外する
			log.Println("deca1name", deca1name)
			log.Println("deca1L=", deca1L)
		}

	// ２－(14)
	case "LLRRLRRRRLRR", "LRRLRRRRLRRL", "LRRRRLRRLLRR", "LRRLLRRLRRRR":
		if lrPtn[1] == "L" && lrPtn[4] == "L" && lrPtn[9] == "L" {
			num0 = lrIdx[9]
		} else if lrPtn[3] == "L" && lrPtn[8] == "L" && lrPtn[11] == "L" {
			num0 = lrIdx[8]
		} else if lrPtn[5] == "L" && lrPtn[8] == "L" && lrPtn[9] == "L" {
			num0 = lrIdx[5]
		} else if lrPtn[3] == "L" && lrPtn[4] == "L" && lrPtn[7] == "L" {
			num0 = lrIdx[0]
		}
		log.Println("２－(14)", lrjoin)

		// num0はL4点のインデックス番号
		// 12角形を１つの四角形と１つの10角形に分割する
		cordz, order2, keyList, areaTag := DodePrepro(num0, cord2, nodDode, order, tp)
		// 大きい耳となる方の四角形を分割し，残りで10角形を作る
		if areaTag == "areaA" {
			// 四角形を作る
			rect5name, _ := MkrectA(cordz, order2, keyList, nodDode, num0, d0Num)
			// rect5name = ['D1', 'L4', 'R7', 'R8']

			log.Println("rect5name", rect5name)
			rect5L = MakeRectList(cordz, order2, rect5name)
			log.Println("rect5L=", rect5L)
			// 10角形を作る
			deca1name, deca1L = MkdecaA(cordz, order2, keyList, nodDode, num0, d0Num)
			// L4点，R7点，R8点を除外する
			log.Println("deca1name", deca1name)
			log.Println("deca1L=", deca1L)
		} else if areaTag == "areaB" {
			// 四角形を作る
			rect5name, _ := MkrectB(cordz, order2, keyList, nodDode, num0, d0Num)
			// rect5name = ['L4', 'D2', 'R5', 'R6']

			log.Println("rect5name", rect5name)
			rect5L = MakeRectList(cordz, order2, rect5name)
			log.Println("rect5L=", rect5L)
			// 10角形を作る
			deca1name, deca1L = MkdecaB(cordz, order2, keyList, nodDode, num0, d0Num)
			// L4点，R5点，R6点を除外する
			log.Println("deca1name", deca1name)
			log.Println("deca1L=", deca1L)
		}

	// ２－(15)
	case "LLRRLRRRRRLR", "LRRLRRRRRLRL", "LRRRRRLRLLRR", "LRLLRRLRRRRR":
		if lrPtn[1] == "L" && lrPtn[4] == "L" && lrPtn[10] == "L" {
			num0 = lrIdx[10]
		} else if lrPtn[3] == "L" && lrPtn[9] == "L" && lrPtn[11] == "L" {
			num0 = lrIdx[9]
		} else if lrPtn[6] == "L" && lrPtn[8] == "L" && lrPtn[9] == "L" {
			num0 = lrIdx[6]
		} else if lrPtn[2] == "L" && lrPtn[3] == "L" && lrPtn[6] == "L" {
			num0 = lrIdx[0]
		}
		log.Println("２－(15)", lrjoin)

		// num0はL4点のインデックス番号
		// 12角形を１つの四角形と１つの10角形に分割する
		// 分割線aが対向する辺はL4点から３つ前と４つ前のR点が作る辺
		tp = "a2"
		cordz, order2, keyList, areaTag := DodePrepro(num0, cord2, nodDode, order, tp)
		// 大きい耳となる方の四角形を分割し，残りで10角形を作る
		if areaTag == "areaA" {
			// 四角形を作る
			rect5name, _ := MkrectA3(cordz, order2, keyList, nodDode, num0, d0Num)
			// rect5name = ['R7', 'D1', 'R5', 'R6']

			log.Println("rect5name", rect5name)
			rect5L = MakeRectList(cordz, order2, rect5name)
			log.Println("rect5L=", rect5L)
			// 10角形を作る
			deca1name, deca1L = MkdecaA2(cordz, order2, keyList, nodDode, num0, d0Num)
			// R5点，R6点，R7点を除外する
			log.Println("deca1name", deca1name)
			log.Println("deca1L=", deca1L)
		}
		// areaBはareaAより必ず小さい

	// ２－(16)
	case "LLRRRLRLRRRR", "LRRRLRLRRRRL", "LRLRRRRLLRRR", "LRRRRLLRRRLR":
		if lrPtn[1] == "L" && lrPtn[5] == "L" && lrPtn[7] == "L" {
			num0 = lrIdx[7]
		} else if lrPtn[4] == "L" && lrPtn[6] == "L" && lrPtn[11] == "L" {
			num0 = lrIdx[6]
		} else if lrPtn[2] == "L" && lrPtn[7] == "L" && lrPtn[8] == "L" {
			num0 = lrIdx[2]
		} else if lrPtn[5] == "L" && lrPtn[6] == "L" && lrPtn[10] == "L" {
			num0 = lrIdx[0]
		}
		log.Println("２－(16)", lrjoin)

		// num0はL4点のインデックス番号
		// 12角形を１つの四角形と１つの10角形に分割する
		// 分割線bが対向する辺はL4点から３つ目と４つ目のR点が作る辺
		tp = "b2"
		cordz, order2, keyList, areaTag := DodePrepro(num0, cord2, nodDode, order, tp)
		// 大きい耳となる方の四角形を分割し，残りで10角形を作る
		if areaTag == "areaA" {
			// 四角形を作る
			rect5name, _ := MkrectA(cordz, order2, keyList, nodDode, num0, d0Num)
			// rect5name = ['D1', 'L4', 'R5', 'R6']

			log.Println("rect5name", rect5name)
			rect5L = MakeRectList(cordz, order2, rect5name)
			log.Println("rect5L=", rect5L)
			// 10角形を作る
			deca1name, deca1L = MkdecaA(cordz, order2, keyList, nodDode, num0, d0Num)
			// L4点，R5点，R6点を除外する
			log.Println("deca1name", deca1name)
			log.Println("deca1L=", deca1L)
		} else if areaTag == "areaB" {
			// areaBはareaAより必ず大きい
			// しかし，分割線bはエラーになる可能性がある
			// 四角形を作る
			rect5name, _ := MkrectB2(cordz, order2, keyList, nodDode, num0, d0Num)
			// rect5name = ['D2', 'R5', 'R6', 'R7']

			log.Println("rect5name", rect5name)
			rect5L = MakeRectList(cordz, order2, rect5name)
			log.Println("rect5L=", rect5L)
			// 10角形を作る
			deca1name, deca1L = MkdecaB2(cordz, order2, keyList, nodDode, num0, d0Num)
			// R5点，R6点，R7点を除外する
			log.Println("deca1name", deca1name)
			log.Println("deca1L=", deca1L)
		}

	// ２－(17)
	case "LLRRRLRRLRRR", "LRRRLRRLRRRL", "LRRLRRRLLRRR", "LRRRLLRRRLRR":
		if lrPtn[1] == "L" && lrPtn[5] == "L" && lrPtn[8] == "L" {
			num0 = lrIdx[8]
		} else if lrPtn[4] == "L" && lrPtn[7] == "L" && lrPtn[11] == "L" {
			num0 = lrIdx[7]
		} else if lrPtn[3] == "L" && lrPtn[7] == "L" && lrPtn[8] == "L" {
			num0 = lrIdx[3]
		} else if lrPtn[4] == "L" && lrPtn[5] == "L" && lrPtn[9] == "L" {
			num0 = lrIdx[0]
		}
		log.Println("２－(17)", lrjoin)

		// num0はL4点のインデックス番号
		// 12角形を１つの四角形と１つの10角形に分割する
		cordz, order2, keyList, areaTag := DodePrepro(num0, cord2, nodDode, order, tp)
		// 大きい耳となる方の四角形を分割し，残りで10角形を作る
		if areaTag == "areaA" {
			// 四角形を作る
			rect5name, _ := MkrectA(cordz, order2, keyList, nodDode, num0, d0Num)
			// rect5name = ['D1', 'L4', 'R6', 'R7']

			log.Println("rect5name", rect5name)
			rect5L = MakeRectList(cordz, order2, rect5name)
			log.Println("rect5L=", rect5L)
			// 10角形を作る
			deca1name, deca1L = MkdecaA(cordz, order2, keyList, nodDode, num0, d0Num)
			// L4点，R6点，R7点を除外する
			log.Println("deca1name", deca1name)
			log.Println("deca1L=", deca1L)
		} else if areaTag == "areaB" {
			// 四角形を作る
			rect5name, _ := MkrectB(cordz, order2, keyList, nodDode, num0, d0Num)
			// rect5name = ['L4', 'D2', 'R4', 'R5']

			log.Println("rect5name", rect5name)
			rect5L = MakeRectList(cordz, order2, rect5name)
			log.Println("rect5L=", rect5L)
			// 10角形を作る
			deca1name, deca1L = MkdecaB(cordz, order2, keyList, nodDode, num0, d0Num)
			// L4点，R4点，R5点を除外する
			log.Println("deca1name", deca1name)
			log.Println("deca1L=", deca1L)
		}

	// ２－(18)
	case "LLRRRLRRRLRR", "LRRRLRRRLRRL", "LRRRLRRLLRRR", "LRRLLRRRLRRR":
		if lrPtn[1] == "L" && lrPtn[5] == "L" && lrPtn[9] == "L" {
			num0 = lrIdx[9]
		} else if lrPtn[4] == "L" && lrPtn[8] == "L" && lrPtn[11] == "L" {
			num0 = lrIdx[8]
		} else if lrPtn[4] == "L" && lrPtn[7] == "L" && lrPtn[8] == "L" {
			num0 = lrIdx[4]
		} else if lrPtn[3] == "L" && lrPtn[4] == "L" && lrPtn[8] == "L" {
			num0 = lrIdx[0]
		}
		log.Println("２－(18)", lrjoin)

		// num0はL4点のインデックス番号
		// 12角形を１つの四角形と１つの10角形に分割する
		cordz, order2, keyList, areaTag := DodePrepro(num0, cord2, nodDode, order, tp)
		// 大きい耳となる方の四角形を分割し，残りで10角形を作る
		if areaTag == "areaA" {
			// 四角形を作る
			rect5name, _ := MkrectA(cordz, order2, keyList, nodDode, num0, d0Num)
			// rect5name = ['D1', 'L4', 'R7', 'R8']

			log.Println("rect5name", rect5name)
			rect5L = MakeRectList(cordz, order2, rect5name)
			log.Println("rect5L=", rect5L)
			// 10角形を作る
			deca1name, deca1L = MkdecaA(cordz, order2, keyList, nodDode, num0, d0Num)
			// L4点，R7点，R8点を除外する
			log.Println("deca1name", deca1name)
			log.Println("deca1L=", deca1L)
		} else if areaTag == "areaB" {
			// 四角形を作る
			rect5name, _ := MkrectB(cordz, order2, keyList, nodDode, num0, d0Num)
			// rect5name = ['L4', 'D2', 'R5', 'R6']

			log.Println("rect5name", rect5name)
			rect5L = MakeRectList(cordz, order2, rect5name)
			log.Println("rect5L=", rect5L)
			// 10角形を作る
			deca1name, deca1L = MkdecaB(cordz, order2, keyList, nodDode, num0, d0Num)
			// L4点，R5点，R6点を除外する
			log.Println("deca1name", deca1name)
			log.Println("deca1L=", deca1L)
		}

	// ２－(19)
	case "LLRRRLRRRRLR", "LRRRLRRRRLRL", "LRRRRLRLLRRR", "LRLLRRRLRRRR":
		if lrPtn[1] == "L" && lrPtn[5] == "L" && lrPtn[10] == "L" {
			num0 = lrIdx[10]
		} else if lrPtn[4] == "L" && lrPtn[9] == "L" && lrPtn[11] == "L" {
			num0 = lrIdx[9]
		} else if lrPtn[5] == "L" && lrPtn[7] == "L" && lrPtn[8] == "L" {
			num0 = lrIdx[5]
		} else if lrPtn[2] == "L" && lrPtn[3] == "L" && lrPtn[7] == "L" {
			num0 = lrIdx[0]
		}
		log.Println("２－(19)", lrjoin)

		// num0はL4点のインデックス番号
		// 12角形を１つの四角形と１つの10角形に分割する
		cordz, order2, keyList, areaTag := DodePrepro(num0, cord2, nodDode, order, tp)
		// 大きい耳となる方の四角形を分割し，残りで10角形を作る
		// 分割線aが対向する辺はL4点から３つ前と４つ前のR点が作る辺
		tp = "a2"
		if areaTag == "areaA" {
			// 四角形を作る
			rect5name, _ := MkrectA3(cordz, order2, keyList, nodDode, num0, d0Num)
			// rect5name = ['R7', 'D1', 'R5', 'R6']

			log.Println("rect5name", rect5name)
			rect5L = MakeRectList(cordz, order2, rect5name)
			log.Println("rect5L=", rect5L)
			// 10角形を作る
			deca1name, deca1L = MkdecaA2(cordz, order2, keyList, nodDode, num0, d0Num)
			// R5点，R6点，R7点を除外する
			log.Println("deca1name", deca1name)
			log.Println("deca1L=", deca1L)
		} else if areaTag == "areaB" {
			// areaAはareaBより必ず大きい
			// しかし，分割線aはエラーになる可能性がある
			// 四角形を作る
			rect5name, _ := MkrectB(cordz, order2, keyList, nodDode, num0, d0Num)
			// rect5name = ['L4', 'D2', 'R6', 'R7']

			log.Println("rect5name", rect5name)
			rect5L = MakeRectList(cordz, order2, rect5name)
			log.Println("rect5L=", rect5L)
			// 10角形を作る
			deca1name, deca1L = MkdecaB(cordz, order2, keyList, nodDode, num0, d0Num)
			// L4点，R6点，R7点を除外する
			log.Println("deca1name", deca1name)
			log.Println("deca1L=", deca1L)
		}

	// ２－(20)
	case "LLRRRRLRLRRR", "LRRRRLRLRRRL", "LRLRRRLLRRRR", "LRRRLLRRRRLR":
		if lrPtn[1] == "L" && lrPtn[6] == "L" && lrPtn[8] == "L" {
			num0 = lrIdx[8]
		} else if lrPtn[5] == "L" && lrPtn[7] == "L" && lrPtn[11] == "L" {
			num0 = lrIdx[7]
		} else if lrPtn[2] == "L" && lrPtn[6] == "L" && lrPtn[7] == "L" {
			num0 = lrIdx[2]
		} else if lrPtn[4] == "L" && lrPtn[5] == "L" && lrPtn[10] == "L" {
			num0 = lrIdx[0]
		}
		log.Println("２－(20)", lrjoin)

		// num0はL4点のインデックス番号
		// 12角形を１つの四角形と１つの10角形に分割する
		// 分割線bが対向する辺はL4点から３つ目と４つ目のR点が作る辺
		tp = "b2"
		cordz, order2, keyList, areaTag := DodePrepro(num0, cord2, nodDode, order, tp)
		// 大きい耳となる方の四角形を分割し，残りで10角形を作る
		if areaTag == "areaA" {
			// areaBはareaAより必ず大きい
			// しかし，分割線bはエラーになる可能性がある
			// 四角形を作る
			rect5name, _ := MkrectA(cordz, order2, keyList, nodDode, num0, d0Num)
			// rect5name = ['D1', 'L4', 'R6', 'R7']

			log.Println("rect5name", rect5name)
			rect5L = MakeRectList(cordz, order2, rect5name)
			log.Println("rect5L=", rect5L)
			// 10角形を作る
			deca1name, deca1L = MkdecaA(cordz, order2, keyList, nodDode, num0, d0Num)
			// L4点，R6点，R7点を除外する
			log.Println("deca1name", deca1name)
			log.Println("deca1L=", deca1L)
		} else if areaTag == "areaB" {
			// 四角形を作る
			rect5name, _ := MkrectB2(cordz, order2, keyList, nodDode, num0, d0Num)
			// rect5name = ['D2', 'R6', 'R7', 'R8']

			log.Println("rect5name", rect5name)
			rect5L = MakeRectList(cordz, order2, rect5name)
			log.Println("rect5L=", rect5L)
			// 10角形を作る
			deca1name, deca1L = MkdecaB2(cordz, order2, keyList, nodDode, num0, d0Num)
			// R6点，R7点，R8点を除外する
			log.Println("deca1name", deca1name)
			log.Println("deca1L=", deca1L)
		}

	// ２－(21)
	case "LLRRRRLRRLRR", "LRRRRLRRLRRL", "LRRLRRLLRRRR", "LRRLLRRRRLRR":
		if lrPtn[1] == "L" && lrPtn[6] == "L" && lrPtn[9] == "L" {
			num0 = lrIdx[9]
		} else if lrPtn[5] == "L" && lrPtn[8] == "L" && lrPtn[11] == "L" {
			num0 = lrIdx[8]
		} else if lrPtn[3] == "L" && lrPtn[6] == "L" && lrPtn[7] == "L" {
			num0 = lrIdx[3]
		} else if lrPtn[3] == "L" && lrPtn[4] == "L" && lrPtn[9] == "L" {
			num0 = lrIdx[0]
		}
		log.Println("２－(21)", lrjoin)

		// num0はL4点のインデックス番号
		// 12角形を１つの四角形と１つの10角形に分割する
		cordz, order2, keyList, areaTag := DodePrepro(num0, cord2, nodDode, order, tp)
		// 大きい耳となる方の四角形を分割し，残りで10角形を作る
		if areaTag == "areaA" {
			// 四角形を作る
			rect5name, _ := MkrectA(cordz, order2, keyList, nodDode, num0, d0Num)
			// rect5name = ['D1', 'L4', 'R7', 'R8']

			log.Println("rect5name", rect5name)
			rect5L = MakeRectList(cordz, order2, rect5name)
			log.Println("rect5L=", rect5L)
			// 10角形を作る
			deca1name, deca1L = MkdecaA(cordz, order2, keyList, nodDode, num0, d0Num)
			// L4点，R7点，R8点を除外する
			log.Println("deca1name", deca1name)
			log.Println("deca1L=", deca1L)
		} else if areaTag == "areaB" {
			// 四角形を作る
			rect5name, _ := MkrectB(cordz, order2, keyList, nodDode, num0, d0Num)
			// rect5name = ['L4', 'D2', 'R5', 'R6']

			log.Println("rect5name", rect5name)
			rect5L = MakeRectList(cordz, order2, rect5name)
			log.Println("rect5L=", rect5L)
			// 10角形を作る
			deca1name, deca1L = MkdecaB(cordz, order2, keyList, nodDode, num0, d0Num)
			// L4点，R5点，R6点を除外する
			log.Println("deca1name", deca1name)
			log.Println("deca1L=", deca1L)
		}

	// ２－(22)
	case "LLRRRRLRRRLR", "LRRRRLRRRLRL", "LRRRLRLLRRRR", "LRLLRRRRLRRR":
		if lrPtn[1] == "L" && lrPtn[6] == "L" && lrPtn[10] == "L" {
			num0 = lrIdx[10]
		} else if lrPtn[5] == "L" && lrPtn[9] == "L" && lrPtn[11] == "L" {
			num0 = lrIdx[9]
		} else if lrPtn[4] == "L" && lrPtn[6] == "L" && lrPtn[7] == "L" {
			num0 = lrIdx[4]
		} else if lrPtn[2] == "L" && lrPtn[3] == "L" && lrPtn[8] == "L" {
			num0 = lrIdx[0]
		}
		log.Println("２－(22)", lrjoin)

		// num0はL4点のインデックス番号
		// 12角形を１つの四角形と１つの10角形に分割する
		cordz, order2, keyList, areaTag := DodePrepro(num0, cord2, nodDode, order, tp)
		// 大きい耳となる方の四角形を分割し，残りで10角形を作る
		// 分割線aが対向する辺はL4点から３つ前と４つ前のR点が作る辺
		tp = "a2"
		if areaTag == "areaA" {
			// 四角形を作る
			rect5name, _ := MkrectA3(cordz, order2, keyList, nodDode, num0, d0Num)
			// rect5name = ['R7', 'D1', 'R5', 'R6']

			log.Println("rect5name", rect5name)
			rect5L = MakeRectList(cordz, order2, rect5name)
			log.Println("rect5L=", rect5L)
			// 10角形を作る
			deca1name, deca1L = MkdecaA2(cordz, order2, keyList, nodDode, num0, d0Num)
			// R5点，R6点，R7点を除外する
			log.Println("deca1name", deca1name)
			log.Println("deca1L=", deca1L)
		} else if areaTag == "areaB" {
			// areaAはareaBより必ず大きい
			// しかし，分割線aはエラーになる可能性がある
			// 四角形を作る
			rect5name, _ := MkrectB(cordz, order2, keyList, nodDode, num0, d0Num)
			// rect5name = ['L4', 'D2', 'R6', 'R7']

			log.Println("rect5name", rect5name)
			rect5L = MakeRectList(cordz, order2, rect5name)
			log.Println("rect5L=", rect5L)
			// 10角形を作る
			deca1name, deca1L = MkdecaB(cordz, order2, keyList, nodDode, num0, d0Num)
			// L4点，R6点，R7点を除外する
			log.Println("deca1name", deca1name)
			log.Println("deca1L=", deca1L)
		}

	// ２－(23)
	case "LLRRRRRLRLRR", "LRRRRRLRLRRL", "LRLRRLLRRRRR", "LRRLLRRRRRLR":
		if lrPtn[1] == "L" && lrPtn[7] == "L" && lrPtn[9] == "L" {
			num0 = lrIdx[9]
		} else if lrPtn[6] == "L" && lrPtn[8] == "L" && lrPtn[11] == "L" {
			num0 = lrIdx[8]
		} else if lrPtn[2] == "L" && lrPtn[5] == "L" && lrPtn[6] == "L" {
			num0 = lrIdx[2]
		} else if lrPtn[3] == "L" && lrPtn[4] == "L" && lrPtn[10] == "L" {
			num0 = lrIdx[0]
		}
		log.Println("２－(23)", lrjoin)

		// num0はL4点のインデックス番号
		// 12角形を１つの四角形と１つの10角形に分割する
		cordz, order2, keyList, areaTag := DodePrepro(num0, cord2, nodDode, order, tp)
		// 大きい耳となる方の四角形を分割し，残りで10角形を作る
		if areaTag == "areaA" {
			// 四角形を作る
			rect5name, _ := MkrectA(cordz, order2, keyList, nodDode, num0, d0Num)
			// rect5name = ['D1', 'L4', 'R7', 'R8']

			log.Println("rect5name", rect5name)
			rect5L = MakeRectList(cordz, order2, rect5name)
			log.Println("rect5L=", rect5L)
			// 10角形を作る
			deca1name, deca1L = MkdecaA(cordz, order2, keyList, nodDode, num0, d0Num)
			// L4点，R7点，R8点を除外する
			log.Println("deca1name", deca1name)
			log.Println("deca1L=", deca1L)
		}
		// areaBは成立しない

	// ２－(24)
	case "LLRRRRRLRRLR", "LRRRRRLRRLRL", "LRRLRLLRRRRR", "LRLLRRRRRLRR":
		if lrPtn[1] == "L" && lrPtn[7] == "L" && lrPtn[10] == "L" {
			num0 = lrIdx[10]
		} else if lrPtn[6] == "L" && lrPtn[9] == "L" && lrPtn[11] == "L" {
			num0 = lrIdx[9]
		} else if lrPtn[3] == "L" && lrPtn[5] == "L" && lrPtn[6] == "L" {
			num0 = lrIdx[3]
		} else if lrPtn[2] == "L" && lrPtn[3] == "L" && lrPtn[9] == "L" {
			num0 = lrIdx[0]
		}
		log.Println("２－(24)", lrjoin)

		// num0はL4点のインデックス番号
		// 12角形を１つの四角形と１つの10角形に分割する
		cordz, order2, keyList, areaTag := DodePrepro(num0, cord2, nodDode, order, tp)
		// 大きい耳となる方の四角形を分割し，残りで10角形を作る
		if areaTag == "areaA" {
			// areaAは成立しない
			// TODO:
		} else if areaTag == "areaB" {
			// 四角形を作る
			rect5name, _ := MkrectB(cordz, order2, keyList, nodDode, num0, d0Num)
			// rect5name = ['L4', 'D2', 'R6', 'R7']

			log.Println("rect5name", rect5name)
			rect5L = MakeRectList(cordz, order2, rect5name)
			log.Println("rect5L=", rect5L)
			// 10角形を作る
			deca1name, deca1L = MkdecaB(cordz, order2, keyList, nodDode, num0, d0Num)
			// L4点，R6点，R7点を除外する
			log.Println("deca1name", deca1name)
			log.Println("deca1L=", deca1L)
		}

	// ２－(25)
	case "LLRRRRRRLRLR", "LRRRRRRLRLRL", "LRLRLLRRRRRR", "LRLLRRRRRRLR":
		if lrPtn[1] == "L" && lrPtn[8] == "L" && lrPtn[10] == "L" {
			num0 = lrIdx[8]
		} else if lrPtn[7] == "L" && lrPtn[9] == "L" && lrPtn[11] == "L" {
			num0 = lrIdx[7]
		} else if lrPtn[2] == "L" && lrPtn[4] == "L" && lrPtn[5] == "L" {
			num0 = lrIdx[0]
		} else if lrPtn[2] == "L" && lrPtn[3] == "L" && lrPtn[10] == "L" {
			num0 = lrIdx[10]
		}
		log.Println("２－(25)", lrjoin)

		// L3点を起点に四角形を分割することができない
		// num0はL3点のインデックス番号
		// 12角形を１つの四角形と１つの10角形に分割する
		// 分割線aが対向する辺はL3点から３つ前と４つ前のR点が作る辺
		tp = "a2"
		cordz, order2, keyList, areaTag := DodePrepro(num0, cord2, nodDode, order, tp)
		// 大きい耳となる方の四角形を分割し，残りで10角形を作る
		if areaTag == "areaA" {
			// 四角形を作る
			rect5name, _ := MkrectA3(cordz, order2, keyList, nodDode, num0, d0Num)
			// rect5name = ['R6', 'D1', 'R4', 'R5']

			log.Println("rect5name", rect5name)
			rect5L = MakeRectList(cordz, order2, rect5name)
			log.Println("rect5L=", rect5L)
			// 10角形を作る
			deca1name, deca1L = MkdecaA2(cordz, order2, keyList, nodDode, num0, d0Num)
			// R4点，R5点，R6点を除外する
			log.Println("deca1name", deca1name)
			log.Println("deca1L=", deca1L)
		}
		// areaBはareaAより必ず小さい

	// ３－(1)
	case "LLLRLRRRRRRR", "LLRLRRRRRRRL", "LRLRRRRRRRLL", "LRRRRRRRLLLR":
		if lrPtn[1] == "L" && lrPtn[2] == "L" && lrPtn[4] == "L" {
			num0 = lrIdx[4]
		} else if lrPtn[1] == "L" && lrPtn[3] == "L" && lrPtn[11] == "L" {
			num0 = lrIdx[3]
		} else if lrPtn[2] == "L" && lrPtn[10] == "L" && lrPtn[11] == "L" {
			num0 = lrIdx[2]
		} else if lrPtn[8] == "L" && lrPtn[9] == "L" && lrPtn[10] == "L" {
			num0 = lrIdx[0]
		}
		log.Println("３－(1)", lrjoin)

		// num0はL4点のインデックス番号
		// 12角形を１つの四角形と１つの10角形に分割する
		// 分割線bが対向する辺はL4点から３つ目と４つ目のR点が作る辺
		tp = "b2"
		cordz, order2, keyList, areaTag := DodePrepro(num0, cord2, nodDode, order, tp)
		// 大きい耳となる方の四角形を分割し，残りで10角形を作る
		if areaTag == "areaA" {
			// areaAはareaBより必ず小さい
			// TODO:
		} else if areaTag == "areaB" {
			// 四角形を作る
			rect5name, _ := MkrectB2(cordz, order2, keyList, nodDode, num0, d0Num)
			// rect5name = ['D2', 'R2', 'R3', 'R4']

			log.Println("rect5name", rect5name)
			rect5L = MakeRectList(cordz, order2, rect5name)
			log.Println("rect5L=", rect5L)
			// 10角形を作る
			deca1name, deca1L = MkdecaB2(cordz, order2, keyList, nodDode, num0, d0Num)
			// R2点，R3点，R4点を除外する
			log.Println("deca1name", deca1name)
			log.Println("deca1L=", deca1L)
		}

	// ３－(2)
	case "LLLRRLRRRRRR", "LLRRLRRRRRRL", "LRRLRRRRRRLL", "LRRRRRRLLLRR":
		if lrPtn[1] == "L" && lrPtn[2] == "L" && lrPtn[5] == "L" {
			num0 = lrIdx[5]
		} else if lrPtn[1] == "L" && lrPtn[4] == "L" && lrPtn[11] == "L" {
			num0 = lrIdx[4]
		} else if lrPtn[3] == "L" && lrPtn[10] == "L" && lrPtn[11] == "L" {
			num0 = lrIdx[3]
		} else if lrPtn[7] == "L" && lrPtn[8] == "L" && lrPtn[9] == "L" {
			num0 = lrIdx[0]
		}
		log.Println("３－(2)", lrjoin)

		// num0はL4点のインデックス番号
		// 12角形を１つの四角形と１つの10角形に分割する
		// 分割線bが対向する辺はL4点から３つ目と４つ目のR点が作る辺
		tp = "b2"
		cordz, order2, keyList, areaTag := DodePrepro(num0, cord2, nodDode, order, tp)
		// 大きい耳となる方の四角形を分割し，残りで10角形を作る
		if areaTag == "areaA" {
			// areaBはareaAより必ず大きい
			// しかし，分割線bはエラーになる可能性がある
			// 四角形を作る
			rect5name, _ := MkrectA(cordz, order2, keyList, nodDode, num0, d0Num)
			// rect5name = ['D1', 'L4', 'R3', 'R4']

			log.Println("rect5name", rect5name)
			rect5L = MakeRectList(cordz, order2, rect5name)
			log.Println("rect5L=", rect5L)
			// 10角形を作る
			deca1name, deca1L = MkdecaA(cordz, order2, keyList, nodDode, num0, d0Num)
			// L4点，R3点，R4点を除外する
			log.Println("deca1name", deca1name)
			log.Println("deca1L=", deca1L)
		} else if areaTag == "areaB" {
			// 四角形を作る
			rect5name, _ := MkrectB2(cordz, order2, keyList, nodDode, num0, d0Num)
			// rect5name = ['D2', 'R3', 'R4', 'R5']

			log.Println("rect5name", rect5name)
			rect5L = MakeRectList(cordz, order2, rect5name)
			log.Println("rect5L=", rect5L)
			// 10角形を作る
			deca1name, deca1L = MkdecaB2(cordz, order2, keyList, nodDode, num0, d0Num)
			// R3点，R4点，R5点を除外する
			log.Println("deca1name", deca1name)
			log.Println("deca1L=", deca1L)
		}

	// ３－(3)
	case "LLLRRRLRRRRR", "LLRRRLRRRRRL", "LRRRLRRRRRLL", "LRRRRRLLLRRR":
		if lrPtn[1] == "L" && lrPtn[2] == "L" && lrPtn[6] == "L" {
			num0 = lrIdx[0]
		} else if lrPtn[1] == "L" && lrPtn[5] == "L" && lrPtn[11] == "L" {
			num0 = lrIdx[11]
		} else if lrPtn[4] == "L" && lrPtn[10] == "L" && lrPtn[11] == "L" {
			num0 = lrIdx[10]
		} else if lrPtn[6] == "L" && lrPtn[7] == "L" && lrPtn[8] == "L" {
			num0 = lrIdx[6]
		}
		log.Println("３－(3)", lrjoin)

		// L4点からの分割線a，分割線bは共にエラーとなる可能性がある
		// num0はL1点のインデックス番号
		// 12角形を１つの四角形と１つの10角形に分割する
		// 分割線aが対向する辺はL3点から３つ前と４つ前のR点が作る辺
		tp = "a2"
		cordz, order2, keyList, areaTag := DodePrepro(num0, cord2, nodDode, order, tp)
		// 大きい耳となる方の四角形を分割し，残りで10角形を作る
		if areaTag == "areaA" {
			// 四角形を作る
			rect5name, _ := MkrectA3(cordz, order2, keyList, nodDode, num0, d0Num)
			// rect5name = ['R8', 'D1', 'R6', 'R7']

			log.Println("rect5name", rect5name)
			rect5L = MakeRectList(cordz, order2, rect5name)
			log.Println("rect5L=", rect5L)
			// 10角形を作る
			deca1name, deca1L = MkdecaA2(cordz, order2, keyList, nodDode, num0, d0Num)
			// R6点，R7点，R8点を除外する
			log.Println("deca1name", deca1name)
			log.Println("deca1L=", deca1L)
		}
		// areaBはareaAより必ず小さい

	// ３－(4)
	case "LLLRRRRLRRRR", "LLRRRRLRRRRL", "LRRRRLRRRRLL", "LRRRRLLLRRRR":
		if lrPtn[1] == "L" && lrPtn[2] == "L" && lrPtn[7] == "L" {
			num0 = lrIdx[0]
		} else if lrPtn[1] == "L" && lrPtn[6] == "L" && lrPtn[11] == "L" {
			num0 = lrIdx[11]
		} else if lrPtn[5] == "L" && lrPtn[10] == "L" && lrPtn[11] == "L" {
			num0 = lrIdx[10]
		} else if lrPtn[5] == "L" && lrPtn[6] == "L" && lrPtn[7] == "L" {
			num0 = lrIdx[5]
		}
		log.Println("３－(4)", lrjoin)

		// L4点からの分割線a，分割線bは共にエラーとなる可能性がある
		// num0はL1点のインデックス番号
		// 12角形を１つの四角形と１つの10角形に分割する
		// 分割線aが対向する辺はL3点から３つ前と４つ前のR点が作る辺
		tp = "a2"
		cordz, order2, keyList, areaTag := DodePrepro(num0, cord2, nodDode, order, tp)
		// 大きい耳となる方の四角形を分割し，残りで10角形を作る
		if areaTag == "areaA" {
			// 四角形を作る
			rect5name, _ := MkrectA3(cordz, order2, keyList, nodDode, num0, d0Num)
			// rect5name = ['R8', 'D1', 'R6', 'R7']

			log.Println("rect5name", rect5name)
			rect5L = MakeRectList(cordz, order2, rect5name)
			log.Println("rect5L=", rect5L)
			// 10角形を作る
			deca1name, deca1L = MkdecaA2(cordz, order2, keyList, nodDode, num0, d0Num)
			// R6点，R7点，R8点を除外する
			log.Println("deca1name", deca1name)
			log.Println("deca1L=", deca1L)
		}
		// areaBはareaAより必ず小さい

	// ３－(5)
	case "LLLRRRRRLRRR", "LLRRRRRLRRRL", "LRRRRRLRRRLL", "LRRRLLLRRRRR":
		if lrPtn[1] == "L" && lrPtn[2] == "L" && lrPtn[8] == "L" {
			num0 = lrIdx[8]
		} else if lrPtn[1] == "L" && lrPtn[7] == "L" && lrPtn[11] == "L" {
			num0 = lrIdx[7]
		} else if lrPtn[6] == "L" && lrPtn[10] == "L" && lrPtn[11] == "L" {
			num0 = lrIdx[6]
		} else if lrPtn[4] == "L" && lrPtn[5] == "L" && lrPtn[6] == "L" {
			num0 = lrIdx[0]
		}
		log.Println("３－(5)", lrjoin)

		// num0はL4点のインデックス番号
		// 12角形を１つの四角形と１つの10角形に分割する
		cordz, order2, keyList, areaTag := DodePrepro(num0, cord2, nodDode, order, tp)
		// 大きい耳となる方の四角形を分割し，残りで10角形を作る
		if areaTag == "areaA" {
			// 四角形を作る
			rect5name, _ := MkrectA(cordz, order2, keyList, nodDode, num0, d0Num)
			// rect5name = ['D1', 'L4', 'R6', 'R7']

			log.Println("rect5name", rect5name)
			rect5L = MakeRectList(cordz, order2, rect5name)
			log.Println("rect5L=", rect5L)
			// 10角形を作る
			deca1name, deca1L = MkdecaA(cordz, order2, keyList, nodDode, num0, d0Num)
			// L4点，R6点，R7点を除外する
			log.Println("deca1name", deca1name)
			log.Println("deca1L=", deca1L)
		} else if areaTag == "areaB" {
			// 四角形を作る
			rect5name, _ := MkrectB(cordz, order2, keyList, nodDode, num0, d0Num)
			// rect5name = ['L4', 'D2', 'R4', 'R5']

			log.Println("rect5name", rect5name)
			rect5L = MakeRectList(cordz, order2, rect5name)
			log.Println("rect5L=", rect5L)
			// 10角形を作る
			deca1name, deca1L = MkdecaB(cordz, order2, keyList, nodDode, num0, d0Num)
			// L4点，R4点，R5点を除外する
			log.Println("deca1name", deca1name)
			log.Println("deca1L=", deca1L)
		}

	// ３－(6)
	case "LLLRRRRRRLRR", "LLRRRRRRLRRL", "LRRRRRRLRRLL", "LRRLLLRRRRRR":
		if lrPtn[1] == "L" && lrPtn[2] == "L" && lrPtn[9] == "L" {
			num0 = lrIdx[9]
		} else if lrPtn[1] == "L" && lrPtn[8] == "L" && lrPtn[11] == "L" {
			num0 = lrIdx[8]
		} else if lrPtn[7] == "L" && lrPtn[10] == "L" && lrPtn[11] == "L" {
			num0 = lrIdx[7]
		} else if lrPtn[3] == "L" && lrPtn[4] == "L" && lrPtn[5] == "L" {
			num0 = lrIdx[0]
		}
		log.Println("３－(6)", lrjoin)

		// num0はL4点のインデックス番号
		// 12角形を１つの四角形と１つの10角形に分割する
		cordz, order2, keyList, areaTag := DodePrepro(num0, cord2, nodDode, order, tp)
		// 大きい耳となる方の四角形を分割し，残りで10角形を作る
		if areaTag == "areaA" {
			// 四角形を作る
			rect5name, _ := MkrectA(cordz, order2, keyList, nodDode, num0, d0Num)
			// rect5name = ['D1', 'L4', 'R7', 'R8']

			log.Println("rect5name", rect5name)
			rect5L = MakeRectList(cordz, order2, rect5name)
			log.Println("rect5L=", rect5L)
			// 10角形を作る
			deca1name, deca1L = MkdecaA(cordz, order2, keyList, nodDode, num0, d0Num)
			// L4点，R7点，R8点を除外する
			log.Println("deca1name", deca1name)
			log.Println("deca1L=", deca1L)
		} else if areaTag == "areaB" {
			// areaAはareaBより必ず大きい
			// しかし，分割線aはエラーになる可能性がある
			// 四角形を作る
			rect5name, _ := MkrectB(cordz, order2, keyList, nodDode, num0, d0Num)
			// rect5name = ['L4', 'D2', 'R5', 'R6']

			log.Println("rect5name", rect5name)
			rect5L = MakeRectList(cordz, order2, rect5name)
			log.Println("rect5L=", rect5L)
			// 10角形を作る
			deca1name, deca1L = MkdecaB(cordz, order2, keyList, nodDode, num0, d0Num)
			// L4点，R5点，R6点を除外する
			log.Println("deca1name", deca1name)
			log.Println("deca1L=", deca1L)
		}

	// ３－(7)
	case "LLLRRRRRRRLR", "LLRRRRRRRLRL", "LRRRRRRRLRLL", "LRLLLRRRRRRR":
		if lrPtn[1] == "L" && lrPtn[2] == "L" && lrPtn[10] == "L" {
			num0 = lrIdx[10]
		} else if lrPtn[1] == "L" && lrPtn[9] == "L" && lrPtn[11] == "L" {
			num0 = lrIdx[9]
		} else if lrPtn[8] == "L" && lrPtn[10] == "L" && lrPtn[11] == "L" {
			num0 = lrIdx[8]
		} else if lrPtn[2] == "L" && lrPtn[3] == "L" && lrPtn[4] == "L" {
			num0 = lrIdx[0]
		}
		log.Println("３－(7)", lrjoin)

		// num0はL4点のインデックス番号
		// 12角形を１つの四角形と１つの10角形に分割する
		cordz, order2, keyList, areaTag := DodePrepro(num0, cord2, nodDode, order, tp)
		// 大きい耳となる方の四角形を分割し，残りで10角形を作る
		// 分割線aが対向する辺はL4点から３つ前と４つ前のR点が作る辺
		tp = "a2"
		if areaTag == "areaA" {
			// 四角形を作る
			rect5name, _ := MkrectA3(cordz, order2, keyList, nodDode, num0, d0Num)
			// rect5name = ['R7', 'D1', 'R5', 'R6']

			log.Println("rect5name", rect5name)
			rect5L = MakeRectList(cordz, order2, rect5name)
			log.Println("rect5L=", rect5L)
			// 10角形を作る
			deca1name, deca1L = MkdecaA2(cordz, order2, keyList, nodDode, num0, d0Num)
			// R5点，R6点，R7点を除外する
			log.Println("deca1name", deca1name)
			log.Println("deca1L=", deca1L)
		}
		// areaBはareaAより必ず小さい

	// ４－(1)
	case "LLLLRRRRRRRR", "LLLRRRRRRRRL", "LLRRRRRRRRLL", "LRRRRRRRRLLL":
		if lrPtn[1] == "L" && lrPtn[2] == "L" && lrPtn[3] == "L" {
			num0 = lrIdx[3]
		} else if lrPtn[1] == "L" && lrPtn[2] == "L" && lrPtn[11] == "L" {
			num0 = lrIdx[2]
		} else if lrPtn[1] == "L" && lrPtn[10] == "L" && lrPtn[11] == "L" {
			num0 = lrIdx[1]
		} else if lrPtn[9] == "L" && lrPtn[10] == "L" && lrPtn[11] == "L" {
			num0 = lrIdx[0]
		}
		log.Println("４－(1)", lrjoin)

		// num0はL4点のインデックス番号
		// 12角形を１つの四角形と１つの10角形に分割する
		// 分割線bが対向する辺はL4点から３つ目と４つ目のR点が作る辺
		tp = "b2"
		cordz, order2, keyList, areaTag := DodePrepro(num0, cord2, nodDode, order, tp)
		// 大きい耳となる方の四角形を分割し，残りで10角形を作る
		if areaTag == "areaA" {
			// areaAはareaBより必ず小さい
			// TODO:
		} else if areaTag == "areaB" {
			// 四角形を作る
			rect5name, _ := MkrectB2(cordz, order2, keyList, nodDode, num0, d0Num)
			// rect5name = ['D2', 'R1', 'R2', 'R3']

			log.Println("rect5name", rect5name)
			rect5L = MakeRectList(cordz, order2, rect5name)
			log.Println("rect5L=", rect5L)
			// 10角形を作る
			deca1name, deca1L = MkdecaB2(cordz, order2, keyList, nodDode, num0, d0Num)
			// R1点，R2点，R3点を除外する
			log.Println("deca1name", deca1name)
			log.Println("deca1L=", deca1L)
		}

	default:
		// TODO:
	}

	if deca1L != nil {
		log.Println("deca1L=", deca1L)
		// 頂点データ数の確認
		nod := len(deca1L)
		if nod != 8 {
			// TODO:
		}
		// 追加したD点のために外積の計算をやり直す必要がある
		extL, _, _ := TriVert(nod, deca1L)
		// L点，R点の辞書を作り直す
		_, _, orderN, lrPtn, lrIdx := Lexicogra(nod, deca1L, extL)
		log.Println("orderN=", orderN)
		// 10角形の四角形分割プログラムに渡す
		rect1L, rect2L, rect3L, rect4L, story, yane = DecaDiv(deca1L, orderN, lrPtn, lrIdx)
		log.Println("rectD1L=", rect1L)
		log.Println("rectD2L=", rect2L)
		log.Println("rectD3L=", rect3L)
		log.Println("rectD4L=", rect4L)
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

	return rect1L, rect2L, rect3L, rect4L, rect5L, story, yane
}
