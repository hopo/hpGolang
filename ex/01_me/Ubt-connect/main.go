package main

import (
	//"bytes"
	//"fmt"
	"log"
	"os/exec"
	//"strings"
)

func main() {
	cmd := exec.Command("ssh", "hp@192.168.0.15")
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}
