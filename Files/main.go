package main

import (
	"fmt"
	"os"

	"learninggo.com/go/files/fileutils"
)

func main() {
	rootPath, _ := os.Getwd()
	content, err := fileutils.ReadTextFile(rootPath + "/data/text.txt")
	if err == nil {
		fmt.Println(content)
	} else {
		fmt.Printf("Error Panic !! %v", err)
	}

}
