package main

func reverseDegree(s string) int {

    alphabetValue := make([]int, 26)
    
    alphaBet := make([]byte, 26)

    for i := 0;i < 26;i++ {

        alphabetValue[i] = 26 - i
    }

    for j := 0;j < 26;j++ {

        alphaBet[j] = byte('a' + j)
    }

    product := make([]int, len(s))

    for k := 0;k < len(s);k++ {

        for l,h := 0,0; l < 26 && h < 26; l,h = l + 1,h + 1 {

            if s[k] == alphaBet[l] {

                product[k] = alphabetValue[l] * (k + 1)
            }
        }
    }

    sum := 0
    
    for index := 0;index < len(product);index++ {

        sum += product[index]
    }

    return sum
    
}©leetcode
