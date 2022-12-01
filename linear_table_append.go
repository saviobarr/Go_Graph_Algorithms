package main

import "fmt"

func append(array []int, value int) []int {
	var length = len(array)
	var tempArray = make([]int, length+1)

	//copy values of array to temp array
	for i := 0; i < length; i++ {
		tempArray[i] = array[i]
	}
	tempArray[length] = value

	return tempArray
}

func main() {
	var scores = []int{90, 70, 50, 80, 60, 85}

	//append 75 to array
	scores = append(scores, 75)
	var length = len(scores)
	for i := 0; i < length; i++ {
		fmt.Printf("%d \n", scores[i])
	}
}
