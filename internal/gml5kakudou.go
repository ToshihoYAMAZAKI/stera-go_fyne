package internal

import (
	"log"
	"math"
	"os"
	"stera/pkg"
)

// ExposKiri は切妻屋根建物のは頂点座標をファイルに書き出す
func Expos5kaku(f *os.File, id, fid string, vcnt int, list [][]float64, elv, btm, toph, hisashi, incline, yaneatu float64, story int) {
	log.Println("五角堂屋根")
	// 屋根モデル（五角堂屋根）
	// 5つの頂点座標の定義
	x1 := list[0][0]
	y1 := list[0][1]
	x2 := list[1][0]
	y2 := list[1][1]
	x3 := list[2][0]
	y3 := list[2][1]
	x4 := list[3][0]
	y4 := list[3][1]
	x5 := list[4][0]
	y5 := list[4][1]

	// 軒鼻の突き出し長さ
	nt := yaneatu / math.Sqrt(math.Pow(incline, 2)+1) * incline
	// 平面から軒庇の上端までの長さ（軒庇＋軒鼻）
	hichtop := hisashi + nt

	// 屋根伏せの5頂点の座標
	var xo1, yo1 float64
	var xo2, yo2 float64
	var xo3, yo3 float64
	var xo4, yo4 float64
	var xo5, yo5 float64

	var xtp1, ytp1 float64
	var xtp2, ytp2 float64
	var xtp3, ytp3 float64
	var xtp4, ytp4 float64
	var xtp5, ytp5 float64

	// 屋根モデルの頂点座標をリストに書き出す
	// 屋根頂点の法線ベクトルを算出してリストに書き出す
	var yanepoly []float64
	// PolyGMLID をfid から設定するための添え字の定義
	sub := 0
	// SurfaceType の定義
	var sftype string

	// 五角堂の棟の中心座標
	xo := (x1 + x2 + x3 + x4 + x5) / 5
	yo := (y1 + y2 + y3 + y4 + y5) / 5
	// 頂点1と頂点2を結ぶ線分に平行な直線の式
	m1, n1 := pkg.ParaLine(x1, y1, x2, y2, xo, yo, hisashi)
	// 頂点2と頂点3を結ぶ線分に平行な直線の式
	m2, n2 := pkg.ParaLine(x2, y2, x3, y3, xo, yo, hisashi)
	// 頂点3と頂点4を結ぶ線分に平行な直線の式
	m3, n3 := pkg.ParaLine(x3, y3, x4, y4, xo, yo, hisashi)
	// 頂点4と頂点5を結ぶ線分に平行な直線の式
	m4, n4 := pkg.ParaLine(x4, y4, x5, y5, xo, yo, hisashi)
	// 頂点5と頂点1を結ぶ線分に平行な直線の式
	m5, n5 := pkg.ParaLine(x5, y5, x1, y1, xo, yo, hisashi)
	// 屋根伏せの5頂点の座標を求める（軒庇の下端）
	xo1, yo1 = pkg.SeekInsec(m5, n5, m1, n1)
	xo2, yo2 = pkg.SeekInsec(m1, n1, m2, n2)
	xo3, yo3 = pkg.SeekInsec(m2, n2, m3, n3)
	xo4, yo4 = pkg.SeekInsec(m3, n3, m4, n4)
	xo5, yo5 = pkg.SeekInsec(m1, n1, m5, n5)
	// 頂点1と頂点2を結ぶ線分に平行な直線の式 // 軒庇の上端
	mtp1, ntp1 := pkg.ParaLine(x1, y1, x2, y2, xo, yo, hichtop)
	// 頂点2と頂点3を結ぶ線分に平行な直線の式 // 軒庇の上端
	mtp2, ntp2 := pkg.ParaLine(x2, y2, x3, y3, xo, yo, hichtop)
	// 頂点3と頂点4を結ぶ線分に平行な直線の式 // 軒庇の上端
	mtp3, ntp3 := pkg.ParaLine(x3, y3, x4, y4, xo, yo, hichtop)
	// 頂点4と頂点5を結ぶ線分に平行な直線の式 // 軒庇の上端
	mtp4, ntp4 := pkg.ParaLine(x4, y4, x5, y5, xo, yo, hichtop)
	// 頂点5と頂点1を結ぶ線分に平行な直線の式（平面）// 軒庇の上端
	mtp5, ntp5 := pkg.ParaLine(x5, y5, x1, y1, xo, yo, hichtop)
	// 屋根伏せの5頂点の座標を求める（軒庇の上端）
	xtp1, ytp1 = pkg.SeekInsec(mtp5, ntp5, mtp1, ntp1)
	xtp2, ytp2 = pkg.SeekInsec(mtp1, ntp1, mtp2, ntp2)
	xtp3, ytp3 = pkg.SeekInsec(mtp2, ntp2, mtp3, ntp3)
	xtp4, ytp4 = pkg.SeekInsec(mtp3, ntp3, mtp4, ntp4)
	xtp5, ytp5 = pkg.SeekInsec(mtp4, ntp4, mtp5, ntp5)

	// 軒先下端高さ（庇×屋根勾配）
	nbt := (toph - hisashi*incline)
	// 軒先上端高さ
	ntp := nbt + yaneatu/math.Sqrt(math.Pow(incline, 2)+1)
	log.Println("nbt, ntp", nbt, ntp)
	// 五角堂の頂点の高さを求める
	// 5頂点から五角堂の頂点までの長さを比較する
	tmo := 0.0
	tm1 := pkg.DistVerts(xo, yo, xo1, yo1)
	if tmo < tm1 {
		tmo = tm1
	}
	tm2 := pkg.DistVerts(xo, yo, xo2, yo2)
	if tmo < tm2 {
		tmo = tm2
	}
	tm3 := pkg.DistVerts(xo, yo, xo3, yo3)
	if tmo < tm3 {
		tmo = tm3
	}
	tm4 := pkg.DistVerts(xo, yo, xo4, yo4)
	if tmo < tm4 {
		tmo = tm4
	}
	tm5 := pkg.DistVerts(xo, yo, xo5, yo5)
	if tmo < tm5 {
		tmo = tm5
	}
	// 五角堂の頂点の下端高さ
	mbt := nbt + tmo*incline
	// 五角堂の頂点の上端高さ
	mtp := mbt + yaneatu*math.Sqrt(1+math.Pow(incline, 2))
	log.Println("mbt, mtp", mbt, mtp)

	// 屋根底面・三角形１のIDと頂点座標の書き出し
	sub = sub + 1
	sftype = "Roof"
	// bldg:boundedBy をファイルに書き出す
	expbldg(f, sftype, id, fid, sub)
	// 頂点座標データ（閉じた図形）
	yanepoly = append(yanepoly, xo1, yo1, nbt)
	yanepoly = append(yanepoly, xo, yo, mbt)
	yanepoly = append(yanepoly, xo2, yo2, nbt)
	yanepoly = append(yanepoly, xo1, yo1, nbt)
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
	yanepoly = append(yanepoly, xo2, yo2, nbt)
	yanepoly = append(yanepoly, xo, yo, mbt)
	yanepoly = append(yanepoly, xo3, yo3, nbt)
	yanepoly = append(yanepoly, xo2, yo2, nbt)
	// bldg:boundedBy の書き出しを終了する
	exitbldg(f, yanepoly, sftype)
	// 屋根データの頂点座標の初期化
	yanepoly = yanepoly[:0]

	// 屋根底面・三角形３のIDと頂点座標の書き出し
	sub = sub + 1
	sftype = "Roof"
	// bldg:boundedBy をファイルに書き出す
	expbldg(f, sftype, id, fid, sub)
	// 頂点座標データ（閉じた図形）
	yanepoly = append(yanepoly, xo3, yo3, nbt)
	yanepoly = append(yanepoly, xo, yo, mbt)
	yanepoly = append(yanepoly, xo4, yo4, nbt)
	yanepoly = append(yanepoly, xo3, yo3, nbt)
	// bldg:boundedBy の書き出しを終了する
	exitbldg(f, yanepoly, sftype)
	// 屋根データの頂点座標の初期化
	yanepoly = yanepoly[:0]

	// 屋根底面・三角形４のIDと頂点座標の書き出し
	sub = sub + 1
	sftype = "Roof"
	// bldg:boundedBy をファイルに書き出す
	expbldg(f, sftype, id, fid, sub)
	// 頂点座標データ（閉じた図形）
	yanepoly = append(yanepoly, xo4, yo4, nbt)
	yanepoly = append(yanepoly, xo, yo, mbt)
	yanepoly = append(yanepoly, xo5, yo5, nbt)
	yanepoly = append(yanepoly, xo4, yo4, nbt)
	// bldg:boundedBy の書き出しを終了する
	exitbldg(f, yanepoly, sftype)
	// 屋根データの頂点座標の初期化
	yanepoly = yanepoly[:0]

	// 屋根底面・三角形５のIDと頂点座標の書き出し
	sub = sub + 1
	sftype = "Roof"
	// bldg:boundedBy をファイルに書き出す
	expbldg(f, sftype, id, fid, sub)
	// 頂点座標データ（閉じた図形）
	yanepoly = append(yanepoly, xo5, yo5, nbt)
	yanepoly = append(yanepoly, xo, yo, mbt)
	yanepoly = append(yanepoly, xo1, yo1, nbt)
	yanepoly = append(yanepoly, xo5, yo5, nbt)
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
	yanepoly = append(yanepoly, xtp1, ytp1, ntp)
	yanepoly = append(yanepoly, xo, yo, mtp)
	yanepoly = append(yanepoly, xtp2, ytp2, ntp)
	yanepoly = append(yanepoly, xtp1, ytp1, ntp)
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
	yanepoly = append(yanepoly, xtp2, ytp2, ntp)
	yanepoly = append(yanepoly, xo, yo, mtp)
	yanepoly = append(yanepoly, xtp3, ytp3, ntp)
	yanepoly = append(yanepoly, xtp2, ytp2, ntp)
	// bldg:boundedBy の書き出しを終了する
	exitbldg(f, yanepoly, sftype)
	// 屋根データの頂点座標の初期化
	yanepoly = yanepoly[:0]

	// 屋根上面・三角形３のIDと頂点座標の書き出し
	sub = sub + 1
	sftype = "Roof"
	// bldg:boundedBy をファイルに書き出す
	expbldg(f, sftype, id, fid, sub)
	// 頂点座標データ（閉じた図形）
	yanepoly = append(yanepoly, xtp3, ytp3, ntp)
	yanepoly = append(yanepoly, xo, yo, mtp)
	yanepoly = append(yanepoly, xtp4, ytp4, ntp)
	yanepoly = append(yanepoly, xtp3, ytp3, ntp)
	// bldg:boundedBy の書き出しを終了する
	exitbldg(f, yanepoly, sftype)
	// 屋根データの頂点座標の初期化
	yanepoly = yanepoly[:0]

	// 屋根上面・三角形４のIDと頂点座標の書き出し
	sub = sub + 1
	sftype = "Roof"
	// bldg:boundedBy をファイルに書き出す
	expbldg(f, sftype, id, fid, sub)
	// 頂点座標データ（閉じた図形）
	yanepoly = append(yanepoly, xtp4, ytp4, ntp)
	yanepoly = append(yanepoly, xo, yo, mtp)
	yanepoly = append(yanepoly, xtp5, ytp5, ntp)
	yanepoly = append(yanepoly, xtp4, ytp4, ntp)
	// bldg:boundedBy の書き出しを終了する
	exitbldg(f, yanepoly, sftype)
	// 屋根データの頂点座標の初期化
	yanepoly = yanepoly[:0]

	// 屋根上面・三角形５のIDと頂点座標の書き出し
	sub = sub + 1
	sftype = "Roof"
	// bldg:boundedBy をファイルに書き出す
	expbldg(f, sftype, id, fid, sub)
	// 頂点座標データ（閉じた図形）
	yanepoly = append(yanepoly, xtp5, ytp5, ntp)
	yanepoly = append(yanepoly, xo, yo, mtp)
	yanepoly = append(yanepoly, xtp1, ytp1, ntp)
	yanepoly = append(yanepoly, xtp5, ytp5, ntp)
	// bldg:boundedBy の書き出しを終了する
	exitbldg(f, yanepoly, sftype)
	// 屋根データの頂点座標の初期化
	yanepoly = yanepoly[:0]

	// 軒端１・三角形1-1のIDと頂点座標の書き出し
	sub = sub + 1
	sftype = "Roof"
	// bldg:boundedBy をファイルに書き出す
	expbldg(f, sftype, id, fid, sub)
	// 頂点座標データ（閉じた図形）
	yanepoly = append(yanepoly, xtp1, ytp1, ntp)
	yanepoly = append(yanepoly, xo2, yo2, nbt)
	yanepoly = append(yanepoly, xo1, yo1, nbt)
	yanepoly = append(yanepoly, xtp1, ytp1, ntp)
	// bldg:boundedBy の書き出しを終了する
	exitbldg(f, yanepoly, sftype)
	// 屋根データの頂点座標の初期化
	yanepoly = yanepoly[:0]

	// 軒端１・三角形1-2のIDと頂点座標の書き出し
	sub = sub + 1
	sftype = "Roof"
	// bldg:boundedBy をファイルに書き出す
	expbldg(f, sftype, id, fid, sub)
	// 頂点座標データ（閉じた図形）
	yanepoly = append(yanepoly, xtp2, ytp2, ntp)
	yanepoly = append(yanepoly, xo2, yo2, nbt)
	yanepoly = append(yanepoly, xtp1, ytp1, ntp)
	yanepoly = append(yanepoly, xtp2, ytp2, ntp)
	// bldg:boundedBy の書き出しを終了する
	exitbldg(f, yanepoly, sftype)
	// 屋根データの頂点座標の初期化
	yanepoly = yanepoly[:0]

	// 軒端２・三角形2-1のIDと頂点座標の書き出し
	sub = sub + 1
	sftype = "Roof"
	// bldg:boundedBy をファイルに書き出す
	expbldg(f, sftype, id, fid, sub)
	// 頂点座標データ（閉じた図形）
	yanepoly = append(yanepoly, xtp2, ytp2, ntp)
	yanepoly = append(yanepoly, xo3, yo3, nbt)
	yanepoly = append(yanepoly, xo2, yo2, nbt)
	yanepoly = append(yanepoly, xtp2, ytp2, ntp)
	// bldg:boundedBy の書き出しを終了する
	exitbldg(f, yanepoly, sftype)
	// 屋根データの頂点座標の初期化
	yanepoly = yanepoly[:0]

	// 軒端２・三角形2-2のIDと頂点座標の書き出し
	sub = sub + 1
	sftype = "Roof"
	// bldg:boundedBy をファイルに書き出す
	expbldg(f, sftype, id, fid, sub)
	// 頂点座標データ（閉じた図形）
	yanepoly = append(yanepoly, xtp3, ytp3, ntp)
	yanepoly = append(yanepoly, xo3, yo3, nbt)
	yanepoly = append(yanepoly, xtp2, ytp2, ntp)
	yanepoly = append(yanepoly, xtp3, ytp3, ntp)
	// bldg:boundedBy の書き出しを終了する
	exitbldg(f, yanepoly, sftype)
	// 屋根データの頂点座標の初期化
	yanepoly = yanepoly[:0]

	// 軒端３・三角形3-1のIDと頂点座標の書き出し
	sub = sub + 1
	sftype = "Roof"
	// bldg:boundedBy をファイルに書き出す
	expbldg(f, sftype, id, fid, sub)
	// 頂点座標データ（閉じた図形）
	yanepoly = append(yanepoly, xtp3, ytp3, ntp)
	yanepoly = append(yanepoly, xo4, yo4, nbt)
	yanepoly = append(yanepoly, xo3, yo3, nbt)
	yanepoly = append(yanepoly, xtp3, ytp3, ntp)
	// bldg:boundedBy の書き出しを終了する
	exitbldg(f, yanepoly, sftype)
	// 屋根データの頂点座標の初期化
	yanepoly = yanepoly[:0]

	// 軒端３・三角形3-2のIDと頂点座標の書き出し
	sub = sub + 1
	sftype = "Roof"
	// bldg:boundedBy をファイルに書き出す
	expbldg(f, sftype, id, fid, sub)
	// 頂点座標データ（閉じた図形）
	yanepoly = append(yanepoly, xtp4, ytp4, ntp)
	yanepoly = append(yanepoly, xo4, yo4, nbt)
	yanepoly = append(yanepoly, xtp3, ytp3, ntp)
	// bldg:boundedBy の書き出しを終了する
	exitbldg(f, yanepoly, sftype)
	// 屋根データの頂点座標の初期化
	yanepoly = yanepoly[:0]

	// 軒端４・三角形4-1のIDと頂点座標の書き出し
	sub = sub + 1
	sftype = "Roof"
	// bldg:boundedBy をファイルに書き出す
	expbldg(f, sftype, id, fid, sub)
	// 頂点座標データ（閉じた図形）
	yanepoly = append(yanepoly, xtp4, ytp4, ntp)
	yanepoly = append(yanepoly, xo5, yo5, nbt)
	yanepoly = append(yanepoly, xo4, yo4, nbt)
	yanepoly = append(yanepoly, xtp4, ytp4, ntp)
	// bldg:boundedBy の書き出しを終了する
	exitbldg(f, yanepoly, sftype)
	// 屋根データの頂点座標の初期化
	yanepoly = yanepoly[:0]

	// 軒端４・三角形4-2のIDと頂点座標の書き出し
	sub = sub + 1
	sftype = "Roof"
	// bldg:boundedBy をファイルに書き出す
	expbldg(f, sftype, id, fid, sub)
	// 頂点座標データ（閉じた図形）
	yanepoly = append(yanepoly, xtp5, ytp5, ntp)
	yanepoly = append(yanepoly, xo5, yo5, nbt)
	yanepoly = append(yanepoly, xtp4, ytp4, ntp)
	yanepoly = append(yanepoly, xtp5, ytp5, ntp)
	// bldg:boundedBy の書き出しを終了する
	exitbldg(f, yanepoly, sftype)
	// 屋根データの頂点座標の初期化
	yanepoly = yanepoly[:0]

	// 軒端５・三角形5-1のIDと頂点座標の書き出し
	sub = sub + 1
	sftype = "Roof"
	// bldg:boundedBy をファイルに書き出す
	expbldg(f, sftype, id, fid, sub)
	// 頂点座標データ（閉じた図形）
	yanepoly = append(yanepoly, xtp5, ytp5, ntp)
	yanepoly = append(yanepoly, xo1, yo1, nbt)
	yanepoly = append(yanepoly, xo5, yo5, nbt)
	// bldg:boundedBy の書き出しを終了する
	exitbldg(f, yanepoly, sftype)
	// 屋根データの頂点座標の初期化
	yanepoly = yanepoly[:0]

	// 軒端５・三角形5-2のIDと頂点座標の書き出し
	sub = sub + 1
	sftype = "Roof"
	// bldg:boundedBy をファイルに書き出す
	expbldg(f, sftype, id, fid, sub)
	// 頂点座標データ（閉じた図形）
	yanepoly = append(yanepoly, xtp1, ytp1, ntp)
	yanepoly = append(yanepoly, xo1, yo1, nbt)
	yanepoly = append(yanepoly, xtp5, ytp5, ntp)
	yanepoly = append(yanepoly, xtp1, ytp1, ntp)
	// bldg:boundedBy の書き出しを終了する
	exitbldg(f, yanepoly, sftype)

	// 傾斜屋根建物の壁面・床面座標を出力する
	bldbody(f, id, fid, vcnt, list, elv, btm, toph, story, sub)
}
