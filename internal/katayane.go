package internal

import (
	"log"
	"math"
	"stera/pkg"
	"strconv"
)

// KataYane は片流れ屋根の頂点座標と頂点の法線ベクトルを求める
func KataYane(list [][]float64, toph, hisashi, keraba, incline,
	yaneatu float64) (yanetxt, yanenor []string) {
	log.Println("片流れ屋根")
	// 屋根モデル（片流れ屋根）
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

	// インチ換算後のケラバ厚さと庇長さ（軒の出）
	kich := keraba / kansan
	hich := hisashi / kansan

	// 頂点1と頂点2を結ぶ線分に平行な直線の式（流れ面上手）
	m1, n1 := pkg.ParaLine(x1, y1, x2, y2, x3, y3, hich/4)
	// log.Println("y = " + strconv.FormatFloat(m1, 'f', -1, 64) + "x + " + strconv.FormatFloat(n1, 'f', -1, 64))
	// 頂点3と頂点4を結ぶ線分に平行な直線の式（流れ面下手）
	m2, n2 := pkg.ParaLine(x3, y3, x4, y4, x1, y1, hich)
	// log.Println("y = " + strconv.FormatFloat(m2, 'f', -1, 64) + "x + " + strconv.FormatFloat(n2, 'f', -1, 64))
	// 頂点2と頂点3を結ぶ線分に平行な直線の式（側面）
	m3, n3 := pkg.ParaLine(x2, y2, x3, y3, x4, y4, kich)
	// log.Println("y = " + strconv.FormatFloat(m3, 'f', -1, 64) + "x + " + strconv.FormatFloat(n3, 'f', -1, 64))
	// 頂点4と頂点1を結ぶ線分に平行な直線の式（側面）
	m4, n4 := pkg.ParaLine(x4, y4, x1, y1, x2, y2, kich)
	// log.Println("y = " + strconv.FormatFloat(m4, 'f', -1, 64) + "x + " + strconv.FormatFloat(n4, 'f', -1, 64))

	// 屋根伏せの4頂点の座標を求める（屋根面の下端）
	xo1, yo1 := pkg.SeekInsec(m4, n4, m1, n1)
	xo2, yo2 := pkg.SeekInsec(m1, n1, m3, n3)
	xo3, yo3 := pkg.SeekInsec(m3, n3, m2, n2)
	xo4, yo4 := pkg.SeekInsec(m2, n2, m4, n4)
	log.Println("xo1, yo1, xo2, yo2, xo3, yo3, xo4, yo4", xo1, yo1, xo2, yo2, xo3, yo3, xo4, yo4)

	// 軒鼻の突き出し長さ
	nt := yaneatu / kansan / math.Sqrt(math.Pow(incline, 2)+1) * incline
	// 壁面から軒庇の上端までの長さ（軒庇＋軒鼻）
	hichtop := hich + nt
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
	nbt := (toph - hisashi/math.Sqrt(math.Pow(incline, 2)+1)*incline) / kansan
	// nbt := (toph - hisashi*incline) / kansan
	// 軒先上端高さ
	ntp := nbt + yaneatu/kansan/math.Sqrt(math.Pow(incline, 2)+1)
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
	// mtp1 := mbt1 + yaneatu/kansan*math.Sqrt(1+math.Pow(incline, 2))
	// mtp2 := mbt2 + yaneatu/kansan*math.Sqrt(1+math.Pow(incline, 2))
	mtp := mbt + yaneatu/kansan*math.Sqrt(1+math.Pow(incline, 2))
	log.Println("mbt, mtp", mbt, mtp)

	// 屋根モデルの頂点座標をリストに書き出す
	// 屋根頂点の法線ベクトルを算出してリストに書き出す
	var yanepoly []float64
	var nor_all [][]float64
	var normal []float64
	var nor []float64
	p := make([][]float64, 3)

	yanepoly = append(yanepoly, xo1, yo1, mbt) // 屋根底面・三角形１
	yanepoly = append(yanepoly, xo4, yo4, nbt)
	yanepoly = append(yanepoly, xo2, yo2, mbt)
	p[0] = []float64{xo1, yo1, mbt}
	p[1] = []float64{xo4, yo4, nbt}
	p[2] = []float64{xo2, yo2, mbt}
	for j := 0; j < 3; j++ {
		nor = pkg.NorVec(p[(0+j)%3], p[(1+j)%3], p[(2+j)%3])
		nor_all = append(nor_all, nor)
	}

	yanepoly = append(yanepoly, xo4, yo4, nbt) // 屋根底面・三角形２
	yanepoly = append(yanepoly, xo3, yo3, nbt)
	yanepoly = append(yanepoly, xo2, yo2, mbt)
	p[0] = []float64{xo4, yo4, nbt}
	p[1] = []float64{xo3, yo3, nbt}
	p[2] = []float64{xo2, yo2, mbt}
	for j := 0; j < 3; j++ {
		nor = pkg.NorVec(p[(0+j)%3], p[(1+j)%3], p[(2+j)%3])
		nor_all = append(nor_all, nor)
	}

	yanepoly = append(yanepoly, xtp1, ytp1, mtp) // 屋根上面・三角形１
	yanepoly = append(yanepoly, xtp4, ytp4, ntp)
	yanepoly = append(yanepoly, xtp2, ytp2, mtp)
	p[0] = []float64{xtp1, ytp1, mtp}
	p[1] = []float64{xtp4, ytp4, ntp}
	p[2] = []float64{xtp2, ytp2, mtp}
	for j := 0; j < 3; j++ {
		nor = pkg.NorVec(p[(0+j)%3], p[(1+j)%3], p[(2+j)%3])
		nor_all = append(nor_all, nor)
	}

	yanepoly = append(yanepoly, xtp4, ytp4, ntp) // 屋根上面・三角形２
	yanepoly = append(yanepoly, xtp3, ytp3, ntp)
	yanepoly = append(yanepoly, xtp2, ytp2, mtp)
	p[0] = []float64{xtp4, ytp4, ntp}
	p[1] = []float64{xtp3, ytp3, ntp}
	p[2] = []float64{xtp2, ytp2, mtp}
	for j := 0; j < 3; j++ {
		nor = pkg.NorVec(p[(0+j)%3], p[(1+j)%3], p[(2+j)%3])
		nor_all = append(nor_all, nor)
	}

	yanepoly = append(yanepoly, xtp4, ytp4, ntp) // 軒端１・三角形-1
	yanepoly = append(yanepoly, xtp3, ytp3, ntp)
	yanepoly = append(yanepoly, xo3, yo3, nbt)
	p[0] = []float64{xtp4, ytp4, ntp}
	p[1] = []float64{xtp3, ytp3, ntp}
	p[2] = []float64{xo3, yo3, nbt}
	for j := 0; j < 3; j++ {
		nor = pkg.NorVec(p[(0+j)%3], p[(1+j)%3], p[(2+j)%3])
		nor_all = append(nor_all, nor)
	}

	yanepoly = append(yanepoly, xtp4, ytp4, ntp) // 軒端２・三角形-2
	yanepoly = append(yanepoly, xo3, yo3, nbt)
	yanepoly = append(yanepoly, xo4, yo4, nbt)
	p[0] = []float64{xtp4, ytp4, ntp}
	p[1] = []float64{xo3, yo3, nbt}
	p[2] = []float64{xo4, yo4, nbt}
	for j := 0; j < 3; j++ {
		nor = pkg.NorVec(p[(0+j)%3], p[(1+j)%3], p[(2+j)%3])
		nor_all = append(nor_all, nor)
	}

	yanepoly = append(yanepoly, xtp1, ytp1, mtp) // 棟端１・三角形-1
	yanepoly = append(yanepoly, xo1, yo1, mbt)
	yanepoly = append(yanepoly, xo2, yo2, mbt)
	p[0] = []float64{xtp1, ytp1, mtp}
	p[1] = []float64{xo1, yo1, mbt}
	p[2] = []float64{xo2, yo2, mbt}
	for j := 0; j < 3; j++ {
		nor = pkg.NorVec(p[(0+j)%3], p[(1+j)%3], p[(2+j)%3])
		nor_all = append(nor_all, nor)
	}

	yanepoly = append(yanepoly, xtp1, ytp1, mtp) // 棟端２・三角形-2
	yanepoly = append(yanepoly, xo2, yo2, mbt)
	yanepoly = append(yanepoly, xtp2, ytp2, mtp)
	p[0] = []float64{xtp1, ytp1, mtp}
	p[1] = []float64{xo2, yo2, mbt}
	p[2] = []float64{xtp2, ytp2, mtp}
	for j := 0; j < 3; j++ {
		nor = pkg.NorVec(p[(0+j)%3], p[(1+j)%3], p[(2+j)%3])
		nor_all = append(nor_all, nor)
	}

	yanepoly = append(yanepoly, xo1, yo1, mbt) // ケラバ１・三角形1-1
	yanepoly = append(yanepoly, xtp1, ytp1, mtp)
	yanepoly = append(yanepoly, xtp4, ytp4, ntp)
	p[0] = []float64{xo1, yo1, mbt}
	p[1] = []float64{xtp1, ytp1, mtp}
	p[2] = []float64{xtp4, ytp4, ntp}
	for j := 0; j < 3; j++ {
		nor = pkg.NorVec(p[(0+j)%3], p[(1+j)%3], p[(2+j)%3])
		nor_all = append(nor_all, nor)
	}

	yanepoly = append(yanepoly, xo1, yo1, mbt) // ケラバ１・三角形1-2
	yanepoly = append(yanepoly, xtp4, ytp4, ntp)
	yanepoly = append(yanepoly, xo4, yo4, nbt)
	p[0] = []float64{xo1, yo1, mbt}
	p[1] = []float64{xtp4, ytp4, ntp}
	p[2] = []float64{xo4, yo4, nbt}
	for j := 0; j < 3; j++ {
		nor = pkg.NorVec(p[(0+j)%3], p[(1+j)%3], p[(2+j)%3])
		nor_all = append(nor_all, nor)
	}

	yanepoly = append(yanepoly, xtp3, ytp3, ntp) // ケラバ２・三角形2-1
	yanepoly = append(yanepoly, xtp2, ytp2, mtp)
	yanepoly = append(yanepoly, xo2, yo2, mbt)
	p[0] = []float64{xtp3, ytp3, ntp}
	p[1] = []float64{xtp2, ytp2, mtp}
	p[2] = []float64{xo2, yo2, mbt}
	for j := 0; j < 3; j++ {
		nor = pkg.NorVec(p[(0+j)%3], p[(1+j)%3], p[(2+j)%3])
		nor_all = append(nor_all, nor)
	}

	yanepoly = append(yanepoly, xo2, yo2, mbt) // ケラバ２・三角形2-2
	yanepoly = append(yanepoly, xo3, yo3, nbt)
	yanepoly = append(yanepoly, xtp3, ytp3, ntp)
	p[0] = []float64{xo2, yo2, mbt}
	p[1] = []float64{xo3, yo3, nbt}
	p[2] = []float64{xtp3, ytp3, ntp}
	for j := 0; j < 3; j++ {
		nor = pkg.NorVec(p[(0+j)%3], p[(1+j)%3], p[(2+j)%3])
		nor_all = append(nor_all, nor)
	}

	// 側面三角壁の下端高さ
	sh0 := toph / kansan
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

	yanepoly = append(yanepoly, x1, y1, sh0) // 側面・三角形２
	yanepoly = append(yanepoly, x1, y1, sh1)
	yanepoly = append(yanepoly, x4, y4, sh0)
	p[0] = []float64{x1, y1, sh0}
	p[1] = []float64{x1, y1, sh1}
	p[2] = []float64{x4, y4, sh0}
	for j := 0; j < 3; j++ {
		nor = pkg.NorVec(p[(0+j)%3], p[(1+j)%3], p[(2+j)%3])
		nor_all = append(nor_all, nor)
	}

	sh2 := (-vec[0]*(x2-xo1)-vec[1]*(y2-yo1))/vec[2] + mbt
	log.Println("sh2=", sh2)

	yanepoly = append(yanepoly, x2, y2, sh2) // 側面・三角形１
	yanepoly = append(yanepoly, x2, y2, sh0)
	yanepoly = append(yanepoly, x3, y3, sh0)
	p[0] = []float64{x3, y3, sh2}
	p[1] = []float64{x2, y2, sh0}
	p[2] = []float64{x3, y3, sh0}
	for j := 0; j < 3; j++ {
		nor = pkg.NorVec(p[(0+j)%3], p[(1+j)%3], p[(2+j)%3])
		nor_all = append(nor_all, nor)
	}

	yanepoly = append(yanepoly, x1, y1, sh1) // 背面・三角形-1
	yanepoly = append(yanepoly, x1, y1, sh0)
	yanepoly = append(yanepoly, x2, y2, sh0)
	p[0] = []float64{x1, y1, sh1}
	p[1] = []float64{x1, y1, sh0}
	p[2] = []float64{x2, y2, sh0}
	for j := 0; j < 3; j++ {
		nor = pkg.NorVec(p[(0+j)%3], p[(1+j)%3], p[(2+j)%3])
		nor_all = append(nor_all, nor)
	}

	yanepoly = append(yanepoly, x1, y1, sh1) // 背面・三角形-2
	yanepoly = append(yanepoly, x2, y2, sh0)
	yanepoly = append(yanepoly, x2, y2, sh2)
	p[0] = []float64{x1, y1, sh1}
	p[1] = []float64{x2, y2, sh0}
	p[2] = []float64{x2, y2, sh2}
	for j := 0; j < 3; j++ {
		nor = pkg.NorVec(p[(0+j)%3], p[(1+j)%3], p[(2+j)%3])
		nor_all = append(nor_all, nor)
	}

	// TODO:背面の高さを出力してTriYaneに渡したい
	// sh1とsh2をTriYaneに渡す

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
