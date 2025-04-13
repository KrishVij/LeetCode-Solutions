package main

func playGame(scorePlayerOne, scorePlayerTwo int, turn1, turn2 bool, nums []int) bool {

	if len(nums) == 0 {

		if scorePlayerOne > scorePlayerTwo || scorePlayerOne == scorePlayerTwo {

			return true
		}

		return false
	}

	if turn1 {

		option1 := playGame(scorePlayerOne+nums[0], scorePlayerTwo, false, true, nums[1:])

		option2 := playGame(scorePlayerOne+nums[len(nums)-1], scorePlayerTwo, false, true, nums[:len(nums)-1])

		return option1 || option2
	}

	if turn2 {

		option3 := playGame(scorePlayerOne, scorePlayerTwo+nums[len(nums)-1], true, false, nums[:len(nums)-1])

		option4 := playGame(scorePlayerOne, scorePlayerTwo+nums[0], true, false, nums[1:])

		return option3 && option4

	}

	return false
}

func predictTheWinner(nums []int) bool {

	return playGame(0, 0, true, false, nums)
}
