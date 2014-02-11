package main

import "fmt"
import "github.com/gallend/golang-file-backup/DirectoryReader"

func main() {
	s := DirectoryReader.ReadAll("C:\\Users\\gregd_000\\Pictures\\Emily")

	for i := 0; i < 2; i++ {
		fmt.Println(s[i].Name())
	}

}
