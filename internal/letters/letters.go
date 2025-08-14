package letters

var vowels = []rune{'a', 'a', 'a', 'a', 'a', 'a', 'e', 'e', 'e', 'e', 'e', 'e', 'e', 'e', 'e', 'e', 'e', 'e', 'e', 'e', 'e', 'e', 'e', 'e', 'i', 'i', 'i', 'i', 'o', 'o', 'o', 'o', 'o', 'o', 'u', 'u', 'u', 'y'}
var consonants = []rune{'b', 'b', 'c', 'c', 'd', 'd', 'd', 'd', 'd', 'f', 'f', 'g', 'g', 'g', 'h', 'h', 'j', 'j', 'k', 'k', 'k', 'l', 'l', 'l', 'm', 'm', 'm', 'n', 'n', 'n', 'n', 'n', 'n', 'n', 'n', 'n', 'n', 'p', 'p', 'q', 'r', 'r', 'r', 'r', 'r', 's', 's', 's', 's', 's', 't', 't', 't', 't', 't', 'v', 'v', 'w', 'w', 'x', 'z', 'z'}


// TODO 1 letter per HTTP request probably, so we need to keep track of the letters already picked....
func Pick(vowelOrConsonant string) rune {
	return vowels[4]
}
