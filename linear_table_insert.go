package main

import "fmt"

func insert(array []int, length int, tempArray []int, value int, insertIndex int) {

	for i := 0; i < length; i++ {
		if i < insertIndex {
			//copy value that before i = insertIndex to tempArray
			tempArray[i] = array[i]
		} else {
			//copy remaining elements to tempArray
			tempArray[i+1] = array[i]
		}
	}

	//insert value to tempArray
	tempArray[insertIndex] = value

}

func main() {
	var scores = []int{90, 70, 50, 80, 60, 85}
	var length = len(scores)
	var tempArray = make([]int, length+1)

	//insert 75 to the index 2
	insert(scores, length, tempArray, 75, 2)

	scores = tempArray

	for i := 0; i < length+1; i++ {
		fmt.Printf("%d\n", scores[i])
	}
}
