package pkg

import (
	"math"
)

// 法線ベクトルを求める
func NorVec(p1, p2, p3 []float64) (nor []float64) {
	// p1, p2, p3は三角メッシュの頂点座標
	v1 := []float64{0.0, 0.0, 0.0}
	v2 := []float64{0.0, 0.0, 0.0}
	crs := []float64{0.0, 0.0, 0.0}
	var len float64

	// ベクトルV1
	for i := 0; i < 3; i++ {
		v1[i] = p2[i] - p1[i]
	}
	// ベクトルV2
	for i := 0; i < 3; i++ {
		v2[i] = p3[i] - p1[i]
	}
	// 法線ベクトルの計算（外積計算）
	for i := 0; i < 3; i++ {
		crs[i] = v1[(i+1)%3]*v2[(i+2)%3] - v1[(i+2)%3]*v2[(i+1)%3]
	}
	len = math.Sqrt(crs[0]*crs[0] + crs[1]*crs[1] + crs[2]*crs[2])

	// 法線ベクトルの正規化
	for i := 0; i < 3; i++ {
		if len == 0 {
			nor = append(nor, 0.0)
		} else {
			nor = append(nor, crs[i]/len)
		}
	}
	return nor
}
