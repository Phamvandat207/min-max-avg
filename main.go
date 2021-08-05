package main

import "fmt"

func findMinAndMax(arr []int) (min int, max int) {
	min = arr[0]
	max = arr[0]
	for _, value := range arr {
		if value < min {
			min = value
		}
		if value > max {
			max = value
		}
	}
	return min, max
}

func findAvg (arr []int) int{
	var sum, avg int
	for i :=0; i < len(arr); i++ {
		sum += arr[i]
	}
	avg = sum/(len(arr)-1)
	return avg
}

func main() {

	var arr = []int{1, 2, 3}
	min, max := findMinAndMax(arr)
	avg := findAvg(arr)
	fmt.Println("Min: ", min)
	fmt.Println("Max: ", max)
	fmt.Println("Average: ", avg)
}