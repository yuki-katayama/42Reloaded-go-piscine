package piscine

import (
	"ft"
)

func PrintComb() {
	for x := '0'; x <= '7'; x++ {
		for y := x + 1; y <= '8'; y++ {
			for z := y + 1; z <= '9'; z++ {
				ft.PrintRune(x);
				ft.PrintRune(y);
				ft.PrintRune(z);
				if x == '7' {
					break;
				}
				ft.PrintRune(',');
				ft.PrintRune(' ');
			}
		}
	}
}