package internal

// Bfcalc は建物地下部分のモデリングのための地下深さを設定する
func Bfcalc(level float64, story, bf int) (btm float64) {
	// 地下階がある場合は階数に応じてZ座標を地下階部分の深さだけ下げる
	if bf == 0 {
		// 地下階がない場合は，３階建て以下は根入れ深さを30cmとし，４階建て以上は階高の1/3の1ｍとする
		if story < 3 {
			btm = level - 0.3
		} else {
			btm = level - 1.0
		}
	} else {
		btm = float64(bf)*3.0 + 0.3
	}
	return btm
}
