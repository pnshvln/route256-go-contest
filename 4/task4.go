package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var count int
	fmt.Fscan(in, &count)

	for i := 0; i < count; i++ {
		var password string
		fmt.Fscan(in, &password)

		withDigit := false
		withSmall := false
		withBig := false
		withVowel := false
		withConsonant := false
		for k := 0; k < len(password); k++ {
			c := password[k]
			r := rune(c)
			if unicode.IsDigit(r) {
				withDigit = true
			} else {
				if unicode.IsLower(r) {
					withSmall = true
				} else {
					withBig = true
				}

				lower := unicode.ToLower(r)
				if lower == 'e' || lower == 'u' || lower == 'i' || lower == 'o' || lower == 'a' || lower == 'y' {
					withVowel = true
				} else {
					withConsonant = true
				}
			}
		}

		var sb strings.Builder
		sb.WriteString(password)
		if !withDigit {
			sb.WriteByte('1')
		}

		if !withBig {
			if !withVowel {
				sb.WriteByte('A')
				withVowel = true
			} else {
				sb.WriteByte('B')
				withConsonant = true
			}
		}

		if !withSmall {
			if !withVowel {
				sb.WriteByte('a')
				withVowel = true
			} else {
				sb.WriteByte('b')
				withConsonant = true
			}
		}

		if !withVowel {
			if !withBig {
				sb.WriteByte('A')
			} else {
				sb.WriteByte('a')
			}
		}

		if !withConsonant {
			if !withBig {
				sb.WriteByte('B')
			} else {
				sb.WriteByte('b')
			}
		}

		fmt.Fprintln(out, sb.String())
	}
}
