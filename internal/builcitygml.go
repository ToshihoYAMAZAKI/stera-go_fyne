package internal

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// expapp はappearanceMember をファイルに書き出す
func expapp(f *os.File, anum int, fid string, sfcnt int) {
	f.WriteString("\t<app:appearanceMember>\n")
	f.WriteString("\t\t<app:Appearance>\n")
	theme := "\t\t\t<app:theme>AreaOfUse:" + strconv.Itoa(anum) + "</app:theme>\n"
	f.WriteString(theme)
	f.WriteString("\t\t\t<app:surfaceDataMember>\n")
	f.WriteString("\t\t\t\t<app:X3DMaterial>\n")
	rgba := AreaRGB(anum)
	var t_rgb []string
	for i := 0; i < 3; i++ {
		t_rgb = append(t_rgb, strconv.FormatFloat(rgba[i], 'f', -1, 64))
	}
	diff_col := "\t\t\t\t\t<app:diffuseColor>" + strings.Join(t_rgb, " ") + "</app:diffuseColor>\n"
	f.WriteString(diff_col)
	alp := 1.0 - rgba[3]
	trans := "\t\t\t\t\t<app:transparency>" + strconv.FormatFloat(alp, 'f', -1, 64) + "</app:transparency>\n"
	f.WriteString(trans)
	for i := 0; i < sfcnt; i++ {
		pgid := "\t\t\t\t\t<app:target>#PolyGMLID_" + fid + "_" + strconv.Itoa(i+1) + "</app:target>\n"
		f.WriteString(pgid)
	}
	f.WriteString("\t\t\t\t</app:X3DMaterial>\n")
	f.WriteString("\t\t\t</app:surfaceDataMember>\n")
	f.WriteString("\t\t</app:Appearance>\n")
	f.WriteString("\t</app:appearanceMember>\n")
}

// expcore はcityObjectMember をファイルに書き出す
func expcore(f *os.File, id string) {
	f.WriteString("\t<core:cityObjectMember>\n")
	uuid := "\t\t<bldg:Building gml:id=\"UUID_Building_" + id + "\">\n"
	f.WriteString(uuid)
	f.WriteString("\t\t\t<bldg:lod2Solid>\n")
	f.WriteString("\t\t\t\t<gml:Solid>\n")
	f.WriteString("\t\t\t\t\t<gml:exterior>\n")
	f.WriteString("\t\t\t\t\t\t<gml:CompositeSurface>\n")
}

// exphref はsurfaceMember にxlink:href をファイルに書き出す
func exphref(f *os.File, fid string, sfcnt int) {
	for i := 0; i < sfcnt; i++ {
		pgid := "\t\t\t\t\t\t\t<gml:surfaceMember xlink:href=\"PolyGMLID_" + fid + "_" + strconv.Itoa(i+1) + "\"/>\n"
		f.WriteString(pgid)
	}
	f.WriteString("\t\t\t\t\t\t</gml:CompositeSurface>\n")
	f.WriteString("\t\t\t\t\t</gml:exterior>\n")
	f.WriteString("\t\t\t\t</gml:Solid>\n")
	f.WriteString("\t\t\t</bldg:lod2Solid>\n")
}

// expbndl はbldg:boundedBy をファイルに書き出す
func expbndl(f *os.File, sftype string, id string, i int) {
	f.WriteString("\t\t\t<bldg:boundedBy>\n")
	var sfid string
	if sftype == "Wall" {
		sfid = "\t\t\t\t<bldg:WallSurface gml:id=\"UUID_WallSurface_" + id + "_" + strconv.Itoa(i) + "\">\n"
	} else if sftype == "Roof" {
		sfid = "\t\t\t\t<bldg:RoofSurface gml:id=\"UUID_RoofSurface_" + id + "_" + strconv.Itoa(i) + "\">\n"
	} else if sftype == "Floor" {
		sfid = "\t\t\t\t<bldg:FloorSurface gml:id=\"UUID_FloorSurface_" + id + "_" + strconv.Itoa(i) + "\">\n"
	} else if sftype == "Ground" {
		sfid = "\t\t\t\t<bldg:GroundSurface gml:id=\"UUID_GroundSurface_" + id + "_" + strconv.Itoa(i) + "\">\n"
	} else if sftype == "Ceiling" {
		sfid = "\t\t\t\t<bldg:CeilingSurface gml:id=\"UUID_CeilingSurface_" + id + "_" + strconv.Itoa(i) + "\">\n"
	}
	f.WriteString(sfid)
	f.WriteString("\t\t\t\t\t<bldg:lod2MultiSurface>\n")
	f.WriteString("\t\t\t\t\t\t<gml:MultiSurface>\n")
}

