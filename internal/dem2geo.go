package internal

import (
	"bufio"
	"log"
	"math"
	"os"
	"stera/pkg"
	"strings"
)

type Zahyo struct {
	x float64
	y float64
	z float64
}

// TinMesh は標高メッシュデータから地形モデリングのための座標データを作成する
func TinMesh(z_zero float64) (x_matrix, y_matrix, z_matrix [][]float64, x_len, y_len float64, x_dot, y_dot int, x_max, x_min, y_max, y_min, z_max, z_min float64) {
	// 変数の宣言
	var kansan = 0.0254 // ｍをインチ換算
	var x_zahyo []float64
	var y_zahyo []float64
	var hyoukou []float64

	// データ数のカウント
	fl, er := os.Open("data/dem_data.txt")
	_, l, _, _ := pkg.FileCount(fl)
	var counter = l
	log.Println("counter=", counter)
	if er != nil {
		log.Fatal()
	}
	defer fl.Close()

	// 標高メッシュデータの読み込み
	fp, er := os.Open("data/dem_data.txt")
	if er != nil {
		log.Fatal()
	}
	defer fp.Close()

	// 一行ずつデータを読み込む
	scanner := bufio.NewScanner(fp)
	zahyo := []*Zahyo{}
	var lines = 0
	for scanner.Scan() {
		// ここで一行ずつ処理
		gstr := scanner.Text()
		// 右端の「,」を削除，「,」がない行末でもエラーにならない
		gstr = strings.TrimRight(gstr, ",")

		// MultiPointをPointに置換する
		if strings.Contains(gstr, "MultiPoint") {
			gstr = strings.Replace(gstr, "[ [", "[", 1)
			gstr = strings.Replace(gstr, "] ]", "]", 1)
		}

		// GeoJSON構造体の変数stcDataを宣言
		geo := pkg.ParseDEM(gstr)

		// データを構造体のスライスに読み込む
		xyz := Zahyo{x: geo[0], y: geo[1], z: geo[2]}
		zahyo = append(zahyo, &xyz)
		lines += 1
	}
	log.Println("lines=", lines)

	// 各座標データを各スライスに読み込む
	for i := 0; i < lines; i++ {
		x_zahyo = append(x_zahyo, zahyo[i].x/kansan) // X座標の読み込み
		y_zahyo = append(y_zahyo, zahyo[i].y/kansan) // Y座標の読み込み
		hyoukou = append(hyoukou, zahyo[i].z/kansan) // 地盤高データの読み込み
	}

	// 配列（マトリックス）の大きさを決める処理
	x_max = x_zahyo[0] // X座標の最大値の初期化
	y_max = y_zahyo[0] // Y座標の最大値の初期化
	x_min = x_zahyo[0] // X座標の最小値の初期化
	y_min = y_zahyo[0] // Y座標の最小値の初期化

	// X座標値・Y座標値の最大値と最小値を求める
	// 標高データの最大値と最小値を求める
	for i := 1; i < counter; i++ {
		if x_max < x_zahyo[i] {
			x_max = x_zahyo[i] // X座標値の最大値の更新
		}
		if y_max < y_zahyo[i] {
			y_max = y_zahyo[i] // Y座標値の最大値の更新
		}
		if x_min > x_zahyo[i] {
			x_min = x_zahyo[i] // X座標値の最小値の更新
		}
		if y_min > y_zahyo[i] {
			y_min = y_zahyo[i] // Y座標値の最小値の更新
		}
		if z_max < hyoukou[i] {
			z_max = hyoukou[i] // Z座標値の最大値の更新
		}
		if z_min > hyoukou[i] {
			z_min = hyoukou[i] // Z座標値の最小値の更新
		}
	}
	// 最低高さの設定
	if z_zero < z_min {
		z_min = z_zero
	}

	// マトリックスの大きさを決めるためにX（東西）方向とY（南北）方向の大きさを求める
	x_len = math.Abs(x_max - x_min) // X（東西）方向の幅
	y_len = math.Abs(y_max - y_min) // Y（南北）方向の高さ

	// X（東西）方向のマス目の大きさ（最小値）を求める
	// X座標の間隔は上（北）の方は狭く，下（南）の方は広く，同じ段は等間隔になっている
	// X（東西）方向のマス目の大きさの初期値
	x_dist := math.Abs(x_zahyo[1] - x_zahyo[0])

	var x_dist_temp float64 // X座標間隔の暫定値
	var x_cnt_temp int      // X座標のデータ数の暫定値
	row_cnt := 0            // 行（段）数の初期値
	x_countmax := 0         // X座標のデータ数の最大値の初期値
	for i := 2; i < counter; i++ {
		// X座標値は東／右に行くほど値が大きくなる
		// X座標値が小さくなっている箇所で改行（段）されたと判断できる
		if x_zahyo[i] > x_zahyo[i-1] {
			// X（東西）方向のマス目の計算過程での大きさ
			x_dist_temp = math.Abs(x_zahyo[i] - x_zahyo[i-1])
			if x_dist > x_dist_temp {
				x_dist = x_dist_temp
			}
			// X座標値が大きくなっている間はデータ数を1つ増やす
			x_cnt_temp = x_cnt_temp + 1
		} else {
			// 改行（段）されたので行（段）数を1つ増やす
			row_cnt = row_cnt + 1
			// X座標のデータ数の最大値を更新する
			if x_countmax < x_cnt_temp {
				x_countmax = x_cnt_temp
			}
			// X座標のデータ数をリセットする
			x_cnt_temp = 0
		}
	}
	log.Println("x_dist=", x_dist)
	log.Println("row_cnt=", row_cnt)
	log.Println("x_countmax=", x_countmax)

	// X（東西）方向の頂点データ数
	x_dot = x_countmax + 1
	log.Println("x_dot=", x_dot)
	// Y（南北）方向の頂点データ数
	y_dot = row_cnt + 1
	log.Println("y_dot=", y_dot)

	// マトリックスの大きさ（２次元配列）を決める
	// X座標
	x_matrix = make([][]float64, y_dot)
	for k := 0; k < y_dot; k++ {
		x_matrix[k] = make([]float64, x_dot)
	}
	// Y座標
	y_matrix = make([][]float64, y_dot)
	for k := 0; k < y_dot; k++ {
		y_matrix[k] = make([]float64, x_dot)
	}
	// Z座標
	z_matrix = make([][]float64, y_dot)
	for k := 0; k < y_dot; k++ {
		z_matrix[k] = make([]float64, x_dot)
	}

	// マトリックスにデータを割り付ける
	col_num := 0        // X座標のインデックス，左端が0
	row_num := 0        // Y座標のインデックス，下端が0
	blunk := 0          // 空白部分の桁数，初期値は0
	var x_row []float64 // X座標の配列
	var y_row []float64 // Y座標の配列
	var z_row []float64 // Z座標の配列

	// Y座標値のデータは南から北（下から上）へ並ぶ
	// X座標値のデータは西から東（左から右）へ並ぶ
	// 一行ずつX･Y･Zデータを配列に割り付ける
	for i := 0; i < counter; i++ {
		// 1番目のデータの格納
		if i == 0 {
			// 各行（段）の1番目のデータの配列（配置場所）を決める
			// 左端からの距離は許容範囲か？
			if (x_zahyo[i] - x_min) < x_dist*0.9 {
				x_row = append(x_row, x_zahyo[i])
				y_row = append(y_row, y_zahyo[i])
				z_row = append(z_row, hyoukou[i])
				col_num = col_num + 1
			} else if (x_zahyo[i] - x_min) >= x_dist*0.9 {
				// 左端との間に空白部分がある
				// 空白部分の桁数を求める
				blunk = int(math.Round((x_zahyo[i] - x_min) / x_dist))
				log.Println("blunk(1)=", blunk)
				// 空白部分を最低標高値としてダミーで埋める
				for j := 0; j < blunk; j++ {
					x_row = append(x_row, x_min+x_dist*float64(j))
					y_row = append(y_row, y_zahyo[i])
					z_row = append(z_row, z_min)
					col_num = col_num + 1
				}
				// 空白データの後ろに1番目のデータを追加する
				x_row = append(x_row, x_zahyo[i])
				y_row = append(y_row, y_zahyo[i])
				z_row = append(z_row, hyoukou[i])
				col_num = col_num + 1
			}
			log.Println("col_num-1=", col_num)
		}

		if i > 0 && i < (counter-1) {
			// 2番目以降のデータの格納
			if x_zahyo[i] > x_zahyo[i-1] {
				// 1番目のデータからの距離は許容範囲か？
				if (x_zahyo[i] - x_zahyo[i-1]) < x_dist*1.5 {
					x_row = append(x_row, x_zahyo[i])
					y_row = append(y_row, y_zahyo[i])
					z_row = append(z_row, hyoukou[i])
					col_num = col_num + 1
				} else if (x_zahyo[i] - x_zahyo[i-1]) >= x_dist*1.5 {
					// 左側のデータとの間に空白部分がある
					// 空白部分の桁数を求める
					blunk = int(math.Ceil((x_zahyo[i] - x_zahyo[i-1]) / x_dist))
					log.Println("blunk(2)=", blunk)
					// 空白部分を最低標高値としてダミーで埋める
					for j := 1; j < blunk; j++ {
						x_row = append(x_row, x_zahyo[i-1]+x_dist*float64(j))
						y_row = append(y_row, y_zahyo[i-1])
						z_row = append(z_row, z_min)
						col_num = col_num + 1
					}
					// 空白データの後ろに2番目のデータを追加する
					x_row = append(x_row, x_zahyo[i])
					y_row = append(y_row, y_zahyo[i])
					z_row = append(z_row, hyoukou[i])
					col_num = col_num + 1
				}
				log.Println("col_num-2=", col_num)

			} else if x_zahyo[i] < x_zahyo[i-1] {
				// X座標値が小さくなった箇所で改行（段）処理を行う
				log.Println("改行（段）処理")
				// 改行（段）される前に空白部分がないか調べる
				if col_num < x_dot {
					// 空白部分の桁数を求める
					blunk = x_dot - col_num
					log.Println("blunk(3)=", blunk)
					// 空白部分を最低標高値としてダミーで埋める
					for j := 1; j <= blunk; j++ {
						x_row = append(x_row, x_zahyo[i-1]+x_dist*float64(j))
						y_row = append(y_row, y_zahyo[i-1])
						z_row = append(z_row, z_min)
						col_num = col_num + 1
					}
				}
				log.Println("col_num-3=", col_num)

				x_row_len := len(x_row)
				log.Println("x_row_len=", x_row_len)
				y_row_len := len(y_row)
				log.Println("y_row_len=", y_row_len)
				z_row_len := len(z_row)
				log.Println("z_row_len=", z_row_len)

				// マトリックスへデータを格納する
				// X座標を格納する
				for x := 0; x < x_dot; x++ {
					x_matrix[row_num][x] = x_row[x]
				}
				// X座標のスライスを空白にする
				x_row = x_row[:0]
				// Y座標を格納する
				for y := 0; y < x_dot; y++ {
					y_matrix[row_num][y] = y_row[y]
				}
				// Y座標のスライスを空白にする
				y_row = y_row[:0]
				// Z座標を格納する
				for z := 0; z < x_dot; z++ {
					z_matrix[row_num][z] = z_row[z]
				}
				// Z座標のスライスを空白にする
				z_row = z_row[:0]
				log.Println("x_matrix=", x_matrix[row_num][x_dot-1])
				log.Println("y_matrix=", y_matrix[row_num][x_dot-1])
				log.Println("z_matrix=", z_matrix[row_num][x_dot-1])

				// 改行（段）されたので行（row）番号を追加する
				row_num = row_num + 1
				// 改行（段）されたので桁（col）番号を0に戻す
				col_num = 0
				log.Println("row_num=", row_num)
				log.Println("col_num=", col_num)

				// 左端からの距離は許容範囲か？
				if (x_zahyo[i] - x_min) <= x_dist*1.0 {
					x_row = append(x_row, x_zahyo[i])
					y_row = append(y_row, y_zahyo[i])
					z_row = append(z_row, hyoukou[i])
					col_num = col_num + 1
				} else if (x_zahyo[i] - x_min) > x_dist*1.0 {
					// 左端との間に空白部分がある
					// 空白部分の桁数を求める
					blunk = int(math.Round((x_zahyo[i] - x_min) / x_dist))
					log.Println("blunk(4)=", blunk)
					// 空白部分を最低標高値としてダミーで埋める
					for j := 0; j < blunk; j++ {
						x_row = append(x_row, x_min+x_dist*float64(j))
						y_row = append(y_row, y_zahyo[i])
						z_row = append(z_row, z_min)
						col_num = col_num + 1
					}
					// 空白データの後ろに1番目のデータを追加する
					x_row = append(x_row, x_zahyo[i])
					y_row = append(y_row, y_zahyo[i])
					z_row = append(z_row, hyoukou[i])
					col_num = col_num + 1
				}
			}
		}

		// 最後のデータに対してマトリックスを完成させる
		if i == (counter - 1) {
			// 最後のデータが右端（東端）に達していない場合の処理
			if col_num < x_dot {
				// 空白部分の桁数を求める
				blunk = x_dot - col_num
				log.Println("blunk(5)=", blunk)
				// 空白部分を最低標高値としてダミーで埋める
				for j := 1; j <= blunk; j++ {
					x_row = append(x_row, x_zahyo[i]+x_dist*float64(j))
					y_row = append(y_row, y_zahyo[i])
					z_row = append(z_row, z_min)
					col_num = col_num + 1
				}
			}

			// マトリックスへデータを格納する
			// X座標を格納する
			for x := 0; x < x_dot; x++ {
				x_matrix[row_num][x] = x_row[x]
			}
			// Y座標を格納する
			for y := 0; y < x_dot; y++ {
				y_matrix[row_num][y] = y_row[y]
			}
			// Z座標を格納する
			for z := 0; z < x_dot; z++ {
				z_matrix[row_num][z] = z_row[z]
			}
			// 最後までデータが追加されたので行（row）番号を追加する
			row_num = row_num + 1
			log.Println("row_num=", row_num)
		}
	}

	return x_matrix, y_matrix, z_matrix, x_len, y_len, x_dot, y_dot, x_max, x_min, y_max, y_min, z_max, z_min
}
