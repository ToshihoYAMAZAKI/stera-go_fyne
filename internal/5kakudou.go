package internal

import (
	"log"
	"math"
	"stera/pkg"
	"strconv"
)

// Kakudou5 は五角堂屋根の頂点座標と頂点の法線ベクトルを求める
func Kakudou5(list [][]float64, toph, hisashi, incline, yaneatu float64) (yanetxt, yanenor []string) {
	// 屋根モデル（五角堂屋根）
	log.Println("五角堂屋根")

	// 5つの頂点座標の定義
	kansan := 0.0254
	x1 := list[0][0] / kansan
	y1 := list[0][1] / kansan
	x2 := list[1][0] / kansan
	y2 := list[1][1] / kansan
	x3 := list[2][0] / kansan
	y3 := list[2][1] / kansan
	x4 := list[3][0] / kansan
	y4 := list[3][1] / kansan
	x5 := list[4][0] / kansan
	y5 := list[4][1] / kansan

	// インチ換算後の庇長さ（軒の出）
	hich := hisashi / kansan
	// 軒鼻の突き出し長さ
	nt := yaneatu / kansan / math.Sqrt(math.Pow(incline, 2)+1) * incline
	// 平面から軒庇の上端までの長さ（軒庇＋軒鼻）
	hichtop := hich + nt

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
	var nor_all [][]float64
	var normal []float64
	var nor []float64
	p := make([][]float64, 3)

	// 五角堂の棟の中心座標
	xo := (x1 + x2 + x3 + x4 + x5) / 5
	yo := (y1 + y2 + y3 + y4 + y5) / 5
	// 頂点1と頂点2を結ぶ線分に平行な直線の式
	m1, n1 := pkg.ParaLine(x1, y1, x2, y2, xo, yo, hich)
	// 頂点2と頂点3を結ぶ線分に平行な直線の式
	m2, n2 := pkg.ParaLine(x2, y2, x3, y3, xo, yo, hich)
	// 頂点3と頂点4を結ぶ線分に平行な直線の式
	m3, n3 := pkg.ParaLine(x3, y3, x4, y4, xo, yo, hich)
	// 頂点4と頂点5を結ぶ線分に平行な直線の式
	m4, n4 := pkg.ParaLine(x4, y4, x5, y5, xo, yo, hich)
	// 頂点5と頂点1を結ぶ線分に平行な直線の式
	m5, n5 := pkg.ParaLine(x5, y5, x1, y1, xo, yo, hich)
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
	nbt := (toph - hisashi*incline) / kansan
	// 軒先上端高さ
	ntp := nbt + yaneatu/kansan/math.Sqrt(math.Pow(incline, 2)+1)
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
	mtp := mbt + yaneatu/kansan*math.Sqrt(1+math.Pow(incline, 2))
	log.Println("mbt, mtp", mbt, mtp)

	yanepoly = append(yanepoly, xo1, yo1, nbt) // 屋根底面・三角形１
	yanepoly = append(yanepoly, xo, yo, mbt)
	yanepoly = append(yanepoly, xo2, yo2, nbt)
	p[0] = []float64{xo1, yo1, nbt}
	p[1] = []float64{xo, yo, mbt}
	p[2] = []float64{xo4, yo4, nbt}
	for j := 0; j < 3; j++ {
		nor = pkg.NorVec(p[(0+j)%3], p[(1+j)%3], p[(2+j)%3])
		nor_all = append(nor_all, nor)
	}

	yanepoly = append(yanepoly, xo2, yo2, nbt) // 屋根底面・三角形２
	yanepoly = append(yanepoly, xo, yo, mbt)
	yanepoly = append(yanepoly, xo3, yo3, nbt)
	p[0] = []float64{xo2, yo2, nbt}
	p[1] = []float64{xo, yo, mbt}
	p[2] = []float64{xo3, yo3, nbt}
	for j := 0; j < 3; j++ {
		nor = pkg.NorVec(p[(0+j)%3], p[(1+j)%3], p[(2+j)%3])
		nor_all = append(nor_all, nor)
	}

	yanepoly = append(yanepoly, xo3, yo3, nbt) // 屋根底面・三角形３
	yanepoly = append(yanepoly, xo, yo, mbt)
	yanepoly = append(yanepoly, xo4, yo4, nbt)
	p[0] = []float64{xo3, yo3, nbt}
	p[1] = []float64{xo, yo, mbt}
	p[2] = []float64{xo4, yo4, nbt}
	for j := 0; j < 3; j++ {
		nor = pkg.NorVec(p[(0+j)%3], p[(1+j)%3], p[(2+j)%3])
		nor_all = append(nor_all, nor)
	}

	yanepoly = append(yanepoly, xo4, yo4, nbt) // 屋根底面・三角形４
	yanepoly = append(yanepoly, xo, yo, mbt)
	yanepoly = append(yanepoly, xo5, yo5, nbt)
	p[0] = []float64{xo4, yo4, nbt}
	p[1] = []float64{xo, yo, mbt}
	p[2] = []float64{xo5, yo5, nbt}
	for j := 0; j < 3; j++ {
		nor = pkg.NorVec(p[(0+j)%3], p[(1+j)%3], p[(2+j)%3])
		nor_all = append(nor_all, nor)
	}

	yanepoly = append(yanepoly, xo5, yo5, nbt) // 屋根底面・三角形５
	yanepoly = append(yanepoly, xo, yo, mbt)
	yanepoly = append(yanepoly, xo1, yo1, nbt)
	p[0] = []float64{xo5, yo5, nbt}
	p[1] = []float64{xo, yo, mbt}
	p[2] = []float64{xo1, yo1, nbt}
	for j := 0; j < 3; j++ {
		nor = pkg.NorVec(p[(0+j)%3], p[(1+j)%3], p[(2+j)%3])
		nor_all = append(nor_all, nor)
	}

	yanepoly = append(yanepoly, xtp1, ytp1, ntp) // 屋根上面・三角形１
	yanepoly = append(yanepoly, xo, yo, mtp)
	yanepoly = append(yanepoly, xtp2, ytp2, ntp)
	p[0] = []float64{xtp1, ytp1, ntp}
	p[1] = []float64{xo, yo, mtp}
	p[2] = []float64{xtp4, ytp4, ntp}
	for j := 0; j < 3; j++ {
		nor = pkg.NorVec(p[(0+j)%3], p[(1+j)%3], p[(2+j)%3])
		nor_all = append(nor_all, nor)
	}

	yanepoly = append(yanepoly, xtp2, ytp2, ntp) // 屋根上面・三角形２
	yanepoly = append(yanepoly, xo, yo, mtp)
	yanepoly = append(yanepoly, xtp3, ytp3, ntp)
	p[0] = []float64{xtp2, ytp2, ntp}
	p[1] = []float64{xo, yo, mtp}
	p[2] = []float64{xtp3, ytp3, ntp}
	for j := 0; j < 3; j++ {
		nor = pkg.NorVec(p[(0+j)%3], p[(1+j)%3], p[(2+j)%3])
		nor_all = append(nor_all, nor)
	}

	yanepoly = append(yanepoly, xtp3, ytp3, ntp) // 屋根上面・三角形３
	yanepoly = append(yanepoly, xo, yo, mtp)
	yanepoly = append(yanepoly, xtp4, ytp4, ntp)
	p[0] = []float64{xtp3, ytp3, ntp}
	p[1] = []float64{xo, yo, mtp}
	p[2] = []float64{xtp4, ytp4, ntp}
	for j := 0; j < 3; j++ {
		nor = pkg.NorVec(p[(0+j)%3], p[(1+j)%3], p[(2+j)%3])
		nor_all = append(nor_all, nor)
	}

	yanepoly = append(yanepoly, xtp4, ytp4, ntp) // 屋根上面・三角形４
	yanepoly = append(yanepoly, xo, yo, mtp)
	yanepoly = append(yanepoly, xtp5, ytp5, ntp)
	p[0] = []float64{xtp4, ytp4, ntp}
	p[1] = []float64{xo, yo, mtp}
	p[2] = []float64{xtp5, ytp5, ntp}
	for j := 0; j < 3; j++ {
		nor = pkg.NorVec(p[(0+j)%3], p[(1+j)%3], p[(2+j)%3])
		nor_all = append(nor_all, nor)
	}

	yanepoly = append(yanepoly, xtp5, ytp5, ntp) // 屋根上面・三角形５
	yanepoly = append(yanepoly, xo, yo, mtp)
	yanepoly = append(yanepoly, xtp1, ytp1, ntp)
	p[0] = []float64{xtp5, ytp5, ntp}
	p[1] = []float64{xo, yo, mtp}
	p[2] = []float64{xtp1, ytp1, ntp}
	for j := 0; j < 3; j++ {
		nor = pkg.NorVec(p[(0+j)%3], p[(1+j)%3], p[(2+j)%3])
		nor_all = append(nor_all, nor)
	}

	yanepoly = append(yanepoly, xtp1, ytp1, ntp) // 軒端１・三角形1-1
	yanepoly = append(yanepoly, xo2, yo2, nbt)
	yanepoly = append(yanepoly, xo1, yo1, nbt)
	p[0] = []float64{xtp1, ytp1, ntp}
	p[1] = []float64{xo2, yo2, ntp}
	p[2] = []float64{xo1, yo1, mtp1}
	for j := 0; j < 3; j++ {
		nor = pkg.NorVec(p[(0+j)%3], p[(1+j)%3], p[(2+j)%3])
		nor_all = append(nor_all, nor)
	}

	yanepoly = append(yanepoly, xtp2, ytp2, ntp) // 軒端１・三角形1-2
	yanepoly = append(yanepoly, xo2, yo2, nbt)
	yanepoly = append(yanepoly, xtp1, ytp1, ntp)
	p[0] = []float64{xtp2, ytp2, ntp}
	p[1] = []float64{xo2, yo2, ntp}
	p[2] = []float64{xtp1, ytp1, ntp}
	for j := 0; j < 3; j++ {
		nor = pkg.NorVec(p[(0+j)%3], p[(1+j)%3], p[(2+j)%3])
		nor_all = append(nor_all, nor)
	}

	yanepoly = append(yanepoly, xtp2, ytp2, ntp) // 軒端２・三角形2-1
	yanepoly = append(yanepoly, xo3, yo3, nbt)
	yanepoly = append(yanepoly, xo2, yo2, nbt)
	p[0] = []float64{xtp2, ytp2, ntp}
	p[1] = []float64{xo3, yo3, ntp}
	p[2] = []float64{xo2, yo2, mtp1}
	for j := 0; j < 3; j++ {
		nor = pkg.NorVec(p[(0+j)%3], p[(1+j)%3], p[(2+j)%3])
		nor_all = append(nor_all, nor)
	}

	yanepoly = append(yanepoly, xtp3, ytp3, ntp) // 軒端２・三角形2-2
	yanepoly = append(yanepoly, xo3, yo3, nbt)
	yanepoly = append(yanepoly, xtp2, ytp2, ntp)
	p[0] = []float64{xtp3, ytp3, ntp}
	p[1] = []float64{xo3, yo3, ntp}
	p[2] = []float64{xtp2, ytp2, ntp}
	for j := 0; j < 3; j++ {
		nor = pkg.NorVec(p[(0+j)%3], p[(1+j)%3], p[(2+j)%3])
		nor_all = append(nor_all, nor)
	}

	yanepoly = append(yanepoly, xtp3, ytp3, ntp) // 軒端３・三角形3-1
	yanepoly = append(yanepoly, xo4, yo4, nbt)
	yanepoly = append(yanepoly, xo3, yo3, nbt)
	p[0] = []float64{xtp3, ytp3, ntp}
	p[1] = []float64{xo4, yo4, ntp}
	p[2] = []float64{xo3, yo3, mtp1}
	for j := 0; j < 3; j++ {
		nor = pkg.NorVec(p[(0+j)%3], p[(1+j)%3], p[(2+j)%3])
		nor_all = append(nor_all, nor)
	}

	yanepoly = append(yanepoly, xtp4, ytp4, ntp) // 軒端３・三角形3-2
	yanepoly = append(yanepoly, xo4, yo4, nbt)
	yanepoly = append(yanepoly, xtp3, ytp3, ntp)
	p[0] = []float64{xtp4, ytp4, ntp}
	p[1] = []float64{xo4, yo4, ntp}
	p[2] = []float64{xtp3, ytp3, ntp}
	for j := 0; j < 3; j++ {
		nor = pkg.NorVec(p[(0+j)%3], p[(1+j)%3], p[(2+j)%3])
		nor_all = append(nor_all, nor)
	}

	yanepoly = append(yanepoly, xtp4, ytp4, ntp) // 軒端４・三角形4-1
	yanepoly = append(yanepoly, xo5, yo5, nbt)
	yanepoly = append(yanepoly, xo4, yo4, nbt)
	p[0] = []float64{xtp4, ytp4, ntp}
	p[1] = []float64{xo5, yo5, ntp}
	p[2] = []float64{xo4, yo4, mtp1}
	for j := 0; j < 3; j++ {
		nor = pkg.NorVec(p[(0+j)%3], p[(1+j)%3], p[(2+j)%3])
		nor_all = append(nor_all, nor)
	}

	yanepoly = append(yanepoly, xtp5, ytp5, ntp) // 軒端４・三角形4-2
	yanepoly = append(yanepoly, xo5, yo5, nbt)
	yanepoly = append(yanepoly, xtp4, ytp4, ntp)
	p[0] = []float64{xtp5, ytp5, ntp}
	p[1] = []float64{xo5, yo5, ntp}
	p[2] = []float64{xtp4, ytp4, ntp}
	for j := 0; j < 3; j++ {
		nor = pkg.NorVec(p[(0+j)%3], p[(1+j)%3], p[(2+j)%3])
		nor_all = append(nor_all, nor)
	}

	yanepoly = append(yanepoly, xtp5, ytp5, ntp) // 軒端５・三角形5-1
	yanepoly = append(yanepoly, xo1, yo1, nbt)
	yanepoly = append(yanepoly, xo5, yo5, nbt)
	p[0] = []float64{xtp5, ytp5, ntp}
	p[1] = []float64{xo1, yo1, ntp}
	p[2] = []float64{xo5, yo5, mtp1}
	for j := 0; j < 3; j++ {
		nor = pkg.NorVec(p[(0+j)%3], p[(1+j)%3], p[(2+j)%3])
		nor_all = append(nor_all, nor)
	}

	yanepoly = append(yanepoly, xtp1, ytp1, ntp) // 軒端５・三角形5-2
	yanepoly = append(yanepoly, xo1, yo1, nbt)
	yanepoly = append(yanepoly, xtp5, ytp5, ntp)
	p[0] = []float64{xtp1, ytp1, ntp}
	p[1] = []float64{xo1, yo1, ntp}
	p[2] = []float64{xtp5, ytp5, ntp}
	for j := 0; j < 3; j++ {
		nor = pkg.NorVec(p[(0+j)%3], p[(1+j)%3], p[(2+j)%3])
		nor_all = append(nor_all, nor)
	}

	// 頂点座標法線ベクトルの書き出し
	for j := range nor_all {
		for k := 0; k < 3; k++ {
			normal = append(normal, nor_all[j][k])
		}
	}

	// 頂点座標リストのテキスト化
	for y := range yanepoly {
		yanetxt = append(yanetxt, strconv.FormatFloat(yanepoly[y], 'f', -1, 64))
	}

	// 頂点座標法線ベクトルのテキスト化
	for n := range normal {
		yanenor = append(yanenor, strconv.FormatFloat(normal[n], 'f', -1, 64))
	}

	return yanetxt, yanenor
}
