package pkg

import (
	"math"
)

// CrossAngl はベクトルの交差角度を求める
func CrossAngl(ax, ay, bx, by float64) (deg float64) {
	// ２つのベクトルのなす角を内積の式より求める
	cosT := (ax*bx + ay*by) / (math.Sqrt(math.Pow(ax, 2)+math.Pow(ay, 2)) * math.Sqrt(math.Pow(bx, 2)+math.Pow(by, 2)))
	// cosθからアークコサインで角度（ラジアン→度）を求める
	deg = math.Acos(cosT) * (180 / math.Pi)
	// log.Println("角度", deg) // Ctrl+/

	return deg
}
