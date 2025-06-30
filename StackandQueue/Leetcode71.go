package main

import "strings"

func simplifyPath(path string) string {
    
    if len(path) == 0 {

        return ""
    }

    Path := strings.Split(path,"/")
    Stack := []string{}

    for _,part := range Path {

        if part == "" || part == "." {

            continue
        }

        if part == ".." {

            if len(Stack) > 0 {

                Stack = Stack[:len(Stack) - 1]
            }else {

                Stack = append(Stack, part)
            }
        }
    }

    return "/" + strings.Join(Stack, "/")
}
