package internal

import (
	"log"
	"os"
	"strconv"
)

// TerrDXFはDXF形式で地形モデルを作成する
func TerrDXF(filename string, x_matrix, y_matrix, z_matrix [][]float64, x_dot, y_dot int, x_max, x_min, y_max, y_min, z_max, z_min float64) {
	// DXF形式で出力するためのファイルを開く
	// f, err := os.Create("data/outputgeo.dxf")
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

	// DXFファイルの出力開始
	// TIN（不等辺三角形網）で地形の３次元形状を表現
	// Rowの設定
	row := y_dot - 1
	// Columnの設定
	col := x_dot - 1

	f.WriteString("  0\n")
	f.WriteString("SECTION\n")
	f.WriteString("  2\n")
	f.WriteString("ENTITIES\n")

	// 点（X座標，Y座標，Z座標）
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			// 左下三角形の出力
			// 3DFACEのセクション毎に3頂点の座標を出力する
			f.WriteString("  0\n")
			f.WriteString("3DFACE\n")
			f.WriteString("  8\n")
			f.WriteString("0\n")
			// 第1頂点の出力
			f.WriteString(" 10\n")
			xl1 := x_matrix[i][j] * 0.0254
			f.WriteString(strconv.FormatFloat(xl1, 'f', -1, 64) + "\n")
			f.WriteString(" 20\n")
			yl1 := y_matrix[i][j] * 0.0254
			f.WriteString(strconv.FormatFloat(yl1, 'f', -1, 64) + "\n")
			f.WriteString(" 30\n")
			zl1 := z_matrix[i][j] * 0.0254
			f.WriteString(strconv.FormatFloat(zl1, 'f', -1, 64) + "\n")
			// 第2頂点の出力
			f.WriteString(" 11\n")
			xl2 := x_matrix[i][j+1] * 0.0254
			f.WriteString(strconv.FormatFloat(xl2, 'f', -1, 64) + "\n")
			f.WriteString(" 21\n")
			yl2 := y_matrix[i][j+1] * 0.0254
			f.WriteString(strconv.FormatFloat(yl2, 'f', -1, 64) + "\n")
			f.WriteString(" 31\n")
			zl2 := z_matrix[i][j+1] * 0.0254
			f.WriteString(strconv.FormatFloat(zl2, 'f', -1, 64) + "\n")
			// 第3頂点の出力
			f.WriteString(" 12\n")
			xl3 := x_matrix[i+1][j] * 0.0254
			f.WriteString(strconv.FormatFloat(xl3, 'f', -1, 64) + "\n")
			f.WriteString(" 22\n")
			yl3 := y_matrix[i+1][j] * 0.0254
			f.WriteString(strconv.FormatFloat(yl3, 'f', -1, 64) + "\n")
			f.WriteString(" 32\n")
			zl3 := z_matrix[i+1][j] * 0.0254
			f.WriteString(strconv.FormatFloat(zl3, 'f', -1, 64) + "\n")
			// 第4頂点の出力
			f.WriteString(" 13\n")
			xl4 := x_matrix[i][j] * 0.0254
			f.WriteString(strconv.FormatFloat(xl4, 'f', -1, 64) + "\n")
			f.WriteString(" 23\n")
			yl4 := y_matrix[i][j] * 0.0254
			f.WriteString(strconv.FormatFloat(yl4, 'f', -1, 64) + "\n")
			f.WriteString(" 33\n")
			zl4 := z_matrix[i][j] * 0.0254
			f.WriteString(strconv.FormatFloat(zl4, 'f', -1, 64) + "\n")
			// 右上三角形の出力
			// 3DFACEのセクション毎に3頂点の座標を出力する
			f.WriteString("  0\n")
			f.WriteString("3DFACE\n")
			f.WriteString("  8\n")
			f.WriteString("0\n")
			// 第1頂点の出力
			f.WriteString(" 10\n")
			xu1 := x_matrix[i][j+1] * 0.0254
			f.WriteString(strconv.FormatFloat(xu1, 'f', -1, 64) + "\n")
			f.WriteString(" 20\n")
			yu1 := y_matrix[i][j+1] * 0.0254
			f.WriteString(strconv.FormatFloat(yu1, 'f', -1, 64) + "\n")
			f.WriteString(" 30\n")
			zu1 := z_matrix[i][j+1] * 0.0254
			f.WriteString(strconv.FormatFloat(zu1, 'f', -1, 64) + "\n")
			// 第2頂点の出力
			f.WriteString(" 11\n")
			xu2 := x_matrix[i+1][j] * 0.0254
			f.WriteString(strconv.FormatFloat(xu2, 'f', -1, 64) + "\n")
			f.WriteString(" 21\n")
			yu2 := y_matrix[i+1][j] * 0.0254
			f.WriteString(strconv.FormatFloat(yu2, 'f', -1, 64) + "\n")
			f.WriteString(" 31\n")
			zu2 := z_matrix[i+1][j] * 0.0254
			f.WriteString(strconv.FormatFloat(zu2, 'f', -1, 64) + "\n")
			// 第3頂点の出力
			f.WriteString(" 12\n")
			xu3 := x_matrix[i+1][j+1] * 0.0254
			f.WriteString(strconv.FormatFloat(xu3, 'f', -1, 64) + "\n")
			f.WriteString(" 22\n")
			yu3 := y_matrix[i+1][j+1] * 0.0254
			f.WriteString(strconv.FormatFloat(yu3, 'f', -1, 64) + "\n")
			f.WriteString(" 32\n")
			zu3 := z_matrix[i+1][j+1] * 0.0254
			f.WriteString(strconv.FormatFloat(zu3, 'f', -1, 64) + "\n")
			// 第4頂点の出力
			f.WriteString(" 13\n")
			xu4 := x_matrix[i][j+1] * 0.0254
			f.WriteString(strconv.FormatFloat(xu4, 'f', -1, 64) + "\n")
			f.WriteString(" 23\n")
			yu4 := y_matrix[i][j+1] * 0.0254
			f.WriteString(strconv.FormatFloat(yu4, 'f', -1, 64) + "\n")
			f.WriteString(" 33\n")
			zu4 := z_matrix[i][j+1] * 0.0254
			f.WriteString(strconv.FormatFloat(zu4, 'f', -1, 64) + "\n")
		}
	}

	// エンティティを書き出す（終了部分）
	f.WriteString("  0\n")
	f.WriteString("ENDSEC\n")
	f.WriteString("  0\n")
	f.WriteString("EOF\n")
}
