package main

import(
  "io"
  "fmt"
  "bufio"
  "strings"
)

func main(){
  fmt.Println("main start")

  s := strings.NewReader("10058 85001")
  br := bufio.NewReader(s)

  k1 := make([]byte, 1)
  ret, err := io.ReadFull(br, k1)
  if err != nil {
    panic(err)
  }

  fmt.Println("read full ", ret, k1)


  k2 := make([]byte, 20)
  ret, err = io.ReadFull(br, k2)
  if err != nil {
  
    fmt.Println("read full ", ret, k2)

    panic(err)
  }

}