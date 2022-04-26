package piscine

func UltimateDivMod(a *int, b *int) {
	var tmp_rem int = *a % *b
	*a = *a / *b
	*b = tmp_rem
}
