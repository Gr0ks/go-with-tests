package main

func Sum(numbers []int) (result int) {
	for _, v := range numbers {
		result += v
	}
	return
}

func SumAll(arraysToSum ...[]int) (sums []int) {
	for _, numbers := range arraysToSum {
		sums = append(sums, Sum(numbers))
	}
	return
}

func SumAllTails(arraysToSum ...[]int) (sums []int) {
	for _, numbers := range arraysToSum {
		tail := numbers[1:]
		sums = append(sums, Sum(tail))
	}
	return
}