package pkg

import (
	"encoding/json"
	"log"
)

// GeoJSON is Convert JSON to Go struct
type GeoJSON struct {
	Type       string `json:"type"`
	Properties struct {
		ID          string      `json:"id"`
		Fid         string      `json:"fid"`
		SeibiData   int         `json:"登録日"`
		SeibiA      int         `json:"削除日"`
		SeibiFinish string      `json:"整備完了日"`
		OrgGILvl    string      `json:"orgGILvl"`
		OrgMDID     interface{} `json:"orgMDId"`
		HyoujiKubun string      `json:"表示区分"`
		KIND        string      `json:"種別"`
		NAME        interface{} `json:"名称"`
		Elev        float64     `json:"標高"`
		Story       int         `json:"階層"`
		Use         string      `json:"建物用途"`
		Area        string      `json:"用途地域"`
		Build       int         `json:"建ぺい率"`
		Floor       int         `json:"容積率"`
		RoofF       float64     `json:"屋上高さ"`
	} `json:"properties"`
	Geometry struct {
		Type        string      `json:"type"`
		Coordinates [][]float64 `json:"coordinates"`
	} `json:"geometry"`
}

// ParseJSON はGeoJSONデータの読み込み処理をする．
func ParseJSON(jStr string) (id, fid string, elv, rfh float64, story int, area string, bcr, far int, cordnts [][]float64) {

	// GeoJson構造体の変数stcDataを宣言
	var stcData GeoJSON

	// エラー処理のためjsonStrを[]byte型に変換？
	// Unmarshalで[]byte型で受け取ったJSON形式のファイルをポインタに保存
	if err := json.Unmarshal([]byte(jStr), &stcData); err != nil {
		log.Println(err)
	}

	// pj := &stcData
	// log.Printf("pointerJSON:%p\n", pj)
	// log.Println("stcData", unsafe.Sizeof(stcData))

	id = stcData.Properties.ID
	fid = stcData.Properties.Fid
	elv = stcData.Properties.Elev
	story = stcData.Properties.Story
	area = stcData.Properties.Area
	bcr = stcData.Properties.Build
	far = stcData.Properties.Floor
	rfh = stcData.Properties.RoofF
	// nullの場合は0になる
	cords := stcData.Geometry.Coordinates

	// 頂点座標の並びを反時計回りに変える
	num := len(cords)
	for i := num - 1; i >= 0; i-- {
		cordnts = append(cordnts, cords[i])
	}

	return id, fid, elv, rfh, story, area, bcr, far, cordnts
}
