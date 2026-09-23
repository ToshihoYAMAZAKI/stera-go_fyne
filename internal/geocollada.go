package internal

import (
	"fmt"
	"log"
	"os"
	"stera/pkg"
	"strconv"
	"strings"
)

// TerrDAEはCOLLADA形式で地形モデルを作成する
func TerrDAE(filename string, x_matrix, y_matrix, z_matrix [][]float64, x_dot, y_dot int) {
	// COLLADA形式で出力するためのファイルを開く
	// f, err := os.Create("data/outputgeo.dae")
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
	f.WriteString("\t\t\t<node name=\"Terrain\">\n")
	// <instance_geometry> ～ </instance_geometry>
	// データ毎に必要な部分（url(#ID）のみが変わる）
	// IDの設定
	id := y_dot - 1
	// id = y_countMax
	// 関数によるIDのみを変更した繰り返し部分の書き出し
	// lib_vis_scene(id)
	for i := 1; i <= id; i++ {
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

	// 地形の３角メッシュのモデリング
	// 四角ポリゴンの数
	// polycount := (x_dot - 1) * (y_dot - 1)
	// log.Println("polycount", polycount)
	// 各行(段)の四角ポリゴンの３角メッシュの数
	meshcount := 2 * (x_dot - 1)
	// log.Println("meshcount", meshcount)
	// 各行(段)の面頂点数の総和（３角メッシュ数×頂点数）
	ver_num := meshcount * 3
	// log.Println("ver_num", ver_num)
	// 各行(段)の配列データ数（３角メッシュ数×頂点数×軸数（X･Y･Z））
	count := meshcount * 3 * 3
	// log.Println("count", count)
	// 各行(段)の面の総数
	poly_mate_cnt := meshcount
	// log.Println("poly_mate_cnt", poly_mate_cnt)

	// ３角メッシュの頂点座標の定義
	var mesh []float64
	// ３角メッシュの頂点座標の文字列
	var t_mesh []string
	// 法線ベクトルの定義
	var nor_all [][]float64
	// 法線ベクトルの配列の定義
	var normal []float64
	//  法線ベクトルの文字列
	var t_normal []string

	// 行(段)単位にモデリングする
	for row := 0; row < id; row++ {
		// // ID番号の設定（iをidに変更）
		gid := row + 1

		// library_geometries(Position)の書き出し
		// geometry idの設定
		geo_id := "\t\t<geometry id=\"ID" + strconv.Itoa(gid) + "\">\n"
		f.WriteString(geo_id)
		f.WriteString("\t\t\t<mesh>\n")
		// source idの設定
		sour_id := "\t\t\t\t<source id=\"ID" + strconv.Itoa(gid) + "-Pos\">\n"
		f.WriteString(sour_id)
		fl_arr_id := "\t\t\t\t\t<float_array id=\"ID" + strconv.Itoa(gid) + "-Pos-array\" count=\"" + strconv.Itoa(count) + "\">\n"
		f.WriteString(fl_arr_id)

		// マトリックス（３角メッシュ／データ部分）の定義
		// 左下から右へ（X方向），そして左に戻って１行上へ（Y方向）
		// 相対する３角メッシュを左下側→右上側→左下側→右上側の順に定義する

		// ３角メッシュの頂点座標の初期化
		mesh = mesh[:0]
		// ３角メッシュの頂点座標の文字列の初期化
		t_mesh = t_mesh[:0]
		// 法線ベクトルの初期化
		nor_all = nor_all[:0][:0]
		// 法線ベクトルの配列の初期化
		normal = normal[:0]
		//  法線ベクトルの文字列の初期化
		t_normal = t_normal[:0]

		// 列ごとに四角ポリゴン(３角メッシュ×２)を最終列までモデリングする
		// 列番号（X番号）
		// log.Println("row", row)
		for col := 0; col < x_dot-1; col++ {
			// log.Println("col", col)
			// 左下側３角メッシュを定義する
			// X･Y座標，標高の割り当て
			x1 := x_matrix[row][col] // 頂点(1)
			y1 := y_matrix[row][col]
			z1 := z_matrix[row][col]
			x2 := x_matrix[row][col+1] // 頂点(2)
			y2 := y_matrix[row][col+1]
			z2 := z_matrix[row][col+1]
			x3 := x_matrix[row+1][col] // 頂点(3)
			y3 := y_matrix[row+1][col]
			z3 := z_matrix[row+1][col]
			// 右上側３角メッシュを定義する
			// X･Y座標，標高の割り当て
			x4 := x_matrix[row][col+1] // 頂点(2)
			y4 := y_matrix[row][col+1]
			z4 := z_matrix[row][col+1]
			x5 := x_matrix[row+1][col+1] // 頂点(4)
			y5 := y_matrix[row+1][col+1]
			z5 := z_matrix[row+1][col+1]
			x6 := x_matrix[row+1][col] // 頂点(3)
			y6 := y_matrix[row+1][col]
			z6 := z_matrix[row+1][col]

			mesh = append(mesh, x1, y1, z1, x2, y2, z2, x3, y3, z3, x4, y4, z4, x5, y5, z5, x6, y6, z6)
			// if col == x_dot-2 {
			// 	log.Println("mesh", mesh)
			// }

			p11 := []float64{x1, y1, z1}
			p12 := []float64{x3, y3, z3}
			p13 := []float64{x2, y2, z2}
			nor1 := pkg.NorVec(p11, p12, p13)
			nor_all = append(nor_all, nor1)
			p21 := []float64{x2, y2, z2}
			p22 := []float64{x1, y1, z1}
			p23 := []float64{x3, y3, z3}
			nor2 := pkg.NorVec(p21, p22, p23)
			nor_all = append(nor_all, nor2)
			p31 := []float64{x3, y3, z3}
			p32 := []float64{x2, y2, z2}
			p33 := []float64{x1, y1, z1}
			nor3 := pkg.NorVec(p31, p32, p33)
			nor_all = append(nor_all, nor3)
			p41 := []float64{x4, y4, z4}
			p42 := []float64{x6, y6, z6}
			p43 := []float64{x5, y5, z5}
			nor4 := pkg.NorVec(p41, p42, p43)
			nor_all = append(nor_all, nor4)
			p51 := []float64{x5, y5, z5}
			p52 := []float64{x4, y4, z4}
			p53 := []float64{x6, y6, z6}
			nor5 := pkg.NorVec(p51, p52, p53)
			nor_all = append(nor_all, nor5)
			p61 := []float64{x6, y6, z6}
			p62 := []float64{x5, y5, z5}
			p63 := []float64{x4, y4, z4}
			nor6 := pkg.NorVec(p61, p62, p63)
			nor_all = append(nor_all, nor6)

			// log.Println("nor_all", nor_all)

			for i := 0; i < 6; i++ {
				for j := 0; j < 3; j++ {
					normal = append(normal, nor_all[i][j])
				}
			}
			// // log.Println("normal", normal)
		}
		// log.Println("row", row)
		// log.Println("mesh", mesh)

		for m := range mesh {
			t_mesh = append(t_mesh, strconv.FormatFloat(mesh[m], 'f', -1, 64))
		}
		// log.Println("t_mesh", t_mesh)

		for n := range normal {
			t_normal = append(t_normal, strconv.FormatFloat(normal[n], 'f', -1, 64))
		}
		// log.Println("normal", normal)

		// ３角メッシュの頂点座標の書き出し
		tin_mesh := "\t\t\t\t\t\t" + strings.Join(t_mesh, " ") + "\n"
		f.WriteString(tin_mesh)

		// Pos_technique_commonの書き出し
		f.WriteString("\t\t\t\t\t</float_array>\n")
		f.WriteString("\t\t\t\t\t<technique_common>\n")
		// accessor count, sourceの設定
		accsor := "\t\t\t\t\t\t<accessor count=\"" + strconv.Itoa(ver_num) + "\" source=\"#ID" + strconv.Itoa(gid) + "-Pos-array\" stride=\"3\">\n"
		f.WriteString(accsor)
		f.WriteString("\t\t\t\t\t\t\t<param name=\"X\" type=\"float\" />\n")
		f.WriteString("\t\t\t\t\t\t\t<param name=\"Y\" type=\"float\" />\n")
		f.WriteString("\t\t\t\t\t\t\t<param name=\"Z\" type=\"float\" />\n")
		f.WriteString("\t\t\t\t\t\t</accessor>\n")
		f.WriteString("\t\t\t\t\t</technique_common>\n")
		f.WriteString("\t\t\t\t</source>\n")

		// library_geometries(Normal)の書き出し
		sour_id2 := "\t\t\t\t<source id=\"ID" + strconv.Itoa(gid) + "-Normal\">\n"
		f.WriteString(sour_id2)
		fl_arr_id2 := "\t\t\t\t\t<float_array id=\"ID" + strconv.Itoa(gid) + "-Normal-array\" count=\"" + strconv.Itoa(count) + "\">\n"
		f.WriteString(fl_arr_id2)

		// 法線ベクトルの書き出し
		nor_mesh := "\t\t\t\t\t\t" + strings.Join(t_normal, " ") + "\n"
		f.WriteString(nor_mesh)

		// Nor_technique_commonの書き出し
		f.WriteString("\t\t\t\t\t</float_array>\n")
		f.WriteString("\t\t\t\t\t<technique_common>\n")
		// accessor coount, sourceの設定
		accsor2 := "\t\t\t\t\t\t<accessor count=\"" + strconv.Itoa(ver_num) + "\" source=\"#ID" + strconv.Itoa(gid) + "-Noamal-array\" stride=\"3\">\n"
		f.WriteString(accsor2)
		f.WriteString("\t\t\t\t\t\t\t<param name=\"X\" type=\"float\" />\n")
		f.WriteString("\t\t\t\t\t\t\t<param name=\"Y\" type=\"float\" />\n")
		f.WriteString("\t\t\t\t\t\t\t<param name=\"Z\" type=\"float\" />\n")
		f.WriteString("\t\t\t\t\t\t</accessor>\n")
		f.WriteString("\t\t\t\t\t</technique_common>\n")
		f.WriteString("\t\t\t\t</source>\n")

		// vertices + polilystの書き出し
		// verticesの書き出し
		id_vtx := "\t\t\t\t<vertices id=\"ID" + strconv.Itoa(gid) + "-Vtx\">\n"
		f.WriteString(id_vtx)
		id_pos := "\t\t\t\t\t<input semantic=\"POSITION\" source=\"#ID" + strconv.Itoa(gid) + "-Pos\" />\n"
		f.WriteString(id_pos)
		id_nor := "\t\t\t\t\t<input semantic=\"NORMAL\" source=\"#ID" + strconv.Itoa(gid) + "-Normal\" />\n"
		f.WriteString(id_nor)
		f.WriteString("\t\t\t\t</vertices>\n")

		poly_mate := "\t\t\t\t<polylist count=\"" + strconv.Itoa(poly_mate_cnt) + "\" material=\"Material2\">\n"
		f.WriteString(poly_mate)
		sour_id3 := "\t\t\t\t\t<input offset=\"0\" semantic=\"VERTEX\" source=\"#ID" + strconv.Itoa(gid) + "-Vtx\" />\n"
		f.WriteString(sour_id3)

		// <vcount>～</vcount>
		var vcnt []int
		for j := 0; j < meshcount; j++ {
			vcnt = append(vcnt, 3)
		}
		var t_vcnt []string
		for v := range vcnt {
			t_vcnt = append(t_vcnt, strconv.Itoa(vcnt[v]))
		}
		f.WriteString("\t\t\t\t\t<vcount> " + strings.Join(t_vcnt, " ") + " </vcount>\n")

		// <p>～</p>
		p_num := ver_num
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

	// library_materialsの書き出し
	f.WriteString("\t</library_geometries>\n")
	f.WriteString("\t<library_materials>\n")
	for i := 1; i <= id; i++ {
		mate_id := "\t\t<material id=\"ID" + strconv.Itoa(i) + "-material\">\n"
		f.WriteString(mate_id)
		effe_url := "\t\t\t<instance_effect url=\"#ID" + strconv.Itoa(i) + "-surface\" />\n"
		f.WriteString(effe_url)
		f.WriteString("\t\t</material>\n")
	}
	f.WriteString("\t</library_materials>\n")
	f.WriteString("\t<library_effects>\n")

	// library_effectsの書き出し
	for i := 1; i <= id; i++ {
		effe_id := "\t\t<effect id=\"ID" + strconv.Itoa(i) + "-surface\">\n"
		f.WriteString(effe_id)
		f.WriteString("\t\t\t<profile_COMMON>\n")
		f.WriteString("\t\t\t\t<technique sid=\"COMMON\">\n")
		f.WriteString("\t\t\t\t\t<lambert>\n")
		f.WriteString("\t\t\t\t\t\t<diffuse>\n")
		col_rgb := "0.41 0.41 0.41 1"
		effe_col := "\t\t\t\t\t\t\t<color>" + col_rgb + "</color>\n"
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
