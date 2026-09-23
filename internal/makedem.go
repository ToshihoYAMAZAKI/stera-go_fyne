package internal

import (
	"bufio"
	"io"
	"os"
	"strings"

	"log"
)

// MakeDem は基盤地図情報の数値標高モデルから標高データを行単位で抽出する
func MakeDem(fn string) error {
	rfl, er := os.Open(fn)
	if er != nil {
		log.Fatal(er)
	}
	defer rfl.Close()

	if f, err := os.Stat("data"); os.IsNotExist(err) || !f.IsDir() {
		os.MkdirAll("data", os.ModePerm)
	} else {
		log.Println("存在します")
	}

	wfl, er := os.Create("data/dem_data.txt")
	if er != nil {
		log.Fatal(er)
	}
	defer wfl.Close()

	r := bufio.NewReader(rfl)
	w := bufio.NewWriter(wfl)

	for {
		row, er := r.ReadString('\n')
		if er != nil && er != io.EOF {
			log.Fatal(er)
		}
		if er == io.EOF && len(row) == 0 {
			break
		}
		if strings.Contains(row, "geometry") {
			_, er = w.WriteString(row)
			if er != nil {
				log.Fatal(er)
			}
			w.Flush()
		}
		// log.Println(row)
	}
	if er != nil {
		log.Fatal(er)
	}
	return nil
}
