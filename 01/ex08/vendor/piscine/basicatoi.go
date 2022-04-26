package piscine

func BasicAtoi(s string) int {
	if s == "" {
		return 0
	}
	runes := []rune(s)
	var res int = 0
	var neg bool = false
	for i, r := range runes {
		if (!(r >= '0' && r <= '9')) {
			return (0);
		}
		res -= int(r - '0');
		if (i == len(runes) - 1) {
			break ;
		}
		res *= 10;
	}
	if (neg == false) {
		res *= -1;
	}
	return (res);
}