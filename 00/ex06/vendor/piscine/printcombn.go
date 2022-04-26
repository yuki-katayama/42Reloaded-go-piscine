package piscine

import "ft"

func CombNum(first int, num int, comb string){
	if num!=0{
		for i := first; i<=9; i++ {
			numbers := comb + (string(i + '0'))
			CombNum(i + 1, num - 1, numbers)
		}
	} else if num == 0{
		for _, char := range comb {
			ft.PrintRune(char)
		}
		ft.PrintRune(',')
		ft.PrintRune(' ')
	}
}

func PrintCombN(n int) {
	if n < 0 {
		return
	}
	combination := ""
	CombNum(0 ,n , combination);
	ft.PrintRune('\n');
}