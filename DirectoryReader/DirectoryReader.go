package DirectoryReader

import "io/ioutil"
import "os"

func ReadAll(directoryName string) []os.FileInfo {
	listings, _ := ioutil.ReadDir(directoryName)

	return listings
}
