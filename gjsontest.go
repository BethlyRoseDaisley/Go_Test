package main

import "github.com/tidwall/gjson"

const json = `{
    "name":{
        "first":"Janet",
        "last":"Prichard"
    },
    "age":47
}`

func main() {
  value := gjson.Get(json, "ag")
  if !value.Exists() {
    println("no such key!")
    return
  }
  println(value.String())
}
