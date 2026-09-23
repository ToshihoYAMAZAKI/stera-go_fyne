package pkg

import (
	"math"
)

// TriAngle は3つの頂点からなる角度を求める
func TriAngle(x1, y1, x2, y2, x3, y3 float64) (deg float64) {
	// ３点の座標より２つのベクトルを求める
	vecAx := (x2 - x1)
	// log.Println("vecAx=", vecAx)
	vecAy := (y2 - y1)
	// log.Println("vecAy=", vecAy)
	vecBx := (x3 - x1)
	// log.Println("vecBx=", vecBx)
	vecBy := (y3 - y1)
	// log.Println("vecBy=", vecBy)

	// ２つのベクトルのなす角を内積の式より求める
	cosT := (vecAx*vecBx + vecAy*vecBy) / (math.Sqrt(math.Pow(vecAx, 2)+math.Pow(vecAy, 2)) * math.Sqrt(math.Pow(vecBx, 2)+math.Pow(vecBy, 2)))
	// log.Println("cosT=", cosT)
	// cosθからアークコサインで角度（ラジアン→度）を求める
	deg = math.Acos(cosT) * (180 / math.Pi)

	return deg
}
