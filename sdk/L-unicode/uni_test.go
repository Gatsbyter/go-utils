package L_unicode

import (
	"fmt"
	"testing"
	"unicode"
)

func TestUnicode(t *testing.T) {
	fmt.Printf("%c - %c - %c - %c\n\n", unicode.MaxRune, unicode.ReplacementChar, unicode.MaxASCII, unicode.MaxLatin1)

	// ⚠️ 这里牛逼
	fmt.Println(unicode.Is(unicode.Scripts["Han"], '国'))
}

// 很多IsXXX 用的时候看看就行了
func TestIs(t *testing.T) {
	for _, r := range "Hello ＡＢＣ！" {
		if unicode.IsUpper(r) {
			fmt.Printf("%c", r)
		}

	}
	fmt.Println()

	for _, r := range "Hello ａｂｃ！" {
		if unicode.IsLower(r) {
			fmt.Printf("%c", r)
		}
	}
	fmt.Println()

	for _, r := range "Hello ᾏᾟᾯ！" {
		if unicode.IsTitle(r) {
			fmt.Printf("%c", r)
		}
	}
	fmt.Println()

	for _, cr := range unicode.Lt.R16 {
		for i := cr.Lo; i <= cr.Hi; i += cr.Stride {
			fmt.Printf("%c", i)
		}
	}
	fmt.Println()
}
