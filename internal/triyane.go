package internal

import (
	"log"
	"math"
	"stera/pkg"
	"strconv"
)

// TriYane は３角形の片流れ屋根の頂点座標と頂点の法線ベクトルを求める
func TriYane(list [][]float64, toph, hisashi, keraba, incline,
	yaneatu float64) (yanetxt, yanenor []string, tripts map[string]float64) {
	// 屋根モデル（片流れ屋根）
	// 3つの頂点座標の定義
	kansan := 0.0254
	x1 := list[0][0] / kansan
	y1 := list[0][1] / kansan
	x2 := list[1][0] / kansan
	y2 := list[1][1] / kansan
	x3 := list[2][0] / kansan
	y3 := list[2][1] / kansan

	// インチ換算後のケラバ厚さと庇長さ（軒の出）
	kich := keraba / kansan
	hich := hisashi / kansan

	// 頂点2と頂点3を結ぶ線分に平行な直線の式（流れ面下手）
	m1, n1 := pkg.ParaLine(x2, y2, x3, y3, x1, y1, hich)
	// 頂点1と頂点2を結ぶ線分に平行な直線の式（側面）
	m2, n2 := pkg.ParaLine(x1, y1, x2, y2, x3, y3, kich)
	// 頂点3と頂点1を結ぶ線分に平行な直線の式（側面）
	m3, n3 := pkg.ParaLine(x3, y3, x1, y1, x2, y2, kich)

	// 屋根伏せの3頂点の座標を求める（屋根面の下端・上端）
	xo1, yo1 := pkg.SeekInsec(m2, n2, m3, n3)
	xo2, yo2 := pkg.SeekInsec(m1, n1, m2, n2)
	xo3, yo3 := pkg.SeekInsec(m3, n3, m1, n1)
	log.Println("xo1, yo1, xo2, yo2, xo3, yo3", xo1, yo1, xo2, yo2, xo3, yo3)

	// 軒鼻の突き出し長さ
	nt := yaneatu / kansan / math.Sqrt(math.Pow(incline, 2)+1) * incline
	// 壁面から軒庇の上端までの長さ（軒庇＋軒鼻）
	hichtop := hich + nt
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
	nbt := (toph - hisashi/math.Sqrt(math.Pow(incline, 2)+1)*incline) / kansan
	// 軒先上端高さ
	ntp := nbt + yaneatu/kansan/math.Sqrt(math.Pow(incline, 2)+1)
	log.Println("nbt, ntp", nbt, ntp)

	// 屋根の流れ面上手の下端高さ
	mbt := nbt + ((sf1+sf2)/2)*incline

	// 屋根の流れ面上手の上端高さ
	mtp := mbt + yaneatu/kansan*math.Sqrt(1+math.Pow(incline, 2))
	log.Println("mbt, mtp", mbt, mtp)

	// 屋根モデルの頂点座標をリストに書き出す
	// 屋根頂点の法線ベクトルを算出してリストに書き出す
	var yanepoly []float64
	var nor_all [][]float64
	var normal []float64
	var nor []float64
	p := make([][]float64, 3)

	yanepoly = append(yanepoly, xo1, yo1, mbt) // 屋根底面・三角形
	yanepoly = append(yanepoly, xo3, yo3, nbt)
	yanepoly = append(yanepoly, xo2, yo2, nbt)
	p[0] = []float64{xo1, yo1, mbt}
	p[1] = []float64{xo3, yo3, nbt}
	p[2] = []float64{xo2, yo2, nbt}
	for j := 0; j < 3; j++ {
		nor = pkg.NorVec(p[(0+j)%3], p[(1+j)%3], p[(2+j)%3])
		nor_all = append(nor_all, nor)
	}

	yanepoly = append(yanepoly, xo1, yo1, mtp) // 屋根上面・三角形
	yanepoly = append(yanepoly, xtp3, ytp3, ntp)
	yanepoly = append(yanepoly, xtp2, ytp2, ntp)
	p[0] = []float64{x1, y1, mtp}
	p[1] = []float64{xtp3, ytp3, ntp}
	p[2] = []float64{xtp2, ytp2, ntp}
	for j := 0; j < 3; j++ {
		nor = pkg.NorVec(p[(0+j)%3], p[(1+j)%3], p[(2+j)%3])
		nor_all = append(nor_all, nor)
	}

	yanepoly = append(yanepoly, xtp3, ytp3, ntp) // 軒端１・三角形-1
	yanepoly = append(yanepoly, xtp2, ytp2, ntp)
	yanepoly = append(yanepoly, xo3, yo3, nbt)
	p[0] = []float64{xtp3, ytp3, ntp}
	p[1] = []float64{xtp2, ytp2, ntp}
	p[2] = []float64{xo3, yo3, nbt}
	for j := 0; j < 3; j++ {
		nor = pkg.NorVec(p[(0+j)%3], p[(1+j)%3], p[(2+j)%3])
		nor_all = append(nor_all, nor)
	}

	yanepoly = append(yanepoly, xo3, yo3, nbt) // 軒端２・三角形-2
	yanepoly = append(yanepoly, xtp2, ytp2, ntp)
	yanepoly = append(yanepoly, xo2, yo2, nbt)
	p[0] = []float64{xo3, yo3, nbt}
	p[1] = []float64{x2, y2, ntp}
	p[2] = []float64{xo2, yo2, nbt}
	for j := 0; j < 3; j++ {
		nor = pkg.NorVec(p[(0+j)%3], p[(1+j)%3], p[(2+j)%3])
		nor_all = append(nor_all, nor)
	}

	yanepoly = append(yanepoly, xo1, yo1, mbt) // ケラバ１・三角形1-1
	yanepoly = append(yanepoly, xo2, yo2, ntp)
	yanepoly = append(yanepoly, xo1, yo1, mtp)
	p[0] = []float64{xo1, yo1, mbt}
	p[1] = []float64{xo2, yo2, ntp}
	p[2] = []float64{xo1, yo1, mtp}
	for j := 0; j < 3; j++ {
		nor = pkg.NorVec(p[(0+j)%3], p[(1+j)%3], p[(2+j)%3])
		nor_all = append(nor_all, nor)
	}

	yanepoly = append(yanepoly, xo1, yo1, mbt) // ケラバ１・三角形1-2
	yanepoly = append(yanepoly, xo2, yo2, nbt)
	yanepoly = append(yanepoly, xo2, yo2, ntp)
	p[0] = []float64{xo1, yo1, mbt}
	p[1] = []float64{xo2, yo2, nbt}
	p[2] = []float64{xo2, yo2, ntp}
	for j := 0; j < 3; j++ {
		nor = pkg.NorVec(p[(0+j)%3], p[(1+j)%3], p[(2+j)%3])
		nor_all = append(nor_all, nor)
	}

	yanepoly = append(yanepoly, xo3, yo3, ntp) // ケラバ２・三角形2-1
	yanepoly = append(yanepoly, xo1, yo1, mbt)
	yanepoly = append(yanepoly, xo1, yo1, mtp)
	p[0] = []float64{xo3, yo3, ntp}
	p[1] = []float64{xo1, yo1, mbt}
	p[2] = []float64{xo1, yo1, mtp}
	for j := 0; j < 3; j++ {
		nor = pkg.NorVec(p[(0+j)%3], p[(1+j)%3], p[(2+j)%3])
		nor_all = append(nor_all, nor)
	}

	yanepoly = append(yanepoly, xo3, yo3, nbt) // ケラバ２・三角形2-2
	yanepoly = append(yanepoly, xo1, yo1, mbt)
	yanepoly = append(yanepoly, xo3, yo3, ntp)
	p[0] = []float64{xo3, yo3, nbt}
	p[1] = []float64{xo1, yo1, mbt}
	p[2] = []float64{xo3, yo3, ntp}
	for j := 0; j < 3; j++ {
		nor = pkg.NorVec(p[(0+j)%3], p[(1+j)%3], p[(2+j)%3])
		nor_all = append(nor_all, nor)
	}

	// 側面三角壁の下端高さ
	sh0 := toph / kansan
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

	yanepoly = append(yanepoly, x1, y1, sh0) // 側面・三角形１
	yanepoly = append(yanepoly, x2, y2, sh0)
	yanepoly = append(yanepoly, x1, y1, sh1)
	p[0] = []float64{x1, y1, sh0}
	p[1] = []float64{x2, y2, sh0}
	p[2] = []float64{x1, y1, mbt}
	for j := 0; j < 3; j++ {
		nor = pkg.NorVec(p[(0+j)%3], p[(1+j)%3], p[(2+j)%3])
	}
	nor_all = append(nor_all, nor)

	sh2 := (-vec[0]*(xo3-xo1)-vec[1]*(yo3-yo1))/vec[2] + mbt
	log.Println("sh2=", sh2)

	yanepoly = append(yanepoly, x3, y3, sh0) // 側面・三角形２
	yanepoly = append(yanepoly, x1, y1, sh0)
	yanepoly = append(yanepoly, x1, y1, sh2)
	p[0] = []float64{x3, y3, sh0}
	p[1] = []float64{x1, y1, sh0}
	p[2] = []float64{x1, y1, mbt}
	for j := 0; j < 3; j++ {
		nor = pkg.NorVec(p[(0+j)%3], p[(1+j)%3], p[(2+j)%3])
	}
	nor_all = append(nor_all, nor)

	// 頂点座標法線ベクトルの書き出し
	for j := range nor_all {
		for k := 0; k < 3; k++ {
			normal = append(normal, nor_all[j][k])
		}
	}

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

	// 頂点座標リストのテキスト化
	for y := range yanepoly {
		yanetxt = append(yanetxt, strconv.FormatFloat(yanepoly[y], 'f', -1, 64))
	}

	// 頂点座標法線ベクトルのテキスト化
	for n := range normal {
		yanenor = append(yanenor, strconv.FormatFloat(normal[n], 'f', -1, 64))
	}

	return yanetxt, yanenor, tripts
}
