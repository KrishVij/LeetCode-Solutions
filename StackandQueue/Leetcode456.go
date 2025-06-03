package main

func find132pattern(nums []int) bool {
    
    second := -1 << 31
    stack := []int{}

    for i := len(nums) - 1;i >= 0;i-- {

        if nums[i] < second {

            return true
        }

        for len(stack) > 0 && stack[len(stack) - 1] < nums[i] {

            topEle := stack[len(stack) - 1]
            
            second = topEle

            stack = stack[:len(stack) - 1]
        }

        stack = append(stack,nums[i])
    }

    return false
}
