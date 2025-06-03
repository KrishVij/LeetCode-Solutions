func countSubarrays(nums []int, k int64) int64 {

	if len(nums) == 0 {

		return 0
	}

	var Left int64
	var Right int64
	var Count int64
	var Sum int64
	var Length int64
	var Score int64

	for Right < int64(len(nums)) {

		Sum += int64(nums[Right])

		Length = Right - Left + 1

		Score = Sum * Length

		for Score >= k {

			Sum -= int64(nums[Left])
			Left++
			Length = Right - Left + 1
			Score = Sum * Length

		}
		
//We do Count += Length instead of Count++ as for each valid right, all subarrays ending at right and starting between left and right (inclusive) are valid, so we add (right - left + 1) to the result.
		Count += Length 
		Right++

	}

	return Count
}
