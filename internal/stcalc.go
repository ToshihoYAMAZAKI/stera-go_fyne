package internal

import (
	"log"
	"math"
)

// Stcalc は用途地域から建物階数を設定する
func Stcalc(st, bcr, far, anum int) (story int) {
	// 容積率／建ぺい率で建物階数の上限値を求める
	var maxst int
	if far > 0 && bcr > 0 {
		maxst = int(math.Round(float64(far)/float64(bcr) + 0.5))
		if anum == 1 || anum == 2 {
			maxst = 3
		}
	} else {
		maxst = 3
	}
	log.Println("maxst=", maxst)
	// 建物の階数が設定されていない場合は乱数で建物階数を設定する
	if st > 0 {
		// toph = float64(st) * 3.3
		story = st
	} else {
		snum := RandStory(anum)
		// 堅ろう建物の場合，全ての建物は３階建て以上とする
		// 商業地域の場合は４階建て以上とする
		if snum < 3 {
			snum = 3
			if anum == 9 {
				snum = 4
			}
		}
		if snum > maxst {
			snum = maxst
		}
		log.Println("snum=", snum)
		// toph = float64(snum) * 3.3
		story = snum
	}

	return story
}
