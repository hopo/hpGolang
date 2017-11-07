package main

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func main() {
	command := "htop"
	exec_cmd(command)
}

func exec_cmd(c string) {
	cmdsl := strings.Split(c, " ")
	out, err := exec.Command(cmdsl[0]).Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", out)
	fmt.Println("*D*O*N*E*")
}
