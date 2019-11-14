package main 

import(
	"os/exec"
	"fmt"
	"time"
)

func main() {
	cmd := exec.Command("print.exe")
	fmt.Println(cmd.Process)
	fmt.Println(cmd.Args)
	err  := cmd.Start()
	if err != nil {
		fmt.Println("exec failed.")
		return
	}

	fmt.Println(cmd.Process)

	time.Sleep(5*time.Second)

	cmd.Process.Kill()

	fmt.Println(cmd.Process)

	time.Sleep(time.Hour)
}