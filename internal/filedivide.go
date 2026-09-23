package internal

import (
	"bufio"
	"io"
	"os"
	"strings"

	"log"
)

// KenroBuil は堅ろう建物用の構造体の定義
type KenroBuil struct {
	ID       string
	Fid      string
	Elv      float64
	Story    int
	Basement int
	Area     string
	Build    int
	Floor    int
	Roof     float64
	Cords    [][]float64
}

// MuhekiBuil は無壁舎建物用の構造体の定義
type MuhekiBuil struct {
	ID    string
	Fid   string
	Elv   float64
	Story int
	Area  string
	Roof  float64
	Cords [][]float64
}

// DivideLine は建物データを３つに分類する
func DivideLine(fn string) error {
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

	wfl1, er := os.Create("data/hutsu_list.txt")
	if er != nil {
		log.Fatal(er)
	}
	defer wfl1.Close()

	wfl2, er := os.Create("data/kenro_list.txt")
	if er != nil {
		log.Fatal(er)
	}
	defer wfl2.Close()

	wfl3, er := os.Create("data/other_list.txt")
	if er != nil {
		log.Fatal(er)
	}
	defer wfl3.Close()

	r := bufio.NewReader(rfl)
	w1 := bufio.NewWriter(wfl1)
	w2 := bufio.NewWriter(wfl2)
	w3 := bufio.NewWriter(wfl3)

	for {
		row, er := r.ReadString('\n')
		if er != nil && er != io.EOF {
			log.Fatal(er)
		}
		if er == io.EOF && len(row) == 0 {
			break
		}
		if strings.Contains(row, "普通建物") {
			log.Println("普通建物")
			_, er = w1.WriteString(row)
			if er != nil {
				log.Fatal(er)
			}
			w1.Flush()
		}
		if strings.Contains(row, "堅ろう建物") {
			log.Println("堅ろう建物")
			_, er = w2.WriteString(row)
			if er != nil {
				log.Fatal(er)
			}
			w2.Flush()
		}
		if strings.Contains(row, "無壁舎") {
			log.Println("無壁舎")
			_, er = w3.WriteString(row)
			if er != nil {
				log.Fatal(er)
			}
			w3.Flush()
		}
	}
	if er != nil {
		log.Fatal(er)
	}
	return nil
}
