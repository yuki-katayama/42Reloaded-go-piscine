package piscine

import (
	"ft"
)

func PrintThreeDigits(first rune, second rune, third rune) {
	ft.PrintRune(first);
	ft.PrintRune(second);
	ft.PrintRune(third);
}

func PrintSeparate() {
	ft.PrintRune(',');
	ft.PrintRune(' ');
}

func PrintComb() {
	for x := '0'; x <= '7'; x++ {
		for y := x + 1; y <= '8'; y++ {
			for z := y + 1; z <= '9'; z++ {
				PrintThreeDigits(x, y, z);
				if x == '7' {
					break;
				}
				PrintSeparate();
			}
		}
	}
	ft.PrintRune('\n');
}