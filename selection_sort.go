package main

import "time"

func Selection_Sort(items []int) ([]int, int, float64) {
	var count int
	var n = len(items)

	inicio := time.Now().UnixNano() / int64(time.Millisecond)

	for i := 0; i < n; i++ {
		count++
		var minIdx = i
		for j := i; j < n; j++ {
			count++
			if items[j] < items[minIdx] {
				count++
				minIdx = j
			}
		}
		items[i], items[minIdx] = items[minIdx], items[i]
	}
	final := time.Now().UnixNano() / int64(time.Millisecond)

	tempo := float64(final - inicio)

	return items, count, tempo
}
