package internal

import (
	"log"
	"math"
	"stera/pkg"
	"strconv"
)

// YoseYane は寄棟屋根の頂点座標と頂点の法線ベクトルを求める
func YoseYane(list [][]float64, toph, hisashi, keraba, incline,
	yaneatu float64) (yanetxt, yanenor []string) {
	// 屋根モデル（寄棟屋根）
	// 4つの頂点座標の定義
	kansan := 0.0254
	x1 := list[0][0] / kansan
	y1 := list[0][1] / kansan
	x2 := list[1][0] / kansan
	y2 := list[1][1] / kansan
	x3 := list[2][0] / kansan
	y3 := list[2][1] / kansan
	x4 := list[3][0] / kansan
	y4 := list[3][1] / kansan

	// 寄棟屋根の棟方向を妻面と平面の長さを比較してチェックする
	d1 := pkg.DistVerts(x1, y1, x2, y2)
	d2 := pkg.DistVerts(x4, y4, x1, y1)
	log.Println("d1, d2", d1, d2)
	if d1 > d2 {
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

	// インチ換算後のケラバ厚さと庇長さ（軒の出）
	kich := keraba / kansan
	hich := hisashi / kansan

	// 頂点1と頂点2を結ぶ線分に平行な直線の式（妻面）
	m1, n1 := pkg.ParaLine(x1, y1, x2, y2, x3, y3, kich)
	// log.Println("y = " + strconv.FormatFloat(m1, 'f', -1, 64) + "x + " + strconv.FormatFloat(n1, 'f', -1, 64))
	// 頂点3と頂点4を結ぶ線分に平行な直線の式（妻面）
	m2, n2 := pkg.ParaLine(x3, y3, x4, y4, x1, y1, kich)
	// log.Println("y = " + strconv.FormatFloat(m2, 'f', -1, 64) + "x + " + strconv.FormatFloat(n2, 'f', -1, 64))
	// 頂点2と頂点3を結ぶ線分に平行な直線の式（平面）
	m3, n3 := pkg.ParaLine(x2, y2, x3, y3, x4, y4, hich)
	// log.Println("y = " + strconv.FormatFloat(m3, 'f', -1, 64) + "x + " + strconv.FormatFloat(n3, 'f', -1, 64))
	// 頂点4と頂点1を結ぶ線分に平行な直線の式（平面）
	m4, n4 := pkg.ParaLine(x4, y4, x1, y1, x2, y2, hich)
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

	// 軒鼻の突き出し長さ（妻面）
	nt := yaneatu / kansan / math.Sqrt(math.Pow(incline, 2)+1) * incline
	// 妻面から軒庇の上端までの長さ（ケラバ＋軒鼻）
	kichtop := kich + nt
	// 頂点1と頂点2を結ぶ線分に平行な直線の式（妻面）// 軒庇の上端
	mtp1, ntp1 := pkg.ParaLine(x1, y1, x2, y2, x3, y3, kichtop)
	// 頂点3と頂点4を結ぶ線分に平行な直線の式（妻面）// 軒庇の上端
	mtp2, ntp2 := pkg.ParaLine(x3, y3, x4, y4, x1, y1, kichtop)

	// 妻面から軒庇の上端までの長さ（軒庇＋軒鼻）
	hichtop := hich + nt
	// 頂点2と頂点3を結ぶ線分に平行な直線の式（平面）// 軒庇の上端
	mtp3, ntp3 := pkg.ParaLine(x2, y2, x3, y3, x1, y1, hichtop)
	// 頂点4と頂点1を結ぶ線分に平行な直線の式（平面）// 軒庇の上端
	mtp4, ntp4 := pkg.ParaLine(x4, y4, x1, y1, x2, y2, hichtop)

	// 屋根伏せの4頂点の座標を求める（軒庇の上端）
	xtp1, ytp1 := pkg.SeekInsec(mtp4, ntp4, mtp1, ntp1)
	xtp2, ytp2 := pkg.SeekInsec(mtp1, ntp1, mtp3, ntp3)
	xtp3, ytp3 := pkg.SeekInsec(mtp3, ntp3, mtp2, ntp2)
	xtp4, ytp4 := pkg.SeekInsec(mtp2, ntp2, mtp4, ntp4)
	log.Println("xtp1, ytp1, xtp2, ytp2, xtp3, ytp3, xtp4, ytp4", xtp1, ytp1, xtp2, ytp2, xtp3, ytp3, xtp4, ytp4)

	// 屋根の棟端の妻面での座標を求める
	xm1 := (xo1 + xo2) / 2
	ym1 := (yo1 + yo2) / 2
	xm2 := (xo3 + xo4) / 2
	ym2 := (yo3 + yo4) / 2
	// 屋根の棟端の座標を通る直線の方程式
	line := pkg.LineEquat(xm1, ym1, xm2, ym2)
	m := line["m"]
	n := line["n"]
	// 妻面の長さ
	l1 := pkg.DistVerts(x1, y1, x2, y2)
	l2 := pkg.DistVerts(x3, y3, x4, y4)
	// 妻面に平行で棟端となる直線の方程式
	my1, ny1 := pkg.ParaLine2(x1, y1, x2, y2, x3, y3, l1/2)
	my2, ny2 := pkg.ParaLine2(x3, y3, x4, y4, x1, y1, l2/2)
	// ２つの直線の交点から寄棟屋根の棟端の座標を求める
	xy1, yy1 := pkg.SeekInsec(m, n, my1, ny1)
	xy2, yy2 := pkg.SeekInsec(m, n, my2, ny2)

	// 軒先下端高さ（庇×屋根勾配）
	nbt := (toph - hisashi*incline) / kansan
	// 軒先上端高さ
	ntp := nbt + yaneatu/kansan/math.Sqrt(math.Pow(incline, 2)+1)
	log.Println("nbt, ntp", nbt, ntp)

	// 寄棟屋根の棟端の下端高さ
	mbt01 := nbt + l1/2*incline
	mbt02 := nbt + l2/2*incline
	mbt1 := (mbt01 + mbt02) / 2
	mbt2 := (mbt01 + mbt02) / 2
	// 寄棟屋根の棟端の上端高さ
	myp01 := mbt1 + yaneatu/kansan*math.Sqrt(1+math.Pow(incline, 2))
	myp02 := mbt2 + yaneatu/kansan*math.Sqrt(1+math.Pow(incline, 2))
	myp1 := (myp01 + myp02) / 2
	myp2 := (myp01 + myp02) / 2
	log.Println("mbt1, myp1, mbt2, myp2", mbt1, myp1, mbt2, myp2)

	// 屋根モデルの頂点座標をリストに書き出す
	// 屋根頂点の法線ベクトルを算出してリストに書き出す
	var yanepoly []float64
	var nor_all [][]float64
	var normal []float64
	var nor []float64
	p := make([][]float64, 3)

	yanepoly = append(yanepoly, xo1, yo1, nbt) // 屋根底面・三角形１
	yanepoly = append(yanepoly, xy1, yy1, mbt1)
	yanepoly = append(yanepoly, xo2, yo2, nbt)
	p[0] = []float64{xo1, yo1, nbt}
	p[1] = []float64{xy1, yy1, mbt1}
	p[2] = []float64{xo2, yo2, nbt}
	for j := 0; j < 3; j++ {
		nor = pkg.NorVec(p[(0+j)%3], p[(1+j)%3], p[(2+j)%3])
		nor_all = append(nor_all, nor)
	}

	yanepoly = append(yanepoly, xo3, yo3, nbt) // 屋根底面・三角形２
	yanepoly = append(yanepoly, xy2, yy2, mbt2)
	yanepoly = append(yanepoly, xo4, yo4, nbt)
	p[0] = []float64{xo3, yo3, nbt}
	p[1] = []float64{xy2, yy2, mbt1}
	p[2] = []float64{xo4, yo4, nbt}
	for j := 0; j < 3; j++ {
		nor = pkg.NorVec(p[(0+j)%3], p[(1+j)%3], p[(2+j)%3])
		nor_all = append(nor_all, nor)
	}

	yanepoly = append(yanepoly, xo2, yo2, nbt) // 屋根底面・三角形3-1
	yanepoly = append(yanepoly, xy1, yy1, mbt1)
	yanepoly = append(yanepoly, xo3, yo3, nbt)
	p[0] = []float64{xo2, yo2, nbt}
	p[1] = []float64{xy1, yy1, mbt1}
	p[2] = []float64{xo3, yo3, nbt}
	for j := 0; j < 3; j++ {
		nor = pkg.NorVec(p[(0+j)%3], p[(1+j)%3], p[(2+j)%3])
		nor_all = append(nor_all, nor)
	}

	yanepoly = append(yanepoly, xy1, yy1, mbt1) // 屋根底面・三角形3-2
	yanepoly = append(yanepoly, xy2, yy2, mbt2)
	yanepoly = append(yanepoly, xo3, yo3, nbt)
	p[0] = []float64{xy1, yy1, mbt1}
	p[1] = []float64{xy2, yy2, mbt2}
	p[2] = []float64{xo3, yo3, nbt}
	for j := 0; j < 3; j++ {
		nor = pkg.NorVec(p[(0+j)%3], p[(1+j)%3], p[(2+j)%3])
		nor_all = append(nor_all, nor)
	}

	yanepoly = append(yanepoly, xo4, yo4, nbt) // 屋根底面・三角形4-1
	yanepoly = append(yanepoly, xy2, yy2, mbt2)
	yanepoly = append(yanepoly, xo1, yo1, nbt)
	p[0] = []float64{xo4, yo4, nbt}
	p[1] = []float64{xy2, yy2, mbt2}
	p[2] = []float64{xo1, yo1, nbt}
	for j := 0; j < 3; j++ {
		nor = pkg.NorVec(p[(0+j)%3], p[(1+j)%3], p[(2+j)%3])
		nor_all = append(nor_all, nor)
	}

	yanepoly = append(yanepoly, xy2, yy2, mbt2) // 屋根底面・三角形4-2
	yanepoly = append(yanepoly, xy1, yy1, mbt1)
	yanepoly = append(yanepoly, xo1, yo1, nbt)
	p[0] = []float64{xy2, yy2, mbt2}
	p[1] = []float64{xy1, yy1, mbt1}
	p[2] = []float64{xo1, yo1, nbt}
	for j := 0; j < 3; j++ {
		nor = pkg.NorVec(p[(0+j)%3], p[(1+j)%3], p[(2+j)%3])
		nor_all = append(nor_all, nor)
	}

	yanepoly = append(yanepoly, xtp1, ytp1, ntp) // 屋根上面・三角形１
	yanepoly = append(yanepoly, xy1, yy1, myp1)
	yanepoly = append(yanepoly, xtp2, ytp2, ntp)
	p[0] = []float64{xtp1, ytp1, ntp}
	p[1] = []float64{xy1, yy1, myp1}
	p[2] = []float64{xtp2, ytp2, ntp}
	for j := 0; j < 3; j++ {
		nor = pkg.NorVec(p[(0+j)%3], p[(1+j)%3], p[(2+j)%3])
		nor_all = append(nor_all, nor)
	}

	yanepoly = append(yanepoly, xtp3, ytp3, ntp) // 屋根上面・三角形２
	yanepoly = append(yanepoly, xy2, yy2, myp2)
	yanepoly = append(yanepoly, xtp4, ytp4, ntp)
	p[0] = []float64{xtp3, ytp3, ntp}
	p[1] = []float64{xy2, yy2, myp1}
	p[2] = []float64{xtp4, ytp4, ntp}
	for j := 0; j < 3; j++ {
		nor = pkg.NorVec(p[(0+j)%3], p[(1+j)%3], p[(2+j)%3])
		nor_all = append(nor_all, nor)
	}

	yanepoly = append(yanepoly, xtp2, ytp2, ntp) // 屋根上面・三角形3-1
	yanepoly = append(yanepoly, xy1, yy1, myp1)
	yanepoly = append(yanepoly, xtp3, ytp3, ntp)
	p[0] = []float64{xtp2, ytp2, ntp}
	p[1] = []float64{xy1, yy1, myp1}
	p[2] = []float64{xtp3, ytp3, ntp}
	for j := 0; j < 3; j++ {
		nor = pkg.NorVec(p[(0+j)%3], p[(1+j)%3], p[(2+j)%3])
		nor_all = append(nor_all, nor)
	}

	yanepoly = append(yanepoly, xy1, yy1, myp1) // 屋根上面・三角形3-2
	yanepoly = append(yanepoly, xy2, yy2, myp2)
	yanepoly = append(yanepoly, xtp3, ytp3, ntp)
	p[0] = []float64{xy1, yy1, myp1}
	p[1] = []float64{xy2, yy2, myp2}
	p[2] = []float64{xtp3, ytp3, ntp}
	for j := 0; j < 3; j++ {
		nor = pkg.NorVec(p[(0+j)%3], p[(1+j)%3], p[(2+j)%3])
		nor_all = append(nor_all, nor)
	}

	yanepoly = append(yanepoly, xtp4, ytp4, ntp) // 屋根上面・三角形4-1
	yanepoly = append(yanepoly, xy2, yy2, myp2)
	yanepoly = append(yanepoly, xtp1, ytp1, ntp)
	p[0] = []float64{xtp4, ytp4, ntp}
	p[1] = []float64{xy2, yy2, myp2}
	p[2] = []float64{xtp1, ytp1, ntp}
	for j := 0; j < 3; j++ {
		nor = pkg.NorVec(p[(0+j)%3], p[(1+j)%3], p[(2+j)%3])
		nor_all = append(nor_all, nor)
	}

	yanepoly = append(yanepoly, xy2, yy2, myp2) // 屋根上面・三角形4-2
	yanepoly = append(yanepoly, xy1, yy1, myp1)
	yanepoly = append(yanepoly, xtp1, ytp1, ntp)
	p[0] = []float64{xy2, yy2, myp2}
	p[1] = []float64{xy1, yy1, myp1}
	p[2] = []float64{xtp1, ytp1, ntp}
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

	yanepoly = append(yanepoly, xo3, yo3, nbt) //  軒端・三角形2-1
	yanepoly = append(yanepoly, xo4, yo4, nbt)
	yanepoly = append(yanepoly, xtp3, ytp3, ntp)
	p[0] = []float64{xo3, yo3, nbt}
	p[1] = []float64{xo4, yo4, nbt}
	p[2] = []float64{xtp3, ytp3, ntp}
	for j := 0; j < 3; j++ {
		nor = pkg.NorVec(p[(0+j)%3], p[(1+j)%3], p[(2+j)%3])
		nor_all = append(nor_all, nor)
	}

	yanepoly = append(yanepoly, xo4, yo4, nbt) //  軒端・三角形2-2
	yanepoly = append(yanepoly, xtp3, ytp3, ntp)
	yanepoly = append(yanepoly, xtp4, ytp4, ntp)
	p[0] = []float64{xo4, yo4, nbt}
	p[1] = []float64{xtp3, ytp3, ntp}
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
	yanepoly = append(yanepoly, xtp2, ytp2, ntp)
	yanepoly = append(yanepoly, xtp3, ytp3, ntp)
	p[0] = []float64{xo3, yo3, nbt}
	p[1] = []float64{xtp2, ytp2, ntp}
	p[2] = []float64{xtp3, ytp3, ntp}
	for j := 0; j < 3; j++ {
		nor = pkg.NorVec(p[(0+j)%3], p[(1+j)%3], p[(2+j)%3])
		nor_all = append(nor_all, nor)
	}

	yanepoly = append(yanepoly, xo4, yo4, nbt) //  軒端・三角形4-1
	yanepoly = append(yanepoly, xo1, yo1, nbt)
	yanepoly = append(yanepoly, xtp4, ytp4, ntp)
	p[0] = []float64{xo4, yo4, nbt}
	p[1] = []float64{xo1, yo1, nbt}
	p[2] = []float64{xtp4, ytp4, ntp}
	for j := 0; j < 3; j++ {
		nor = pkg.NorVec(p[(0+j)%3], p[(1+j)%3], p[(2+j)%3])
		nor_all = append(nor_all, nor)
	}

	yanepoly = append(yanepoly, xo1, yo1, nbt) //  軒端・三角形4-2
	yanepoly = append(yanepoly, xtp4, ytp4, ntp)
	yanepoly = append(yanepoly, xtp1, ytp1, ntp)
	p[0] = []float64{xo1, yo1, nbt}
	p[1] = []float64{xtp4, ytp4, ntp}
	p[2] = []float64{xtp1, ytp1, ntp}
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
