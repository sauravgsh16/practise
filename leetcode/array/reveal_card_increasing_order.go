package main

import (
	"sort"
)

func deckRevealedIncreasing(deck []int) []int {
	// index skip implementation
	n := len(deck)

	sort.Ints(deck)

	result := make([]int, n)
	result[0] = deck[0]

	for i := 1; i < n; i++ {

	}
}

func deckRevealedIncreasingQueueImplementation(deck []int) []int {
	queue := []int{}
	cards := make([]int, 1000001)

	n := len(deck)

	// Update index for performing counting sort
	for i, d := range deck {
		queue = append(queue, i)
		cards[d]++
	}

	result := make([]int, n)
	for i, c := range cards {
		if c != 0 {
			// remove top index from queue
			index := queue[0]
			queue = queue[1:]

			// update result
			result[index] = i

			if len(queue) == 0 {
				break
			}
			// remove top index from queue again and add to last of queue
			queue = append(queue, queue[0])
			queue = queue[1:]
		}
	}
	return result
}
