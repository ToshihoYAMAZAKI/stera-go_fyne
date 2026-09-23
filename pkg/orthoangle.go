package pkg

import "log"

// OrthoAngle はL点を含む直線と対向する辺との交差角度を求める
func OrthoAngle(orthoXY [][]float64, faceXY [][]float64) (isecX, isecY, theta float64) {
	// 直交する辺の両端座標（一方がL点）
	x1 := orthoXY[0][0] // X座標
	y1 := orthoXY[0][1] // Y座標
	x2 := orthoXY[1][0] // X座標
	y2 := orthoXY[1][1] // Y座標
	// log.Println("x1=", x1)
	// log.Println("y1=", y1)
	// log.Println("x2=", x2)
	// log.Println("y2=", y2)
	// 直交する直線の方程式
	line1 := LineEquat(x1, y1, x2, y2)
	// log.Println(line1)
	m1 := line1["m"] // 傾き
	n1 := line1["n"] // 切片
	// log.Printf("m1: %g\n", m1)
	// log.Printf("n1: %g\n", n1)
	// 対向する辺の両端座標
	x3 := faceXY[0][0] // X座標
	y3 := faceXY[0][1] // Y座標
	x4 := faceXY[1][0] // X座標
	y4 := faceXY[1][1] // Y座標
	// log.Println("x3=", x3)
	// log.Println("y3=", y3)
	// log.Println("x4=", x4)
	// log.Println("y4=", y4)
	// 対向する直線の方程式
	line2 := LineEquat(x3, y3, x4, y4)
	// log.Println(line2)
	m2 := line2["m"]
	n2 := line2["n"]
	// log.Printf("m2: %g\n", m2)
	// log.Printf("n2: %g\n", n2)
	// ２直線の交点を求める
	isecX, isecY = SeekInsec(m1, n1, m2, n2)
	log.Println("isecX=", isecX)
	log.Println("isecY=", isecY)

	// ベクトルA　交点とL1点を結ぶベクトル
	// vectAX, vectAY, vectA := VectLen(faceXY[0][0], faceXY[0][1], isecX, isecY)
	vectAX, vectAY, _ := VectLen(faceXY[1][0], faceXY[1][1], isecX, isecY)
	// log.Println("vectAX=", vectAX) // Ctrl+/
	// log.Println("vectAY=", vectAY) // Ctrl+/
	// log.Println("vectA=", vectA)   // Ctrl+/
	// ベクトルB　交点と３つ目の点を結ぶベクトル
	// vectBX, vectBY, vectB := VectLen(orthoXY[0][0], orthoXY[0][1], isecX, isecY)
	vectBX, vectBY, _ := VectLen(orthoXY[0][0], orthoXY[0][1], isecX, isecY)
	// log.Println("vectBX=", vectBX) // Ctrl+/
	// log.Println("vectBY=", vectBY) // Ctrl+/
	// log.Println("vectB=", vectB)   // Ctrl+/
	// 内積を計算して交差する角度を求める
	// theta = CalcTheta(vectAX, vectAY, vectBX, vectBY, vectA, vectB)
	theta = CrossAngl(vectAX, vectAY, vectBX, vectBY)
	log.Println("theta=", theta)

	return isecX, isecY, theta
}
