package iteration

const repeatCount = 5

// Repeat the string n-times and return
func Repeat(character string, repeatCount int) string {
	var repeated string
	for i := 0; i < repeatCount; i++ {
		repeated += character
	}
	return repeated
}