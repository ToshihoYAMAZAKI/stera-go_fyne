package pkg

import (
	"log"
	"reflect"
	"strings"
)

// OctaDiv は８角形を３つの４角形に分割する
func OctaDiv(XY [][]float64, order map[string]int, pttrn []string,
	indx []int) (cord [][]float64, rect1List [][]float64, rect2List [][]float64,
	rect3List [][]float64, story []int, yane []string) {
	// 頂点データ数の確認
	nodOct := len(XY)
	if nodOct != 8 {
		return
	}
	// 頂点並びからL点の数を確認する
	lrtxt := strings.Join(pttrn, "")
	cnt := strings.Count(lrtxt, "L")
	if cnt < 2 {
		log.Println("cnt", cnt)
		return
	}

	var chkptn []string
	var octT string

	chkptn = []string{"L", "R", "L", "R", "R", "R", "R", "R"}
	octT = "歯型1"
	if reflect.DeepEqual(pttrn, chkptn) {
		log.Println(octT)
		cord, rect1List, rect2List, rect3List, story, yane = MakeOct1(XY, order)
	}

	chkptn = []string{"L", "R", "R", "R", "R", "R", "L", "R"}
	octT = "歯型2"
	if reflect.DeepEqual(pttrn, chkptn) {
		log.Println(octT)
		cord, rect1List, rect2List, rect3List, story, yane = MakeOct2(XY, order)
	}

	chkptn = []string{"L", "L", "R", "R", "R", "R", "R", "R"}
	octT = "凹型1"
	if reflect.DeepEqual(pttrn, chkptn) {
		log.Println(octT)
		cord, rect1List, rect2List, rect3List, story, yane = MakeOct3(XY, order)
	}

	chkptn = []string{"L", "R", "R", "R", "R", "R", "R", "L"}
	octT = "凹型2"
	if reflect.DeepEqual(pttrn, chkptn) {
		log.Println(octT)
		cord, rect1List, rect2List, rect3List, story, yane = MakeOct4(XY, order)
	}

	chkptn = []string{"L", "R", "R", "L", "R", "R", "R", "R"}
	octT = "凸型1"
	if reflect.DeepEqual(pttrn, chkptn) {
		log.Println(octT)
		cord, rect1List, rect2List, rect3List, story, yane = MakeOct5(XY, order)
	}

	chkptn = []string{"L", "R", "R", "R", "R", "L", "R", "R"}
	octT = "凸型2"
	if reflect.DeepEqual(pttrn, chkptn) {
		log.Println(octT)
		cord, rect1List, rect2List, rect3List, story, yane = MakeOct6(XY, order)
	}

	chkptn = []string{"L", "R", "R", "R", "L", "R", "R", "R"}
	octT = "Ｓ型"
	if reflect.DeepEqual(pttrn, chkptn) {
		log.Println(octT)
		cord, rect1List, rect2List, rect3List, story, yane = MakeOct7(XY, order)
	}

	return cord, rect1List, rect2List, rect3List, story, yane
}
