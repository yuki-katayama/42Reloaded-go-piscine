package piscine

func StrLen(s []int) int {
	len := 0;
	for range s {
		len++;
	}
	return len;
}

func SortIntegerTable(table []int) {
	for x := 0; x < StrLen(table); x++ {
		for y := x + 1; y < StrLen(table); y++ {
			if table[x] > table[y] {
				table[x], table[y] = table[y], table[x]
			}
		}
	}
}
