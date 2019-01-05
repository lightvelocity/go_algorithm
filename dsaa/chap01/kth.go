// 第k个最大者

package main

import (
	"math/rand"
	"time"
	"fmt"
)

// random int array
func randomIntArray(n int, max int) []int {
	numbers := make([]int, n)
	for i := 0; i < n; i++ {
		numbers[i] = rand.Intn(max)
	}
	return numbers
}

// bubble sort
func bubbleSort(numbers *[]int, asc bool) {
	array := *numbers
	if asc {
		for i := len(array)-1; i >= 1; i--{
			for j := 1; j <= i; j++ {
				if array[j-1] > array[j] {
					array[j-1], array[j] = array[j], array[j-1]
				}
			}
		}
	} else {
		for i := len(array)-1; i >= 1; i--{
			for j := 1; j <= i; j++ {
				if array[j-1] < array[j] {
					array[j-1], array[j] = array[j], array[j-1]
				}
			}
		}
	}
}

func kth1(numbers []int, k int) int {
	bubbleSort(&numbers, false)
	return numbers[k-1]
}

func kth2(numbers []int, k int) int {
	knums := make([]int, k)
	for i := 0; i < k; i++ {
		knums[i] = numbers[i]
	}
	bubbleSort(&knums, false)

	for i := k; i < len(numbers); i++ {
		num := numbers[i]
		for j := 0; j < k; j++ {
			if num > knums[j] {
				for m := k-1; m > j; m-- {
					knums[m] = knums[m-1]
				}
				knums[j] = num;
			}
		}
	}

	return knums[k-1]
}

func kth(n int, max int, k int) {
	start := time.Now()
	interval := time.Now().Sub(start)

	start = time.Now()
	numbers := randomIntArray(n, max)
	fmt.Printf("generate random int array cost %v\n", time.Duration(int64(time.Now().Sub(start)) - int64(interval)))

	start = time.Now()
	kth := kth1(numbers, k)
	fmt.Printf("1: kth element is %d, cost %v\n", kth, time.Duration(int64(time.Now().Sub(start)) - int64(interval)))

	start = time.Now()
	kth = kth2(numbers, k)
	fmt.Printf("2: kth element is %d, cost %v\n", kth, time.Duration(int64(time.Now().Sub(start)) - int64(interval)))
}

func main() {
	n := 1000000
	kth(n, n * 100, n/10)
}