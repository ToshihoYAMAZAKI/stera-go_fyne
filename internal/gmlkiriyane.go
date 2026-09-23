package internal

import (
	"log"
	"math"
	"os"
	"stera/pkg"
	"strconv"
	"strings"
)

// bldg:boundedBy をファイルに書き出す
func expbldg(f *os.File, sftype, id, fid string, sub int) {
	// id をUUID に書き換えてテキスト変換して書き出す
	expbndl(f, sftype, id, sub)
	// fid をPolyGMLID に書き換えてテキスト変換して書き出す
	expsrfc(f, fid, sub)
}

// bldg:boundedBy の書き出しを終了する
func exitbldg(f *os.File, polist []float64, sftype string) {
	var listxt []string
	// 頂点データのテキスト化
	for m := range polist {
		listxt = append(listxt, strconv.FormatFloat(polist[m], 'f', -1, 64))
	}
	bscortxt := "\t\t\t\t\t\t\t\t\t\t\t<gml:posList srsDimension=\"3\"> " + strings.Join(listxt, " ") + "</gml:posList>\n"
	f.WriteString(bscortxt)
	// surfaceMember の書き出しを終了する
	exitsrfc(f)
	// bldg:boundedBy の書き出しを終了する
	exitbndl(f, sftype)
}

// bldbody は傾斜屋根建物の壁面・床面座標を出力する
func bldbody(f *os.File, id, fid string, vcnt int, list [][]float64, elv, btm, toph float64, story, sub int) {
	// 基面データの頂点座標の定義
	var basepoly []float64
	// 底面データの頂点座標の定義
	var bttmpoly []float64
	// 上面データの頂点座標の定義
	var toppoly []float64
	// 各階データの頂点座標の定義
	var flrpoly []float64
	// 側面データの頂点座標の定義
	var sidepoly []float64
	// SurfaceType の定義
	var sftype string

	// 基面IDと頂点座標の書き出し
	sub = sub + 1
	sftype = "Floor"
	// bldg:boundedBy をファイルに書き出す
	expbldg(f, sftype, id, fid, sub)
	// 頂点座標データ（閉じた図形）
	for j := 0; j < vcnt; j++ {
		basepoly = append(basepoly, list[j][0])
		basepoly = append(basepoly, list[j][1])
		basepoly = append(basepoly, elv)
	}
	// bldg:boundedBy の書き出しを終了する
	exitbldg(f, basepoly, sftype)

	// 底面IDと頂点座標の書き出し
	sub = sub + 1
	sftype = "Ground"
	// bldg:boundedBy をファイルに書き出す
	expbldg(f, sftype, id, fid, sub)
	// 頂点座標データ（閉じた図形）
	for j := vcnt - 1; j >= 0; j-- {
		bttmpoly = append(bttmpoly, list[j][0])
		bttmpoly = append(bttmpoly, list[j][1])
		bttmpoly = append(bttmpoly, btm)
	}
	// bldg:boundedBy の書き出しを終了する
	exitbldg(f, bttmpoly, sftype)

	// 上面IDと頂点座標の書き出し
	sub = sub + 1
	sftype = "Ceiling"
	// bldg:boundedBy をファイルに書き出す
	expbldg(f, sftype, id, fid, sub)
	// 頂点座標データ（閉じた図形）
	for j := 0; j < vcnt; j++ {
		toppoly = append(toppoly, list[j][0])
		toppoly = append(toppoly, list[j][1])
		toppoly = append(toppoly, toph)
	}
	// bldg:boundedBy の書き出しを終了する
	exitbldg(f, toppoly, sftype)

	// 各階IDと頂点座標の書き出し
	for j := 1; j < story; j++ {
		sub = sub + 1
		sftype = "Floor"
		// bldg:boundedBy をファイルに書き出す
		expbldg(f, sftype, id, fid, sub)
		// 頂点座標データ（閉じた図形）
		for k := 0; k < vcnt; k++ {
			flrpoly = append(flrpoly, list[k][0])
			flrpoly = append(flrpoly, list[k][1])
			flrpoly = append(flrpoly, (elv + 3.3*float64(j)))
		}
		// bldg:boundedBy の書き出しを終了する
		exitbldg(f, flrpoly, sftype)

		// 各階データの頂点座標の初期化
		flrpoly = flrpoly[:0]
	}

	// 側面頂点座標の書き出し
	for j := 0; j < vcnt-1; j++ {
		sub = sub + 1
		sftype = "Wall"
		// bldg:boundedBy をファイルに書き出す
		expbldg(f, sftype, id, fid, sub)
		// 頂点座標データ（閉じた図形）
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
		// bldg:boundedBy の書き出しを終了する
		exitbldg(f, sidepoly, sftype)

		// 側面データの頂点座標の初期化
		sidepoly = sidepoly[:0]
	}
}

