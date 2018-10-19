package main

import (
    "bufio"
    "fmt"
    "os"
)


func toInt(buf []byte) (n int) {
    for _, v := range buf {
        n = n*10 + int(v-'0')
    }
    return
}

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    var a, b int
    fmt.Sscanf(scanner.Text(), "%d %d", &a, &b)
    if a == b {
        fmt.Println("too late")
    } else {
        sl := make([]bool, a+1)
        for i:=0; i<b; i++ {
            scanner.Scan()
            sl[toInt(scanner.Bytes())] = true
        }
        for i:=1; i<a+1; i++ {
            if sl[i] == false {
                fmt.Println(i)
                break
            }
        }
    }
}