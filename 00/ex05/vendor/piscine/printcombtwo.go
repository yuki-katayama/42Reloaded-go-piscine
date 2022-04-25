package piscine

import (
	"ft"
)

func PrintFourDigits(first rune, second rune, third rune, fourth rune) {
	ft.PrintRune(first);
	ft.PrintRune(second);
	ft.PrintRune(' ');
	ft.PrintRune(third);
	ft.PrintRune(fourth);
}

func PrintSeparate() {
	ft.PrintRune(',');
	ft.PrintRune(' ');
}

func PrintCombTwo() {
	for x1 := '0'; x1 <= '9'; x1++ {
		for y1 := '0'; y1 <= '8'; y1++ {
			for x2 := x1; x2 <= '9'; x2++ {
				for y2 := y1 + 1; y2 <= '9'; y2++ {
					PrintFourDigits(x1, y1, x2, y2);
					if x1 == '9' && y1 == '8' && x2 == '9' && y2 == '9' {
						break;
					}
					PrintSeparate();
				}
			}
		}
	}
	ft.PrintRune('\n');
}