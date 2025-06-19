package main

import (

	"strings"
	"strconv"
)

func exclusiveTime(n int, logs []string) []int {
    
    if n == 0 || len(logs) == 0 {

        return []int{0}
    }
    
    callStack := []int{}
    execTimeFns := make([]int,n)
    prevTime := 0

    for _,log := range logs {

        parts := strings.Split(log,":")
        parseID,_ := strconv.Atoi(parts[0])
        parseType := parts[1]
        parseTime,_ := strconv.Atoi(parts[2])

       if parseType == "start" {
        
            if len(callStack) > 0 {

            execTimeFns[callStack[len(callStack) - 1]] += parseTime - prevTime
           }    
            callStack = append(callStack,parseID)
            prevTime = parseTime
        }else {

            top := callStack[len(callStack) - 1]
            callStack = callStack[:len(callStack) - 1]
            execTimeFns[top] += parseTime - prevTime + 1
            prevTime = parseTime + 1 
        }
       
    }

    return execTimeFns
    
}