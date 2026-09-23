package internal

import (
	"log"
	"math"
	"stera/pkg"
	"strconv"
)

// KiriYane は切妻屋根の頂点座標と頂点の法線ベクトルを求める
func KiriYane(list [][]float64, toph, hisashi, keraba, incline,
	yaneatu float64) (yanetxt, yanenor []string) {
	log.Println("切妻屋根")
	// 屋根モデル（切妻屋根）
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

	// インチ換算後のケラバ厚さと庇長さ（軒の出）
	kich := keraba / kansan
	hich := hisashi / kansan

	// 頂点1と頂点2を結ぶ線分に平行な直線の式（妻面）
	m1, n1 := pkg.ParaLine(x1, y1, x2, y2, x3, y3, kich)
	// log.Println("y = " + strconv.FormatFloat(m1, 'f', -1, 64) + "x + " + strconv.FormatFloat(n1, 'f', -1, 64))
	// 頂点3と頂点4を結ぶ線分に平行な直線の式（妻面）
	m2, n2 := pkg.ParaLine(x3, y3, x4, y4, x1, y1, kich)
	// log.Println("y = " + strconv.FormatFloat(m2, 'f', -1, 64) + "x + " + strconv.FormatFloat(n2, 'f', -1, 64))
	// 頂点2と頂点3を結ぶ線分に平行な直線の式（平面）// 軒庇の下端
	m3, n3 := pkg.ParaLine(x2, y2, x3, y3, x4, y4, hich)
	// log.Println("y = " + strconv.FormatFloat(m3, 'f', -1, 64) + "x + " + strconv.FormatFloat(n3, 'f', -1, 64))
	// 頂点4と頂点1を結ぶ線分に平行な直線の式（平面）// 軒庇の下端
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

	// 軒鼻の突き出し長さ
	nt := yaneatu / kansan / math.Sqrt(math.Pow(incline, 2)+1) * incline
	// 平面から軒庇の上端までの長さ（軒庇＋軒鼻）
	hichtop := hich + nt
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
	nbt := (toph - hisashi*incline) / kansan
	// 軒先上端高さ
	ntp := nbt + yaneatu/kansan/math.Sqrt(math.Pow(incline, 2)+1)
	log.Println("nbt, ntp", nbt, ntp)

	// 屋根の棟端の下端高さ
	mbt1 := nbt + tm1/2*incline
	mbt2 := nbt + tm2/2*incline
	// 屋根の棟端の上端高さ
	mtp1 := mbt1 + yaneatu/kansan*math.Sqrt(1+math.Pow(incline, 2))
	mtp2 := mbt2 + yaneatu/kansan*math.Sqrt(1+math.Pow(incline, 2))
	log.Println("mbt1, mtp1, mbt2, mtp2", mbt1, mtp1, mbt2, mtp2)

	// 屋根モデルの頂点座標をリストに書き出す
	// 屋根頂点の法線ベクトルを算出してリストに書き出す
	var yanepoly []float64
	var nor_all [][]float64
	var normal []float64
	var nor []float64
	p := make([][]float64, 3)

	yanepoly = append(yanepoly, xo1, yo1, nbt) // 屋根底面・三角形１
	yanepoly = append(yanepoly, xo4, yo4, nbt)
	yanepoly = append(yanepoly, xm1, ym1, mbt1)
	p[0] = []float64{xo1, yo1, nbt}
	p[1] = []float64{xo4, yo4, nbt}
	p[2] = []float64{xm1, ym1, mbt1}
	for j := 0; j < 3; j++ {
		nor = pkg.NorVec(p[(0+j)%3], p[(1+j)%3], p[(2+j)%3])
		nor_all = append(nor_all, nor)
	}

	yanepoly = append(yanepoly, xo4, yo4, nbt) // 屋根底面・三角形２
	yanepoly = append(yanepoly, xm2, ym2, mbt2)
	yanepoly = append(yanepoly, xm1, ym1, mbt1)
	p[0] = []float64{xo4, yo4, nbt}
	p[1] = []float64{xm2, ym2, mbt2}
	p[2] = []float64{xm1, ym1, mbt1}
	for j := 0; j < 3; j++ {
		nor = pkg.NorVec(p[(0+j)%3], p[(1+j)%3], p[(2+j)%3])
		nor_all = append(nor_all, nor)
	}

	yanepoly = append(yanepoly, xm1, ym1, mbt1) // 屋根底面・三角形３
	yanepoly = append(yanepoly, xm2, ym2, mbt2)
	yanepoly = append(yanepoly, xo2, yo2, nbt)
	p[0] = []float64{xm1, ym1, mbt1}
	p[1] = []float64{xm2, ym2, mbt2}
	p[2] = []float64{xo2, yo2, nbt}
	for j := 0; j < 3; j++ {
		nor = pkg.NorVec(p[(0+j)%3], p[(1+j)%3], p[(2+j)%3])
		nor_all = append(nor_all, nor)
	}

	yanepoly = append(yanepoly, xm2, ym2, mbt2) // 屋根底面・三角形４
	yanepoly = append(yanepoly, xo3, yo3, nbt)
	yanepoly = append(yanepoly, xo2, yo2, nbt)
	p[0] = []float64{xm2, ym2, mbt2}
	p[1] = []float64{xo3, yo3, nbt}
	p[2] = []float64{xo2, yo2, nbt}
	for j := 0; j < 3; j++ {
		nor = pkg.NorVec(p[(0+j)%3], p[(1+j)%3], p[(2+j)%3])
		nor_all = append(nor_all, nor)
	}

	yanepoly = append(yanepoly, xtp1, ytp1, ntp) // 屋根上面・三角形１
	yanepoly = append(yanepoly, xtp4, ytp4, ntp)
	yanepoly = append(yanepoly, xm1, ym1, mtp1)
	p[0] = []float64{xtp1, ytp1, ntp}
	p[1] = []float64{xtp4, ytp4, ntp}
	p[2] = []float64{xm1, ym1, mtp1}
	for j := 0; j < 3; j++ {
		nor = pkg.NorVec(p[(0+j)%3], p[(1+j)%3], p[(2+j)%3])
		nor_all = append(nor_all, nor)
	}

	yanepoly = append(yanepoly, xtp4, ytp4, ntp) // 屋根上面・三角形２
	yanepoly = append(yanepoly, xm2, ym2, mtp2)
	yanepoly = append(yanepoly, xm1, ym1, mtp1)
	p[0] = []float64{xtp4, ytp4, ntp}
	p[1] = []float64{xm2, ym2, mtp2}
	p[2] = []float64{xm1, ym1, mtp1}
	for j := 0; j < 3; j++ {
		nor = pkg.NorVec(p[(0+j)%3], p[(1+j)%3], p[(2+j)%3])
		nor_all = append(nor_all, nor)
	}

	yanepoly = append(yanepoly, xm1, ym1, mtp1) // 屋根上面・三角形３
	yanepoly = append(yanepoly, xm2, ym2, mtp2)
	yanepoly = append(yanepoly, xtp2, ytp2, ntp)
	p[0] = []float64{xm1, ym1, mtp1}
	p[1] = []float64{xm2, ym2, mtp2}
	p[2] = []float64{xtp2, ytp2, ntp}
	for j := 0; j < 3; j++ {
		nor = pkg.NorVec(p[(0+j)%3], p[(1+j)%3], p[(2+j)%3])
		nor_all = append(nor_all, nor)
	}

	yanepoly = append(yanepoly, xm2, ym2, mtp2) // 屋根上面・三角形４
	yanepoly = append(yanepoly, xtp3, ytp3, ntp)
	yanepoly = append(yanepoly, xtp2, ytp2, ntp)
	p[0] = []float64{xm2, ym2, mtp2}
	p[1] = []float64{xtp3, ytp3, ntp}
	p[2] = []float64{xtp2, ytp2, ntp}
	for j := 0; j < 3; j++ {
		nor = pkg.NorVec(p[(0+j)%3], p[(1+j)%3], p[(2+j)%3])
		nor_all = append(nor_all, nor)
	}

	yanepoly = append(yanepoly, xtp2, ytp2, ntp) //  軒端１・三角形-1
	yanepoly = append(yanepoly, xo2, yo2, nbt)
	yanepoly = append(yanepoly, xo3, yo3, nbt)
	p[0] = []float64{xtp2, ytp2, ntp}
	p[1] = []float64{xo2, yo2, nbt}
	p[2] = []float64{xo3, yo3, nbt}
	for j := 0; j < 3; j++ {
		nor = pkg.NorVec(p[(0+j)%3], p[(1+j)%3], p[(2+j)%3])
		nor_all = append(nor_all, nor)
	}

	yanepoly = append(yanepoly, xtp2, ytp2, ntp) //  軒端１・三角形-2
	yanepoly = append(yanepoly, xo3, yo3, nbt)
	yanepoly = append(yanepoly, xtp3, ytp3, ntp)
	p[0] = []float64{xtp2, ytp2, ntp}
	p[1] = []float64{xo3, yo3, nbt}
	p[2] = []float64{xtp3, ytp3, ntp}
	for j := 0; j < 3; j++ {
		nor = pkg.NorVec(p[(0+j)%3], p[(1+j)%3], p[(2+j)%3])
		nor_all = append(nor_all, nor)
	}

	yanepoly = append(yanepoly, xtp4, ytp4, ntp) //  軒端２・三角形-1
	yanepoly = append(yanepoly, xo4, yo4, nbt)
	yanepoly = append(yanepoly, xo1, yo1, nbt)
	p[0] = []float64{xtp4, ytp4, ntp}
	p[1] = []float64{xo4, yo4, nbt}
	p[2] = []float64{xo1, yo1, nbt}
	for j := 0; j < 3; j++ {
		nor = pkg.NorVec(p[(0+j)%3], p[(1+j)%3], p[(2+j)%3])
		nor_all = append(nor_all, nor)
	}

	yanepoly = append(yanepoly, xtp4, ytp4, ntp) //  軒端２・三角形-2
	yanepoly = append(yanepoly, xo1, yo1, nbt)
	yanepoly = append(yanepoly, xtp1, ytp1, ntp)
	p[0] = []float64{xtp4, ytp4, ntp}
	p[1] = []float64{xo1, yo1, nbt}
	p[2] = []float64{xtp1, ytp1, ntp}
	for j := 0; j < 3; j++ {
		nor = pkg.NorVec(p[(0+j)%3], p[(1+j)%3], p[(2+j)%3])
		nor_all = append(nor_all, nor)
	}

	yanepoly = append(yanepoly, xtp1, ytp1, ntp) // ケラバ１・三角形-1
	yanepoly = append(yanepoly, xo1, yo1, nbt)
	yanepoly = append(yanepoly, xm1, ym1, mbt1)
	p[0] = []float64{xtp1, ytp1, ntp}
	p[1] = []float64{xo1, yo1, nbt}
	p[2] = []float64{xm1, ym1, mbt1}
	for j := 0; j < 3; j++ {
		nor = pkg.NorVec(p[(0+j)%3], p[(1+j)%3], p[(2+j)%3])
		nor_all = append(nor_all, nor)
	}

	yanepoly = append(yanepoly, xtp1, ytp1, ntp) // ケラバ１・三角形-2
	yanepoly = append(yanepoly, xm1, ym1, mbt1)
	yanepoly = append(yanepoly, xm1, ym1, mtp1)
	p[0] = []float64{xtp1, ytp1, ntp}
	p[1] = []float64{xm1, ym1, mbt1}
	p[2] = []float64{xm1, ym1, mtp1}
	for j := 0; j < 3; j++ {
		nor = pkg.NorVec(p[(0+j)%3], p[(1+j)%3], p[(2+j)%3])
		nor_all = append(nor_all, nor)
	}

	yanepoly = append(yanepoly, xm1, ym1, mtp1) // ケラバ２・三角形-1
	yanepoly = append(yanepoly, xm1, ym1, mbt1)
	yanepoly = append(yanepoly, xo2, yo2, nbt)
	p[0] = []float64{xm1, ym1, mtp1}
	p[1] = []float64{xm1, ym1, mbt1}
	p[2] = []float64{xo2, yo2, nbt}
	for j := 0; j < 3; j++ {
		nor = pkg.NorVec(p[(0+j)%3], p[(1+j)%3], p[(2+j)%3])
		nor_all = append(nor_all, nor)
	}

	yanepoly = append(yanepoly, xm1, ym1, mtp1) // ケラバ２・三角形-2
	yanepoly = append(yanepoly, xo2, yo2, nbt)
	yanepoly = append(yanepoly, xtp2, ytp2, ntp)
	p[0] = []float64{xm1, ym1, mtp1}
	p[1] = []float64{xo2, yo2, nbt}
	p[2] = []float64{xtp2, ytp2, ntp}
	for j := 0; j < 3; j++ {
		nor = pkg.NorVec(p[(0+j)%3], p[(1+j)%3], p[(2+j)%3])
		nor_all = append(nor_all, nor)
	}

	yanepoly = append(yanepoly, xtp3, ytp3, ntp) // ケラバ３・三角形-1
	yanepoly = append(yanepoly, xo3, yo3, nbt)
	yanepoly = append(yanepoly, xm2, ym2, mbt2)
	p[0] = []float64{xtp3, ytp3, ntp}
	p[1] = []float64{xo3, yo3, nbt}
	p[2] = []float64{xm2, ym2, mbt2}
	for j := 0; j < 3; j++ {
		nor = pkg.NorVec(p[(0+j)%3], p[(1+j)%3], p[(2+j)%3])
		nor_all = append(nor_all, nor)
	}

	yanepoly = append(yanepoly, xtp3, ytp3, ntp) // ケラバ３・三角形-2
	yanepoly = append(yanepoly, xm2, ym2, mbt2)
	yanepoly = append(yanepoly, xm2, ym2, mtp2)
	p[0] = []float64{xtp3, ytp3, ntp}
	p[1] = []float64{xm2, ym2, mbt2}
	p[2] = []float64{xm2, ym2, mtp2}
	for j := 0; j < 3; j++ {
		nor = pkg.NorVec(p[(0+j)%3], p[(1+j)%3], p[(2+j)%3])
		nor_all = append(nor_all, nor)
	}

	yanepoly = append(yanepoly, xm2, ym2, mtp2) // ケラバ４・三角形-1
	yanepoly = append(yanepoly, xm2, ym2, mbt2)
	yanepoly = append(yanepoly, xo4, yo4, nbt)
	p[0] = []float64{xm2, ym2, mtp2}
	p[1] = []float64{xm2, ym2, mbt2}
	p[2] = []float64{xo4, yo4, nbt}
	for j := 0; j < 3; j++ {
		nor = pkg.NorVec(p[(0+j)%3], p[(1+j)%3], p[(2+j)%3])
		nor_all = append(nor_all, nor)
	}

	yanepoly = append(yanepoly, xm2, ym2, mtp2) // ケラバ４・三角形-2
	yanepoly = append(yanepoly, xo4, yo4, nbt)
	yanepoly = append(yanepoly, xtp4, ytp4, ntp)
	p[0] = []float64{xm2, ym2, mtp2}
	p[1] = []float64{xo4, yo4, nbt}
	p[2] = []float64{xtp4, ytp4, ntp}
	for j := 0; j < 3; j++ {
		nor = pkg.NorVec(p[(0+j)%3], p[(1+j)%3], p[(2+j)%3])
		nor_all = append(nor_all, nor)
	}

	// 妻面三角壁の下端高さ
	h := toph / kansan
	// 妻面三角壁1の下端長さ
	l1 := pkg.DistVerts(x1, y1, x2, y2)
	// 妻面三角壁1の上端高さ
	mh1 := h + l1/2*incline
	// 妻面三角壁1の頂点座標
	xc1 := (x1 + x2) / 2
	yc1 := (y1 + y2) / 2

	yanepoly = append(yanepoly, x1, y1, h) // 妻壁・三角形１
	yanepoly = append(yanepoly, x2, y2, h)
	yanepoly = append(yanepoly, xc1, yc1, mh1)
	p[0] = []float64{x1, y1, h}
	// log.Println("p[0]=", p[0])
	p[1] = []float64{x2, y2, h}
	// log.Println("p[2]=", p[2])
	p[2] = []float64{xc1, yc1, mh1}
	// log.Println("p[1]=", p[1])
	for j := 0; j < 3; j++ {
		nor = pkg.NorVec(p[(0+j)%3], p[(1+j)%3], p[(2+j)%3])
		nor_all = append(nor_all, nor)
	}

	// 妻面三角壁2の下端長さ
	l2 := pkg.DistVerts(x3, y3, x4, y4)
	// 妻面三角壁2の上端高さ
	mh2 := h + l2/2*incline
	// 妻面三角壁2の頂点座標
	xc2 := (x3 + x4) / 2
	yc2 := (y3 + y4) / 2

	yanepoly = append(yanepoly, x3, y3, h) // 妻壁・三角形２
	yanepoly = append(yanepoly, x4, y4, h)
	yanepoly = append(yanepoly, xc2, yc2, mh2)
	p[0] = []float64{x3, y3, h}
	// log.Println("p[0]=", p[0])
	p[1] = []float64{x4, y4, h}
	// log.Println("p[2]=", p[2])
	p[2] = []float64{xc2, yc2, mh2}
	// log.Println("p[1]=", p[1])
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
