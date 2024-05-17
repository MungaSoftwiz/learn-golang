package main

// size embedded with type so only accepts [5]int not [4]int
//func Sum(numbers [5]int) int {
//	sum := 0 	-> Used only within  functions
//	for i := 0; i < 5; i++ { uses a loop, range below
//	for _, number := range numbers {
//		sum  += number
//	}
//	return sum
//}

func Sum(numbers []int) int {
        sum := 0
        for _, number := range numbers {
                sum  += number
        }
        return sum
}

// uses variadic param to accept any no of slices
func SumAll(numbersToSum ...[]int) []int {
//	lengthOfNumbers := len(numbersToSum)
//	sums := make([]int, lengthOfNumbers) allows creating
//	a slice with a starting capacity of lengthOfNumbers

	var sums []int
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}
	return sums
}

func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}
	return sums
}
