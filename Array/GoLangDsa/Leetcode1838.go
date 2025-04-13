import "sort"

func maxFrequency(nums []int, k int) int {

	descOrder := append([]int{}, nums...)
	sort.Sort(sort.Reverse(sort.IntSlice(descOrder)))
	comparison := descOrder[0]
	count := 0
	countStorage := 1

	for i := 1; i < len(descOrder); i++ {

		diff := comparison - descOrder[i]

		if k >= diff {

			k -= diff
			descOrder[i] = comparison
			count++
		} else {

			break
		}
		if count > countStorage {

			countStorage = count
			comparison = descOrder[i]
		}

	}

	return countStorage

}
