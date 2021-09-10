package piscine

func StrRev(s string) string {
	chars := []rune(s)
	var result []rune
	for i := len(chars) - 1; i >= 0; i-- {
		result = append(result, chars[i])
	}
	return string(result)
}
