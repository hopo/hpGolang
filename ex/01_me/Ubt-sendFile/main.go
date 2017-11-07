package main

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
	//"sync"
)

// scp ~/Desktop/Ubt-sendFile/* hp@192.168.0.15:home/hp/Desktop/

func main() {
	command := "brew ls"
	//command := "scp ~/Desktop/Ubt-sendFile/* hp@192.168.0.15:home/hp/Desktop/"
	exec_cmd(command)
	fmt.Println("\n")
	log.Print("** D * O * N * E **")
}

func exec_cmd(c string) {
	cmdsl := strings.Split(c, " ")
	out, err := exec.Command(cmdsl[0], cmdsl[1]).Output()
	chkerr(err)
	fmt.Printf("%s", out)
}

func chkerr(e error) {
	if e != nil {
		panic(e)
	}
}
