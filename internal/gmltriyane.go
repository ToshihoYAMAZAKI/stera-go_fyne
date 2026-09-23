package internal

import (
	"log"
	"math"
	"os"
	"stera/pkg"
)

// ExposTri は３角形の片流れ屋根の頂点座標をファイルに書き出す
func ExposTri(f *os.File, id, fid string, vcnt int, list [][]float64, elv, btm, toph, hisashi, keraba, incline, yaneatu float64, story int) (tripts map[string]float64) {
	log.Println("３角形片流れ屋根")
	// 屋根モデル（３角形片流れ屋根）
	// 3つの頂点座標の定義
	x1 := list[0][0]
	y1 := list[0][1]
	x2 := list[1][0]
	y2 := list[1][1]
	x3 := list[2][0]
	y3 := list[2][1]

	// 頂点2と頂点3を結ぶ線分に平行な直線の式（流れ面下手）
	m1, n1 := pkg.ParaLine(x2, y2, x3, y3, x1, y1, hisashi)
	// 頂点1と頂点2を結ぶ線分に平行な直線の式（側面）
	m2, n2 := pkg.ParaLine(x1, y1, x2, y2, x3, y3, keraba)
	// 頂点3と頂点1を結ぶ線分に平行な直線の式（側面）
	m3, n3 := pkg.ParaLine(x3, y3, x1, y1, x2, y2, keraba)

	// 屋根伏せの3頂点の座標を求める（屋根面の下端・上端）
	xo1, yo1 := pkg.SeekInsec(m2, n2, m3, n3)
	xo2, yo2 := pkg.SeekInsec(m1, n1, m2, n2)
	xo3, yo3 := pkg.SeekInsec(m3, n3, m1, n1)
	log.Println("xo1, yo1, xo2, yo2, xo3, yo3", xo1, yo1, xo2, yo2, xo3, yo3)

	// 軒鼻の突き出し長さ
	nt := yaneatu / math.Sqrt(math.Pow(incline, 2)+1) * incline
	// 壁面から軒庇の上端までの長さ（軒庇＋軒鼻）
	hichtop := hisashi + nt
	// 頂点2と頂点3を結ぶ線分に平行な直線の式（流れ面下手上端）
	mkt1, nkt1 := pkg.ParaLine(x2, y2, x3, y3, x1, y1, hichtop)

	// 屋根伏せの3頂点の座標を求める（屋根面の上端）
	// xtp1, ytp1 := pkg.SeekInsec(m2, n2, m3, n3)
	xtp2, ytp2 := pkg.SeekInsec(mkt1, nkt1, m2, n2)
	xtp3, ytp3 := pkg.SeekInsec(m3, n3, mkt1, nkt1)
	// log.Println("xtp1, ytp1, xtp2, ytp2, xtp3, ytp3", xtp1, ytp1, xtp2, ytp2, xtp3, ytp3)

	// 屋根の側面の長さを求める
	sf1 := pkg.DistVerts(x1, y1, xo2, yo2)
	sf2 := pkg.DistVerts(xo3, yo3, x1, y1)
	log.Println("sf1, sf2", sf1, sf2)

	// 軒先下端高さ（庇×屋根勾配）
	nbt := (toph - hisashi/math.Sqrt(math.Pow(incline, 2)+1)*incline)
	// 軒先上端高さ
	ntp := nbt + yaneatu/math.Sqrt(math.Pow(incline, 2)+1)
	log.Println("nbt, ntp", nbt, ntp)

	// 屋根の流れ面上手の下端高さ
	mbt := nbt + ((sf1+sf2)/2)*incline

	// 屋根の流れ面上手の上端高さ
	mtp := mbt + yaneatu*math.Sqrt(1+math.Pow(incline, 2))
	log.Println("mbt, mtp", mbt, mtp)

	// 屋根モデルの頂点座標をリストに書き出す
	// 屋根頂点の法線ベクトルを算出してリストに書き出す
	var yanepoly []float64
	// PolyGMLID をfid から設定するための添え字の定義
	sub := 0
	// SurfaceType の定義
	var sftype string

	// 屋根底面・三角形のIDと頂点座標の書き出し
	sub = sub + 1
	sftype = "Roof"
	// bldg:boundedBy をファイルに書き出す
	expbldg(f, sftype, id, fid, sub)
	// 頂点座標データ（閉じた図形）
	yanepoly = append(yanepoly, xo1, yo1, mbt)
	yanepoly = append(yanepoly, xo3, yo3, nbt)
	yanepoly = append(yanepoly, xo2, yo2, nbt)
	yanepoly = append(yanepoly, xo1, yo1, mbt)
	// bldg:boundedBy の書き出しを終了する
	exitbldg(f, yanepoly, sftype)
	// 屋根データの頂点座標の初期化
	yanepoly = yanepoly[:0]

	// 屋根上面・三角形のIDと頂点座標の書き出し
	sub = sub + 1
	sftype = "Roof"
	// bldg:boundedBy をファイルに書き出す
	expbldg(f, sftype, id, fid, sub)
	// 頂点座標データ（閉じた図形）
	yanepoly = append(yanepoly, xo1, yo1, mtp)
	yanepoly = append(yanepoly, xtp3, ytp3, ntp)
	yanepoly = append(yanepoly, xtp2, ytp2, ntp)
	yanepoly = append(yanepoly, xo1, yo1, mtp)
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
	yanepoly = append(yanepoly, xtp3, ytp3, ntp)
	yanepoly = append(yanepoly, xtp2, ytp2, ntp)
	yanepoly = append(yanepoly, xo3, yo3, nbt)
	yanepoly = append(yanepoly, xtp3, ytp3, ntp)
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
	yanepoly = append(yanepoly, xo3, yo3, nbt)
	yanepoly = append(yanepoly, xtp2, ytp2, ntp)
	yanepoly = append(yanepoly, xo2, yo2, nbt)
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
	yanepoly = append(yanepoly, xo2, yo2, ntp)
	yanepoly = append(yanepoly, xo1, yo1, mtp)
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
	yanepoly = append(yanepoly, xo2, yo2, nbt)
	yanepoly = append(yanepoly, xo2, yo2, ntp)
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
	yanepoly = append(yanepoly, xo3, yo3, ntp)
	yanepoly = append(yanepoly, xo1, yo1, mbt)
	yanepoly = append(yanepoly, xo1, yo1, mtp)
	yanepoly = append(yanepoly, xo3, yo3, ntp)
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
	yanepoly = append(yanepoly, xo3, yo3, nbt)
	yanepoly = append(yanepoly, xo1, yo1, mbt)
	yanepoly = append(yanepoly, xo3, yo3, ntp)
	yanepoly = append(yanepoly, xo3, yo3, nbt)
	// bldg:boundedBy の書き出しを終了する
	exitbldg(f, yanepoly, sftype)
	// 屋根データの頂点座標の初期化
	yanepoly = yanepoly[:0]

	// 側面三角壁の下端高さ
	sh0 := toph
	log.Println("sh0=", sh0)

	// (xo1, yo1, mbt),(xo2, yo2, nbt),(xo3, yo3, nbt)を通る平面の式
	p1 := []float64{xo1, yo1, mbt}
	p2 := []float64{xo2, yo2, nbt}
	p3 := []float64{xo3, yo3, nbt}

	// 法線ベクトル（外積）を求める
	vec := pkg.NorVec(p1, p2, p3)

	// 平面上の点の座標(xo1, yo1, mbt)
	// 平面の式 a(x-xo1)+b(y-yo1)+c(z-mbt)=0
	// vec[0] * (xo2 - xo1) + vec[1] * (yo2 - yo1) + vec[2] * (sh1 - mbt) = 0
	// vec[2] * (sh1 - mbt) = -vec[0] * (xo2 - xo1) - vec[1] * (yo2 - yo1)
	// (sh1 - mbt) = (-vec[0] * (xo2 - xo1) - vec[1] * (yo2 - yo1)) / vec[2]
	sh1 := (-vec[0]*(xo2-xo1)-vec[1]*(yo2-yo1))/vec[2] + mbt
	log.Println("sh1=", sh1)

	// 側面・三角形１のIDと頂点座標の書き出し
	sub = sub + 1
	sftype = "Roof"
	// bldg:boundedBy をファイルに書き出す
	expbldg(f, sftype, id, fid, sub)
	// 頂点座標データ（閉じた図形）
	yanepoly = append(yanepoly, x1, y1, sh0)
	yanepoly = append(yanepoly, x2, y2, sh0)
	yanepoly = append(yanepoly, x1, y1, sh1)
	yanepoly = append(yanepoly, x1, y1, sh0)
	// bldg:boundedBy の書き出しを終了する
	exitbldg(f, yanepoly, sftype)
	// 屋根データの頂点座標の初期化
	yanepoly = yanepoly[:0]

	sh2 := (-vec[0]*(xo3-xo1)-vec[1]*(yo3-yo1))/vec[2] + mbt
	log.Println("sh2=", sh2)

	// 側面・三角形２のIDと頂点座標の書き出し
	sub = sub + 1
	sftype = "Roof"
	// bldg:boundedBy をファイルに書き出す
	expbldg(f, sftype, id, fid, sub)
	// 頂点座標データ（閉じた図形）
	yanepoly = append(yanepoly, x3, y3, sh0)
	yanepoly = append(yanepoly, x1, y1, sh0)
	yanepoly = append(yanepoly, x1, y1, sh2)
	yanepoly = append(yanepoly, x3, y3, sh0)
	// bldg:boundedBy の書き出しを終了する
	exitbldg(f, yanepoly, sftype)

	// 傾斜屋根建物の壁面・床面座標を出力する
	bldbody(f, id, fid, vcnt, list, elv, btm, toph, story, sub)

	triverts := make(map[string]float64)
	triverts["xo1"] = xo1
	triverts["yo1"] = yo1
	triverts["mbt"] = mbt
	triverts["mtp"] = mtp
	triverts["xo2"] = xo2
	triverts["yo2"] = yo2
	triverts["xtp2"] = xtp2
	triverts["ytp2"] = ytp2
	triverts["xo3"] = xo3
	triverts["yo3"] = yo3
	triverts["xtp3"] = xtp3
	triverts["ytp3"] = ytp3
	triverts["nbt"] = nbt
	triverts["ntp"] = ntp
	tripts = triverts

	return tripts
}
