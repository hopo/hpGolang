package main

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
	//"sync"
)

func main() {
	command := "scp ~/Desktop/Ubt-sendFile/* hp@192.168.0.15:home/hp/Desktop/"
	exec_cmd(command)
	fmt.Println("\n")
	log.Print("*** D * O * N * E ***")
}

func exec_cmd(c string) {
	cmdsl := strings.Split(c, " ")
	cmd := exec.Command(cmdsl[0], cmdsl[1], cmdsl[2])
	err := cmd.Run()
	chkerr(err)
	fmt.Printf("%s", cmd)
}

func chkerr(e error) {
	if e != nil {
		panic(e)
	}
}
