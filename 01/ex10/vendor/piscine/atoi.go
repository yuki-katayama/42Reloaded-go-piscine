package piscine

func StrLen(s []rune) int {
	len := 0
	for range s {
		len++
	}
	return len
}

func Atoi(s string) int {
	if s == "" {
		return 0
	}
	runes := []rune(s)
	var res int = 0
	var neg bool = false
	if runes[0] == '-' || runes[0] == '+' {
		if runes[0] == '-' {
			neg = true
		}
		runes[0] = '0'
	}
	for i, r := range runes {
		if !(r >= '0' && r <= '9') {
			return (0)
		}
		res += int(r - '0')
		if i == StrLen(runes)-1 {
			break
		}
		res *= 10
	}
	if neg == true {
		res *= -1
	}
	return (res)
}
