package main

import (
	"fmt"
	"sort"
	"strconv"
)

func findRelativeRanks(score []int) []string {

	length := len(score)
	ranked := make([]string, length)

	sortedscore := append([]int{}, score...)
	sort.Sort(sort.Reverse(sort.IntSlice(sortedscore)))

	rankMap := make(map[int]int)
	for i, s := range sortedscore {

		rankMap[s] = i + 1
	}

	for i, s := range score {

		rank := rankMap[s]
		if rank == 1 {

			ranked[i] = "Gold Medal"
		} else if rank == 2 {

			ranked[i] = "Silver Medal"
		} else if rank == 3 {

			ranked[i] = "Bronze Medal"
		} else {

			ranked[i] = strconv.Itoa(rank)
		}

	}

	return ranked
}
