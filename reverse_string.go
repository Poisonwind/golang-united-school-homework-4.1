package reverse_string

func ReverseString(input string) (output string) {

	runeArray := []rune(input)

	for i, j := 0, len(runeArray)-1; i < j; i, j = i+1, j-1 {

		runeArray[i], runeArray[j] = runeArray[j], runeArray[i]

	}

	output = string(runeArray)

	return output
}
