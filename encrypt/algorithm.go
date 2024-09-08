package encrypt

func Encrypting(str string) string {
	word := ""
	for _, c := range str {
		asciiCode := int(c)
		char := string(asciiCode + 3)
		word += char
	}
	return word
}
