package main

import (
	"fmt"
	"time"
)

func bubbleSort(array []int) {
	for i := 0; i < len(array); i++ {
		for j := 0; j < len(array)-1-i; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}
}

func makeTimestamp() int64 {
	return time.Now().UnixNano() / (int64(time.Millisecond)/int64(time.Nanosecond))
}

func main() {
	arrayRange := 100
	array := make([]int, arrayRange)

	for i := 0; i < arrayRange ; i++ {
		array[i] = i
	}

	fmt.Println("Arrays unorganized: ", array)

	startTimeStamp := makeTimestamp()
	fmt.Println("Starting Timestamp: ", startTimeStamp)
	bubbleSort(array)
	finishTimeStamp := makeTimestamp()
	fmt.Println("Finish Timestamp: ", finishTimeStamp)
	fmt.Println("Arrays organized: ", array)

	fmt.Println("Total in ms: ", finishTimeStamp - startTimeStamp)

}

//Coded by LuccasDev
//GoLang is amazing
