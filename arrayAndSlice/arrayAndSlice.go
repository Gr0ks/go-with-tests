package main

func Sum(numbers [5]int) (result int) {
	for _, v := range numbers {
		result += v
	}
	return
}