// exitbndl はbldg:boundedBy の書き出しを終了する
func exitbndl(f *os.File, sftype string) {
	f.WriteString("\t\t\t\t\t\t</gml:MultiSurface>\n")
	f.WriteString("\t\t\t\t\t</bldg:lod2MultiSurface>\n")
	if sftype == "Wall" {
		f.WriteString("\t\t\t\t</bldg:WallSurface>\n")
	} else if sftype == "Roof" {
		f.WriteString("\t\t\t\t</bldg:RoofSurface>\n")
	} else if sftype == "Floor" {
		f.WriteString("\t\t\t\t</bldg:FloorSurface>\n")
	} else if sftype == "Ground" {
		f.WriteString("\t\t\t\t</bldg:GroundSurface>\n")
	} else if sftype == "Ceiling" {
		f.WriteString("\t\t\t\t</bldg:CeilingSurface>\n")
	}
	f.WriteString("\t\t\t</bldg:boundedBy>\n")
}

// expsrfc はsurfaceMember をファイルに書き出す
// gml:surfaceMember の書き出し
// fid をPolyGMLID に書き換えてテキスト変換して書き出す
// PolyGMLID とX/Y/Z座標をfor文で書き出す
func expsrfc(f *os.File, fid string, i int) {
	f.WriteString("\t\t\t\t\t\t\t<gml:surfaceMember>\n")
	pgid := "\t\t\t\t\t\t\t\t<gml:Polygon gml:id=\"PolyGMLID_" + fid + "_" + strconv.Itoa(i) + "\">\n"
	f.WriteString(pgid)
	f.WriteString("\t\t\t\t\t\t\t\t\t<gml:exterior>\n")
	lrid := "\t\t\t\t\t\t\t\t\t\t<gml:LinearRing gml:id=\"PolyGMLID_" + fid + "_" + strconv.Itoa(i) + "_0\">\n"
	f.WriteString(lrid)
}

// exitsrfc はsurfaceMember の書き出しを終了する
func exitsrfc(f *os.File) {
	f.WriteString("\t\t\t\t\t\t\t\t\t\t</gml:LinearRing>\n")
	f.WriteString("\t\t\t\t\t\t\t\t\t</gml:exterior>\n")
	f.WriteString("\t\t\t\t\t\t\t\t</gml:Polygon>\n")
	f.WriteString("\t\t\t\t\t\t\t</gml:surfaceMember>\n")
}

// exitcore はcityObjectMember の書き出しを終了する
func exitcore(f *os.File) {
	f.WriteString("\t\t\t\t\t\t</gml:CompositeSurface>\n")
	f.WriteString("\t\t\t\t\t</gml:exterior>\n")
	f.WriteString("\t\t\t\t</gml:Solid>\n")
	f.WriteString("\t\t\t</bldg:lod2Solid>\n")
	f.WriteString("\t\t</bldg:Building>\n")
	f.WriteString("\t</core:cityObjectMember>\n")
}

// exitcore2 はbldg:boundedBy の書き出しを終了する
func exitcore2(f *os.File) {
	f.WriteString("\t\t</bldg:Building>\n")
	f.WriteString("\t</core:cityObjectMember>\n")
}

