package internal

import (
	"log"
	"math"
	"os"
	"stera/pkg"
)

// ExposYose は片流れ屋根建物のは頂点座標をファイルに書き出す
func ExposKata(f *os.File, id, fid string, vcnt int, list [][]float64, elv, btm, toph, hisashi, keraba, incline, yaneatu float64, story int) {
	log.Println("片流れ屋根")
	// 屋根モデル（片流れ屋根）
	// 4つの頂点座標の定義
	x1 := list[0][0]
	y1 := list[0][1]
	x2 := list[1][0]
	y2 := list[1][1]
	x3 := list[2][0]
	y3 := list[2][1]
	x4 := list[3][0]
	y4 := list[3][1]

	// 片流れ屋根の棟方向を妻面と平面の長さを比較してチェックする
	d1 := pkg.DistVerts(x1, y1, x2, y2)
	d2 := pkg.DistVerts(x4, y4, x1, y1)
	log.Println("d1, d2", d1, d2)
	if d1 < d2 {
		xp := x1
		yp := y1
		x1 = x2
		y1 = y2
		x2 = x3
		y2 = y3
		x3 = x4
		y3 = y4
		x4 = xp
		y4 = yp
	}

	// 頂点1と頂点2を結ぶ線分に平行な直線の式（流れ面上手）
	m1, n1 := pkg.ParaLine(x1, y1, x2, y2, x3, y3, hisashi/4)
	// log.Println("y = " + strconv.FormatFloat(m1, 'f', -1, 64) + "x + " + strconv.FormatFloat(n1, 'f', -1, 64))
	// 頂点3と頂点4を結ぶ線分に平行な直線の式（流れ面下手）
	m2, n2 := pkg.ParaLine(x3, y3, x4, y4, x1, y1, hisashi)
	// log.Println("y = " + strconv.FormatFloat(m2, 'f', -1, 64) + "x + " + strconv.FormatFloat(n2, 'f', -1, 64))
	// 頂点2と頂点3を結ぶ線分に平行な直線の式（側面）
	m3, n3 := pkg.ParaLine(x2, y2, x3, y3, x4, y4, keraba)
	// log.Println("y = " + strconv.FormatFloat(m3, 'f', -1, 64) + "x + " + strconv.FormatFloat(n3, 'f', -1, 64))
	// 頂点4と頂点1を結ぶ線分に平行な直線の式（側面）
	m4, n4 := pkg.ParaLine(x4, y4, x1, y1, x2, y2, keraba)
	// log.Println("y = " + strconv.FormatFloat(m4, 'f', -1, 64) + "x + " + strconv.FormatFloat(n4, 'f', -1, 64))

	// 屋根伏せの4頂点の座標を求める（屋根面の下端）
	xo1, yo1 := pkg.SeekInsec(m4, n4, m1, n1)
	xo2, yo2 := pkg.SeekInsec(m1, n1, m3, n3)
	xo3, yo3 := pkg.SeekInsec(m3, n3, m2, n2)
	xo4, yo4 := pkg.SeekInsec(m2, n2, m4, n4)
	log.Println("xo1, yo1, xo2, yo2, xo3, yo3, xo4, yo4", xo1, yo1, xo2, yo2, xo3, yo3, xo4, yo4)

	// 軒鼻の突き出し長さ
	nt := yaneatu / math.Sqrt(math.Pow(incline, 2)+1) * incline
	// 壁面から軒庇の上端までの長さ（軒庇＋軒鼻）
	hichtop := hisashi + nt
	// 頂点3と頂点4を結ぶ線分に平行な直線の式（流れ面下手上端）
	mkt2, nkt2 := pkg.ParaLine(x3, y3, x4, y4, x1, y1, hichtop)

	// 屋根伏せの4頂点の座標を求める（屋根面の上端）
	xtp1, ytp1 := pkg.SeekInsec(m4, n4, m1, n1)
	xtp2, ytp2 := pkg.SeekInsec(m1, n1, m3, n3)
	xtp3, ytp3 := pkg.SeekInsec(m3, n3, mkt2, nkt2)
	xtp4, ytp4 := pkg.SeekInsec(mkt2, nkt2, m4, n4)
	log.Println("xtp1, ytp1, xtp2, ytp2, xtp3, ytp3, xtp4, ytp4", xtp1, ytp1, xtp2, ytp2, xtp3, ytp3, xtp4, ytp4)

	// 屋根の側面の長さを求める
	sf1 := pkg.DistVerts(xo4, yo4, xo1, yo1)
	sf2 := pkg.DistVerts(xo2, yo2, xo3, yo3)
	log.Println("sf1, sf2", sf1, sf2)

	// 軒先下端高さ（庇×屋根勾配）
	nbt := (toph - hisashi/math.Sqrt(math.Pow(incline, 2)+1)*incline)
	// nbt := (toph - hisashi*incline)
	// 軒先上端高さ
	ntp := nbt + yaneatu/math.Sqrt(math.Pow(incline, 2)+1)
	log.Println("nbt, ntp", nbt, ntp)

	// 屋根の流れ面上手の下端高さ
	mbt1 := nbt + sf1*incline
	mbt2 := nbt + sf2*incline
	var mbt float64
	if mbt1 < mbt2 {
		mbt = mbt1
	} else {
		mbt = mbt2
	}
	// 屋根の流れ面上手の上端高さ
	// mtp1 := mbt1 + yaneatu*math.Sqrt(1+math.Pow(incline, 2))
	// mtp2 := mbt2 + yaneatu*math.Sqrt(1+math.Pow(incline, 2))
	mtp := mbt + yaneatu*math.Sqrt(1+math.Pow(incline, 2))
	log.Println("mbt, mtp", mbt, mtp)

	// 屋根モデルの頂点座標をリストに書き出す
	// 屋根頂点の法線ベクトルを算出してリストに書き出す
	var yanepoly []float64
	// PolyGMLID をfid から設定するための添え字の定義
	sub := 0
	// SurfaceType の定義
	var sftype string

	// 屋根底面・三角形１のIDと頂点座標の書き出し
	sub = sub + 1
	sftype = "Roof"
	// bldg:boundedBy をファイルに書き出す
	expbldg(f, sftype, id, fid, sub)
	// 頂点座標データ（閉じた図形）
	yanepoly = append(yanepoly, xo1, yo1, mbt)
	yanepoly = append(yanepoly, xo4, yo4, nbt)
	yanepoly = append(yanepoly, xo2, yo2, mbt)
	yanepoly = append(yanepoly, xo1, yo1, mbt)
	// bldg:boundedBy の書き出しを終了する
	exitbldg(f, yanepoly, sftype)
	// 屋根データの頂点座標の初期化
	yanepoly = yanepoly[:0]

	// 屋根底面・三角形２のIDと頂点座標の書き出し
	sub = sub + 1
	sftype = "Roof"
	// bldg:boundedBy をファイルに書き出す
	expbldg(f, sftype, id, fid, sub)
	// 頂点座標データ（閉じた図形）
	yanepoly = append(yanepoly, xo4, yo4, nbt)
	yanepoly = append(yanepoly, xo3, yo3, nbt)
	yanepoly = append(yanepoly, xo2, yo2, mbt)
	yanepoly = append(yanepoly, xo4, yo4, nbt)
	// bldg:boundedBy の書き出しを終了する
	exitbldg(f, yanepoly, sftype)
	// 屋根データの頂点座標の初期化
	yanepoly = yanepoly[:0]

	// 屋根上面・三角形１のIDと頂点座標の書き出し
	sub = sub + 1
	sftype = "Roof"
	// bldg:boundedBy をファイルに書き出す
	expbldg(f, sftype, id, fid, sub)
	// 頂点座標データ（閉じた図形）
	yanepoly = append(yanepoly, xtp1, ytp1, mtp)
	yanepoly = append(yanepoly, xtp4, ytp4, ntp)
	yanepoly = append(yanepoly, xtp2, ytp2, mtp)
	yanepoly = append(yanepoly, xtp1, ytp1, mtp)
	// bldg:boundedBy の書き出しを終了する
	exitbldg(f, yanepoly, sftype)
	// 屋根データの頂点座標の初期化
	yanepoly = yanepoly[:0]

	// 屋根上面・三角形２のIDと頂点座標の書き出し
	sub = sub + 1
	sftype = "Roof"
	// bldg:boundedBy をファイルに書き出す
	expbldg(f, sftype, id, fid, sub)
	// 頂点座標データ（閉じた図形）
	yanepoly = append(yanepoly, xtp4, ytp4, ntp)
	yanepoly = append(yanepoly, xtp3, ytp3, ntp)
	yanepoly = append(yanepoly, xtp2, ytp2, mtp)
	yanepoly = append(yanepoly, xtp4, ytp4, ntp)
	// bldg:boundedBy の書き出しを終了する
	exitbldg(f, yanepoly, sftype)
	// 屋根データの頂点座標の初期化
	yanepoly = yanepoly[:0]

	// 軒端１・三角形-1のIDと頂点座標の書き出し
	sub = sub + 1
	sftype = "Roof"
	// bldg:boundedBy をファイルに書き出す
	expbldg(f, sftype, id, fid, sub)
	// 頂点座標データ（閉じた図形）
	yanepoly = append(yanepoly, xtp4, ytp4, ntp)
	yanepoly = append(yanepoly, xtp3, ytp3, ntp)
	yanepoly = append(yanepoly, xo3, yo3, nbt)
	yanepoly = append(yanepoly, xtp4, ytp4, ntp)
	// bldg:boundedBy の書き出しを終了する
	exitbldg(f, yanepoly, sftype)
	// 屋根データの頂点座標の初期化
	yanepoly = yanepoly[:0]

	// 軒端２・三角形-2のIDと頂点座標の書き出し
	sub = sub + 1
	sftype = "Roof"
	// bldg:boundedBy をファイルに書き出す
	expbldg(f, sftype, id, fid, sub)
	// 頂点座標データ（閉じた図形）
	yanepoly = append(yanepoly, xtp4, ytp4, ntp)
	yanepoly = append(yanepoly, xo3, yo3, nbt)
	yanepoly = append(yanepoly, xo4, yo4, nbt)
	yanepoly = append(yanepoly, xtp4, ytp4, ntp)
	// bldg:boundedBy の書き出しを終了する
	exitbldg(f, yanepoly, sftype)
	// 屋根データの頂点座標の初期化
	yanepoly = yanepoly[:0]

	// 棟端１・三角形-1のIDと頂点座標の書き出し
	sub = sub + 1
	sftype = "Roof"
	// bldg:boundedBy をファイルに書き出す
	expbldg(f, sftype, id, fid, sub)
	// 頂点座標データ（閉じた図形）
	yanepoly = append(yanepoly, xtp1, ytp1, mtp)
	yanepoly = append(yanepoly, xo1, yo1, mbt)
	yanepoly = append(yanepoly, xo2, yo2, mbt)
	yanepoly = append(yanepoly, xtp1, ytp1, mtp)
	// bldg:boundedBy の書き出しを終了する
	exitbldg(f, yanepoly, sftype)
	// 屋根データの頂点座標の初期化
	yanepoly = yanepoly[:0]

	// 棟端２・三角形-2のIDと頂点座標の書き出し
	sub = sub + 1
	sftype = "Roof"
	// bldg:boundedBy をファイルに書き出す
	expbldg(f, sftype, id, fid, sub)
	// 頂点座標データ（閉じた図形）
	yanepoly = append(yanepoly, xtp1, ytp1, mtp)
	yanepoly = append(yanepoly, xo2, yo2, mbt)
	yanepoly = append(yanepoly, xtp2, ytp2, mtp)
	yanepoly = append(yanepoly, xtp1, ytp1, mtp)
	// bldg:boundedBy の書き出しを終了する
	exitbldg(f, yanepoly, sftype)
	// 屋根データの頂点座標の初期化
	yanepoly = yanepoly[:0]

	// ケラバ１・三角形1-1のIDと頂点座標の書き出し
	sub = sub + 1
	sftype = "Roof"
	// bldg:boundedBy をファイルに書き出す
	expbldg(f, sftype, id, fid, sub)
	// 頂点座標データ（閉じた図形）
	yanepoly = append(yanepoly, xo1, yo1, mbt)
	yanepoly = append(yanepoly, xtp1, ytp1, mtp)
	yanepoly = append(yanepoly, xtp4, ytp4, ntp)
	yanepoly = append(yanepoly, xo1, yo1, mbt)
	// bldg:boundedBy の書き出しを終了する
	exitbldg(f, yanepoly, sftype)
	// 屋根データの頂点座標の初期化
	yanepoly = yanepoly[:0]

	// ケラバ１・三角形1-2のIDと頂点座標の書き出し
	sub = sub + 1
	sftype = "Roof"
	// bldg:boundedBy をファイルに書き出す
	expbldg(f, sftype, id, fid, sub)
	// 頂点座標データ（閉じた図形）
	yanepoly = append(yanepoly, xo1, yo1, mbt)
	yanepoly = append(yanepoly, xtp4, ytp4, ntp)
	yanepoly = append(yanepoly, xo4, yo4, nbt)
	yanepoly = append(yanepoly, xo1, yo1, mbt)
	// bldg:boundedBy の書き出しを終了する
	exitbldg(f, yanepoly, sftype)
	// 屋根データの頂点座標の初期化
	yanepoly = yanepoly[:0]

	// ケラバ２・三角形2-1のIDと頂点座標の書き出し
	sub = sub + 1
	sftype = "Roof"
	// bldg:boundedBy をファイルに書き出す
	expbldg(f, sftype, id, fid, sub)
	// 頂点座標データ（閉じた図形）
	yanepoly = append(yanepoly, xtp3, ytp3, ntp)
	yanepoly = append(yanepoly, xtp2, ytp2, mtp)
	yanepoly = append(yanepoly, xo2, yo2, mbt)
	yanepoly = append(yanepoly, xtp3, ytp3, ntp)
	// bldg:boundedBy の書き出しを終了する
	exitbldg(f, yanepoly, sftype)
	// 屋根データの頂点座標の初期化
	yanepoly = yanepoly[:0]

	// ケラバ２・三角形2-2のIDと頂点座標の書き出し
	sub = sub + 1
	sftype = "Roof"
	// bldg:boundedBy をファイルに書き出す
	expbldg(f, sftype, id, fid, sub)
	// 頂点座標データ（閉じた図形）
	yanepoly = append(yanepoly, xo2, yo2, mbt)
	yanepoly = append(yanepoly, xo3, yo3, nbt)
	yanepoly = append(yanepoly, xtp3, ytp3, ntp)
	yanepoly = append(yanepoly, xo2, yo2, mbt)
	// bldg:boundedBy の書き出しを終了する
	exitbldg(f, yanepoly, sftype)
	// 屋根データの頂点座標の初期化
	yanepoly = yanepoly[:0]

	// 側面三角壁の下端高さ
	sh0 := toph
	log.Println("sh0=", sh0)

	// (xo1, yo1, mbt),(xo2, yo2, mbt),(xo4, yo4, nbt)を通る平面の式
	p1 := []float64{xo1, yo1, mbt}
	p2 := []float64{xo2, yo2, mbt}
	p3 := []float64{xo4, yo4, nbt}

	// 法線ベクトル（外積）を求める
	vec := pkg.NorVec(p1, p2, p3)
	log.Println("vec=", vec)

	// 平面上の点の座標(xo1, yo1, mbt)
	// 平面の式 a(x-xo1)+b(y-yo1)+c(z-mbt)=0
	// vec[0] * (x1 - xo1) + vec[1] * (y1 - yo1) + vec[2] * (sh1 - mbt) = 0
	// vec[2] * (sh1 - mbt) = -vec[0] * (x1 - xo1) - vec[1] * (y1 - yo1)
	// (sh1 - mbt) = (-vec[0] * (x1 - xo1) - vec[1] * (y1 - yo1)) / vec[2]
	sh1 := (-vec[0]*(x1-xo1)-vec[1]*(y1-yo1))/vec[2] + mbt
	log.Println("sh1=", sh1)

	// 側面・三角形２のIDと頂点座標の書き出し
	sub = sub + 1
	sftype = "Roof"
	// bldg:boundedBy をファイルに書き出す
	expbldg(f, sftype, id, fid, sub)
	// 頂点座標データ（閉じた図形）
	yanepoly = append(yanepoly, x1, y1, sh0)
	yanepoly = append(yanepoly, x1, y1, sh1)
	yanepoly = append(yanepoly, x4, y4, sh0)
	yanepoly = append(yanepoly, x1, y1, sh0)
	// bldg:boundedBy の書き出しを終了する
	exitbldg(f, yanepoly, sftype)
	// 屋根データの頂点座標の初期化
	yanepoly = yanepoly[:0]

	sh2 := (-vec[0]*(x2-xo1)-vec[1]*(y2-yo1))/vec[2] + mbt
	log.Println("sh2=", sh2)

	// 側面・三角形１のIDと頂点座標の書き出し
	sub = sub + 1
	sftype = "Roof"
	// bldg:boundedBy をファイルに書き出す
	expbldg(f, sftype, id, fid, sub)
	// 頂点座標データ（閉じた図形）
	yanepoly = append(yanepoly, x2, y2, sh2)
	yanepoly = append(yanepoly, x2, y2, sh0)
	yanepoly = append(yanepoly, x3, y3, sh0)
	yanepoly = append(yanepoly, x2, y2, sh2)
	// bldg:boundedBy の書き出しを終了する
	exitbldg(f, yanepoly, sftype)
	// 屋根データの頂点座標の初期化
	yanepoly = yanepoly[:0]

	// 背面・三角形-1のIDと頂点座標の書き出し
	sub = sub + 1
	sftype = "Roof"
	// bldg:boundedBy をファイルに書き出す
	expbldg(f, sftype, id, fid, sub)
	// 頂点座標データ（閉じた図形）
	yanepoly = append(yanepoly, x1, y1, sh1)
	yanepoly = append(yanepoly, x1, y1, sh0)
	yanepoly = append(yanepoly, x2, y2, sh0)
	yanepoly = append(yanepoly, x1, y1, sh1)
	// bldg:boundedBy の書き出しを終了する
	exitbldg(f, yanepoly, sftype)
	// 屋根データの頂点座標の初期化
	yanepoly = yanepoly[:0]

	// 背面・三角形-2のIDと頂点座標の書き出し
	sub = sub + 1
	sftype = "Roof"
	// bldg:boundedBy をファイルに書き出す
	expbldg(f, sftype, id, fid, sub)
	// 頂点座標データ（閉じた図形）
	yanepoly = append(yanepoly, x1, y1, sh1)
	yanepoly = append(yanepoly, x2, y2, sh0)
	yanepoly = append(yanepoly, x2, y2, sh2)
	yanepoly = append(yanepoly, x1, y1, sh1)
	// bldg:boundedBy の書き出しを終了する
	exitbldg(f, yanepoly, sftype)

	// 傾斜屋根建物の壁面・床面座標を出力する
	bldbody(f, id, fid, vcnt, list, elv, btm, toph, story, sub)
}
