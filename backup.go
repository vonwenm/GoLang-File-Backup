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

	originDirectory := os.Args[1]
	backupDirectory := os.Args[2]

	s := DirectoryReader.ReadAll(originDirectory)

	for i := 0; i < len(s); i++ {
		fmt.Println(s[i].Name())
	}

	b := DirectoryReader.ReadAll(backupDirectory)

	for i := 0; i < len(b); i++ {
		fmt.Println(b[i].Name())
	}

}		

func CompareLists() (isEqual bool, added []string, removed []string)
{

}

func CopyDirectory(){}

func CopyFile(){}

func Archive(){}
