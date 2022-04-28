package piscine

func StrLen(s []rune) int {
	len := 0
	for range s {
		len++
	}
	return len
}

func BasicAtoi(s string) int {
	if s == "" {
		return 0
	}
	runes := []rune(s)
	var res int = 0
	for i, r := range runes {
		res += int(r - '0')
		if i == StrLen(runes)-1 {
			break
		}
		res *= 10
	}
	return (res)
}
