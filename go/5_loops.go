// p5. Loops
package main

import "fmt"

func main() {
    i := 1
    for i <= 10 {
        fmt.Println(i)
        i = i + 1
    }

    for j := 0; j <= 5; j++ {
        fmt.Println(j)
    }

    for k := 0; k <= 10; k++ {
        if k % 2 == 0 {
            continue
        }
        fmt.Println(k)
    }

    for {
        fmt.Println("keep printing")
    }
}
