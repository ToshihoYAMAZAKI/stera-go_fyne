package internal

import (
	"fmt"
	"log"
	"os"
	"stera/pkg"
	"strconv"
	"strings"
)

// Yanedtail は屋根の詳細設定の構造体の定義
type Yanedtail struct {
	Kirinoki float64
	Kirikera float64
	Kiriincl float64
	Kiriroof float64
	Yosenoki float64
	Yoseincl float64
	Yoseroof float64
	Katanoki float64
	Katakera float64
	Kataincl float64
	Kataroof float64
}

// counter は面の頂点数，座標データ数を計算する
func counter(vcnt, story int) (v_num, cnt, pm_cnt int) {
	// 基面の頂点数
	base := vcnt
	// 底面の頂点数
	bttm := vcnt
	// 上面の頂点数
	top := vcnt
	// 各階の頂点数
	flr := vcnt

	// 面頂点数の総和
	v_num = vcnt*4 + bttm + top + base + flr*(story-1)
	// 座標データ数の総和
	cnt = v_num * 3
	// 面の総数（辺の数+底面+上面+基面+各階）
	pm_cnt = vcnt + 2 + 1 + (story - 1)

	return v_num, cnt, pm_cnt
}

// lib_geo_pos はlibrary_geometries(Position)をファイルに書き出す
func lib_geo_pos(f *os.File, i, cnt int) {
	// geometry idの設定
	geo_id := "\t\t<geometry id=\"ID" + strconv.Itoa(i) + "\">\n"
	f.WriteString(geo_id)
	f.WriteString("\t\t\t<mesh>\n")
	// source idの設定
	sour_id := "\t\t\t\t<source id=\"ID" + strconv.Itoa(i) + "-Pos\">\n"
	f.WriteString(sour_id)
	fl_arr_id := "\t\t\t\t\t<float_array id=\"ID" + strconv.Itoa(i) + "-Pos-array\" count=\"" + strconv.Itoa(cnt) + "\">\n"
	f.WriteString(fl_arr_id)
}

// expcords は頂点座標をファイルに書き出す
func expcords(f *os.File, vcnt int, list [][]float64, elv, btm, toph float64, story int) {
	var kansan = 0.0254 // ｍをインチ換算
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

	// 基面頂点座標の書き出し
	for j := 0; j < vcnt; j++ {
		basepoly = append(basepoly, list[j][0]/kansan)
		basepoly = append(basepoly, list[j][1]/kansan)
		basepoly = append(basepoly, elv/kansan)
	}
	// 基面データのテキスト化
	for m := range basepoly {
		basetxt = append(basetxt, strconv.FormatFloat(basepoly[m], 'f', -1, 64))
	}
	bscortxt := "\t\t\t\t\t\t" + strings.Join(basetxt, " ") + "\n"
	f.WriteString(bscortxt)

	// 底面頂点座標の書き出し
	for j := 0; j < vcnt; j++ {
		bttmpoly = append(bttmpoly, list[j][0]/kansan)
		bttmpoly = append(bttmpoly, list[j][1]/kansan)
		bttmpoly = append(bttmpoly, btm/kansan)
	}
	// 底面データのテキスト化
	for m := range bttmpoly {
		bttmtxt = append(bttmtxt, strconv.FormatFloat(bttmpoly[m], 'f', -1, 64))
	}
	btcortxt := "\t\t\t\t\t\t" + strings.Join(bttmtxt, " ") + "\n"
	f.WriteString(btcortxt)

	// 上面頂点座標の書き出し
	for j := 0; j < vcnt; j++ {
		toppoly = append(toppoly, list[j][0]/kansan)
		toppoly = append(toppoly, list[j][1]/kansan)
		toppoly = append(toppoly, toph/kansan)
	}
	// 上面データのテキスト化
	for m := range toppoly {
		toptxt = append(toptxt, strconv.FormatFloat(toppoly[m], 'f', -1, 64))
	}
	tpcortxt := "\t\t\t\t\t\t" + strings.Join(toptxt, " ") + "\n"
	f.WriteString(tpcortxt)

	// 各階頂点座標の書き出し
	for j := 1; j < story; j++ {
		for k := 0; k < vcnt; k++ {
			flrpoly = append(flrpoly, list[k][0]/kansan)
			flrpoly = append(flrpoly, list[k][1]/kansan)
			flrpoly = append(flrpoly, (elv+3.5*float64(j))/kansan)
		}
		// 各階データのテキスト化
		for m := range flrpoly {
			flrtxt = append(flrtxt, strconv.FormatFloat(flrpoly[m], 'f', -1, 64))
		}
		flcortxt := "\t\t\t\t\t\t" + strings.Join(flrtxt, " ") + "\n"
		f.WriteString(flcortxt)

		// 各階データの頂点座標の初期化
		flrpoly = flrpoly[:0]
		// 各階データの頂点座標の文字列の初期化
		flrtxt = flrtxt[:0]
	}

	// 側面頂点座標の書き出し
	for j := 0; j < vcnt; j++ {
		// 頂点下1
		sidepoly = append(sidepoly, list[j][0]/kansan)
		sidepoly = append(sidepoly, list[j][1]/kansan)
		sidepoly = append(sidepoly, btm/kansan)
		// 頂点下2
		sidepoly = append(sidepoly, list[(j+1)%vcnt][0]/kansan)
		sidepoly = append(sidepoly, list[(j+1)%vcnt][1]/kansan)
		sidepoly = append(sidepoly, btm/kansan)
		// 頂点上2
		sidepoly = append(sidepoly, list[(j+1)%vcnt][0]/kansan)
		sidepoly = append(sidepoly, list[(j+1)%vcnt][1]/kansan)
		sidepoly = append(sidepoly, toph/kansan)
		// 頂点上1
		sidepoly = append(sidepoly, list[j][0]/kansan)
		sidepoly = append(sidepoly, list[j][1]/kansan)
		sidepoly = append(sidepoly, toph/kansan)

		// 側面データのテキスト化
		for m := range sidepoly {
			sidetxt = append(sidetxt, strconv.FormatFloat(sidepoly[m], 'f', -1, 64))
		}
		sdcortxt := "\t\t\t\t\t\t" + strings.Join(sidetxt, " ") + "\n"
		f.WriteString(sdcortxt)

		// 側面データの頂点座標の初期化
		sidepoly = sidepoly[:0]
		// 側面データの頂点座標の文字列の初期化
		sidetxt = sidetxt[:0]
	}
	// f.WriteString("\t\t\t\t\t</float_array>\n")
}

