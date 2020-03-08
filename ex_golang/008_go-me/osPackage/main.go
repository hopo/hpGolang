package main

import (
	"fmt"
	"os"
)

func main() {

	// Get HOME environment
	home := os.Getenv("HOME")

	filename := fmt.Sprintf("%s/.gdrive/gdrive", home)

	// file status
	_, err := os.Stat(filename)

	// checked file exist?
	if os.IsNotExist(err) {
		fmt.Println("Nonono File")
	} else {
		fmt.Println("File Exist")
	}

	// make new folder
	newFolder := fmt.Sprintf("%s/TestTestTest", home)
	err = os.Mkdir(newFolder, 0777)
	if err != nil {
		panic(err)
	}

	fmt.Println("Make Folder: ", newFolder)

}
