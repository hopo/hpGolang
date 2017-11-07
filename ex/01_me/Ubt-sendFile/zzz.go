package main

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func main() {
	command := "git add ./"
	exec_cmd(command)
}

func exec_cmd(c string) {
	cmdsl := strings.Split(c, " ")
	cmd := exec.Command(cmdsl[0], cmdsl[1], cmdsl[2])
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("*D*O*N*E*")
}
