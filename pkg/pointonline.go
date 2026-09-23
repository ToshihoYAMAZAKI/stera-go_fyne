package pkg

// PointonLine は点が線分上にあるかどうか判定する
func PointonLine(ax, ay, bx, by, px, py float64) (chk bool) {
	if (ax <= px && px <= bx) || (bx <= px && px <= ax) {
		if (ay <= py && py <= by) || (by <= py && py <= ay) {
			// if (py*(ax-bx))+(ay*(bx-px))+(by*(px-ax)) == 0 {
			// 	// 点Pが線分AB上にある
			// 	return true
			// }
			return true
		}
	}
	// 点Pが線分AB上にない
	return false
}
