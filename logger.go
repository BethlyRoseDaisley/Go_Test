package main

import(
  "fmt"
  "log"
  "os"
)

var g_logFile *os.File
var g_err error
var g_log *log.Logger

func RedictionLogger() {
  
}


func main() {
  fmt.Println("-------main start--------\n")
  arr := []int {1,2,3,4,5,6,7}
  log.Print("Print arrary ", arr, "\n")
  log.Println("PrintIn arrart ", arr)
  log.Printf("Printf arrary with item [%d,%d]\n", arr[0], arr[6])


  log.Println("test for Print")
  //log.Fatalln("test for Fatal")
  //log.Panicln("test for Panic")

  g_logFile, g_err := os.Open("./test.log")

  if g_err != nil {
    log.Fatalln("./test.log file open error ", g_err, ".")
  }

  defer g_logFile.Close()


  g_log = log.New(g_logFile, "", log.Lmicroseconds|log.Llongfile)

  g_log.Println("test for Print")

  fmt.Println("\n-------main end--------")

}