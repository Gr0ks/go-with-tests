package iteration

func Repeat(letter string) (repeated string) {
	for i := 0; i < 5; i++ {
		repeated += letter
	}
	return
}