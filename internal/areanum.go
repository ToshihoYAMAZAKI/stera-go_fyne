package internal

// AreaNum は用途地域の名称を番号に変換する
func AreaNum(youto string) (num int) {
	if youto == "第1種低層住居専用地域" || youto == "第一種低層住居専用地域" {
		num = 1
	} else if youto == "第2種低層住居専用地域" || youto == "第二種低層住居専用地域" {
		num = 2
	} else if youto == "第1種中高層住居専用地域" || youto == "第一種中高層住居専用地域" {
		num = 3
	} else if youto == "第2種中高層住居専用地域" || youto == "第二種中高層住居専用地域" {
		num = 4
	} else if youto == "第1種住居地域" || youto == "第一種住居地域" {
		num = 5
	} else if youto == "第2種住居地域" || youto == "第二種住居地域" {
		num = 6
	} else if youto == "準住居地域" {
		num = 7
	} else if youto == "近隣商業地域" {
		num = 8
	} else if youto == "商業地域" {
		num = 9
	} else if youto == "準工業地域" {
		num = 10
	} else if youto == "工業地域" {
		num = 11
	} else if youto == "工業専用地域" {
		num = 12
	} else {
		num = 13
	}

	return num
}
