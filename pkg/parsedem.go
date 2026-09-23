package pkg

import (
	"encoding/json"
	"log"
)

// DemJSON is Convert JSON to Go struct
type DemJSON struct {
	Geometry struct {
		Coordinates []float64 `json:"coordinates"`
		Type        string    `json:"type"`
	} `json:"geometry"`
	Properties struct {
		標高 float64 `json:"標高"`
	} `json:"properties"`
	Type string `json:"type"`
}

// ParseDEM はDemJSON(標高メッシュ)データの読み込み処理をする．
func ParseDEM(jStr string) (cordnts []float64) {

	// GeoJson構造体の変数stcDataを宣言
	var stcData DemJSON

	// エラー処理のためjsonStrを[]byte型に変換？
	// Unmarshalで[]byte型で受け取ったJSON形式のファイルをポインタに保存
	if err := json.Unmarshal([]byte(jStr), &stcData); err != nil {
		log.Println(err)
		// return nil
	}

	// nullの場合は0になる
	cordnts = stcData.Geometry.Coordinates

	return cordnts
}