// ExposKiri は切妻屋根建物のは頂点座標をファイルに書き出す
func ExposKiri(f *os.File, id, fid string, vcnt int, list [][]float64, elv, btm, toph, hisashi, keraba, incline, yaneatu float64, story int) {
	log.Println("切妻屋根")
	// 屋根モデル（切妻屋根）
	// 4つの頂点座標の定義
	x1 := list[0][0]
	y1 := list[0][1]
	x2 := list[1][0]
	y2 := list[1][1]
	x3 := list[2][0]
	y3 := list[2][1]
	x4 := list[3][0]
	y4 := list[3][1]
	// log.Println("x1, y1, x2, y2, x3, y3, x4, y4", x1, y1, x2, y2, x3, y3, x4, y4)

	// 妻面の長さと平面の長さを比較する
	tuma := pkg.DistVerts(x1, y1, x2, y2)
	hira := pkg.DistVerts(x2, y2, x3, y3)
	if tuma > hira*1.2 {
		x0 := x1
		y0 := y1
		x1 = x2
		y1 = y2
		x2 = x3
		y2 = y3
		x3 = x4
		y3 = y4
		x4 = x0
		y4 = y0
	}

	// 頂点1と頂点2を結ぶ線分に平行な直線の式（妻面）
	m1, n1 := pkg.ParaLine(x1, y1, x2, y2, x3, y3, keraba)
	// log.Println("y = " + strconv.FormatFloat(m1, 'f', -1, 64) + "x + " + strconv.FormatFloat(n1, 'f', -1, 64))
	// 頂点3と頂点4を結ぶ線分に平行な直線の式（妻面）
	m2, n2 := pkg.ParaLine(x3, y3, x4, y4, x1, y1, keraba)
	// log.Println("y = " + strconv.FormatFloat(m2, 'f', -1, 64) + "x + " + strconv.FormatFloat(n2, 'f', -1, 64))
	// 頂点2と頂点3を結ぶ線分に平行な直線の式（平面）// 軒庇の下端
	m3, n3 := pkg.ParaLine(x2, y2, x3, y3, x4, y4, hisashi)
	// log.Println("y = " + strconv.FormatFloat(m3, 'f', -1, 64) + "x + " + strconv.FormatFloat(n3, 'f', -1, 64))
	// 頂点4と頂点1を結ぶ線分に平行な直線の式（平面）// 軒庇の下端
	m4, n4 := pkg.ParaLine(x4, y4, x1, y1, x2, y2, hisashi)
	// log.Println("y = " + strconv.FormatFloat(m4, 'f', -1, 64) + "x + " + strconv.FormatFloat(n4, 'f', -1, 64))
	log.Println("m1, n1=", m1, n1)
	log.Println("m2, n2=", m2, n2)
	log.Println("m3, n3=", m3, n3)
	log.Println("m4, n4=", m4, n4)

	// 屋根伏せの4頂点の座標を求める（軒庇の下端）
	xo1, yo1 := pkg.SeekInsec(m4, n4, m1, n1)
	xo2, yo2 := pkg.SeekInsec(m1, n1, m3, n3)
	xo3, yo3 := pkg.SeekInsec(m3, n3, m2, n2)
	xo4, yo4 := pkg.SeekInsec(m2, n2, m4, n4)
	log.Println("xo1, yo1, xo2, yo2, xo3, yo3, xo4, yo4", xo1, yo1, xo2, yo2, xo3, yo3, xo4, yo4)

	// 軒鼻の突き出し長さ
	nt := yaneatu / math.Sqrt(math.Pow(incline, 2)+1) * incline
	// 平面から軒庇の上端までの長さ（軒庇＋軒鼻）
	hichtop := hisashi + nt
	// 頂点2と頂点3を結ぶ線分に平行な直線の式（平面）// 軒庇の上端
	mtp3, ntp3 := pkg.ParaLine(x2, y2, x3, y3, x1, y1, hichtop)
	// 頂点4と頂点1を結ぶ線分に平行な直線の式（平面）// 軒庇の上端
	mtp4, ntp4 := pkg.ParaLine(x4, y4, x1, y1, x2, y2, hichtop)

	// 屋根伏せの4頂点の座標を求める（軒庇の上端）
	xtp1, ytp1 := pkg.SeekInsec(mtp4, ntp4, m1, n1)
	xtp2, ytp2 := pkg.SeekInsec(m1, n1, mtp3, ntp3)
	xtp3, ytp3 := pkg.SeekInsec(mtp3, ntp3, m2, n2)
	xtp4, ytp4 := pkg.SeekInsec(m2, n2, mtp4, ntp4)
	log.Println("xtp1, ytp1, xtp2, ytp2, xtp3, ytp3, xtp4, ytp4", xtp1, ytp1, xtp2, ytp2, xtp3, ytp3, xtp4, ytp4)

	// 屋根の棟端の座標を求める
	xm1 := (xo1 + xo2) / 2
	ym1 := (yo1 + yo2) / 2
	xm2 := (xo3 + xo4) / 2
	ym2 := (yo3 + yo4) / 2
	log.Println("xm1, ym1, xm2, ym2", xm1, ym1, xm2, ym2)

	// 屋根の妻面の長さを求める
	tm1 := pkg.DistVerts(xo1, yo1, xo2, yo2)
	tm2 := pkg.DistVerts(xo3, yo3, xo4, yo4)
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

	// 屋根データの頂点座標の定義
	var yanepoly []float64
	// PolyGMLID をfid から設定するための添え字の定義
	sub := 0
	// SurfaceType の定義
	var sftype string

	// 屋根底面・三角形１のIDと頂点座標の書き出し
	sub = sub + 1
	sftype = "Roof"
	// // bldg:boundedBy をファイルに書き出す
	// // id をUUID に書き換えてテキスト変換して書き出す
	// expbndl(f, sftype, id, sub)
	// // fid をPolyGMLID に書き換えてテキスト変換して書き出す
	// expsrfc(f, fid, sub)
	// bldg:boundedBy をファイルに書き出す
	expbldg(f, sftype, id, fid, sub)
	// 頂点座標データ（閉じた図形）
	yanepoly = append(yanepoly, xo1, yo1, nbt)
	yanepoly = append(yanepoly, xo4, yo4, nbt)
	yanepoly = append(yanepoly, xm1, ym1, mbt1)
	yanepoly = append(yanepoly, xo1, yo1, nbt)
	// // 頂点データのテキスト化
	// for m := range yanepoly {
	// 	yanetxt = append(yanetxt, strconv.FormatFloat(yanepoly[m], 'f', -1, 64))
	// }
	// bscortxt = "\t\t\t\t\t\t\t\t\t\t\t<gml:posList srsDimension=\"3\"> " + strings.Join(yanetxt, " ") + "</gml:posList>\n"
	// f.WriteString(bscortxt)
	// // surfaceMember の書き出しを終了する
	// exitsrfc(f)
	// // bldg:boundedBy の書き出しを終了する
	// exitbndl(f, sftype)
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
	yanepoly = append(yanepoly, xm2, ym2, mbt2)
	yanepoly = append(yanepoly, xm1, ym1, mbt1)
	yanepoly = append(yanepoly, xo4, yo4, nbt)
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
	yanepoly = append(yanepoly, xo2, yo2, nbt)
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
	yanepoly = append(yanepoly, xo2, yo2, nbt)
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
	yanepoly = append(yanepoly, xtp4, ytp4, ntp)
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
	yanepoly = append(yanepoly, xtp4, ytp4, ntp)
	yanepoly = append(yanepoly, xm2, ym2, mtp2)
	yanepoly = append(yanepoly, xm1, ym1, mtp1)
	yanepoly = append(yanepoly, xtp4, ytp4, ntp)
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
	yanepoly = append(yanepoly, xtp2, ytp2, ntp)
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
	yanepoly = append(yanepoly, xtp2, ytp2, ntp)
	yanepoly = append(yanepoly, xm2, ym2, mtp2)
	// bldg:boundedBy の書き出しを終了する
	exitbldg(f, yanepoly, sftype)
	// 屋根データの頂点座標の初期化
	yanepoly = yanepoly[:0]

	//  軒端１・三角形-1のIDと頂点座標の書き出し
	sub = sub + 1
	sftype = "Roof"
	// bldg:boundedBy をファイルに書き出す
	expbldg(f, sftype, id, fid, sub)
	// 頂点座標データ（閉じた図形）
	yanepoly = append(yanepoly, xtp2, ytp2, ntp)
	yanepoly = append(yanepoly, xo2, yo2, nbt)
	yanepoly = append(yanepoly, xo3, yo3, nbt)
	yanepoly = append(yanepoly, xtp2, ytp2, ntp)
	// bldg:boundedBy の書き出しを終了する
	exitbldg(f, yanepoly, sftype)
	// 屋根データの頂点座標の初期化
	yanepoly = yanepoly[:0]

	//  軒端１・三角形-2のIDと頂点座標の書き出し
	sub = sub + 1
	sftype = "Roof"
	// bldg:boundedBy をファイルに書き出す
	expbldg(f, sftype, id, fid, sub)
	// 頂点座標データ（閉じた図形）
	yanepoly = append(yanepoly, xtp2, ytp2, ntp)
	yanepoly = append(yanepoly, xo3, yo3, nbt)
	yanepoly = append(yanepoly, xtp3, ytp3, ntp)
	yanepoly = append(yanepoly, xtp2, ytp2, ntp)
	// bldg:boundedBy の書き出しを終了する
	exitbldg(f, yanepoly, sftype)
	// 屋根データの頂点座標の初期化
	yanepoly = yanepoly[:0]

	//  軒端２・三角形-1のIDと頂点座標の書き出し
	sub = sub + 1
	sftype = "Roof"
	// bldg:boundedBy をファイルに書き出す
	expbldg(f, sftype, id, fid, sub)
	// 頂点座標データ（閉じた図形）
	yanepoly = append(yanepoly, xtp4, ytp4, ntp)
	yanepoly = append(yanepoly, xo4, yo4, nbt)
	yanepoly = append(yanepoly, xo1, yo1, nbt)
	yanepoly = append(yanepoly, xtp4, ytp4, ntp)
	// bldg:boundedBy の書き出しを終了する
	exitbldg(f, yanepoly, sftype)
	// 屋根データの頂点座標の初期化
	yanepoly = yanepoly[:0]

	//  軒端２・三角形-2のIDと頂点座標の書き出し
	sub = sub + 1
	sftype = "Roof"
	// bldg:boundedBy をファイルに書き出す
	expbldg(f, sftype, id, fid, sub)
	// 頂点座標データ（閉じた図形）
	yanepoly = append(yanepoly, xtp4, ytp4, ntp)
	yanepoly = append(yanepoly, xo1, yo1, nbt)
	yanepoly = append(yanepoly, xtp1, ytp1, ntp)
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
	yanepoly = append(yanepoly, xo2, yo2, nbt)
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
	yanepoly = append(yanepoly, xo2, yo2, nbt)
	yanepoly = append(yanepoly, xtp2, ytp2, ntp)
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
	yanepoly = append(yanepoly, xo4, yo4, nbt)
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
	yanepoly = append(yanepoly, xo4, yo4, nbt)
	yanepoly = append(yanepoly, xtp4, ytp4, ntp)
	yanepoly = append(yanepoly, xm2, ym2, mtp2)
	// bldg:boundedBy の書き出しを終了する
	exitbldg(f, yanepoly, sftype)
	// 屋根データの頂点座標の初期化
	yanepoly = yanepoly[:0]

	// 妻面三角壁の下端高さ
	h := toph
	// 妻面三角壁1の下端長さ
	l1 := pkg.DistVerts(x1, y1, x2, y2)
	// 妻面三角壁1の上端高さ
	mh1 := h + l1/2*incline
	// 妻面三角壁1の頂点座標
	xc1 := (x1 + x2) / 2
	yc1 := (y1 + y2) / 2

	// 妻壁・三角形１のIDと頂点座標の書き出し
	sub = sub + 1
	sftype = "Wall"
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

	// 妻面三角壁2の下端長さ
	l2 := pkg.DistVerts(x3, y3, x4, y4)
	// 妻面三角壁2の上端高さ
	mh2 := h + l2/2*incline
	// 妻面三角壁2の頂点座標
	xc2 := (x3 + x4) / 2
	yc2 := (y3 + y4) / 2

	// 妻壁・三角形２のIDと頂点座標の書き出し
	sub = sub + 1
	sftype = "Wall"
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
	// yanepoly = yanepoly[:0]

	// 傾斜屋根建物の壁面・床面座標を出力する
	bldbody(f, id, fid, vcnt, list, elv, btm, toph, story, sub)

	// // 基面データの頂点座標の定義
	// var basepoly []float64
	// // 基面データの頂点座標の文字列
	// // var basetxt []string
	// // 底面データの頂点座標の定義
	// var bttmpoly []float64
	// // 底面データの頂点座標の文字列
	// // var bttmtxt []string
	// // 上面データの頂点座標の定義
	// var toppoly []float64
	// // 上面データの頂点座標の文字列
	// // var toptxt []string
	// // 各階データの頂点座標の定義
	// var flrpoly []float64
	// // 各階データの頂点座標の文字列
	// // var flrtxt []string
	// // 側面データの頂点座標の定義
	// var sidepoly []float64
	// // 側面データの頂点座標の文字列
	// // var sidetxt []string

	// // 基面IDと頂点座標の書き出し
	// sub = sub + 1
	// sftype = "Floor"
	// // bldg:boundedBy をファイルに書き出す
	// expbldg(f, sftype, id, fid, sub)
	// // 頂点座標データ（閉じた図形）
	// for j := 0; j < vcnt; j++ {
	// 	basepoly = append(basepoly, list[j][0])
	// 	basepoly = append(basepoly, list[j][1])
	// 	basepoly = append(basepoly, elv)
	// }
	// // bldg:boundedBy の書き出しを終了する
	// exitbldg(f, basepoly, sftype)

	// // 底面IDと頂点座標の書き出し
	// sub = sub + 1
	// sftype = "Ground"
	// // bldg:boundedBy をファイルに書き出す
	// expbldg(f, sftype, id, fid, sub)
	// // 頂点座標データ（閉じた図形）
	// for j := vcnt - 1; j >= 0; j-- {
	// 	bttmpoly = append(bttmpoly, list[j][0])
	// 	bttmpoly = append(bttmpoly, list[j][1])
	// 	bttmpoly = append(bttmpoly, btm)
	// }
	// // bldg:boundedBy の書き出しを終了する
	// exitbldg(f, bttmpoly, sftype)

	// // 上面IDと頂点座標の書き出し
	// sub = sub + 1
	// sftype = "Ceiling"
	// // bldg:boundedBy をファイルに書き出す
	// expbldg(f, sftype, id, fid, sub)
	// // 頂点座標データ（閉じた図形）
	// for j := 0; j < vcnt; j++ {
	// 	toppoly = append(toppoly, list[j][0])
	// 	toppoly = append(toppoly, list[j][1])
	// 	toppoly = append(toppoly, toph)
	// }
	// // bldg:boundedBy の書き出しを終了する
	// exitbldg(f, toppoly, sftype)

	// // 各階IDと頂点座標の書き出し
	// for j := 1; j < story; j++ {
	// 	sub = sub + 1
	// 	sftype = "Floor"
	// 	// bldg:boundedBy をファイルに書き出す
	// 	expbldg(f, sftype, id, fid, sub)
	// 	// 頂点座標データ（閉じた図形）
	// 	for k := 0; k < vcnt; k++ {
	// 		flrpoly = append(flrpoly, list[k][0])
	// 		flrpoly = append(flrpoly, list[k][1])
	// 		flrpoly = append(flrpoly, (elv + 3.3*float64(j)))
	// 	}
	// 	// bldg:boundedBy の書き出しを終了する
	// 	exitbldg(f, flrpoly, sftype)

	// 	// 各階データの頂点座標の初期化
	// 	flrpoly = flrpoly[:0]
	// }

	// // 側面頂点座標の書き出し
	// for j := 0; j < vcnt-1; j++ {
	// 	sub = sub + 1
	// 	sftype = "Wall"
	// 	// bldg:boundedBy をファイルに書き出す
	// 	expbldg(f, sftype, id, fid, sub)
	// 	// 頂点座標データ（閉じた図形）
	// 	// 頂点下1
	// 	sidepoly = append(sidepoly, list[j][0])
	// 	sidepoly = append(sidepoly, list[j][1])
	// 	sidepoly = append(sidepoly, btm)
	// 	// 頂点下2
	// 	sidepoly = append(sidepoly, list[(j+1)%vcnt][0])
	// 	sidepoly = append(sidepoly, list[(j+1)%vcnt][1])
	// 	sidepoly = append(sidepoly, btm)
	// 	// 頂点上2
	// 	sidepoly = append(sidepoly, list[(j+1)%vcnt][0])
	// 	sidepoly = append(sidepoly, list[(j+1)%vcnt][1])
	// 	sidepoly = append(sidepoly, toph)
	// 	// 頂点上1
	// 	sidepoly = append(sidepoly, list[j][0])
	// 	sidepoly = append(sidepoly, list[j][1])
	// 	sidepoly = append(sidepoly, toph)
	// 	// bldg:boundedBy の書き出しを終了する
	// 	exitbldg(f, sidepoly, sftype)

	// 	// 側面データの頂点座標の初期化
	// 	sidepoly = sidepoly[:0]
	// }
}
