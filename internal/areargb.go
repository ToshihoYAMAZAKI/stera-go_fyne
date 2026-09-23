package internal

import "fmt"

type Areargb struct {
	Area1  []float64
	Area2  []float64
	Area3  []float64
	Area4  []float64
	Area5  []float64
	Area6  []float64
	Area7  []float64
	Area8  []float64
	Area9  []float64
	Area10 []float64
	Area11 []float64
	Area12 []float64
}

var Checked bool
var Youtorgb Areargb

// AreaRGB は用途地域ごとに着色する色のRGB値を設定する
func AreaRGB(num int) (rgb []float64) {
	fmt.Println("Checked=", Checked)
	if Checked {
		if num == 1 {
			// rgb = append(rgb, 0)
			// rgb = append(rgb, 0.647)
			// rgb = append(rgb, 0.408)
			// rgb = append(rgb, 1)
			fmt.Println("Area1=", Youtorgb.Area1)
			rgb = Youtorgb.Area1
		} else if num == 2 {
			// rgb = append(rgb, 0.467)
			// rgb = append(rgb, 0.718)
			// rgb = append(rgb, 0.620)
			// rgb = append(rgb, 1)
			fmt.Println("Area2=", Youtorgb.Area2)
			rgb = Youtorgb.Area2
		} else if num == 3 {
			// rgb = append(rgb, 0.314)
			// rgb = append(rgb, 0.686)
			// rgb = append(rgb, 0.424)
			// rgb = append(rgb, 1)
			fmt.Println("Area3=", Youtorgb.Area3)
			rgb = Youtorgb.Area3
		} else if num == 4 {
			// rgb = append(rgb, 0.745)
			// rgb = append(rgb, 0.804)
			// rgb = append(rgb, 0)
			// rgb = append(rgb, 1)
			fmt.Println("Area4=", Youtorgb.Area4)
			rgb = Youtorgb.Area4
		} else if num == 5 {
			// rgb = append(rgb, 0)
			// rgb = append(rgb, 0.937)
			// rgb = append(rgb, 0.267)
			// rgb = append(rgb, 1)
			fmt.Println("Area5=", Youtorgb.Area5)
			rgb = Youtorgb.Area5
		} else if num == 6 {
			// rgb = append(rgb, 0.976)
			// rgb = append(rgb, 0.698)
			// rgb = append(rgb, 0)
			// rgb = append(rgb, 1)
			fmt.Println("Area6=", Youtorgb.Area6)
			rgb = Youtorgb.Area6
		} else if num == 7 {
			// rgb = append(rgb, 0.933)
			// rgb = append(rgb, 0.498)
			// rgb = append(rgb, 0)
			// rgb = append(rgb, 1)
			fmt.Println("Area7=", Youtorgb.Area7)
			rgb = Youtorgb.Area7
		} else if num == 8 {
			// rgb = append(rgb, 0.941)
			// rgb = append(rgb, 0.569)
			// rgb = append(rgb, 0.604)
			// rgb = append(rgb, 1)
			fmt.Println("Area8=", Youtorgb.Area8)
			rgb = Youtorgb.Area8
		} else if num == 9 {
			// rgb = append(rgb, 0.910)
			// rgb = append(rgb, 0.345)
			// rgb = append(rgb, 0.522)
			// rgb = append(rgb, 1)
			fmt.Println("Area9=", Youtorgb.Area9)
			rgb = Youtorgb.Area9
		} else if num == 10 {
			// rgb = append(rgb, 0.820)
			// rgb = append(rgb, 0.741)
			// rgb = append(rgb, 0.851)
			// rgb = append(rgb, 1)
			fmt.Println("Area10=", Youtorgb.Area10)
			rgb = Youtorgb.Area10
		} else if num == 11 {
			// rgb = append(rgb, 0.749)
			// rgb = append(rgb, 0.886)
			// rgb = append(rgb, 0.906)
			// rgb = append(rgb, 1)
			fmt.Println("Area11=", Youtorgb.Area11)
			rgb = Youtorgb.Area11
		} else if num == 12 {
			// rgb = append(rgb, 0.314)
			// rgb = append(rgb, 0.420)
			// rgb = append(rgb, 0.678)
			// rgb = append(rgb, 1)
			fmt.Println("Area12=", Youtorgb.Area12)
			rgb = Youtorgb.Area12
		} else if num == 13 {
			rgb = append(rgb, 0.410)
			rgb = append(rgb, 0.410)
			rgb = append(rgb, 0.410)
			rgb = append(rgb, 1)
		}
	} else {
		rgb = append(rgb, 0.410)
		rgb = append(rgb, 0.410)
		rgb = append(rgb, 0.410)
		rgb = append(rgb, 1)
	}
	fmt.Println(rgb)

	return rgb
}
