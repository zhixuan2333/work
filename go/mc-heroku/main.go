package main

import (
	"fmt"
	"io/ioutil"
	"os/exec"
	"time"
)

// Version minecraft version
const Version = "1.16.3"

func main() {

	go Sync()
	mcstart()

}

// Sync is mcserver
func Sync() {
	fmt.Println("Sync now")

	for {
		time.Sleep(time.Second * 30)

		fmt.Println("-----------")

	}

}

func mcstart() {

	cmd := exec.Command("java", "-jar", "server.jar")

	stdin, _ := cmd.StdinPipe()
	stdout, _ := cmd.StdoutPipe()

	if err := cmd.Start(); err != nil {
		fmt.Println("Execute failed when Start:" + err.Error())
		return
	}

	// stdin.Write([]byte("go text for grep\n"))
	// stdin.Write([]byte("go test text for grep\n"))
	stdin.Close()

	outbytes, _ := ioutil.ReadAll(stdout)
	stdout.Close()

	if err := cmd.Wait(); err != nil {
		fmt.Println("Execute failed when Wait:" + err.Error())
		return
	}

	fmt.Println("Execute finished:" + string(outbytes))

}
