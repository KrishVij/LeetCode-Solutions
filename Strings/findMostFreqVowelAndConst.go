package main

import "fmt"


func maxFreqSum(s string) int {

    vowels := "aeiou"
    Sum := 0
    maxVowel := 0
    maxConsonant := 0
    
    vowelStore := make(map[byte]int)
    consonantStore := make(map[byte]int)

    for i := 0;i < len(s);i++ {

        if strings.Contains(vowels,string(s[i])) {

            vowelStore[s[i]]++
        }else {

             consonantStore[s[i]]++
        }
    
    }
    
    for _,val := range vowelStore {

        if val > maxVowel {

            maxVowel = val
        }
    }

    for _,val := range consonantStore {

        if len(consonantStore) == 0 { break }
        if val > maxConsonant {

            maxConsonant = val
        }
    }

    Sum = maxVowel + maxConsonant
    
    return Sum
}
