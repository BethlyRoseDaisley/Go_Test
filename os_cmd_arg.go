package main

import (
    "os"
    "fmt"
)

func main()  {
    list := os.Args
    n := len(list)
    fmt.Printf("n = %d\n", n)
    for i:=0;i<n;i++  {
        fmt.Printf("args[%d]=%s\n", i, list[i])
    }
}