// exposlst は頂点座標をファイルに書き出す
func exposlst(f *os.File, fid string, vcnt int, list [][]float64, elv, btm, toph float64, story int) {
	// PolyGMLID をfid から設定するための添え字の定義
	var sub int
	// 基面データの頂点座標の定義
	var basepoly []float64
	// 基面データの頂点座標の文字列
	var basetxt []string
	// 底面データの頂点座標の定義
	var bttmpoly []float64
	// 底面データの頂点座標の文字列
	var bttmtxt []string
	// 上面データの頂点座標の定義
	var toppoly []float64
	// 上面データの頂点座標の文字列
	var toptxt []string
	// 各階データの頂点座標の定義
	var flrpoly []float64
	// 各階データの頂点座標の文字列
	var flrtxt []string
	// 側面データの頂点座標の定義
	var sidepoly []float64
	// 側面データの頂点座標の文字列
	var sidetxt []string

	// 基面IDと頂点座標の書き出し
	sub = 1
	expsrfc(f, fid, sub)
	for j := 0; j < vcnt; j++ {
		basepoly = append(basepoly, list[j][0])
		basepoly = append(basepoly, list[j][1])
		basepoly = append(basepoly, elv)
	}
	// 基面データのテキスト化
	for m := range basepoly {
		basetxt = append(basetxt, strconv.FormatFloat(basepoly[m], 'f', -1, 64))
	}
	bscortxt := "\t\t\t\t\t\t\t\t\t\t\t<gml:posList srsDimension=\"3\"> " + strings.Join(basetxt, " ") + "</gml:posList>\n"
	f.WriteString(bscortxt)
	exitsrfc(f)

	// 底面IDと頂点座標の書き出し
	sub = sub + 1
	expsrfc(f, fid, sub)
	for j := vcnt - 1; j >= 0; j-- {
		bttmpoly = append(bttmpoly, list[j][0])
		bttmpoly = append(bttmpoly, list[j][1])
		bttmpoly = append(bttmpoly, btm)
	}
	// 底面データのテキスト化
	for m := range bttmpoly {
		bttmtxt = append(bttmtxt, strconv.FormatFloat(bttmpoly[m], 'f', -1, 64))
	}
	btcortxt := "\t\t\t\t\t\t\t\t\t\t\t<gml:posList srsDimension=\"3\"> " + strings.Join(bttmtxt, " ") + "</gml:posList>\n"
	f.WriteString(btcortxt)
	exitsrfc(f)

	// 上面IDと頂点座標の書き出し
	sub = sub + 1
	expsrfc(f, fid, sub)
	for j := 0; j < vcnt; j++ {
		toppoly = append(toppoly, list[j][0])
		toppoly = append(toppoly, list[j][1])
		toppoly = append(toppoly, toph)
	}
	// 上面データのテキスト化
	for m := range toppoly {
		toptxt = append(toptxt, strconv.FormatFloat(toppoly[m], 'f', -1, 64))
	}
	tpcortxt := "\t\t\t\t\t\t\t\t\t\t\t<gml:posList srsDimension=\"3\"> " + strings.Join(toptxt, " ") + "</gml:posList>\n"
	f.WriteString(tpcortxt)
	exitsrfc(f)

	// 各階頂点座標の書き出し
	for j := 1; j < story; j++ {
		sub = sub + 1
		expsrfc(f, fid, sub)
		for k := 0; k < vcnt; k++ {
			flrpoly = append(flrpoly, list[k][0])
			flrpoly = append(flrpoly, list[k][1])
			flrpoly = append(flrpoly, (elv + 3.3*float64(j)))
		}
		// 各階データのテキスト化
		for m := range flrpoly {
			flrtxt = append(flrtxt, strconv.FormatFloat(flrpoly[m], 'f', -1, 64))
		}
		flcortxt := "\t\t\t\t\t\t\t\t\t\t\t<gml:posList srsDimension=\"3\"> " + strings.Join(flrtxt, " ") + "</gml:posList>\n"
		f.WriteString(flcortxt)
		exitsrfc(f)

		// 各階データの頂点座標の初期化
		flrpoly = flrpoly[:0]
		// 各階データの頂点座標の文字列の初期化
		flrtxt = flrtxt[:0]
	}

	// 側面頂点座標の書き出し
	for j := 0; j < vcnt-1; j++ {
		sub = sub + 1
		expsrfc(f, fid, sub)
		// 頂点下1
		sidepoly = append(sidepoly, list[j][0])
		sidepoly = append(sidepoly, list[j][1])
		sidepoly = append(sidepoly, btm)
		// 頂点下2
		sidepoly = append(sidepoly, list[(j+1)%vcnt][0])
		sidepoly = append(sidepoly, list[(j+1)%vcnt][1])
		sidepoly = append(sidepoly, btm)
		// 頂点上2
		sidepoly = append(sidepoly, list[(j+1)%vcnt][0])
		sidepoly = append(sidepoly, list[(j+1)%vcnt][1])
		sidepoly = append(sidepoly, toph)
		// 頂点上1
		sidepoly = append(sidepoly, list[j][0])
		sidepoly = append(sidepoly, list[j][1])
		sidepoly = append(sidepoly, toph)

		// 側面データのテキスト化
		for m := range sidepoly {
			sidetxt = append(sidetxt, strconv.FormatFloat(sidepoly[m], 'f', -1, 64))
		}
		sdcortxt := "\t\t\t\t\t\t\t\t\t\t\t<gml:posList srsDimension=\"3\"> " + strings.Join(sidetxt, " ") + "</gml:posList>\n"
		f.WriteString(sdcortxt)
		exitsrfc(f)

		// 側面データの頂点座標の初期化
		sidepoly = sidepoly[:0]
		// 側面データの頂点座標の文字列の初期化
		sidetxt = sidetxt[:0]
	}
}

