package piscine

import "github.com/01-edu/z01"

func PrintComb() {
	for i := 48; i < 58; i++ {
		for j := 48; j < 58 && i < 58; j++ {
			for k := 48; k < 58 && i < 58; k++ {
				if j > i {
					if k > j {
						z01.PrintRune(rune(i))
						z01.PrintRune(rune(j))
						z01.PrintRune(rune(k))
						if i == 55 {
							z01.PrintRune('\n')
						} else {
							z01.PrintRune(',')
							z01.PrintRune(rune(32))
						}
					}
				}
			}
		}
	}
}
