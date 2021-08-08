package iteration

func Repeat(character string, repeatTime int) string {
	var repeated string
	for i := 0; i < repeatTime; i++ {
		repeated += character
	}
	return repeated
}
