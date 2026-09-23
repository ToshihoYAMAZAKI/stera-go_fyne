package pkg

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
)

// FileCount は読み込んだファイルの行数，文字数などを数える
func FileCount(file *os.File) (int, int, int, int) {
	rd := bufio.NewReader(file)
	c, l, w, m := 0, 0, 0, 0 // バイト数, 行数, 単語数, 文字数
	inword := false
	for {
		r, s, err := rd.ReadRune()
		if s == 0 {
			if err == io.EOF {
				return c, l, w, m
			} else {
				fmt.Fprintf(os.Stderr, "%d バイト目で不正なコードを検出しました\n", c+1)
				os.Exit(1)
			}
		} else if unicode.IsSpace(r) {
			inword = false
			if r == '\n' {
				l++
			}
		} else if !inword {
			inword = true
			w++
		}
		c += s
		m++
	}
}
