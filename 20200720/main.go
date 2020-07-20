package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func bubble_sort(a []int) {
	is_sorted := true

	start_time := time.Now()

	for {
		for i := 0; i < len(a)-1; i++ {
			for j := i; j < len(a)-1; j++ {
				if a[j] > a[j+1] {
					a[j], a[j+1] = a[j+1], a[j]
					is_sorted = true
				}
			}
		}
		if is_sorted {
			is_sorted = false
		} else {
			break
		}
	}

	end_time := time.Now()

	fmt.Println("bubble_sort", a)
	fmt.Println(end_time.Sub(start_time).Seconds(), "s")
}

func bucket_sort(a []int) {
	buckets := make([][]int, 0)
	for i := 0; i < 100; i++ {
		bucket := make([]int, 0)
		buckets = append(buckets, bucket)
	}

	start_time := time.Now()

	// push bucket
	for i := 0; i < len(a); i++ {
		a_elem := a[i]
		buckets[a_elem] = append(buckets[a_elem], a_elem)
	}

	// pull bucket
	sorted_a := make([]int, 0)
	for i := 0; i < len(buckets); i++ {
		for j := 0; j < len(buckets[i]); j++ {
			sorted_a = append(sorted_a, buckets[i][j])
		}
	}

	end_time := time.Now()

	fmt.Println("bucket_sort", sorted_a)
	fmt.Println(end_time.Sub(start_time).Seconds(), "s")
}

func slice_search_index(number int, a []int) int {
	for i := 0; i < len(a); i++ {
		if a[i] == number {
			return i
		}
	}
	return -1
}

func slice_remove(a []int, index int) []int {
	result := a[:index]
	if index+1 < len(a) {
		for i := index + 1; i < len(a); i++ {
			result = append(result, a[i])
		}
	}
	return result
}

func radix_sort(a []int) {
	buckets := make([][]int, 0)
	for i := 0; i < 10; i++ {
		bucket := make([]int, 0)
		buckets = append(buckets, bucket)
	}

	start_time := time.Now()

	// 1回目
	for i := 0; i < len(a); i++ {
		number := a[i]
		number_str := strconv.Itoa(number)
		number_slice := strings.Split(number_str, "")
		bucket_index, _ := strconv.Atoi(number_slice[len(number_slice)-1])
		buckets[bucket_index] = append(buckets[bucket_index], number)
	}

	// 2回目以降
	for i := 1; i >= 0; i-- {
		for j := 0; j < len(a); j++ {
			number := a[j]
			number_str := strconv.Itoa(number)
			number_slice := strings.Split(number_str, "")

			if len(number_slice) < i+1 {
				continue
			}
			bucket_index, _ := strconv.Atoi(number_slice[i])

			before_bucket_index := 0
			for k := 0; k < len(buckets); k++ {
				for l := 0; l < len(buckets[k]); l++ {
					if buckets[k][l] == number {
						before_bucket_index = k
					}
				}
			}

			delete_index := slice_search_index(number, buckets[before_bucket_index])
			buckets[before_bucket_index] = slice_remove(buckets[before_bucket_index], delete_index)

			buckets[bucket_index] = append(buckets[bucket_index], number)
		}
	}

	sorted_a := make([]int, 0)

	for i := 0; i < len(buckets); i++ {
		for j := 0; j < len(buckets[i]); j++ {
			sorted_a = append(sorted_a, buckets[i][j])
		}
	}

	end_time := time.Now()

	fmt.Println("radix_sort", sorted_a)
	fmt.Println(end_time.Sub(start_time).Seconds(), "s")
}

func main() {
	a := make([]int, 0)

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 20; i++ {
		elem := rand.Intn(1000)
		a = append(a, elem)
	}

	radix_sort(a)
}
