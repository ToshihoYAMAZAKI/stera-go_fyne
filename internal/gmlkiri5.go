package internal

import (
	"log"
	"math"
	"os"
	"stera/pkg"
)

// ExposKiri は切妻屋根建物のは頂点座標をファイルに書き出す
func ExposKiri5(f *os.File, id, fid string, vcnt int, list [][]float64, elv, btm, toph, hisashi, keraba, incline, yaneatu float64, story int) {
	log.Println("５角形切妻屋根")
	// 屋根モデル（５角形切妻屋根）
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

	// 屋根伏せの4頂点の座標
	var xo1, yo1 float64
	var xo2, yo2 float64
	var xo3, yo3 float64
	var xo4, yo4 float64

	var xtp1, ytp1 float64
	var xtp2, ytp2 float64
	var xtp3, ytp3 float64
	var xtp4, ytp4 float64

	// ５角形屋根の頂点の直角と広角の配置から棟方向を決める
	// 頂点の角度を調べる

	// var deg []float64
	// for i := 1; i < 4; i++ {
	// 	d := pkg.TriAngle(list[i-1][0], list[i-1][1], list[i][0], list[i][1], list[i+1][0], list[i+1][1])
	// 	deg = append(deg, d)
	// }

	// 屋根モデルの頂点座標をリストに書き出す
	// 屋根頂点の法線ベクトルを算出してリストに書き出す
	var yanepoly []float64
	// PolyGMLID をfid から設定するための添え字の定義
	sub := 0
	// SurfaceType の定義
	var sftype string

	d1 := pkg.DistVerts(x1, y1, x2, y2)
	d2 := pkg.DistVerts(x5, y5, x1, y1)
	log.Println("d1, d2", d1, d2)
	if d1 < d2 {
		// 頂点1と頂点2を結ぶ線分に平行な直線の式（妻面）
		m1, n1 := pkg.ParaLine(x1, y1, x2, y2, x5, y5, keraba)
		// 頂点5と頂点1を結ぶ線分に平行な直線の式（平面）// 軒庇の下端
		m4, n4 := pkg.ParaLine(x5, y5, x1, y1, x2, y2, hisashi)
		// 頂点2と頂点3を結ぶ線分に平行な直線の式（平面）
		mh, nh := pkg.ParaLine(x2, y2, x3, y3, x5, y5, hisashi)
		// 頂点4と頂点5を結ぶ線分に平行な直線の式（妻面）
		mk, nk := pkg.ParaLine(x4, y4, x5, y5, x2, y2, keraba)
		// 頂点3と頂点4の合成頂点を求める
		xo3, yo3 = pkg.SeekInsec(mh, nh, mk, nk)
		// 屋根伏せの4頂点の座標を求める（軒庇の下端）
		xo1, yo1 = pkg.SeekInsec(m4, n4, m1, n1)
		xo2, yo2 = pkg.SeekInsec(m1, n1, mh, nh)
		xo4, yo4 = pkg.SeekInsec(mk, nk, m4, n4)
		// 頂点2と頂点3を結ぶ線分に平行な直線の式（平面）// 軒庇の上端
		mtp3, ntp3 := pkg.ParaLine(x2, y2, x3, y3, x1, y1, hichtop)
		// 頂点5と頂点1を結ぶ線分に平行な直線の式（平面）// 軒庇の上端
		mtp4, ntp4 := pkg.ParaLine(x5, y5, x1, y1, x2, y2, hichtop)
		// 屋根伏せの4頂点の座標を求める（軒庇の上端）
		xtp1, ytp1 = pkg.SeekInsec(mtp4, ntp4, m1, n1)
		xtp2, ytp2 = pkg.SeekInsec(m1, n1, mtp3, ntp3)
		xtp3, ytp3 = pkg.SeekInsec(mtp3, ntp3, mk, nk)
		xtp4, ytp4 = pkg.SeekInsec(mk, nk, mtp4, ntp4)
	} else if d1 > d2 {
		// 頂点5と頂点1を結ぶ線分に平行な直線の式（妻面）
		m4, n4 := pkg.ParaLine(x5, y5, x1, y1, x2, y2, keraba)
		// 頂点1と頂点2を結ぶ線分に平行な直線の式（平面）// 軒庇の下端
		m1, n1 := pkg.ParaLine(x1, y1, x2, y2, x5, y5, hisashi)
		// 頂点4と頂点5を結ぶ線分に平行な直線の式（妻面）
		mh, nh := pkg.ParaLine(x4, y4, x5, y5, x2, y2, hisashi)
		// 頂点2と頂点3を結ぶ線分に平行な直線の式（平面）
		mk, nk := pkg.ParaLine(x2, y2, x3, y3, x4, y4, keraba)
		// 頂点3と頂点4の合成頂点を求める
		xo3, yo3 = pkg.SeekInsec(mh, nh, mk, nk)
		// 屋根伏せの4頂点の座標を求める（軒庇の下端）
		xo1, yo1 = pkg.SeekInsec(m4, n4, m1, n1)
		xo2, yo2 = pkg.SeekInsec(m1, n1, mk, nk)
		xo4, yo4 = pkg.SeekInsec(mh, nh, m4, n4)
		// 頂点1と頂点2を結ぶ線分に平行な直線の式（平面）// 軒庇の上端
		mtp1, ntp1 := pkg.ParaLine(x1, y1, x2, y2, x5, y5, hichtop)
		// 頂点5と頂点1を結ぶ線分に平行な直線の式（平面）// 軒庇の上端
		mtp2, ntp2 := pkg.ParaLine(x4, y4, x5, y5, x1, y1, hichtop)
		// 屋根伏せの4頂点の座標を求める（軒庇の上端）
		xtp1, ytp1 = pkg.SeekInsec(m4, n4, mtp1, ntp1)
		xtp2, ytp2 = pkg.SeekInsec(mtp1, ntp1, mk, nk)
		xtp3, ytp3 = pkg.SeekInsec(mk, nk, mtp2, ntp2)
		xtp4, ytp4 = pkg.SeekInsec(mtp2, ntp2, m4, n4)
	}

	// 屋根の棟端の座標を求める
	xm1 := (xo4 + xo1) / 2
	ym1 := (yo4 + yo1) / 2
	xm2 := (xo2 + xo3) / 2
	ym2 := (yo2 + yo3) / 2

	// 屋根の妻面の長さを求める
	tm1 := pkg.DistVerts(xo4, yo4, xo1, yo1)
	tm2 := pkg.DistVerts(xo2, yo2, xo3, yo3)
	log.Println("tm1, tm2", tm1, tm2)

	// 軒先下端高さ（庇×屋根勾配）
	nbt := (toph - hisashi*incline)
	// 軒先上端高さ
	ntp := nbt + yaneatu/math.Sqrt(math.Pow(incline, 2)+1)
	log.Println("nbt, ntp", nbt, ntp)

	// 屋根の棟端の下端高さ
	mbt1 := nbt + tm1/2*incline
	mbt2 := nbt + tm2/2*incline
	// 屋根の棟端の上端高さ
	mtp1 := mbt1 + yaneatu*math.Sqrt(1+math.Pow(incline, 2))
	mtp2 := mbt2 + yaneatu*math.Sqrt(1+math.Pow(incline, 2))
	log.Println("mbt1, mtp1, mbt2, mtp2", mbt1, mtp1, mbt2, mtp2)

	// 屋根モデルの頂点座標をリストに書き出す
	// 屋根底面・三角形１のIDと頂点座標の書き出し
	sub = sub + 1
	sftype = "Roof"
	// bldg:boundedBy をファイルに書き出す
	expbldg(f, sftype, id, fid, sub)
	// 頂点座標データ（閉じた図形）
	yanepoly = append(yanepoly, xo1, yo1, nbt)
	yanepoly = append(yanepoly, xo2, yo2, nbt)
	yanepoly = append(yanepoly, xm1, ym1, mbt1)
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
	yanepoly = append(yanepoly, xm2, ym2, mbt2)
	yanepoly = append(yanepoly, xm1, ym1, mbt1)
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
	yanepoly = append(yanepoly, xm1, ym1, mbt1)
	yanepoly = append(yanepoly, xm2, ym2, mbt2)
	yanepoly = append(yanepoly, xo4, yo4, nbt)
	yanepoly = append(yanepoly, xm1, ym1, mbt1)
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
	yanepoly = append(yanepoly, xm2, ym2, mbt2)
	yanepoly = append(yanepoly, xo3, yo3, nbt)
	yanepoly = append(yanepoly, xo4, yo4, nbt)
	yanepoly = append(yanepoly, xm2, ym2, mbt2)
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
	yanepoly = append(yanepoly, xtp2, ytp2, ntp)
	yanepoly = append(yanepoly, xm1, ym1, mtp1)
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
	yanepoly = append(yanepoly, xm2, ym2, mtp2)
	yanepoly = append(yanepoly, xm1, ym1, mtp1)
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
	yanepoly = append(yanepoly, xm1, ym1, mtp1)
	yanepoly = append(yanepoly, xm2, ym2, mtp2)
	yanepoly = append(yanepoly, xtp4, ytp4, ntp)
	yanepoly = append(yanepoly, xm1, ym1, mtp1)
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
	yanepoly = append(yanepoly, xm2, ym2, mtp2)
	yanepoly = append(yanepoly, xtp3, ytp3, ntp)
	yanepoly = append(yanepoly, xtp4, ytp4, ntp)
	yanepoly = append(yanepoly, xm2, ym2, mtp2)
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
	yanepoly = append(yanepoly, xtp2, ytp2, ntp)
	yanepoly = append(yanepoly, xo2, yo2, nbt)
	yanepoly = append(yanepoly, xo1, yo1, nbt)
	yanepoly = append(yanepoly, xtp2, ytp2, ntp)
	// bldg:boundedBy の書き出しを終了する
	exitbldg(f, yanepoly, sftype)
	// 屋根データの頂点座標の初期化
	yanepoly = yanepoly[:0]

	// 軒端１・三角形-2のIDと頂点座標の書き出し
	sub = sub + 1
	sftype = "Roof"
	// bldg:boundedBy をファイルに書き出す
	expbldg(f, sftype, id, fid, sub)
	// 頂点座標データ（閉じた図形）
	yanepoly = append(yanepoly, xtp2, ytp2, ntp)
	yanepoly = append(yanepoly, xo1, yo1, nbt)
	yanepoly = append(yanepoly, xtp1, ytp1, ntp)
	yanepoly = append(yanepoly, xtp2, ytp2, ntp)
	// bldg:boundedBy の書き出しを終了する
	exitbldg(f, yanepoly, sftype)
	// 屋根データの頂点座標の初期化
	yanepoly = yanepoly[:0]

	// 軒端２・三角形-1のIDと頂点座標の書き出し
	sub = sub + 1
	sftype = "Roof"
	// bldg:boundedBy をファイルに書き出す
	expbldg(f, sftype, id, fid, sub)
	// 頂点座標データ（閉じた図形）
	yanepoly = append(yanepoly, xtp4, ytp4, ntp)
	yanepoly = append(yanepoly, xo4, yo4, nbt)
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
	yanepoly = append(yanepoly, xtp3, ytp3, ntp)
	yanepoly = append(yanepoly, xtp4, ytp4, ntp)
	// bldg:boundedBy の書き出しを終了する
	exitbldg(f, yanepoly, sftype)
	// 屋根データの頂点座標の初期化
	yanepoly = yanepoly[:0]

	// ケラバ１・三角形-1のIDと頂点座標の書き出し
	sub = sub + 1
	sftype = "Roof"
	// bldg:boundedBy をファイルに書き出す
	expbldg(f, sftype, id, fid, sub)
	// 頂点座標データ（閉じた図形）
	yanepoly = append(yanepoly, xtp1, ytp1, ntp)
	yanepoly = append(yanepoly, xo1, yo1, nbt)
	yanepoly = append(yanepoly, xm1, ym1, mbt1)
	yanepoly = append(yanepoly, xtp1, ytp1, ntp)
	// bldg:boundedBy の書き出しを終了する
	exitbldg(f, yanepoly, sftype)
	// 屋根データの頂点座標の初期化
	yanepoly = yanepoly[:0]

	// ケラバ１・三角形-2のIDと頂点座標の書き出し
	sub = sub + 1
	sftype = "Roof"
	// bldg:boundedBy をファイルに書き出す
	expbldg(f, sftype, id, fid, sub)
	// 頂点座標データ（閉じた図形）
	yanepoly = append(yanepoly, xtp1, ytp1, ntp)
	yanepoly = append(yanepoly, xm1, ym1, mbt1)
	yanepoly = append(yanepoly, xm1, ym1, mtp1)
	yanepoly = append(yanepoly, xtp1, ytp1, ntp)
	// bldg:boundedBy の書き出しを終了する
	exitbldg(f, yanepoly, sftype)
	// 屋根データの頂点座標の初期化
	yanepoly = yanepoly[:0]

	// ケラバ２・三角形-1のIDと頂点座標の書き出し
	sub = sub + 1
	sftype = "Roof"
	// bldg:boundedBy をファイルに書き出す
	expbldg(f, sftype, id, fid, sub)
	// 頂点座標データ（閉じた図形）
	yanepoly = append(yanepoly, xm1, ym1, mtp1)
	yanepoly = append(yanepoly, xm1, ym1, mbt1)
	yanepoly = append(yanepoly, xo4, yo4, nbt)
	yanepoly = append(yanepoly, xm1, ym1, mtp1)
	// bldg:boundedBy の書き出しを終了する
	exitbldg(f, yanepoly, sftype)
	// 屋根データの頂点座標の初期化
	yanepoly = yanepoly[:0]

	// ケラバ２・三角形-2のIDと頂点座標の書き出し
	sub = sub + 1
	sftype = "Roof"
	// bldg:boundedBy をファイルに書き出す
	expbldg(f, sftype, id, fid, sub)
	// 頂点座標データ（閉じた図形）
	yanepoly = append(yanepoly, xm1, ym1, mtp1)
	yanepoly = append(yanepoly, xo4, yo4, nbt)
	yanepoly = append(yanepoly, xtp4, ytp4, ntp)
	yanepoly = append(yanepoly, xm1, ym1, mtp1)
	// bldg:boundedBy の書き出しを終了する
	exitbldg(f, yanepoly, sftype)
	// 屋根データの頂点座標の初期化
	yanepoly = yanepoly[:0]

	// ケラバ３・三角形-1のIDと頂点座標の書き出し
	sub = sub + 1
	sftype = "Roof"
	// bldg:boundedBy をファイルに書き出す
	expbldg(f, sftype, id, fid, sub)
	// 頂点座標データ（閉じた図形）
	yanepoly = append(yanepoly, xtp3, ytp3, ntp)
	yanepoly = append(yanepoly, xo3, yo3, nbt)
	yanepoly = append(yanepoly, xm2, ym2, mbt2)
	yanepoly = append(yanepoly, xtp3, ytp3, ntp)
	// bldg:boundedBy の書き出しを終了する
	exitbldg(f, yanepoly, sftype)
	// 屋根データの頂点座標の初期化
	yanepoly = yanepoly[:0]

	// ケラバ３・三角形-2のIDと頂点座標の書き出し
	sub = sub + 1
	sftype = "Roof"
	// bldg:boundedBy をファイルに書き出す
	expbldg(f, sftype, id, fid, sub)
	// 頂点座標データ（閉じた図形）
	yanepoly = append(yanepoly, xtp3, ytp3, ntp)
	yanepoly = append(yanepoly, xm2, ym2, mbt2)
	yanepoly = append(yanepoly, xm2, ym2, mtp2)
	yanepoly = append(yanepoly, xtp3, ytp3, ntp)
	// bldg:boundedBy の書き出しを終了する
	exitbldg(f, yanepoly, sftype)
	// 屋根データの頂点座標の初期化
	yanepoly = yanepoly[:0]

	// ケラバ４・三角形-1のIDと頂点座標の書き出し
	sub = sub + 1
	sftype = "Roof"
	// bldg:boundedBy をファイルに書き出す
	expbldg(f, sftype, id, fid, sub)
	// 頂点座標データ（閉じた図形）
	yanepoly = append(yanepoly, xm2, ym2, mtp2)
	yanepoly = append(yanepoly, xm2, ym2, mbt2)
	yanepoly = append(yanepoly, xo2, yo2, nbt)
	yanepoly = append(yanepoly, xm2, ym2, mtp2)
	// bldg:boundedBy の書き出しを終了する
	exitbldg(f, yanepoly, sftype)
	// 屋根データの頂点座標の初期化
	yanepoly = yanepoly[:0]

	// ケラバ４・三角形-2のIDと頂点座標の書き出し
	sub = sub + 1
	sftype = "Roof"
	// bldg:boundedBy をファイルに書き出す
	expbldg(f, sftype, id, fid, sub)
	// 頂点座標データ（閉じた図形）
	yanepoly = append(yanepoly, xm2, ym2, mtp2)
	yanepoly = append(yanepoly, xo2, yo2, nbt)
	yanepoly = append(yanepoly, xtp2, ytp2, ntp)
	yanepoly = append(yanepoly, xm2, ym2, mtp2)
	// bldg:boundedBy の書き出しを終了する
	exitbldg(f, yanepoly, sftype)
	// 屋根データの頂点座標の初期化
	yanepoly = yanepoly[:0]

	// 妻面三角壁の下端高さ
	h := toph
	// 頂点2と頂点3を結ぶ直線の式
	line1 := pkg.LineEquat(x2, y2, x3, y3)
	// 頂点4と頂点5を結ぶ直線の式
	line2 := pkg.LineEquat(x4, y4, x5, y5)
	// 頂点3と頂点4の合成頂点を求める
	x0, y0 := pkg.SeekInsec(line1["m"], line1["n"], line2["m"], line2["n"])

	if d1 < d2 {
		// 妻面三角壁1の下端長さ
		l1 := pkg.DistVerts(x1, y1, x2, y2)
		// 妻面三角壁1の上端高さ
		mh1 := h + l1/2*incline
		// 妻面三角壁1の頂点座標
		xc1 := (x1 + x2) / 2
		yc1 := (y1 + y2) / 2

		// 妻壁・三角形１のIDと頂点座標の書き出し
		sub = sub + 1
		sftype = "Roof"
		// bldg:boundedBy をファイルに書き出す
		expbldg(f, sftype, id, fid, sub)
		// 頂点座標データ（閉じた図形）
		yanepoly = append(yanepoly, x1, y1, h)
		yanepoly = append(yanepoly, x2, y2, h)
		yanepoly = append(yanepoly, xc1, yc1, mh1)
		yanepoly = append(yanepoly, x1, y1, h)
		// bldg:boundedBy の書き出しを終了する
		exitbldg(f, yanepoly, sftype)
		// 屋根データの頂点座標の初期化
		yanepoly = yanepoly[:0]

		// // 頂点2と頂点3を結ぶ直線の式
		// line1 := pkg.LineEquat(x2, y2, x3, y3)
		// // 頂点4と頂点5を結ぶ直線の式
		// line2 := pkg.LineEquat(x4, y4, x5, y5)
		// // 頂点3と頂点4の合成頂点を求める
		// x0, y0 := pkg.SeekInsec(line1["m"], line1["n"], line2["m"], line2["n"])

		// 妻面三角壁2の下端長さ
		l2 := pkg.DistVerts(x0, y0, x5, y5)
		// 妻面三角壁2の上端高さ
		mh2 := h + l2/2*incline
		// 妻面三角壁2の頂点座標
		xc2 := (x0 + x5) / 2
		yc2 := (y0 + y5) / 2

		// 妻壁・三角形２-1のIDと頂点座標の書き出し
		sub = sub + 1
		sftype = "Roof"
		// bldg:boundedBy をファイルに書き出す
		expbldg(f, sftype, id, fid, sub)
		// 頂点座標データ（閉じた図形）
		yanepoly = append(yanepoly, x3, y3, h)
		yanepoly = append(yanepoly, x4, y4, h)
		yanepoly = append(yanepoly, xc2, yc2, mh2)
		yanepoly = append(yanepoly, x3, y3, h)
		// bldg:boundedBy の書き出しを終了する
		exitbldg(f, yanepoly, sftype)
		// 屋根データの頂点座標の初期化
		yanepoly = yanepoly[:0]

		// 妻壁・三角形２-2のIDと頂点座標の書き出し
		sub = sub + 1
		sftype = "Roof"
		// bldg:boundedBy をファイルに書き出す
		expbldg(f, sftype, id, fid, sub)
		// 頂点座標データ（閉じた図形）
		yanepoly = append(yanepoly, x5, y5, h)
		yanepoly = append(yanepoly, x4, y4, h)
		yanepoly = append(yanepoly, xc2, yc2, mh2)
		yanepoly = append(yanepoly, x5, y5, h)
		// bldg:boundedBy の書き出しを終了する
		exitbldg(f, yanepoly, sftype)

	} else if d1 > d2 {
		// 妻面三角壁1の下端長さ
		l1 := pkg.DistVerts(x5, y5, x1, y1)
		// 妻面三角壁1の上端高さ
		mh1 := h + l1/2*incline
		// 妻面三角壁1の頂点座標
		xc1 := (x5 + x1) / 2
		yc1 := (y5 + y1) / 2

		// 妻壁・三角形１のIDと頂点座標の書き出し
		sub = sub + 1
		sftype = "Roof"
		// bldg:boundedBy をファイルに書き出す
		expbldg(f, sftype, id, fid, sub)
		// 頂点座標データ（閉じた図形）
		yanepoly = append(yanepoly, x1, y1, h)
		yanepoly = append(yanepoly, x5, y5, h)
		yanepoly = append(yanepoly, xc1, yc1, mh1)
		yanepoly = append(yanepoly, x1, y1, h)
		// bldg:boundedBy の書き出しを終了する
		exitbldg(f, yanepoly, sftype)
		// 屋根データの頂点座標の初期化
		yanepoly = yanepoly[:0]

		// // 頂点2と頂点3を結ぶ直線の式
		// line1 := pkg.LineEquat(x2, y2, x3, y3)
		// // 頂点4と頂点5を結ぶ直線の式
		// line2 := pkg.LineEquat(x4, y4, x5, y5)
		// // 頂点3と頂点4の合成頂点を求める
		// x0, y0 := pkg.SeekInsec(line1["m"], line1["n"], line2["m"], line2["n"])

		// 妻面三角壁2の下端長さ
		l2 := pkg.DistVerts(x2, y2, x0, y0)
		// 妻面三角壁2の上端高さ
		mh2 := h + l2/2*incline
		// 妻面三角壁2の頂点座標
		xc2 := (x2 + x0) / 2
		yc2 := (y2 + y0) / 2

		// 妻壁・三角形２-1のIDと頂点座標の書き出し
		sub = sub + 1
		sftype = "Roof"
		// bldg:boundedBy をファイルに書き出す
		expbldg(f, sftype, id, fid, sub)
		// 頂点座標データ（閉じた図形）
		yanepoly = append(yanepoly, x3, y3, h)
		yanepoly = append(yanepoly, x2, y2, h)
		yanepoly = append(yanepoly, xc2, yc2, mh2)
		yanepoly = append(yanepoly, x3, y3, h)
		// bldg:boundedBy の書き出しを終了する
		exitbldg(f, yanepoly, sftype)
		// 屋根データの頂点座標の初期化
		yanepoly = yanepoly[:0]

		// 妻壁・三角形２-2のIDと頂点座標の書き出し
		sub = sub + 1
		sftype = "Roof"
		// bldg:boundedBy をファイルに書き出す
		expbldg(f, sftype, id, fid, sub)
		// 頂点座標データ（閉じた図形）
		yanepoly = append(yanepoly, x3, y3, h)
		yanepoly = append(yanepoly, x4, y4, h)
		yanepoly = append(yanepoly, xc2, yc2, mh2)
		// bldg:boundedBy の書き出しを終了する
		exitbldg(f, yanepoly, sftype)
	}

	// 傾斜屋根建物の壁面・床面座標を出力する
	bldbody(f, id, fid, vcnt, list, elv, btm, toph, story, sub)
}
