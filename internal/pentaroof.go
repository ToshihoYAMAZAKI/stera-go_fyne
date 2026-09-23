package internal

import (
	"log"
	"math"
	"stera/pkg"
	"strconv"
)

// PentaYane は５角形屋根の頂点座標と頂点の法線ベクトルを求める
func PentaYane(list [][]float64, toph, hisashi, keraba, incline,
	yaneatu float64) (yanetxt, yanenor []string) {
	// 屋根モデル（５角形屋根）
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
	xc := (x1 + x5) / 2
	yc := (y1 + y5) / 2

	// ５角形屋根の棟方向を妻面と平面の長さを比較してチェックする
	d1 := pkg.DistVerts(x5, y5, x1, y1)
	d2 := pkg.DistVerts(x3, y3, xc, yc)
	log.Println("d1, d2", d1, d2)

	// インチ換算後のケラバ厚さと庇長さ（軒の出）
	kich := keraba / kansan
	hich := hisashi / kansan

	// 壁面上端の高さ
	wh := toph / kansan

	// 頂点1と頂点2を結ぶ線分に平行な直線の式（平面）
	m1, n1 := pkg.ParaLine(x1, y1, x2, y2, x5, y5, hich)
	// log.Println("y = " + strconv.FormatFloat(m1, 'f', -1, 64) + "x + " + strconv.FormatFloat(n1, 'f', -1, 64))
	// 頂点4と頂点5を結ぶ線分に平行な直線の式（平面）
	m2, n2 := pkg.ParaLine(x4, y4, x5, y5, x1, y1, hich)
	// log.Println("y = " + strconv.FormatFloat(m2, 'f', -1, 64) + "x + " + strconv.FormatFloat(n2, 'f', -1, 64))
	// 頂点2と頂点3を結ぶ線分に平行な直線の式（正面1）
	m3, n3 := pkg.ParaLine(x2, y2, x3, y3, x5, y5, kich)
	// log.Println("y = " + strconv.FormatFloat(m3, 'f', -1, 64) + "x + " + strconv.FormatFloat(n3, 'f', -1, 64))
	// 頂点3と頂点4を結ぶ線分に平行な直線の式（正面2）
	m4, n4 := pkg.ParaLine(x3, y3, x4, y4, x1, y1, kich)
	// log.Println("y = " + strconv.FormatFloat(m4, 'f', -1, 64) + "x + " + strconv.FormatFloat(n4, 'f', -1, 64))
	// 頂点5と頂点1を結ぶ線分に平行な直線の式（妻面）
	m5, n5 := pkg.ParaLine(x5, y5, x1, y1, x3, y3, kich)
	// log.Println("y = " + strconv.FormatFloat(m4, 'f', -1, 64) + "x + " + strconv.FormatFloat(n4, 'f', -1, 64))
	log.Println("m1, n1=", m1, n1)
	log.Println("m2, n2=", m2, n2)
	log.Println("m3, n3=", m3, n3)
	log.Println("m4, n4=", m4, n4)
	log.Println("m5, n5=", m5, n5)

	// 屋根伏せの5頂点の座標を求める（軒庇の下端）
	xo1, yo1 := pkg.SeekInsec(m5, n5, m1, n1)
	xo2, yo2 := pkg.SeekInsec(m1, n1, m3, n3)
	xo3, yo3 := pkg.SeekInsec(m3, n3, m4, n4)
	xo4, yo4 := pkg.SeekInsec(m4, n4, m2, n2)
	xo5, yo5 := pkg.SeekInsec(m2, n2, m5, n5)
	log.Println("xo1, yo1, xo2, yo2, xo3, yo3, xo4, yo4, xo5, yo5", xo1, yo1, xo2, yo2, xo3, yo3, xo4, yo4, xo5, yo5)

	// 軒鼻の突き出し長さ（平面）
	nt := yaneatu / kansan / math.Sqrt(math.Pow(incline, 2)+1) * incline
	// 平面から軒庇の上端までの長さ（軒庇＋軒鼻）
	hichtop := hich + nt
	// 頂点1と頂点2を結ぶ線分に平行な直線の式（平面）// 軒庇の上端
	mtp1, ntp1 := pkg.ParaLine(x1, y1, x2, y2, x5, y5, hichtop)
	// 頂点4と頂点5を結ぶ線分に平行な直線の式（平面）// 軒庇の上端
	mtp2, ntp2 := pkg.ParaLine(x4, y4, x5, y5, x1, y1, hichtop)

	// 正面から軒庇の上端までの長さ（ケラバ＋軒鼻）
	kichtop := kich + nt
	// 頂点2と頂点3を結ぶ線分に平行な直線の式（正面）// 軒庇の上端
	mtp3, ntp3 := pkg.ParaLine(x2, y2, x3, y3, x5, y5, kichtop)
	// 頂点3と頂点4を結ぶ線分に平行な直線の式（正面）// 軒庇の上端
	mtp4, ntp4 := pkg.ParaLine(x3, y3, x4, y4, x1, y1, kichtop)

	// 屋根伏せの5頂点の座標を求める（軒庇の上端）
	xtp1, ytp1 := pkg.SeekInsec(m5, n5, mtp1, ntp1)
	xtp2, ytp2 := pkg.SeekInsec(mtp1, ntp1, mtp3, ntp3)
	xtp3, ytp3 := pkg.SeekInsec(mtp3, ntp3, mtp4, ntp4)
	xtp4, ytp4 := pkg.SeekInsec(mtp4, ntp4, mtp2, ntp2)
	xtp5, ytp5 := pkg.SeekInsec(mtp2, ntp2, m5, n5)
	log.Println("xtp1, ytp1, xtp2, ytp2, xtp3, ytp3, xtp4, ytp4, xtp5, ytp5", xtp1, ytp1, xtp2, ytp2, xtp3, ytp3, xtp4, ytp4, xtp5, ytp5)

	// 屋根の棟端の正面/妻面での座標を求める
	// xm0 := (xo2 + xo4) / 2
	// ym0 := (yo2 + yo4) / 2
	xm0 := xo3
	ym0 := yo3
	xm2 := (xo5 + xo1) / 2
	ym2 := (yo5 + yo1) / 2
	// 屋根の棟端の座標を通る直線の方程式
	line := pkg.LineEquat(xm0, ym0, xm2, ym2)
	m := line["m"]
	n := line["n"]
	// 正面の突き出し長さ
	l1 := pkg.DistVerts(x3, y3, xm0, ym0)
	log.Println("xm0=", xm0)
	log.Println("ym0=", ym0)
	log.Println("l1=", l1)
	// 正面に平行で棟端となる直線の方程式
	my1, ny1 := pkg.ParaLine2(x2, y2, x4, y4, x3, y3, l1/2)
	// ２つの直線の交点から５角形屋根の棟端の座標を求める
	xy1, yy1 := pkg.SeekInsec(m, n, my1, ny1)
	xy2, yy2 := pkg.SeekInsec(m, n, m5, n5)

	// 軒先下端高さ（庇×屋根勾配）
	nbt := (toph - hisashi*incline) / kansan
	// 軒先上端高さ
	ntp := nbt + yaneatu/kansan/math.Sqrt(math.Pow(incline, 2)+1)
	log.Println("nbt, ntp", nbt, ntp)

	// ５角形屋根の棟端の下端高さ
	mbt1 := nbt + d1/2*incline
	mbt2 := mbt1
	// ５角形屋根の棟端の上端高さ
	myp1 := mbt1 + yaneatu/kansan*math.Sqrt(1+math.Pow(incline, 2))
	myp2 := myp1
	log.Println("mbt1, myp1, mbt2, myp2", mbt1, myp1, mbt2, myp2)

	// 屋根モデルの頂点座標をリストに書き出す
	// 屋根頂点の法線ベクトルを算出してリストに書き出す
	var yanepoly []float64
	var nor_all [][]float64
	var normal []float64
	var nor []float64
	p := make([][]float64, 3)

	yanepoly = append(yanepoly, xo1, yo1, nbt) // 屋根底面・三角形1-1
	yanepoly = append(yanepoly, xy1, yy1, mbt1)
	yanepoly = append(yanepoly, xo2, yo2, nbt)
	p[0] = []float64{xo1, yo1, nbt}
	p[1] = []float64{xy1, yy1, mbt1}
	p[2] = []float64{xo2, yo2, nbt}
	for j := 0; j < 3; j++ {
		nor = pkg.NorVec(p[(0+j)%3], p[(1+j)%3], p[(2+j)%3])
		nor_all = append(nor_all, nor)
	}

	yanepoly = append(yanepoly, xo1, yo1, nbt) // 屋根底面・三角形1-2
	yanepoly = append(yanepoly, xy2, yy2, mbt2)
	yanepoly = append(yanepoly, xy1, yy1, mbt1)
	p[0] = []float64{xo1, yo1, nbt}
	p[1] = []float64{xy2, yy2, mbt2}
	p[2] = []float64{xy1, yy1, mbt1}
	for j := 0; j < 3; j++ {
		nor = pkg.NorVec(p[(0+j)%3], p[(1+j)%3], p[(2+j)%3])
		nor_all = append(nor_all, nor)
	}

	yanepoly = append(yanepoly, xo5, yo5, nbt) // 屋根底面・三角形2-1
	yanepoly = append(yanepoly, xo4, yo4, nbt)
	yanepoly = append(yanepoly, xy1, yy1, mbt1)
	p[0] = []float64{xo5, yo5, nbt}
	p[1] = []float64{xo4, yo4, nbt}
	p[2] = []float64{xy1, yy1, mbt1}
	for j := 0; j < 3; j++ {
		nor = pkg.NorVec(p[(0+j)%3], p[(1+j)%3], p[(2+j)%3])
		nor_all = append(nor_all, nor)
	}

	yanepoly = append(yanepoly, xo5, yo5, nbt) // 屋根底面・三角形2-2
	yanepoly = append(yanepoly, xy1, yy1, mbt1)
	yanepoly = append(yanepoly, xy2, yy2, mbt2)
	p[0] = []float64{xo5, yo5, nbt}
	p[1] = []float64{xy1, yy1, mbt1}
	p[2] = []float64{xy2, yy2, mbt2}
	for j := 0; j < 3; j++ {
		nor = pkg.NorVec(p[(0+j)%3], p[(1+j)%3], p[(2+j)%3])
		nor_all = append(nor_all, nor)
	}

	yanepoly = append(yanepoly, xy1, yy1, mbt1) // 屋根底面・三角形3-1
	yanepoly = append(yanepoly, xo3, yo3, nbt)
	yanepoly = append(yanepoly, xo2, yo2, nbt)
	p[0] = []float64{xy1, yy1, mbt1}
	p[1] = []float64{xo3, yo3, nbt}
	p[2] = []float64{xo2, yo2, nbt}
	for j := 0; j < 3; j++ {
		nor = pkg.NorVec(p[(0+j)%3], p[(1+j)%3], p[(2+j)%3])
		nor_all = append(nor_all, nor)
	}

	yanepoly = append(yanepoly, xy1, yy1, mbt1) // 屋根底面・三角形3-2
	yanepoly = append(yanepoly, xo4, yo4, nbt)
	yanepoly = append(yanepoly, xo3, yo3, nbt)
	p[0] = []float64{xy1, yy1, mbt1}
	p[1] = []float64{xo4, yo4, nbt}
	p[2] = []float64{xo3, yo3, nbt}
	for j := 0; j < 3; j++ {
		nor = pkg.NorVec(p[(0+j)%3], p[(1+j)%3], p[(2+j)%3])
		nor_all = append(nor_all, nor)
	}

	yanepoly = append(yanepoly, xtp1, ytp1, ntp) // 屋根上面・三角形1-1
	yanepoly = append(yanepoly, xy1, yy1, myp1)
	yanepoly = append(yanepoly, xtp2, ytp2, ntp)
	p[0] = []float64{xtp1, ytp1, ntp}
	p[1] = []float64{xy1, yy1, myp1}
	p[2] = []float64{xtp2, ytp2, ntp}
	for j := 0; j < 3; j++ {
		nor = pkg.NorVec(p[(0+j)%3], p[(1+j)%3], p[(2+j)%3])
		nor_all = append(nor_all, nor)
	}

	yanepoly = append(yanepoly, xtp1, ytp1, ntp) // 屋根上面・三角形1-2
	yanepoly = append(yanepoly, xy2, yy2, myp2)
	yanepoly = append(yanepoly, xy1, yy1, myp1)
	p[0] = []float64{xtp1, ytp1, ntp}
	p[1] = []float64{xy2, yy2, myp2}
	p[2] = []float64{xy1, yy1, myp1}
	for j := 0; j < 3; j++ {
		nor = pkg.NorVec(p[(0+j)%3], p[(1+j)%3], p[(2+j)%3])
		nor_all = append(nor_all, nor)
	}

	yanepoly = append(yanepoly, xtp5, ytp5, ntp) // 屋根上面・三角形2-1
	yanepoly = append(yanepoly, xtp4, ytp4, ntp)
	yanepoly = append(yanepoly, xy1, yy1, myp1)
	p[0] = []float64{xtp5, ytp5, ntp}
	p[1] = []float64{xtp4, ytp4, ntp}
	p[2] = []float64{xy1, yy1, myp1}
	for j := 0; j < 3; j++ {
		nor = pkg.NorVec(p[(0+j)%3], p[(1+j)%3], p[(2+j)%3])
		nor_all = append(nor_all, nor)
	}

	yanepoly = append(yanepoly, xtp5, ytp5, ntp) // 屋根上面・三角形2-2
	yanepoly = append(yanepoly, xy1, yy1, myp1)
	yanepoly = append(yanepoly, xy2, yy2, myp2)
	p[0] = []float64{xtp5, ytp5, ntp}
	p[1] = []float64{xy1, yy1, myp1}
	p[2] = []float64{xy2, yy2, myp2}
	for j := 0; j < 3; j++ {
		nor = pkg.NorVec(p[(0+j)%3], p[(1+j)%3], p[(2+j)%3])
		nor_all = append(nor_all, nor)
	}

	yanepoly = append(yanepoly, xy1, yy1, myp1) // 屋根上面・三角形3-1
	yanepoly = append(yanepoly, xtp3, ytp3, ntp)
	yanepoly = append(yanepoly, xtp2, ytp2, ntp)
	p[0] = []float64{xy1, yy1, myp1}
	p[1] = []float64{xtp3, ytp3, ntp}
	p[2] = []float64{xtp2, ytp2, ntp}
	for j := 0; j < 3; j++ {
		nor = pkg.NorVec(p[(0+j)%3], p[(1+j)%3], p[(2+j)%3])
		nor_all = append(nor_all, nor)
	}

	yanepoly = append(yanepoly, xy1, yy1, myp1) // 屋根上面・三角形3-2
	yanepoly = append(yanepoly, xtp4, ytp4, ntp)
	yanepoly = append(yanepoly, xtp3, ytp3, ntp)
	p[0] = []float64{xy1, yy1, myp1}
	p[1] = []float64{xtp4, ytp4, ntp}
	p[2] = []float64{xtp3, ytp3, ntp}
	for j := 0; j < 3; j++ {
		nor = pkg.NorVec(p[(0+j)%3], p[(1+j)%3], p[(2+j)%3])
		nor_all = append(nor_all, nor)
	}

	yanepoly = append(yanepoly, xo1, yo1, nbt) //  軒端・三角形1-1
	yanepoly = append(yanepoly, xo2, yo2, nbt)
	yanepoly = append(yanepoly, xtp1, ytp1, ntp)
	p[0] = []float64{xo1, yo1, nbt}
	p[1] = []float64{xo2, yo2, nbt}
	p[2] = []float64{xtp1, ytp1, ntp}
	for j := 0; j < 3; j++ {
		nor = pkg.NorVec(p[(0+j)%3], p[(1+j)%3], p[(2+j)%3])
		nor_all = append(nor_all, nor)
	}

	yanepoly = append(yanepoly, xo2, yo2, nbt) //  軒端・三角形1-2
	yanepoly = append(yanepoly, xtp1, ytp1, ntp)
	yanepoly = append(yanepoly, xtp2, ytp2, ntp)
	p[0] = []float64{xo2, yo2, nbt}
	p[1] = []float64{xtp1, ytp1, ntp}
	p[2] = []float64{xtp2, ytp2, ntp}
	for j := 0; j < 3; j++ {
		nor = pkg.NorVec(p[(0+j)%3], p[(1+j)%3], p[(2+j)%3])
		nor_all = append(nor_all, nor)
	}

	yanepoly = append(yanepoly, xo5, yo5, nbt) //  軒端・三角形2-1
	yanepoly = append(yanepoly, xtp5, ytp5, ntp)
	yanepoly = append(yanepoly, xo4, yo4, nbt)
	p[0] = []float64{xo5, yo5, nbt}
	p[1] = []float64{xtp5, ytp5, ntp}
	p[2] = []float64{xo4, yo4, nbt}
	for j := 0; j < 3; j++ {
		nor = pkg.NorVec(p[(0+j)%3], p[(1+j)%3], p[(2+j)%3])
		nor_all = append(nor_all, nor)
	}

	yanepoly = append(yanepoly, xo4, yo4, nbt) //  軒端・三角形2-2
	yanepoly = append(yanepoly, xtp5, ytp5, ntp)
	yanepoly = append(yanepoly, xtp4, ytp4, ntp)
	p[0] = []float64{xo4, yo4, nbt}
	p[1] = []float64{xtp5, ytp5, ntp}
	p[2] = []float64{xtp4, ytp4, ntp}
	for j := 0; j < 3; j++ {
		nor = pkg.NorVec(p[(0+j)%3], p[(1+j)%3], p[(2+j)%3])
		nor_all = append(nor_all, nor)
	}

	yanepoly = append(yanepoly, xo2, yo2, nbt) //  軒端・三角形3-1
	yanepoly = append(yanepoly, xo3, yo3, nbt)
	yanepoly = append(yanepoly, xtp2, ytp2, ntp)
	p[0] = []float64{xo2, yo2, nbt}
	p[1] = []float64{xo3, yo3, nbt}
	p[2] = []float64{xtp2, ytp2, ntp}
	for j := 0; j < 3; j++ {
		nor = pkg.NorVec(p[(0+j)%3], p[(1+j)%3], p[(2+j)%3])
		nor_all = append(nor_all, nor)
	}

	yanepoly = append(yanepoly, xo3, yo3, nbt) //  軒端・三角形3-2
	yanepoly = append(yanepoly, xtp3, ytp3, ntp)
	yanepoly = append(yanepoly, xtp2, ytp2, ntp)
	p[0] = []float64{xo3, yo3, nbt}
	p[1] = []float64{xtp3, ytp3, ntp}
	p[2] = []float64{xtp2, ytp2, ntp}
	for j := 0; j < 3; j++ {
		nor = pkg.NorVec(p[(0+j)%3], p[(1+j)%3], p[(2+j)%3])
		nor_all = append(nor_all, nor)
	}

	yanepoly = append(yanepoly, xo3, yo3, nbt) //  軒端・三角形4-1
	yanepoly = append(yanepoly, xtp4, ytp4, ntp)
	yanepoly = append(yanepoly, xtp3, ytp3, ntp)
	p[0] = []float64{xo3, yo3, nbt}
	p[1] = []float64{xtp4, ytp4, ntp}
	p[2] = []float64{xtp3, ytp3, ntp}
	for j := 0; j < 3; j++ {
		nor = pkg.NorVec(p[(0+j)%3], p[(1+j)%3], p[(2+j)%3])
		nor_all = append(nor_all, nor)
	}

	yanepoly = append(yanepoly, xo3, yo3, nbt) //  軒端・三角形4-2
	yanepoly = append(yanepoly, xo4, yo4, nbt)
	yanepoly = append(yanepoly, xtp4, ytp4, ntp)
	p[0] = []float64{xo3, yo3, nbt}
	p[1] = []float64{xo4, yo4, nbt}
	p[2] = []float64{xtp4, ytp4, ntp}
	for j := 0; j < 3; j++ {
		nor = pkg.NorVec(p[(0+j)%3], p[(1+j)%3], p[(2+j)%3])
		nor_all = append(nor_all, nor)
	}

	yanepoly = append(yanepoly, xo1, yo1, nbt) //  妻面軒端・三角形1-1
	yanepoly = append(yanepoly, xtp1, ytp1, ntp)
	yanepoly = append(yanepoly, xy2, yy2, mbt2)
	p[0] = []float64{xo1, yo1, nbt}
	p[1] = []float64{xtp1, ytp1, ntp}
	p[2] = []float64{xy2, yy2, mbt2}
	for j := 0; j < 3; j++ {
		nor = pkg.NorVec(p[(0+j)%3], p[(1+j)%3], p[(2+j)%3])
		nor_all = append(nor_all, nor)
	}

	yanepoly = append(yanepoly, xy2, yy2, myp2) //  妻面軒端・三角形1-2
	yanepoly = append(yanepoly, xy2, yy2, mbt2)
	yanepoly = append(yanepoly, xtp1, ytp1, ntp)
	p[0] = []float64{xy2, yy2, myp2}
	p[1] = []float64{xy2, yy2, mbt2}
	p[2] = []float64{xtp1, ytp1, ntp}
	for j := 0; j < 3; j++ {
		nor = pkg.NorVec(p[(0+j)%3], p[(1+j)%3], p[(2+j)%3])
		nor_all = append(nor_all, nor)
	}

	yanepoly = append(yanepoly, xy2, yy2, mbt2) //  妻面軒端・三角形2-1
	yanepoly = append(yanepoly, xy2, yy2, myp2)
	yanepoly = append(yanepoly, xtp5, ytp5, ntp)
	p[0] = []float64{xy2, yy2, mbt2}
	p[1] = []float64{xy2, yy2, myp2}
	p[2] = []float64{xtp5, ytp5, ntp}
	for j := 0; j < 3; j++ {
		nor = pkg.NorVec(p[(0+j)%3], p[(1+j)%3], p[(2+j)%3])
		nor_all = append(nor_all, nor)
	}

	yanepoly = append(yanepoly, xtp5, ytp5, ntp) //  妻面軒端・三角形2-2
	yanepoly = append(yanepoly, xo5, yo5, nbt)
	yanepoly = append(yanepoly, xy2, yy2, mbt2)
	p[0] = []float64{xtp5, ytp5, ntp}
	p[1] = []float64{xo5, yo5, nbt}
	p[2] = []float64{xy2, yy2, mbt2}
	for j := 0; j < 3; j++ {
		nor = pkg.NorVec(p[(0+j)%3], p[(1+j)%3], p[(2+j)%3])
		nor_all = append(nor_all, nor)
	}

	yanepoly = append(yanepoly, x1, y1, wh) // 妻面・三角形
	yanepoly = append(yanepoly, xc, yc, mbt2)
	yanepoly = append(yanepoly, x5, y5, wh)
	p[0] = []float64{x1, y1, wh}
	p[1] = []float64{xo2, yo2, mbt2}
	p[2] = []float64{x5, y5, wh}
	for j := 0; j < 3; j++ {
		nor = pkg.NorVec(p[(0+j)%3], p[(1+j)%3], p[(2+j)%3])
		nor_all = append(nor_all, nor)
	}

	log.Println("nor_all=", nor_all)

	// 頂点座標法線ベクトルの書き出し
	for j := range nor_all {
		for k := 0; k < 3; k++ {
			normal = append(normal, nor_all[j][k])
		}
	}

	log.Println("yanepoly=", yanepoly)
	log.Println("normal=", normal)

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
