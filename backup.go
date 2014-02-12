package main

import (
	"errors"
	"fmt"
	"github.com/gallend/golang-file-backup/DirectoryReader"
	"log"
	"os"
)

var (
	ErrInvalidArg = errors.New("Missing one or more arguments")
)

func main() {

	if len(os.Args) != 3 {
		log.Fatal(ErrInvalidArg)
	}

	fmt.Println(len(os.Args))

	originDirectory := os.Args[1]
	// backupDirectory := os.Args[1]

	s := DirectoryReader.ReadAll(originDirectory)

	for i := 0; i < 2; i++ {
		fmt.Println(s[i].Name())
	}

}
