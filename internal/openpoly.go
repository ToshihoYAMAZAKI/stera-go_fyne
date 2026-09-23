package internal

import (
	"encoding/gob"
	"fmt"
	"log"
	"os"
)

// OpenPoly は多角形建物のリストを作る
func OpenPoly() (polyList []Polygon, lp int) {
	// セーブファイルから復元
	flp, erp := os.Open("./polyfile.gob")
	if erp != nil {
		log.Fatal(erp)
	}
	defer flp.Close()
	// var polyList = make([]Polygon, 0)
	// デコーダーの作成
	pDecoder := gob.NewDecoder(flp)
	// デコード
	if erp := pDecoder.Decode(&polyList); erp != nil {
		log.Fatal("decode error:", erp)
	}
	// fmt.Println("polyList\n")
	lp = len(polyList)
	fmt.Println("polylist = ", lp)

	return polyList, lp
}
