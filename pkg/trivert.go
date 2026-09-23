package pkg

import (
	"log"
)

// TriVert は３頂点から外積と内角を計算し時計回りかどうか判断する
func TriVert(n int, XY [][]float64) (extLst []float64,
	degLst []float64, rotate string) {

	// 各頂点の外積の値を計算して配列に格納する
	faceR := 0
	faceL := 0
	// 外積を計算する
	for i := 0; i < n; i++ {
		xs := XY[i][0]
		ys := XY[i][1]
		xp := XY[(i-1+n)%n][0]
		yp := XY[(i-1+n)%n][1]
		xn := XY[(i+1)%n][0]
		yn := XY[(i+1)%n][1]

		// 外積を計算する
		// a x b = |a||b|sinθ = ax * by - ay * bx
		s := (xp-xs)*(yn-ys) - (xn-xs)*(yp-ys)
		// 外積の計算結果を頂点順に配列に格納する
		extLst = append(extLst, s)
		// log.Println(i, "外積[i]=", extLst[i]) // Ctrl+/

		// ベクトルAのX座標の差分
		ax := xp - xs
		// ベクトルAのY座標の差分
		ay := yp - ys
		// ベクトルAの長さ
		// al := math.Sqrt(math.Pow(ax, 2) + math.Pow(ay, 2))
		// ベクトルBのX座標の差分
		bx := xn - xs
		// ベクトルBのY座標の差分
		by := yn - ys
		// ベクトルBの長さ
		// bl := math.Sqrt(math.Pow(bx, 2) + math.Pow(by, 2))

		// 交差角度を求める
		deg := CrossAngl(ax, ay, bx, by)
		// log.Println("角度", deg) // Ctrl+/
		// 交差角度を頂点順に配列に格納する
		degLst = append(degLst, deg)
		log.Println(i, "内角[i]=", degLst[i]) // Ctrl+/

		// 外積の結果で左回りか右回りか判断する
		// ベクトルAの向きからベクトルBまで回る向きを右ネジの回る向き
		// ３次元直交座標系はＺアップ
		// 外積の値が正の時XY平面では頂点が反時計回り
		if s >= 0.0 {
			faceL++
		} else {
			faceR++
		}
	}

	log.Println("R= ", faceR) // Ctrl+/
	log.Println("L= ", faceL) // Ctrl+/
	// 時計回りか反時計回りか判断する
	if faceR < faceL {
		// 時計回り
		rotate = "cw"
	} else {
		// 反時計回り
		rotate = "ccw"
	}
	log.Println("Rotate=", rotate) // Ctrl+/

	return extLst, degLst, rotate
}
