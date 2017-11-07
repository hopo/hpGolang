package main1

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
	"sync"
)

// scp ~/Desktop/Ubt-sendFile/* hp@192.168.0.15:home/hp/Desktop/

func main() {
	command := "npm -l --depth 0"
	//command := "scp ~/Desktop/Ubt-sendFile/* hp@192.168.0.15:home/hp/Desktop/"
	wg := new(sync.WaitGroup)
	for {
		wg.Add(1)
		go exec_cmd(command, wg)
	}
	wg.Wait()
}

func exec_cmd(c string, wg *sync.WaitGroup) {
	cmdsl := strings.Split(c, " ")
	cmd := exec.Command(cmdsl[0], cmdsl[1], cmdsl[2])
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("************")
	fmt.Printf("%s", cmd)
	wg.Done()
}
