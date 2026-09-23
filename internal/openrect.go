package internal

import (
	"encoding/gob"
	"fmt"
	"log"
	"os"
)

// OpenRect は四角形建物のリストを作る
func OpenRect() (rectList []SlopeRoof, lr int) {
	// セーブファイルから復元
	flr, err := os.Open("./rectfile.gob")
	if err != nil {
		log.Fatal(err)
	}
	defer flr.Close()
	// var rectList = make([]SlopeRoof, 0)
	// デコーダーの作成
	rDecoder := gob.NewDecoder(flr)
	// デコード
	if err := rDecoder.Decode(&rectList); err != nil {
		log.Fatal("decode error:", err)
	}
	// fmt.Println("rectList\n")
	lr = len(rectList)
	fmt.Println("rectList = ", lr)

	return rectList, lr
}