// メイン関数
// BuildCity はCityGML形式で多角形建物モデルを作成する
func BuildCity(filename string, yanedtail Yanedtail) {
	// ログファイルを新規作成，追記，書き込み専用，パーミションは読むだけ
	file, err := os.OpenFile("citygml.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	// ログの出力先を変更
	log.SetOutput(file)

	// CityGML形式で出力するためのファイルを開く
	// f, err := os.Create("C:/data/output.xml")
	f, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	// defer f.Close()
	defer func() {
		err := f.Close()
		if err != nil {
			log.Println(err)
		}
	}()

	// 四角形リストを開く
	rectList, _ := OpenRect()

	// 多角形リストを開く
	polyList, _ := OpenPoly()

	// 堅ろう建物のファイルを開く
	kList, _ := OpenKenro()

	// 無壁舎建物のファイルを開く
	mList, _ := OpenMuheki()

	// 底面高さの設定
	var btm float64
	// 上面高さの設定
	var toph float64

	// CityGMLファイルの出力開始
	// ファイル出力（ヘッダー部分の書き出し）
	xml := "<?xml version=\"1.0\" encoding=\"utf-8\"?>\n"
	data := []byte(xml)
	cnt, err := f.Write(data)
	if err != nil {
		log.Println(err)
		log.Println("fail to write file")
	}
	fmt.Printf("write %d bytes\n", cnt)
	// f.Write([]byte("<?xml version=\"1.0\" encoding=\"utf-8\"?>\n"))
	f.WriteString("<core:CityModel xmlns:xsi=\"http://www.w3.org/2001/XMLSchema-instance\"\n")
	f.WriteString("\t\t\t\txsi:schemaLocation=\"http://www.opengis.net/citygml/2.0 ./CityGML_2.0/CityGML.xsd\"\n")
	f.WriteString("\t\t\t\txmlns=\"http://www.opengis.net/citygml/2.0\"\n")
	f.WriteString("\t\t\t\txmlns:xAL=\"urn:oasis:names:tc:ciq:xsdschema:xAL:2.0\"\n")
	f.WriteString("\t\t\t\txmlns:app=\"http://www.opengis.net/citygml/appearance/2.0\"\n")
	f.WriteString("\t\t\t\txmlns:xlink=\"http://www.w3.org/1999/xlink\"\n")
	f.WriteString("\t\t\t\txmlns:gml=\"http://www.opengis.net/gml\"\n")
	f.WriteString("\t\t\t\txmlns:core=\"http://www.opengis.net/citygml/2.0\"\n")
	f.WriteString("\t\t\t\txmlns:dem=\"http://www.opengis.net/citygml/relief/2.0\"\n")
	f.WriteString("\t\t\t\txmlns:bldg=\"http://www.opengis.net/citygml/building/2.0\">\n")

	// // 逐次処理（１件ずつデータ読み込み → モデリングを行う）
	// 普通建物（傾斜屋根）の処理

	// 7角形屋根は三角屋根の棟高さと軒高さで隣り合う片流れ屋根の高さを決める
	// 三角屋根の棟と軒の頂点座標を格納する
	// var tripts TriVerts
	var tripts map[string]float64

	for r := range rectList {
		i := r + 1
		log.Println("ID=", rectList[i-1].ID)
		log.Println("Type=", rectList[i-1].Type)

		// UUID の設定
		id := rectList[i-1].ID
		// PolyGMLID の設定
		fid := rectList[i-1].Fid
		log.Println("fid=", fid)
		// 頂点数のチェック
		vcnt := len(rectList[i-1].List)
		// if vcnt != 4 {
		// 	break
		// }

		// 建物階数の定義
		story := rectList[i-1].Story
		log.Println("rectList[i-1].Story=", story)

		// 面頂点数・座標データ数の総和
		// 屋根タイプにより変化する
		// v_num, cnt, pm_cnt := counter(vcnt)
		_, _, pm_cnt := counter(vcnt, story)
		if rectList[i-1].Type == "kiri" {
			// 切妻屋根の面数／(2x2+3x2)×2+2=22
			pm_cnt = pm_cnt + ((2*2+3*2)*2 + 2)
		} else if rectList[i-1].Type == "yose" {
			// 寄棟屋根の面数／(2+2x2)x2+2x4=20
			pm_cnt = pm_cnt + ((2+2*2)*2 + 2*4)
		} else if rectList[i-1].Type == "kata" {
			// 片流れ屋根の面数／2x2+2+2+2x2+2+2=16
			pm_cnt = pm_cnt + (2*2 + 2 + 2 + 2*2 + 2 + 2)
		} else if rectList[i-1].Type == "kata1" || rectList[i-1].Type == "kata2" {
			// 片流れ屋根の面数／2x2+2+2+2x2+4+2=18
			pm_cnt = pm_cnt + (2*2 + 2 + 2 + 2*2 + 4 + 2)
		} else if rectList[i-1].Type == "penta" {
			// ５角形屋根の面数／(2x2+2x2+2x2)+(2+2+2+2+2+2)+1=25
			pm_cnt = pm_cnt + ((2 * 2 * 3) + (2 * 6) + 1)
		} else if rectList[i-1].Type == "kiri5" {
			// 切妻屋根の面数／(2x2+3x2)×2+3=23
			pm_cnt = pm_cnt + ((2*2+3*2)*2 + 3)
		} else if rectList[i-1].Type == "5kakudou" {
			// 三角形片流れ屋根の面数／5x2+2x5=30
			pm_cnt = pm_cnt + (5*2 + 2*5)
		} else if rectList[i-1].Type == "heprec" {
			// 変形片流れ屋根の面数／2x2+2x4+2=14
			pm_cnt = pm_cnt + (2*2 + 2*4 + 2)
		} else if rectList[i-1].Type == "heptri" {
			// 三角形片流れ屋根の面数／2+2x3+2=10
			pm_cnt = pm_cnt + (2 + 2*3 + 2)
		}

		// 建物面数の設定
		sfcnt := pm_cnt
		// 基準面（地盤高さ）の設定
		level := rectList[i-1].Elv
		// 建物の上面高さの設定
		toph = rectList[i-1].Elv + 3.3*float64(story)
		// 建物地下部分のモデリングのための地下深さの設定
		nest := 0.3
		btm = level - nest

		// 用途地域番号の読み込み
		youto := rectList[i-1].Area
		// log.Println("youto=", youto)
		anum := AreaNum(youto)
		// log.Println("anum=", anum)

		// 閉じた図形にするための頂点座標の追加
		rectList[i-1].List = append(rectList[i-1].List, rectList[i-1].List[0])
		vnum := vcnt + 1

		// app:appearanceMember の書き出し
		// 用途地域のColor をテキスト変換して書き出す
		// PolyGMLID をfor文で書き出す
		expapp(f, anum, fid, sfcnt)

		// core:cityObjectMember の書き出し
		// UUID をテキスト変換して書き出す
		// PolyGMLID とX/Y/Z座標をfor文で書き出す
		expcore(f, id)

		// CompositeSurface の書き出し
		// surfaceMember にxlink:href を書き出す
		exphref(f, fid, sfcnt)

		// 屋根モデル別に頂点座標のリストを出力する
		// 建物モデル（直方体）を基面，底面，各階，上面，側面の順に書き出す
		// X/Y/Z座標を配列に格納して，これをテキスト変換して書き出す

		if rectList[i-1].Type == "kiri" {
			// 軒庇の出とケラバの厚さを設定する
			hisashi := yanedtail.Kirinoki
			keraba := yanedtail.Kirikera
			incline := yanedtail.Kiriincl
			yaneatu := yanedtail.Kiriroof

			// 切妻屋根建物の頂点座標をリストにテキスト化して書き出す
			ExposKiri(f, id, fid, vnum, rectList[i-1].List, rectList[i-1].Elv, btm, toph, hisashi, keraba, incline, yaneatu, story)
			// log.Println("rectList[i-1].List=", rectList[i-1].List)

		} else if rectList[i-1].Type == "yose" {
			// 軒庇の出とケラバの厚さを設定する
			hisashi := yanedtail.Yosenoki
			keraba := hisashi
			incline := yanedtail.Yoseincl
			yaneatu := yanedtail.Yoseroof

			// 寄棟屋根の頂点の法線ベクトルをリストにテキスト化して書き出す
			ExposYose(f, id, fid, vnum, rectList[i-1].List, rectList[i-1].Elv, btm, toph, hisashi, keraba, incline, yaneatu, story)

		} else if rectList[i-1].Type == "kata" {
			// 軒庇の出とケラバの厚さを設定する
			hisashi := yanedtail.Katanoki
			keraba := yanedtail.Katakera
			incline := yanedtail.Kataincl
			yaneatu := yanedtail.Kataroof

			// 片流れ屋根の頂点の法線ベクトルをリストにテキスト化して書き出す
			ExposKata(f, id, fid, vnum, rectList[i-1].List, rectList[i-1].Elv, btm, toph, hisashi, keraba, incline, yaneatu, story)

		} else if rectList[i-1].Type == "heptri" {
			// 軒庇の出とケラバの厚さを設定する
			hisashi := 0.60
			keraba := 0.30
			incline := 0.3
			yaneatu := 0.075

			// ３角形の片流れ屋根の頂点の法線ベクトルをリストにテキスト化して書き出す
			ExposTri(f, id, fid, vnum, rectList[i-1].List, rectList[i-1].Elv, btm, toph, hisashi, keraba, incline, yaneatu, story)

		} else if rectList[i-1].Type == "kata1" {
			// 軒庇の出とケラバの厚さを設定する
			hisashi := 0.60
			keraba := 0.30
			incline := 0.3
			yaneatu := 0.075

			// 片流れ屋根の頂点の法線ベクトルをリストにテキスト化して書き出す
			ExposKata1(f, id, fid, vnum, rectList[i-1].List, rectList[i-1].Elv, btm, toph, hisashi, keraba, incline, yaneatu, story, tripts)

		} else if rectList[i-1].Type == "kata2" {
			// 軒庇の出とケラバの厚さを設定する
			hisashi := 0.60
			keraba := 0.30
			incline := 0.3
			yaneatu := 0.075

			// 片流れ屋根の頂点の法線ベクトルをリストにテキスト化して書き出す
			ExposKata2(f, id, fid, vnum, rectList[i-1].List, rectList[i-1].Elv, btm, toph, hisashi, keraba, incline, yaneatu, story, tripts)

		} else if rectList[i-1].Type == "penta" {
			// 軒庇の出とケラバの厚さを設定する
			hisashi := 0.60
			keraba := hisashi
			incline := 0.3
			yaneatu := 0.075

			// ５角形屋根の頂点の法線ベクトルをリストにテキスト化して書き出す
			ExposPenta(f, id, fid, vnum, rectList[i-1].List, rectList[i-1].Elv, btm, toph, hisashi, keraba, incline, yaneatu, story)

		} else if rectList[i-1].Type == "kiri5" {
			// 軒庇の出とケラバの厚さを設定する
			hisashi := 0.60
			keraba := 0.30
			incline := 0.45
			yaneatu := 0.11

			// ５角形屋根の切妻屋根の頂点の法線ベクトルをリストにテキスト化して書き出す
			ExposKiri5(f, id, fid, vnum, rectList[i-1].List, rectList[i-1].Elv, btm, toph, hisashi, keraba, incline, yaneatu, story)

		} else if rectList[i-1].Type == "5kakudou" {
			// 軒庇の出とケラバの厚さを設定する
			hisashi := 0.60
			incline := 0.45
			yaneatu := 0.11

			// ５角形屋根の頂点の法線ベクトルをリストにテキスト化して書き出す
			Expos5kaku(f, id, fid, vnum, rectList[i-1].List, rectList[i-1].Elv, btm, toph, hisashi, incline, yaneatu, story)
		}

		// bldg:boundedBy の書き出しを終了する
		exitcore2(f)
	}

	// 普通建物（多角形）の処理
	for p := range polyList {
		i := p + 1

		// UUID の設定
		id := polyList[i-1].ID
		// PolyGMLID の設定
		fid := polyList[i-1].Fid
		log.Println("fid=", fid)
		// 頂点数
		vcnt := len(polyList[i-1].List)
		// 建物階数の定義
		story := 2
		// 建物面数の設定
		sfcnt := 2 + story + vcnt
		// 基準面（地盤高さ）の設定
		level := polyList[i-1].Elv
		// 建物の上面高さの設定
		toph = polyList[i-1].Elv + 3.3*float64(story)
		// 建物地下部分のモデリングのための地下深さの設定
		nest := 0.3
		btm = level - nest

		// 用途地域番号の読み込み
		youto := polyList[i-1].Area
		// log.Println("youto=", youto)
		anum := AreaNum(youto)
		// log.Println("anum=", anum)

		// 閉じた図形にするための頂点座標の追加
		polyList[i-1].List = append(polyList[i-1].List, polyList[i-1].List[0])
		vnum := vcnt + 1

		// app:appearanceMember の書き出し
		// 用途地域のColor をテキスト変換して書き出す
		// PolyGMLID をfor文で書き出す
		expapp(f, anum, fid, sfcnt)

		// core:cityObjectMember の書き出し
		// UUID をテキスト変換して書き出す
		// PolyGMLID とX/Y/Z座標をfor文で書き出す
		expcore(f, id)

		// 基面，底面，上面，各階，側面の順に書き出す
		// X/Y/Z座標を配列に格納して，これをテキスト変換して書き出す
		exposlst(f, fid, vnum, polyList[i-1].List, polyList[i-1].Elv, btm, toph, story)

		// core:cityObjectMember の書き出しを終了する
		exitcore(f)
	}

	// 堅ろう建物の処理
	for k := range kList {
		i := k + 1

		// UUID の設定
		id := kList[i-1].ID
		// PolyGMLID の設定
		fid := kList[i-1].Fid
		// log.Println("fid=", fid)
		// 頂点数
		vnum := len(kList[i-1].Cords)
		// log.Println("vertex=", vnum)

		// 用途地域番号の読み込み
		youto := kList[i-1].Area
		log.Println("youto=", youto)
		anum := AreaNum(youto)
		log.Println("anum=", anum)

		// 建ぺい率・容積率の設定
		bcr := kList[i-1].Build
		far := kList[i-1].Floor
		// 建物階数の定義
		st := kList[i-1].Story
		log.Println("st=", st)
		// TODO: 読み込んだ建物階数を乱数で生成した階数に置き換えて反映させる必要がある
		// 建物の上面高さの設定
		// 建物の階数が設定されていない場合は乱数で建物階数を設定する
		story := Stcalc(st, bcr, far, anum)
		log.Println("story=", story)

		// 建物面数の設定
		sfcnt := 2 + story + vnum - 1

		// 基準面（地盤高さ）の設定
		level := kList[i-1].Elv
		// 建物の上面高さの設定
		toph = level + 3.3*float64(story)
		// 建物地下部分のモデリングのための地下深さの設定
		// 地下階がある場合は階数に応じてZ座標を地下階部分の深さだけ下げる
		bf := kList[i-1].Basement
		btm = Bfcalc(level, story, bf)

		// app:appearanceMember の書き出し
		// 用途地域のColor をテキスト変換して書き出す
		// PolyGMLID をfor文で書き出す
		expapp(f, anum, fid, sfcnt)

		// core:cityObjectMember の書き出し
		// UUID をテキスト変換して書き出す
		// PolyGMLID とX/Y/Z座標をfor文で書き出す
		expcore(f, id)

		// 基面，底面，上面，各階，側面の順に書き出す
		// X/Y/Z座標を配列に格納して，これをテキスト変換して書き出す
		exposlst(f, fid, vnum, kList[i-1].Cords, kList[i-1].Elv, btm, toph, story)

		// core:cityObjectMember の書き出しを終了する
		exitcore(f)
	}

	// 無壁舎建物の処理
	for m := range mList {
		i := m + 1

		// UUID の設定
		id := mList[i-1].ID
		// PolyGMLID の設定
		fid := mList[i-1].Fid
		// log.Println("fid=", fid)
		// 頂点数
		vnum := len(mList[i-1].Cords)
		// 建物階数の定義
		story := 1
		// 建物面数の設定
		sfcnt := 2 + story + vnum - 1
		// 基準面（地盤高さ）の設定
		level := mList[i-1].Elv
		// 建物の上面高さの設定
		toph = level + 3.3*float64(story)
		// 建物地下部分のモデリングのための地下深さの設定
		btm = level - 0.3

		// 用途地域番号の読み込み
		youto := mList[i-1].Area
		// log.Println("youto=", youto)
		anum := AreaNum(youto)
		// log.Println("anum=", anum)

		// app:appearanceMember の書き出し
		// 用途地域のColor をテキスト変換して書き出す
		// PolyGMLID をfor文で書き出す
		expapp(f, anum, fid, sfcnt)

		// core:cityObjectMember の書き出し
		// UUID をテキスト変換して書き出す
		// PolyGMLID とX/Y/Z座標をfor文で書き出す
		expcore(f, id)

		// 基面，底面，上面，各階，側面の順に書き出す
		// X/Y/Z座標を配列に格納して，これをテキスト変換して書き出す
		exposlst(f, fid, vnum, mList[i-1].Cords, mList[i-1].Elv, btm, toph, story)

		// core:cityObjectMember の書き出しを終了する
		exitcore(f)
	}

	// CityGMLファイルの出力終了
	f.WriteString("</core:CityModel>")
}
