package main

import (
	"fmt"
	"os"

	"learninggo.com/go/files/fileutils"
)

func main() {
	rootPath, _ := os.Getwd()
	filePath := rootPath + "/data/text.txt"
	content, err := fileutils.ReadTextFile(rootPath + "/data/text.txt")
	if err == nil {
		fmt.Println(content)

		newContent := fmt.Sprintf("Original: %v\n Double Original: %v%v", content, content, content)
		fileutils.WriteToFile(filePath+".output.txt", newContent)
	} else {
		fmt.Printf("Error Panic !! %v", err)
	}

}
