package main

import (
	"fmt"
	"io/ioutil"
	"os/exec"
)

func main() {
	exec_cmd()
}

func exec_cmd() {
	lsCmd := exec.Command("ls")
	o, _ := lsCmd.Output()
	grepCmd := exec.Command("grep", ".doc")

	grepIn, _ := grepCmd.StdinPipe()
	grepOut, _ := grepCmd.StdoutPipe()
	grepCmd.Start()
	//grepIn.Write([]byte("hello grep\ngoodbye grep"))
	grepIn.Write(o)
	grepIn.Close()
	grepBytes, _ := ioutil.ReadAll(grepOut)
	grepCmd.Wait()

	fmt.Println("> grep .doc")
	fmt.Println(string(grepBytes))
}