// pos_tech_com はPos_technique_commonをファイルに書き出す
func pos_tech_com(f *os.File, v_num, i int) {
	f.WriteString("\t\t\t\t\t</float_array>\n")
	// Pos_technique_commonの書き出し
	f.WriteString("\t\t\t\t\t<technique_common>\n")
	// accessor count, sourceの設定
	accsor := "\t\t\t\t\t\t<accessor count=\"" + strconv.Itoa(v_num) + "\" source=\"#ID" + strconv.Itoa(i) + "-Pos-array\" stride=\"3\">\n"
	f.WriteString(accsor)
	f.WriteString("\t\t\t\t\t\t\t<param name=\"X\" type=\"float\" />\n")
	f.WriteString("\t\t\t\t\t\t\t<param name=\"Y\" type=\"float\" />\n")
	f.WriteString("\t\t\t\t\t\t\t<param name=\"Z\" type=\"float\" />\n")
	f.WriteString("\t\t\t\t\t\t</accessor>\n")
	f.WriteString("\t\t\t\t\t</technique_common>\n")
	f.WriteString("\t\t\t\t</source>\n")
}

// expnorm は法線ベクトルをファイルに書き出す
func expnorm(f *os.File, i, cnt, vcnt int, list [][]float64, elv, btm, toph float64) {
	var kansan = 0.0254 // ｍをインチ換算
	// 法線ベクトルの定義
	var nor_all [][]float64
	// 法線ベクトルの配列の定義
	var normal []float64
	//  法線ベクトルの文字列
	var t_normal []string
	// library_geometries(Normal)の書き出し
	sour_id2 := "\t\t\t\t<source id=\"ID" + strconv.Itoa(i) + "-Normal\">\n"
	f.WriteString(sour_id2)
	fl_arr_id2 := "\t\t\t\t\t<float_array id=\"ID" + strconv.Itoa(i) + "-Normal-array\" count=\"" + strconv.Itoa(cnt) + "\">\n"
	f.WriteString(fl_arr_id2)

	// for文で１行ずつ各頂点の法線ベクトルを書き出す
	// 基面，底面，上面，各階，側面の順に書き出す
	// 法線ベクトルを配列に格納して，これをテキスト変換して書き出す

	// 基面法線ベクトルの書き出し
	for j := 0; j < vcnt; j++ {
		// var norbase []float64
		// var nortxt []string
		ps1 := []float64{list[j][0], list[j][1], elv}
		ps2 := []float64{list[(j-1+vcnt)%vcnt][0], list[(j-1+vcnt)%vcnt][1], elv}
		ps3 := []float64{list[(j+1)%vcnt][0], list[(j+1)%vcnt][1], elv}
		nors := pkg.NorVec(ps1, ps2, ps3)
		nor_all = append(nor_all, nors)
	}
	for j := 0; j < vcnt; j++ {
		for k := 0; k < 3; k++ {
			normal = append(normal, nor_all[j][k])
		}
	}

	// 底面法線ベクトルの書き出し
	for j := 0; j < vcnt; j++ {
		pb1 := []float64{list[j][0], list[j][1], btm}
		pb2 := []float64{list[(j-1+vcnt)%vcnt][0], list[(j-1+vcnt)%vcnt][1], btm}
		pb3 := []float64{list[(j+1)%vcnt][0], list[(j+1)%vcnt][1], btm}
		norb := pkg.NorVec(pb1, pb2, pb3)
		nor_all = append(nor_all, norb)
	}
	for j := 0; j < vcnt; j++ {
		for k := 0; k < 3; k++ {
			normal = append(normal, nor_all[j][k])
		}
	}

	// 上面法線ベクトルの書き出し
	for j := 0; j < vcnt; j++ {
		pt1 := []float64{list[j][0], list[j][1], toph}
		pt2 := []float64{list[(j-1+vcnt)%vcnt][0], list[(j-1+vcnt)%vcnt][1], toph}
		pt3 := []float64{list[(j+1)%vcnt][0], list[(j+1)%vcnt][1], toph}
		nort := pkg.NorVec(pt1, pt2, pt3)
		nor_all = append(nor_all, nort)
	}
	for j := 0; j < vcnt; j++ {
		for k := 0; k < 3; k++ {
			normal = append(normal, nor_all[j][k])
		}
	}

	// 各階法線ベクトルの書き出し
	for j := 0; j < vcnt; j++ {
		pt1 := []float64{list[j][0], list[j][1], (elv + 3.5*float64(j)) / kansan}
		pt2 := []float64{list[(j-1+vcnt)%vcnt][0], list[(j-1+vcnt)%vcnt][1], (elv + 3.5*float64(j)) / kansan}
		pt3 := []float64{list[(j+1)%vcnt][0], list[(j+1)%vcnt][1], (elv + 3.5*float64(j)) / kansan}
		nort := pkg.NorVec(pt1, pt2, pt3)
		nor_all = append(nor_all, nort)
	}
	for j := 0; j < vcnt; j++ {
		for k := 0; k < 3; k++ {
			normal = append(normal, nor_all[j][k])
		}
	}

	// 側面法線ベクトルの書き出し
	for j := 0; j < vcnt; j++ {
		// 頂点下1
		p11 := []float64{list[j][0], list[j][1], btm}
		p12 := []float64{list[(j+1)%vcnt][0], list[(j+1)%vcnt][1], btm}
		p13 := []float64{list[j][0], list[j][1], toph}
		nor1 := pkg.NorVec(p11, p12, p13)
		nor_all = append(nor_all, nor1)
		for k := 0; k < 3; k++ {
			normal = append(normal, nor_all[j][k])
		}
		// 頂点下2
		p21 := []float64{list[(j+1)%vcnt][0], list[(j+1)%vcnt][1], btm}
		p22 := []float64{list[(j+1)%vcnt][0], list[(j+1)%vcnt][1], toph}
		p23 := []float64{list[j][0], list[j][1], btm}
		nor2 := pkg.NorVec(p21, p22, p23)
		nor_all = append(nor_all, nor2)
		for k := 0; k < 3; k++ {
			normal = append(normal, nor_all[j][k])
		}
		// 頂点下3
		p31 := []float64{list[(j+1)%vcnt][0], list[(j+1)%vcnt][1], toph}
		p32 := []float64{list[j][0], list[j][1], toph}
		p33 := []float64{list[(j+1)%vcnt][0], list[(j+1)%vcnt][1], btm}
		nor3 := pkg.NorVec(p31, p32, p33)
		nor_all = append(nor_all, nor3)
		for k := 0; k < 3; k++ {
			normal = append(normal, nor_all[j][k])
		}
		// 頂点上4
		p41 := []float64{list[j][0], list[j][1], toph}
		p42 := []float64{list[(j-1+vcnt)%vcnt][0], list[(j-1+vcnt)%vcnt][1], toph}
		p43 := []float64{list[j][0], list[j][1], btm}
		nor4 := pkg.NorVec(p41, p42, p43)
		nor_all = append(nor_all, nor4)
		for k := 0; k < 3; k++ {
			normal = append(normal, nor_all[j][k])
		}
	}

	for n := range normal {
		t_normal = append(t_normal, strconv.FormatFloat(normal[n], 'f', -1, 64))
	}

	// 法線ベクトルの書き出し
	nor_mesh := "\t\t\t\t\t\t" + strings.Join(t_normal, " ") + "\n"
	f.WriteString(nor_mesh)
	// f.WriteString("\t\t\t\t\t</float_array>\n")
}

