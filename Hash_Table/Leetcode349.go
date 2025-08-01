package main

import(

    "slices"
)

func intersection(nums1 []int, nums2 []int) []int {
    
    mpp := make(map[int]bool)
    result := []int{}
    
    for _,val := range nums1 {

        mpp[val] = true
    }

    for _,val2 := range nums2 {

        if mpp[val2] == true {

            if (!(slices.Contains(result,val2))) {

                result = append(result,val2)
            }
        }
    }

    return result
}