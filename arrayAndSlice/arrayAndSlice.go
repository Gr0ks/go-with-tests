package main

func Sum(numbers []int) (result int) {
	for _, v := range numbers {
		result += v
	}
	return
}