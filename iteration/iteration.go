package iteration

const repeatCount = 5

func Repeat(letter string) (repeated string) {
	for i := 0; i < repeatCount; i++ {
		repeated += letter
	}
	return
}