// nor_tech_com はNor_technique_commonをファイルに書き出す
func nor_tech_com(f *os.File, v_num, i int) {
	f.WriteString("\t\t\t\t\t</float_array>\n")
	// Nor_technique_commonの書き出し
	f.WriteString("\t\t\t\t\t<technique_common>\n")
	// accessor coount, sourceの設定
	accsor2 := "\t\t\t\t\t\t<accessor count=\"" + strconv.Itoa(v_num) + "\" source=\"#ID" + strconv.Itoa(i) + "-Noamal-array\" stride=\"3\">\n"
	f.WriteString(accsor2)
	f.WriteString("\t\t\t\t\t\t\t<param name=\"X\" type=\"float\" />\n")
	f.WriteString("\t\t\t\t\t\t\t<param name=\"Y\" type=\"float\" />\n")
	f.WriteString("\t\t\t\t\t\t\t<param name=\"Z\" type=\"float\" />\n")
	f.WriteString("\t\t\t\t\t\t</accessor>\n")
	f.WriteString("\t\t\t\t\t</technique_common>\n")
	f.WriteString("\t\t\t\t</source>\n")
}

// exverpoly はvertices + polilystをファイルに書き出す
func exverpoly(f *os.File, i, pm_cnt, vcnt, v_num, story int, yane string) {
	// vertices + polilystの書き出し
	// verticesの書き出し
	id_vtx := "\t\t\t\t<vertices id=\"ID" + strconv.Itoa(i) + "-Vtx\">\n"
	f.WriteString(id_vtx)
	id_pos := "\t\t\t\t\t<input semantic=\"POSITION\" source=\"#ID" + strconv.Itoa(i) + "-Pos\" />\n"
	f.WriteString(id_pos)
	id_nor := "\t\t\t\t\t<input semantic=\"NORMAL\" source=\"#ID" + strconv.Itoa(i) + "-Normal\" />\n"
	f.WriteString(id_nor)
	f.WriteString("\t\t\t\t</vertices>\n")

	poly_mate := "\t\t\t\t<polylist count=\"" + strconv.Itoa(pm_cnt) + "\" material=\"Material2\">\n"
	f.WriteString(poly_mate)
	sour_id3 := "\t\t\t\t\t<input offset=\"0\" semantic=\"VERTEX\" source=\"#ID" + strconv.Itoa(i) + "-Vtx\" />\n"
	f.WriteString(sour_id3)

	// <vcount>～</vcount>
	var vcount []int
	// 建物モデルの平面のvcountの出力
	for j := 0; j < (3 + story - 1); j++ {
		vcount = append(vcount, vcnt)
	}
	// 建物モデルの側面のvcountの出力
	for j := 0; j < vcnt; j++ {
		vcount = append(vcount, 4)
	}
	// 屋根モデルのvcountの出力
	if yane == "kiri" {
		for j := 0; j < 22; j++ {
			vcount = append(vcount, 3)
		}
	}
	if yane == "yose" {
		for j := 0; j < 20; j++ {
			vcount = append(vcount, 3)
		}
	}
	if yane == "kata" {
		for j := 0; j < 16; j++ {
			vcount = append(vcount, 3)
		}
	}
	if yane == "kata1" || yane == "kata2" {
		for j := 0; j < 18; j++ {
			vcount = append(vcount, 3)
		}
	}
	if yane == "penta" {
		for j := 0; j < 25; j++ {
			vcount = append(vcount, 3)
		}
	}
	if yane == "kiri5" {
		for j := 0; j < 23; j++ {
			vcount = append(vcount, 3)
		}
	}
	if yane == "5kakudou" {
		for j := 0; j < 29; j++ {
			vcount = append(vcount, 3)
		}
	}
	if yane == "heprec" {
		for j := 0; j < 14; j++ {
			vcount = append(vcount, 3)
		}
	}
	if yane == "heptri" {
		for j := 0; j < 10; j++ {
			vcount = append(vcount, 3)
		}
	}
	var t_vcnt []string
	for v := range vcount {
		t_vcnt = append(t_vcnt, strconv.Itoa(vcount[v]))
	}
	f.WriteString("\t\t\t\t\t<vcount> " + strings.Join(t_vcnt, " ") + " </vcount>\n")

	// <p>～</p>
	p_num := v_num
	var p_no []int
	for j := 0; j < p_num; j++ {
		p_no = append(p_no, j)
	}
	var t_pnum []string
	for v := range p_no {
		t_pnum = append(t_pnum, strconv.Itoa(p_no[v]))
	}
	f.WriteString("\t\t\t\t\t<p> " + strings.Join(t_pnum, " ") + " </p>\n")
	f.WriteString("\t\t\t\t</polylist>\n")
	f.WriteString("\t\t\t</mesh>\n")
	f.WriteString("\t\t</geometry>\n")
}

