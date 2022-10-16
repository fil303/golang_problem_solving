package main

import "fmt"

func main() {
    a := 1
    b := 1
    for i := 0; i < 10; i++ {
        copy := a + b
        a = b
        fmt.Printf("%d,",a)
        b = copy
    } 
}
// 0,1,1,2,3,5,8,13,21,34
