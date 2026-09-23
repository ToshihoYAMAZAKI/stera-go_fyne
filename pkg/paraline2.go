package pkg

import (
	"math"
)

// ParaLine2 は２点を結ぶ線分に平行な直線の式を求める
func ParaLine2(x1, y1, x2, y2, xo, yo, dist float64) (l, m float64) {
	// ２点を通る直線の式（一般形）ax+by+c=0
	a := (y2 - y1)
	if a == 0.0 {
		a = 0.0001
	}
	b := (x1 - x2)
	if b == 0.0 {
		b = 0.0001
	}
	c := (x2-x1)*y1 - (y2-y1)*x1
	// 平行な直線との距離dist
	d0 := dist * math.Sqrt(math.Pow(a, 2)+math.Pow(b, 2))
	d1 := c - d0
	d2 := c + d0

	// 平行な直線の式 y=-a/b*x-d/b
	l = -a / b
	m1 := -d1 / b
	m2 := -d2 / b

	// 平行な直線の式 ax + by + c = 0
	// 対抗する辺の頂点までの距離を比較する
	// 距離が近い方の直線を採用する
	dst1 := math.Abs(a*xo+b*yo+d1) / math.Sqrt(math.Pow(a, 2)+math.Pow(b, 2))
	// fmt.Println("dst1=", dst1)
	dst2 := math.Abs(a*xo+b*yo+d2) / math.Sqrt(math.Pow(a, 2)+math.Pow(b, 2))
	// fmt.Println("dst2=", dst2)
	if dst1 < dst2 {
		m = m1
	} else {
		m = m2
	}
	return l, m
}
