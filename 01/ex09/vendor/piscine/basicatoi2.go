package piscine

func Atoi(s string) int {
	if s == "" {
		return 0
	}
	runes := []rune(s)
	var res int = 0
	for i, r := range runes {
		if (!(r >= '0' && r <= '9')) {
			return (0);
		}
		res += int(r - '0');
		if (i == len(runes) - 1) {
			break ;
		}
		res *= 10;
	}
	return (res);
}