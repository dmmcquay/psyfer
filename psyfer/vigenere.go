package psyfer

func VigenereCipher(input string, key string, decrypt bool) string {
	alphabet := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	chars := []rune(input)
	k := []rune(key)
	keyPos := 0
	output := ""
	for _, m := range chars {
		index := int(m - 'A')
		offset := int(k[keyPos]-'A') % 26
		if decrypt {
			index -= offset
		} else {
			index += offset
		}
		if index >= 26 {
			index -= 26
		} else if index < 0 {
			index += 26
		}
		output += string(alphabet[index])
		keyPos++
		if keyPos == len(key) {
			keyPos = 0
		}
	}
	return output
}
