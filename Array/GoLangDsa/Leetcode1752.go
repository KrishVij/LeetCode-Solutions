func rotateArray (arr []int) []int {

	rotateFactor := len(arr)

	for i,j := 0,rotateFactor - 1;i<j;i,j = i + 1,j-1 {

		arr[i],arr[j] = arr[j],arr[i]
	}

	return arr

} 

func check(nums []int) bool {
    
	length := len(nums)
	isEqualtoNums := append([]int{},nums...)
	sort.Ints(isEqualtoNums)

	for count := 0; count < length;count++ {

		newArr : = rotateArray(isEqualtoNums)
		if reflect.DeepEqual (nums,newarr) {

			return true
		}
	}

	return false
}
