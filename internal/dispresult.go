package internal

import (
	"log"
	"os"
)

// ３つに分割した建物データを表示する
func DispResult() (wfl1, wfl2, wfl3 string) {
	ba1, er := os.ReadFile("data/hutsu_list.txt")
	if er != nil {
		log.Fatal(er)
	}
	wfl1 = (string(ba1))

	ba2, er := os.ReadFile("data/kenro_list.txt")
	if er != nil {
		log.Fatal(er)
	}
	wfl2 = (string(ba2))

	ba3, er := os.ReadFile("data/other_list.txt")
	if er != nil {
		log.Fatal(er)
	}
	wfl3 = (string(ba3))

	return wfl1, wfl2, wfl3
}
