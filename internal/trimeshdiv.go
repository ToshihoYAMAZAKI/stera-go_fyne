package internal

import (
	"log"
	"stera/pkg"
	"strings"
)

// 三角形の構造体
type Triangle struct {
	vtx1 []float64
	vtx2 []float64
	vtx3 []float64
}

// TriDivは多角形を三角メッシュに分割する
// func TriMeshDiv(cords [][]float64, lrPtn []string, lrIdx []int) () {
func TriMeshDiv() {
	cords := [][]float64{{-52924.99678557923, -141434.5247060217}, {-52930.62579564661, -141433.033723026}, {-52929.910950259844, -141429.46566922666}, {-52934.06967781568, -141428.29606549727}, {-52933.12284861119, -141424.12525117613}, {-52938.80125023466, -141422.8180137251}, {-52936.64424374204, -141413.35331852623}, {-52924.938312263366, -141416.047404978}, {-52926.303104903076, -141422.83272347768}, {-52924.30867775029, -141423.30818000555}, {-52925.511876394405, -141429.0767555994}, {-52923.8653693743, -141429.50155896906}}
	lrPtn := []string{"L", "R", "L", "R", "R", "R", "L", "R", "L", "R", "R", "R"}
	lrIdx := []int{2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 0, 1}

	// 三角形データのスライスを作成する
	triList := []*Triangle{}
	i := 0

	// 検索用にLR並びから半角スペースを除く
	lrjoin := strings.Join(lrPtn, "")
	log.Println("lrjoin=", lrjoin)

	// LR並びにLRRLが含まれているか調べる
	str1 := "LRRL"
	log.Println(strings.Contains(lrjoin, str1))

	// LR並びからLRRLを探して文字列位置を返す
	for strings.Contains(lrjoin, str1) {
		num := len(lrjoin)
		log.Println(strings.Index(lrjoin, str1))
		log.Println("LRRL=", strings.Index(lrjoin, str1))
		// 2つのL点を結んでできる四角形を2つに分割してリストに追加する
		// 2つのL点に挟まれたR点を削除して多角形を作り直す
		// 分離する三角形１の頂点番号
		n := strings.Index(lrjoin, str1)
		idx := lrIdx[n]
		log.Println("idx=", idx)

		tr_11 := idx
		tr_12 := (idx + 1) % num
		tr_13 := (idx + 2) % num
		// 分離する三角形１の頂点座標
		tri1 := Triangle{vtx1: cords[tr_11], vtx2: cords[tr_12], vtx3: cords[tr_13]}
		log.Println(tri1)
		triList = append(triList, &tri1)
		log.Println(triList)
		log.Println("triList=", *triList[i])
		i++

		// 分離する三角形２の頂点番号
		tr_21 := tr_11
		tr_22 := tr_13
		tr_23 := (idx + 3) % num
		tri2 := Triangle{vtx1: cords[tr_21], vtx2: cords[tr_22], vtx3: cords[tr_23]}
		log.Println(tri2)
		triList = append(triList, &tri2)
		log.Println(triList)
		log.Println("triList=", *triList[i])
		i++

		// 三角形を分離した後，多角形を作り直す
		newpoly := append(cords[:tr_12], cords[tr_23:]...)
		log.Println("newpoly=", newpoly)

		// 作り直した多角形のために外積の計算をやり直す必要がある
		extL, _, _ := pkg.TriVert(num-2, newpoly)
		log.Println("extL=", extL)
		// L点，R点の辞書を作り直す
		_, _, _, lrPtn, lrIdx = pkg.Lexicogra(num-2, newpoly, extL)
		// log.Println("order=", order)
		log.Println("lrPtn=", lrPtn)
		log.Println("lrIdx=", lrIdx)

		lrjoin = strings.Join(lrPtn, "")
		log.Println("lrjoin=", lrjoin)

		cords = newpoly

		if !strings.Contains(lrjoin, str1) {
			break
		}
	}

	// LR並びからLRLが含まれているか調べる
	str2 := "LRL"
	log.Println(strings.Index(lrjoin, str2))

	// LR並びからLRLを探して文字列位置を返す
	for strings.Contains(lrjoin, str2) {
		num := len(lrjoin)
		log.Println(strings.Index(lrjoin, str2))
		log.Println("LRL=", strings.Index(lrjoin, str2))
		// 2つのL点を結んでできる三角形を分割してリストに追加する
		// 2つのL点に挟まれたR点を削除して多角形を作り直す
		// 分離する三角形の頂点番号
		n := strings.Index(lrjoin, str2)
		idx := lrIdx[n]
		log.Println("idx=", idx)

		tr_01 := idx
		tr_02 := (idx + 1) % num
		tr_03 := (idx + 2) % num
		// 分離する三角形１の頂点座標
		tri0 := Triangle{vtx1: cords[tr_01], vtx2: cords[tr_02], vtx3: cords[tr_03]}
		log.Println(tri0)
		triList = append(triList, &tri0)
		log.Println(triList)
		log.Println("triList=", *triList[i])
		i++

		// 三角形を分離した後，多角形を作り直す
		newpoly := append(cords[:tr_02], cords[tr_03:]...)
		log.Println("newpoly=", newpoly)

		// 作り直した多角形のために外積の計算をやり直す必要がある
		extL, _, _ := pkg.TriVert(num-1, newpoly)
		log.Println("extL=", extL)
		// L点，R点の辞書を作り直す
		_, _, _, lrPtn, lrIdx = pkg.Lexicogra(num-1, newpoly, extL)
		// log.Println("order=", order)
		log.Println("lrPtn=", lrPtn)
		log.Println("lrIdx=", lrIdx)

		lrjoin = strings.Join(lrPtn, "")
		log.Println("lrjoin=", lrjoin)

		cords = newpoly

		if !strings.Contains(lrjoin, str2) {
			break
		}
	}

	// LR並びにLRRRLが含まれているか調べる
	str5 := "LRRRL"
	log.Println(strings.Contains(lrjoin, str5))

	// LR並びからLRRLを探して文字列位置を返す
	for strings.Contains(lrjoin, str5) {
		num := len(lrjoin)
		log.Println(strings.Index(lrjoin, str5))
		log.Println("LRRRL=", strings.Index(lrjoin, str5))
		// 両端のL点とこれに続く2つのR点を結んでできる2つの三角形を分割する
		// そして，両端のL点と3点のうち中央のR点を結んでできる三角形を分割する
		// 分離する三角形１の頂点番号
		n := strings.Index(lrjoin, str5)
		idx := lrIdx[n]
		log.Println("idx=", idx)

		tr_11 := idx
		tr_12 := (idx + 1) % num
		tr_13 := (idx + 2) % num
		// 分離する三角形１の頂点座標
		tri1 := Triangle{vtx1: cords[tr_11], vtx2: cords[tr_12], vtx3: cords[tr_13]}
		log.Println(tri1)
		triList = append(triList, &tri1)
		log.Println(triList)
		log.Println("triList=", *triList[i])
		i++

		// 分離する三角形２の頂点番号
		tr_21 := tr_13
		tr_22 := (idx + 3) % num
		tr_23 := (idx + 4) % num
		tri2 := Triangle{vtx1: cords[tr_21], vtx2: cords[tr_22], vtx3: cords[tr_23]}
		log.Println(tri2)
		triList = append(triList, &tri2)
		log.Println(triList)
		log.Println("triList=", *triList[i])
		i++

		// 分離する三角形２の頂点番号
		tr_31 := tr_11
		tr_32 := tr_13
		tr_33 := tr_23
		tri3 := Triangle{vtx1: cords[tr_31], vtx2: cords[tr_32], vtx3: cords[tr_33]}
		log.Println(tri3)
		triList = append(triList, &tri3)
		log.Println(triList)
		log.Println("triList=", *triList[i])
		i++

		// 三角形を分離した後，多角形を作り直す
		newpoly := append(cords[:tr_12], cords[tr_23:]...)
		log.Println("newpoly=", newpoly)

		// 作り直した多角形のために外積の計算をやり直す必要がある
		extL, _, _ := pkg.TriVert(num-3, newpoly)
		log.Println("extL=", extL)
		// L点，R点の辞書を作り直す
		_, _, _, lrPtn, lrIdx = pkg.Lexicogra(num-3, newpoly, extL)
		// log.Println("order=", order)
		log.Println("lrPtn=", lrPtn)
		log.Println("lrIdx=", lrIdx)

		lrjoin = strings.Join(lrPtn, "")
		log.Println("lrjoin=", lrjoin)

		cords = newpoly

		if !strings.Contains(lrjoin, str5) {
			break
		}
	}

	// LR並びからRRLが含まれているか調べる
	str3 := "RRL"
	log.Println(strings.Index(lrjoin, str3))

	// LR並びからLLRRを探して文字列位置を返す
	for strings.Contains(lrjoin, str3) {
		num := len(lrjoin)
		if num < 5 {
			break
		}

		log.Println(strings.Index(lrjoin, str3))
		// log.Println("LLRR=", strings.Index(lrjoin, str3))
		log.Println("RRL=", strings.Index(lrjoin, str3))
		// L点と２つ前のR点を結んでできる三角形を分割してリストに追加する
		// L点の１つ前のR点を削除して多角形を作り直す
		// 分離する三角形の頂点番号
		n := strings.Index(lrjoin, str3)
		idx := lrIdx[n]
		log.Println("idx=", idx)

		tr_31 := idx
		tr_32 := (idx + 1) % num
		tr_33 := (idx + 2) % num
		// 分離する三角形１の頂点座標
		tri3 := Triangle{vtx1: cords[tr_31], vtx2: cords[tr_32], vtx3: cords[tr_33]}
		log.Println(tri3)
		triList = append(triList, &tri3)
		log.Println(triList)
		log.Println("triList=", *triList[i])
		i++

		// 三角形を分離した後，多角形を作り直す
		newpoly := append(cords[:tr_32], cords[tr_33:]...)
		log.Println("newpoly=", newpoly)

		// 作り直した多角形のために外積の計算をやり直す必要がある
		extL, _, _ := pkg.TriVert(num-1, newpoly)
		log.Println("extL=", extL)
		// L点，R点の辞書を作り直す
		_, _, _, lrPtn, lrIdx = pkg.Lexicogra(num-1, newpoly, extL)
		// log.Println("order=", order)
		log.Println("lrPtn=", lrPtn)
		log.Println("lrIdx=", lrIdx)

		lrjoin = strings.Join(lrPtn, "")
		log.Println("lrjoin=", lrjoin)

		cords = newpoly

		if !strings.Contains(lrjoin, str3) {
			break
		}
	}

	// LR並びからLRRが含まれているか調べる
	str4 := "LRR"
	log.Println(strings.Index(lrjoin, str4))

	// LR並びからLRRを探して文字列位置を返す
	for strings.Contains(lrjoin, str4) {
		num := len(lrjoin)
		if num < 5 {
			break
		}

		log.Println(strings.Index(lrjoin, str4))
		log.Println("LRR=", strings.Index(lrjoin, str4))
		// L点と２つ次のR点を結んでできる三角形を分割してリストに追加する
		// L点の１つ次のR点を削除して多角形を作り直す
		// 分離する三角形の頂点番号
		n := strings.Index(lrjoin, str4)
		idx := lrIdx[n]
		log.Println("idx=", idx)

		tr_41 := idx
		tr_42 := (idx + 1) % num
		tr_43 := (idx + 2) % num
		// 分離する三角形１の頂点座標
		tri4 := Triangle{vtx1: cords[tr_41], vtx2: cords[tr_42], vtx3: cords[tr_43]}
		log.Println(tri4)
		triList = append(triList, &tri4)
		log.Println(triList)
		log.Println("triList=", *triList[i])
		i++

		// 三角形を分離した後，多角形を作り直す
		newpoly := append(cords[:tr_42], cords[tr_43:]...)
		log.Println("newpoly=", newpoly)

		// 作り直した多角形のために外積の計算をやり直す必要がある
		extL, _, _ := pkg.TriVert(num-1, newpoly)
		log.Println("extL=", extL)
		// L点，R点の辞書を作り直す
		_, _, _, lrPtn, lrIdx = pkg.Lexicogra(num-1, newpoly, extL)
		// log.Println("order=", order)
		log.Println("lrPtn=", lrPtn)
		log.Println("lrIdx=", lrIdx)

		lrjoin = strings.Join(lrPtn, "")
		log.Println("lrjoin=", lrjoin)

		cords = newpoly

		if !strings.Contains(lrjoin, str4) {
			break
		}
	}

	// L点を含む四角形をL点を中心に分割する
	if strings.Contains(lrjoin, "L") && len(lrjoin) == 4 {
		// 分離する三角形の頂点番号
		n := strings.Index(lrjoin, "L")
		idx := lrIdx[n]
		log.Println("idx=", idx)
		// 分離する三角形の頂点座標
		tri51 := Triangle{vtx1: cords[idx], vtx2: cords[(idx+1)%4], vtx3: cords[(idx+2)%4]}
		log.Println(tri51)
		triList = append(triList, &tri51)
		tri52 := Triangle{vtx1: cords[idx], vtx2: cords[(idx+2)%4], vtx3: cords[(idx+3)%4]}
		log.Println(tri52)
		triList = append(triList, &tri52)
	}

	// R点のみで構成される多角形を扇形分割する
	if !strings.Contains(lrjoin, "L") {
		vnum := len(lrjoin)
		for i := 0; i < vnum-2; i++ {
			// tr_61 := 0
			tr_62 := i + 1
			tr_63 := i + 2
			// 分離する三角形の頂点座標
			tri6 := Triangle{vtx1: cords[0], vtx2: cords[tr_62], vtx3: cords[tr_63]}
			log.Println(tri6)
			triList = append(triList, &tri6)
		}
	}
}
