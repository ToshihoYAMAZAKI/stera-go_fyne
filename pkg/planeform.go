package pkg

// PlaneForm は3点を含む平面の式を求める
func PlaneForm(x1, y1, z1, x2, y2, z2, x3, y3, z3 float64) (a, b, c, d float64) {
	a = (y2-y1)*(z3-z1) - (y3-y1)*(z2-z1)
	b = (z2-z1)*(x3-x1) - (z3-z1)*(x2-x1)
	c = (x2-x1)*(y3-y1) - (x3-x1)*(y2-y1)
	d = -(a*x1 + b*y1 + c*z1)

	return a, b, c, d
}
