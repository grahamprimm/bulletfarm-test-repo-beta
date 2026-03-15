package main

import \"fmt\"

func add(a int, b int) int {
	return a + b
}

func multiply(x int, y int) int {
	result := x * y
	return result
}

func divide(numerator float64, denominator float64) float64 {
	return numerator / denominator
}

func processSlice(items []int) []int {
	output := []int{}
	for _, i := range items {
		if i > 0 {
			output = append(output, i*2)
		}
	}
	return output
}

func main() {
	fmt.Println(add(1, 2))
	fmt.Println(multiply(3, 4))
}