// メイン関数
// BuildDAE はCOLLADA形式で多角形建物モデルを作成する
func BuildDAE(filename string, yanedtail Yanedtail) {
	// ログファイルを新規作成，追記，書き込み専用，パーミションは読むだけ
	file, err := os.OpenFile("collada.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	// ログの出力先を変更
	log.SetOutput(file)

	// COLLADA形式で出力するためのファイルを開く
	// f, err := os.Create("data/output.dae")
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

	// // 四角形リストを開く
	rectList, lr := OpenRect()

	// // 多角形リストを開く
	polyList, lp := OpenPoly()

	// 堅ろう建物のファイルを開く
	kList, lk := OpenKenro()

	// 無壁舎建物のファイルを開く
	mList, lm := OpenMuheki()

	// 四角形分割後のファイル数
	gid := lr + lp + lk + lm
	// gid := lk + lm
	log.Println("データ数 = ", gid)

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

	// 法線ベクトルの定義
	var nor_all [][]float64
	// 法線ベクトルの配列の定義
	var normal []float64
	//  法線ベクトルの文字列
	var t_normal []string

	// 底面高さの設定
	var btm float64
	// 上面高さの設定
	var toph float64
	// 用途地域番号の格納
	var areanum []int
	// 表示色の設定
	var col_rgb []float64

	// COLLADAファイルの出力開始
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
	f.WriteString("<COLLADA xmlns=\"http://www.collada.org/2005/11/COLLADASchema\" version=\"1.4.1\">\n")
	f.WriteString("\t<asset>\n")
	f.WriteString("\t\t<unit meter=\"0.0254\" name=\"inch\" />\n")
	f.WriteString("\t\t<up_axis>Z_UP</up_axis>\n")
	f.WriteString("\t</asset>\n")

	// library_visual_scenesの書き出し
	// <library_visual_scenes> ～ <node>
	f.WriteString("\t<library_visual_scenes>\n")
	f.WriteString("\t\t<visual_scene id=\"DefaultScene\">\n")
	f.WriteString("\t\t\t<node name=\"Building\">\n")
	// <instance_geometry> ～ </instance_geometry>
	// データ毎に必要な部分（url(#ID）のみが変わる）

	// 関数によるIDのみを変更した繰り返し部分の書き出し
	// lib_vis_scene(id)
	for i := 1; i <= gid; i++ {
		ig_url := "\t\t\t\t<instance_geometry url=\"#ID" + strconv.Itoa(i) + "\">\n"
		f.WriteString(ig_url)
		f.WriteString("\t\t\t\t\t<bind_material>\n")
		f.WriteString("\t\t\t\t\t\t<technique_common>\n")
		im_tar := "\t\t\t\t\t\t\t<instance_material symbol=\"Material2\" target=\"#ID" + strconv.Itoa(i) + "-material\">\n"
		f.WriteString(im_tar)
		f.WriteString("\t\t\t\t\t\t\t\t<bind_vertex_input semantic=\"UVSET0\" input_semantic=\"TEXCOORD\" input_set=\"0\" />\n")
		f.WriteString("\t\t\t\t\t\t\t</instance_material>\n")
		f.WriteString("\t\t\t\t\t\t</technique_common>\n")
		f.WriteString("\t\t\t\t\t</bind_material>\n")
		f.WriteString("\t\t\t\t</instance_geometry>\n")
	}
	// </node> ～ <library_geometries>
	f.WriteString("\t\t\t</node>\n")
	f.WriteString("\t\t</visual_scene>\n")
	f.WriteString("\t</library_visual_scenes>\n")
	f.WriteString("\t<library_geometries>\n")

	// // 逐次処理（１件ずつデータ読み込み → モデリングを行う）
	// 普通建物（傾斜屋根）の処理

	// 7角形屋根は三角屋根の棟高さと軒高さで隣り合う片流れ屋根の高さを決める
	// 三角屋根の棟と軒の頂点座標を格納する
	// var tripts TriVerts
	var tripts map[string]float64

	// 普通建物（傾斜屋根）の処理
	for r := range rectList {
		i := r + 1
		log.Println("ID=", rectList[i-1].ID)
		log.Println("Type=", rectList[i-1].Type)

		// 頂点数のチェック
		vcnt := len(rectList[i-1].List)

		// 基準面（地盤高さ）の読み込み
		level := rectList[i-1].Elv
		// 階層の読み込み
		story := rectList[i-1].Story
		log.Println("st=", story)

		// 面頂点数・座標データ数の総和
		// 屋根タイプにより変化する
		v_num, cnt, pm_cnt := counter(vcnt, story)
		if rectList[i-1].Type == "kiri" {
			// 切妻屋根の頂点数／(3+3+3+3+3x2+3x2+3x2)×2+3x2=66
			v_num = v_num + ((3+3)*2+3*2*3)*2 + 3*2
			// 切妻屋根の頂点データ数／((3+3+3+3+3x2+3x2+3x2)×2+3x2)x3=198
			cnt = cnt + (((3+3)*2+3*2*3)*2+3*2)*3
			// 切妻屋根の面数／(2x2+3×2)×2+2=22
			pm_cnt = pm_cnt + ((2*2+3*2)*2 + 2)
		} else if rectList[i-1].Type == "yose" {
			// 寄棟屋根の頂点数／(3x2+3x2x2)x2+3x2x4=60
			v_num = v_num + ((3*2+3*2*2)*2 + 3*2*4)
			// 寄棟屋根の頂点データ数／((3x2+3x2x2)x2+3x2x4)x3=180
			cnt = cnt + ((3*2+3*2*2)*2+3*2*4)*3
			// 寄棟屋根の面数／(2+2x2)x2+2x4=20
			pm_cnt = pm_cnt + ((2+2*2)*2 + 2*4)
		} else if rectList[i-1].Type == "kata" {
			// 片流れ屋根の頂点数／3x2x2+3x2+3x2+3x2x2+3x2+3x2=48
			v_num = v_num + (3*2*2 + 3*2 + 3*2 + 3*2*2 + 3*2 + 3*2)
			// 片流れ屋根の頂点データ数／(3x2x2+3x2+3x2+3x2x2+3x2+3x2)x3=144
			cnt = cnt + (3*2*2+3*2+3*2+3*2*2+3*2+3*2)*3
			// 片流れ屋根の面数／2x2+2+2+2x2+2+2=16
			pm_cnt = pm_cnt + (2*2 + 2 + 2 + 2*2 + 2 + 2)
		} else if rectList[i-1].Type == "kata1" || rectList[i-1].Type == "kata2" {
			// 片流れ屋根の頂点数／3x2x2+3x2+3x2+3x2x2+3x4+3x2=54
			v_num = v_num + (3*2*2 + 3*2 + 3*2 + 3*2*2 + 3*4 + 3*2)
			// 片流れ屋根の頂点データ数／(3x2x2+3x2+3x2+3x2x2+3x4+3x2)x3=162
			cnt = cnt + (3*2*2+3*2+3*2+3*2*2+3*4+3*2)*3
			// 片流れ屋根の面数／2x2+2+2+2x2+4+2=18
			pm_cnt = pm_cnt + (2*2 + 2 + 2 + 2*2 + 4 + 2)
		} else if rectList[i-1].Type == "penta" {
			// ５角形屋根の頂点数／((2x2+2x2+2x2)+(2+2+2+2+2+2)+1)x3=75
			v_num = v_num + ((2*2*3)+(2*6)+1)*3
			// ５角形屋根の頂点データ数／((2x2+2x2+2x2)+(2+2+2+2+2+2)+1)x3x3=225
			cnt = cnt + (((2*2*3)+(2*6)+1)*3)*3
			// ５角形屋根の面数／(2x2+2x2+2x2)+(2+2+2+2+2+2)+1=25
			pm_cnt = pm_cnt + ((2 * 2 * 3) + (2 * 6) + 1)
		} else if rectList[i-1].Type == "kiri5" {
			// 切妻屋根の頂点数／(3+3+3+3+3x2+3x2+3x2)×2+3x3=69
			v_num = v_num + ((3+3)*2+3*2*3)*2 + 3*3
			// 切妻屋根の頂点データ数／((3+3+3+3+3x2+3x2+3x2)×2+3x3)x3=207
			cnt = cnt + (((3+3)*2+3*2*3)*2+3*3)*3
			// 切妻屋根の面数／(2x2+3x2)×2+3=23
			pm_cnt = pm_cnt + ((2*2+3*2)*2 + 3)
		} else if rectList[i-1].Type == "5kakudou" {
			// 三角形片流れ屋根の頂点数／3x5x2+3x2x5=60
			v_num = v_num + 3*5*2 + 3*2*5
			// 三角形片流れ屋根の頂点データ数／(3x5x2+3x2x5)x3=180
			cnt = cnt + (3*5*2+3*2*5)*3
			// 三角形片流れ屋根の面数／5x2+2x5=30
			pm_cnt = pm_cnt + (5*2 + 2*5)
		} else if rectList[i-1].Type == "heprec" {
			// 変形片流れ屋根の頂点数／(2x2+2x4+2)x3=42
			v_num = v_num + (2*2+2*4+2)*3
			// 変形片流れ屋根の頂点データ数／((2x2+2x4+2)x3)x3=126
			cnt = cnt + (2*2+2*4+2)*3*3
			// 変形片流れ屋根の面数／2x2+2x4+2=14
			pm_cnt = pm_cnt + (2*2 + 2*4 + 2)
		} else if rectList[i-1].Type == "heptri" {
			// 三角形片流れ屋根の頂点数／(2+2x3+2)x3=30
			v_num = v_num + (2+2*3+2)*3
			// 三角形片流れ屋根の頂点データ数／((2+2x3+2)x3)x3=90
			cnt = cnt + ((2+2*3+2)*3)*3
			// 三角形片流れ屋根の面数／2+2x3+2=10
			pm_cnt = pm_cnt + (2 + 2*3 + 2)
		}

		// 建物の上面高さの設定
		toph = rectList[i-1].Elv + float64(story)*3.3

		// 建物地下部分のモデリングのための地下深さの設定
		nest := 0.3
		btm = level - nest

		// 用途地域番号の読み込み
		youto := rectList[i-1].Area
		// log.Println("youto=", youto)
		anum := AreaNum(youto)
		// log.Println("anum=", anum)
		areanum = append(areanum, anum)

		// library_geometries(Position)の書き出し
		lib_geo_pos(f, i, cnt)

		// 基面データの頂点座標の初期化
		basepoly = basepoly[:0]
		// 基面データの頂点座標の文字列の初期化
		basetxt = basetxt[:0]
		// 底面データの頂点座標の初期化
		bttmpoly = bttmpoly[:0]
		// 底面データの頂点座標の文字列の初期化
		bttmtxt = bttmtxt[:0]
		// 上面データの頂点座標の初期化
		toppoly = toppoly[:0]
		// 上面データの頂点座標の文字列の初期化
		toptxt = toptxt[:0]

		// 法線ベクトルの初期化
		nor_all = nor_all[:0][:0]
		// 法線ベクトルの配列の初期化
		normal = normal[:0]
		//  法線ベクトルの文字列の初期化
		t_normal = t_normal[:0]
		// 屋根頂点データの文字列
		var yanetxt []string
		// 屋根頂点の法線ベクトルの文字列
		var yanenor []string

		if rectList[i-1].Type == "kiri" {
			// 軒庇の出とケラバの厚さを設定する
			hisashi := yanedtail.Kirinoki
			keraba := yanedtail.Kirikera
			incline := yanedtail.Kiriincl
			yaneatu := yanedtail.Kiriroof

			// 切妻屋根の頂点の法線ベクトルをリストにテキスト化して書き出す
			yanetxt, yanenor = KiriYane(rectList[i-1].List, toph, hisashi, keraba, incline, yaneatu)

		} else if rectList[i-1].Type == "yose" {
			// 軒庇の出とケラバの厚さを設定する
			hisashi := yanedtail.Yosenoki
			keraba := hisashi
			incline := yanedtail.Yoseincl
			yaneatu := yanedtail.Yoseroof

			// 寄棟屋根の頂点の法線ベクトルをリストにテキスト化して書き出す
			yanetxt, yanenor = YoseYane(rectList[i-1].List, toph, hisashi, keraba, incline, yaneatu)

		} else if rectList[i-1].Type == "kata" {
			// 軒庇の出とケラバの厚さを設定する
			hisashi := yanedtail.Katanoki
			keraba := yanedtail.Katakera
			incline := yanedtail.Kataincl
			yaneatu := yanedtail.Kataroof

			// 片流れ屋根の頂点の法線ベクトルをリストにテキスト化して書き出す
			yanetxt, yanenor = KataYane(rectList[i-1].List, toph, hisashi, keraba, incline, yaneatu)

		} else if rectList[i-1].Type == "heptri" {
			// 軒庇の出とケラバの厚さを設定する
			hisashi := 0.60
			keraba := 0.30
			incline := 0.3
			yaneatu := 0.075

			// 7角形屋根の頂点の法線ベクトルをリストにテキスト化して書き出す
			yanetxt, yanenor, tripts = TriYane(rectList[i-1].List, toph, hisashi, keraba, incline, yaneatu)
			log.Println("tripts=", tripts)

		} else if rectList[i-1].Type == "kata1" {
			// 軒庇の出とケラバの厚さを設定する
			hisashi := 0.60
			keraba := 0.30
			incline := 0.3
			yaneatu := 0.075

			// 片流れ屋根の頂点の法線ベクトルをリストにテキスト化して書き出す
			yanetxt, yanenor = KataYane1(rectList[i-1].List, toph, hisashi, keraba, incline, yaneatu, tripts)

		} else if rectList[i-1].Type == "kata2" {
			// 軒庇の出とケラバの厚さを設定する
			hisashi := 0.60
			keraba := 0.30
			incline := 0.3
			yaneatu := 0.075

			// 片流れ屋根の頂点の法線ベクトルをリストにテキスト化して書き出す
			yanetxt, yanenor = KataYane2(rectList[i-1].List, toph, hisashi, keraba, incline, yaneatu, tripts)

		} else if rectList[i-1].Type == "penta" {
			// 軒庇の出とケラバの厚さを設定する
			hisashi := 0.60
			keraba := hisashi
			incline := 0.3
			yaneatu := 0.075

			// ５角形屋根の頂点の法線ベクトルをリストにテキスト化して書き出す
			yanetxt, yanenor = PentaYane(rectList[i-1].List, toph, hisashi, keraba, incline, yaneatu)

		} else if rectList[i-1].Type == "kiri5" {
			// 軒庇の出とケラバの厚さを設定する
			hisashi := 0.60
			keraba := 0.30
			incline := 0.45
			yaneatu := 0.11

			// ５角形屋根の切妻屋根の頂点の法線ベクトルをリストにテキスト化して書き出す
			yanetxt, yanenor = Kiri5Yane(rectList[i-1].List, toph, hisashi, keraba, incline, yaneatu)

		} else if rectList[i-1].Type == "5kakudou" {
			// 軒庇の出とケラバの厚さを設定する
			hisashi := 0.60
			incline := 0.45
			yaneatu := 0.11

			// ５角形屋根の頂点の法線ベクトルをリストにテキスト化して書き出す
			yanetxt, yanenor = Kakudou5(rectList[i-1].List, toph, hisashi, incline, yaneatu)
		}

		// 建物モデル（直方体）を基面，底面，各階，上面，側面の順に書き出す
		// X/Y/Z座標を配列に格納して，これをテキスト変換して書き出す
		expcords(f, vcnt, rectList[i-1].List, rectList[i-1].Elv, btm, toph, story)

		// 屋根モデルの頂点座標リストの出力
		cortxt := "\t\t\t\t\t\t" + strings.Join(yanetxt, " ") + "\n"
		f.WriteString(cortxt)

		// Pos_technique_commonの書き出し
		pos_tech_com(f, v_num, i)

		// 法線ベクトルの書き出し
		expnorm(f, i, cnt, vcnt, rectList[i-1].List, rectList[i-1].Elv, btm, toph)

		// 屋根モデルの頂点の法線ベクトルの出力
		nortxt := "\t\t\t\t\t\t" + strings.Join(yanenor, " ") + "\n"
		f.WriteString(nortxt)

		// Nor_technique_commonの書き出し
		nor_tech_com(f, v_num, i)

		// vertices + polilystの書き出し
		exverpoly(f, i, pm_cnt, vcnt, v_num, story, rectList[i-1].Type)
	}

	// 普通建物（多角形）の処理
	for p := range polyList {
		i := p + 1

		// 頂点数
		vcnt := len(polyList[i-1].List)
		log.Println("vertex=", vcnt)

		// 基準面（地盤高さ）の読み込み
		level := polyList[i-1].Elv
		// 屋上高さの読み込み
		rf := polyList[i-1].Roof
		log.Println("rf=", rf)
		// 階層の読み込み
		story := polyList[i-1].Story
		log.Println("st=", story)

		// 屋上高さが判る場合の建物の上面高さの設定
		if rf != 0.0 {
			// 建物階数の設定
			story = int((rf - level) / 3.3)
			// 建物の上面高さの設定
			toph = level + float64(story)*3.3
			// if story != 0 {
			// 	// 建物階数の設定
			// 	story = int((rf - level) / 3.3)
			// 	// 建物の上面高さの設定
			// 	toph = level + float64(story)*3.3
			// } else {
			// 	// 建物階数の設定
			// 	story = 1
			// 	// 建物の上面高さの設定
			// 	toph = rf
			// }
		} else {
			// 屋上高さは判らないが階層が判る場合
			if story != 0 {
				// 建物の上面高さの設定
				toph = level + float64(story)*3.3
			} else {
				// 建物階数の設定
				story = 2
				// 建物の上面高さの設定
				toph = level + float64(story)*3.3
			}
		}

		// 面頂点数・座標データ数の総和
		v_num, cnt, pm_cnt := counter(vcnt, story)

		// 建物地下部分のモデリングのための地下深さの設定
		nest := 0.3
		btm = level - nest

		// 用途地域番号の読み込み
		youto := polyList[i-1].Area
		// log.Println("youto=", youto)
		anum := AreaNum(youto)
		// log.Println("anum=", anum)
		areanum = append(areanum, anum)

		// library_geometries(Position)の書き出し
		lib_geo_pos(f, i+lr, cnt)

		// 基面データの頂点座標の初期化
		basepoly = basepoly[:0]
		// 基面データの頂点座標の文字列の初期化
		basetxt = basetxt[:0]
		// 底面データの頂点座標の初期化
		bttmpoly = bttmpoly[:0]
		// 底面データの頂点座標の文字列の初期化
		bttmtxt = bttmtxt[:0]
		// 上面データの頂点座標の初期化
		toppoly = toppoly[:0]
		// 上面データの頂点座標の文字列の初期化
		toptxt = toptxt[:0]
		// 各階データの頂点座標の初期化
		flrpoly = flrpoly[:0]
		// 各階データの頂点座標の文字列の初期化
		flrtxt = flrtxt[:0]

		// 法線ベクトルの初期化
		nor_all = nor_all[:0][:0]
		// 法線ベクトルの配列の初期化
		normal = normal[:0]
		//  法線ベクトルの文字列の初期化
		t_normal = t_normal[:0]

		// for文で１行ずつ各頂点のX/Y/Z座標を書き出す
		// 基面，底面，上面，各階，側面の順に書き出す
		// X/Y/Z座標を配列に格納して，これをテキスト変換して書き出す
		expcords(f, vcnt, polyList[i-1].List, polyList[i-1].Elv, btm, toph, story)

		// Pos_technique_commonの書き出し
		pos_tech_com(f, v_num, i+lr)

		// 法線ベクトルの書き出し
		expnorm(f, i+lr, cnt, vcnt, polyList[i-1].List, polyList[i-1].Elv, btm, toph)

		// Nor_technique_commonの書き出し
		nor_tech_com(f, v_num, i+lr)

		// vertices + polilystの書き出し
		yane := "flat"
		exverpoly(f, i+lr, pm_cnt, vcnt, v_num, story, yane)
	}

	// 堅ろう建物の処理
	for k := range kList {
		i := k + 1

		// 頂点数
		vnum := len(kList[i-1].Cords)
		vcnt := vnum - 1
		// log.Println("vertex=", vcnt)

		// 建ぺい率・容積率の設定
		bcr := kList[i-1].Build
		far := kList[i-1].Floor

		// 用途地域番号の読み込み
		youto := kList[i-1].Area
		log.Println("youto=", youto)
		anum := AreaNum(youto)
		log.Println("anum=", anum)
		areanum = append(areanum, anum)

		// 基準面（地盤高さ）の読み込み
		level := kList[i-1].Elv
		// 屋上高さの読み込み
		rf := kList[i-1].Roof
		log.Println("rf=", rf)
		// 階層の読み込み
		story := kList[i-1].Story
		log.Println("st=", story)

		// 屋上高さが判る場合の建物の上面高さの設定
		if rf != 0.0 {
			if story != 0 {
				// 建物階数の設定
				story = int((rf - level) / 3.5)
				// 建物の上面高さの設定
				toph = level + float64(story)*3.5
			} else {
				// 建物階数の設定
				story = 1
				// 建物の上面高さの設定
				toph = rf
			}
		} else {
			// 屋上高さは判らないが階層が判る場合
			if story != 0 {
				// 建物の上面高さの設定
				toph = level + float64(story)*3.5
			} else {
				// 建物の階数が設定されていない場合は乱数で建物階数を設定する
				story = Stcalc(story, bcr, far, anum)
				log.Println("story=", story)
				// 建物の上面高さの設定
				toph = level + float64(story)*3.5
				log.Println("toph=", toph)
			}
		}

		// 面頂点数・座標データ数の総和
		v_num, cnt, pm_cnt := counter(vcnt, story)
		// log.Println(v_num)
		// log.Println(cnt)
		// log.Println(pm_cnt)

		// 建物地下部分のモデリングのための地下深さの設定
		// 地下階がある場合は階数に応じてZ座標を地下階部分の深さだけ下げる
		bf := kList[i-1].Basement
		btm = Bfcalc(level, story, bf)
		// log.Println("btm=", btm)

		// library_geometries(Position)の書き出し
		lib_geo_pos(f, i+lr+lp, cnt)

		// 基面データの頂点座標の初期化
		basepoly = basepoly[:0]
		// 基面データの頂点座標の文字列の初期化
		basetxt = basetxt[:0]
		// 底面データの頂点座標の初期化
		bttmpoly = bttmpoly[:0]
		// 底面データの頂点座標の文字列の初期化
		bttmtxt = bttmtxt[:0]
		// 上面データの頂点座標の初期化
		toppoly = toppoly[:0]
		// 上面データの頂点座標の文字列の初期化
		toptxt = toptxt[:0]
		// 各階データの頂点座標の初期化
		flrpoly = flrpoly[:0]
		// 各階データの頂点座標の文字列の初期化
		flrtxt = flrtxt[:0]

		// 法線ベクトルの初期化
		nor_all = nor_all[:0][:0]
		// 法線ベクトルの配列の初期化
		normal = normal[:0]
		//  法線ベクトルの文字列の初期化
		t_normal = t_normal[:0]

		// for文で１行ずつ各頂点のX/Y/Z座標を書き出す
		// 基面，底面，上面，各階，側面の順に書き出す
		// X/Y/Z座標を配列に格納して，これをテキスト変換して書き出す
		expcords(f, vcnt, kList[i-1].Cords, kList[i-1].Elv, btm, toph, story)

		// Pos_technique_commonの書き出し
		// pos_tech_com(f, v_num, i+lr+lp)
		pos_tech_com(f, v_num, i)

		// 法線ベクトルの書き出し
		expnorm(f, i+lr+lp, cnt, vcnt, kList[i-1].Cords, kList[i-1].Elv, btm, toph)

		// Nor_technique_commonの書き出し
		// nor_tech_com(f, v_num, i+lr+lp)
		nor_tech_com(f, v_num, i)

		// vertices + polilystの書き出し
		yane := "flat"
		exverpoly(f, i+lr+lp, pm_cnt, vcnt, v_num, story, yane)
	}

	// 無壁舎建物の処理
	for m := range mList {
		i := m + 1

		// 頂点数
		vnum := len(mList[i-1].Cords)
		vcnt := vnum - 1
		// log.Println("vertex=", vcnt)

		// 基準面（地盤高さ）の読み込み
		level := mList[i-1].Elv
		// 屋上高さの読み込み
		rf := mList[i-1].Roof
		log.Println("rf=", rf)
		// 階層の読み込み
		story := mList[i-1].Story
		log.Println("st=", story)

		// 屋上高さが判る場合の建物の上面高さの設定
		if rf != 0.0 {
			if story != 0 {
				// 建物階数の設定
				story = int((rf - level) / 3.3)
				// 建物の上面高さの設定
				toph = level + float64(story)*3.3
			} else {
				// 建物階数の設定
				story = 1
				// 建物の上面高さの設定
				toph = rf
			}
		} else {
			// 屋上高さは判らないが階層が判る場合
			if story != 0 {
				// 建物の上面高さの設定
				toph = level + float64(story)*3.3
			} else {
				// 建物階数の設定
				story = 1
				// 建物の上面高さの設定
				toph = level + float64(story)*3.3
			}
		}

		// // 面頂点数・座標データ数の総和
		v_num, cnt, pm_cnt := counter(vcnt, story)

		// 建物地下部分のモデリングのための地下深さの設定
		btm = level - 0.3

		// 用途地域番号の読み込み
		youto := mList[i-1].Area
		// log.Println("youto=", youto)
		anum := AreaNum(youto)
		// log.Println("anum=", anum)
		areanum = append(areanum, anum)

		// library_geometries(Position)の書き出し
		lib_geo_pos(f, i+lr+lp+lk, cnt)

		// 基面データの頂点座標の初期化
		basepoly = basepoly[:0]
		// 基面データの頂点座標の文字列の初期化
		basetxt = basetxt[:0]
		// 底面データの頂点座標の初期化
		bttmpoly = bttmpoly[:0]
		// 底面データの頂点座標の文字列の初期化
		bttmtxt = bttmtxt[:0]
		// 上面データの頂点座標の初期化
		toppoly = toppoly[:0]
		// 上面データの頂点座標の文字列の初期化
		toptxt = toptxt[:0]
		// 各階データの頂点座標の初期化
		flrpoly = flrpoly[:0]
		// 各階データの頂点座標の文字列の初期化
		flrtxt = flrtxt[:0]

		// 法線ベクトルの初期化
		nor_all = nor_all[:0][:0]
		// 法線ベクトルの配列の初期化
		normal = normal[:0]
		//  法線ベクトルの文字列の初期化
		t_normal = t_normal[:0]

		// for文で１行ずつ各頂点のX/Y/Z座標を書き出す
		// 基面，底面，上面，各階，側面の順に書き出す
		// X/Y/Z座標を配列に格納して，これをテキスト変換して書き出す
		expcords(f, vcnt, mList[i-1].Cords, mList[i-1].Elv, btm, toph, story)

		// Pos_technique_commonの書き出し
		pos_tech_com(f, v_num, i+lr+lp+lk)
		// pos_tech_com(f, v_num, i+lk)

		// 法線ベクトルの書き出し
		expnorm(f, i+lr+lp+lk, cnt, vcnt, mList[i-1].Cords, mList[i-1].Elv, btm, toph)
		// expnorm(f, i+lk, cnt, vcnt, mList[i-1].Cords, mList[i-1].Elv, btm, toph)

		// Nor_technique_commonの書き出し
		nor_tech_com(f, v_num, i+lr+lp+lk)
		// nor_tech_com(f, v_num, i+lk)

		// vertices + polilystの書き出し
		yane := "flat"
		exverpoly(f, i+lr+lp+lk, pm_cnt, vcnt, v_num, story, yane)
		// exverpoly(f, i+lk, pm_cnt, vcnt, v_num, yane)
	}

	// library_materialsの書き出し
	f.WriteString("\t</library_geometries>\n")
	f.WriteString("\t<library_materials>\n")
	for i := 1; i <= gid; i++ {
		mate_id := "\t\t<material id=\"ID" + strconv.Itoa(i) + "-material\">\n"
		f.WriteString(mate_id)
		effe_url := "\t\t\t<instance_effect url=\"#ID" + strconv.Itoa(i) + "-surface\" />\n"
		f.WriteString(effe_url)
		f.WriteString("\t\t</material>\n")
	}
	f.WriteString("\t</library_materials>\n")
	f.WriteString("\t<library_effects>\n")

	// library_effectsの書き出し
	for i := 1; i <= gid; i++ {
		effe_id := "\t\t<effect id=\"ID" + strconv.Itoa(i) + "-surface\">\n"
		f.WriteString(effe_id)
		f.WriteString("\t\t\t<profile_COMMON>\n")
		f.WriteString("\t\t\t\t<technique sid=\"COMMON\">\n")
		f.WriteString("\t\t\t\t\t<lambert>\n")
		f.WriteString("\t\t\t\t\t\t<diffuse>\n")
		// log.Println("areanum=", areanum[i-1])
		if areanum[i-1] == 0 {
			col_rgb = []float64{0.41, 0.41, 0.41, 1.00}
		} else if areanum[i-1] != 0 {
			col_rgb = AreaRGB(areanum[i-1])
		}
		// log.Println("col_rgb=", col_rgb)
		var t_rgb []string
		for c := range col_rgb {
			t_rgb = append(t_rgb, strconv.FormatFloat(col_rgb[c], 'f', -1, 64))
		}
		// log.Println("col_rgb=", col_rgb)
		effe_col := "\t\t\t\t\t\t\t<color>" + strings.Join(t_rgb, " ") + "</color>\n"
		f.WriteString(effe_col)
		f.WriteString("\t\t\t\t\t\t</diffuse>\n")
		f.WriteString("\t\t\t\t\t</lambert>\n")
		f.WriteString("\t\t\t\t</technique>\n")
		f.WriteString("\t\t\t</profile_COMMON>\n")
		f.WriteString("\t\t</effect>\n")
	}
	// COLLADAファイルの出力終了
	f.WriteString("\t</library_effects>\n")
	f.WriteString("\t<scene>\n")
	f.WriteString("\t\t<instance_visual_scene url=\"#DefaultScene\" />\n")
	f.WriteString("\t</scene>\n")
	f.WriteString("</COLLADA>\n")
}
