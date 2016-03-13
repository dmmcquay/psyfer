package psyfer

import (
	"strings"
	"testing"
)

func TestVigenereCipher(t *testing.T) {
	key := "vig"
	input := "theboyhasthebag"
	input = strings.ToUpper(strings.Replace(input, " ", "", -1))
	key = strings.ToUpper(strings.Replace(key, " ", "", -1))
	expected := "OPKWWECIYOPKWIM"
	actual := VigenereCipher(input, key, false)
	if expected != actual {
		t.Errorf(
			"failed VigenereCipher:\n\texpected: % q\n\t  actual: % q",
			expected,
			actual,
		)
	}
	key = "vig"
	input = "OPKWWECIYOPKWIM"
	input = strings.ToUpper(strings.Replace(input, " ", "", -1))
	key = strings.ToUpper(strings.Replace(key, " ", "", -1))
	expected = "THEBOYHASTHEBAG"
	actual = VigenereCipher(input, key, true)
	if expected != actual {
		t.Errorf(
			"failed VigenereCipher:\n\texpected: % q\n\t  actual: % q",
			expected,
			actual,
		)
	}
}
