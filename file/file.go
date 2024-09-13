package file

import (
	"fmt"
	"sort"
	"os"
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Sort and write items to a file
//
//  1. Get arguments from command-line
//
//  2. Sort them
//
//  3. Write the sorted slice to a file
//
//
// EXPECTED OUTPUT
//
//   go run main.go
//     Send me some items and I will sort them
//
//   go run main.go orange banana apple
//
//   cat sorted.txt
//     apple
//     banana
//     orange
//
//
// HINTS
//
//   + REMEMBER: os.Args is a []string
//
//   + String slices are sortable using `sort.Strings`
//
//   + Use ioutil.WriteFile to write to a file.
//
//   + But you need to convert []string to []byte to be able to
//     write it to a file using the ioutil.WriteFile.
//
//   + To do that, create a new []byte and append the elements of your
//     []string.
//
// ---------------------------------------------------------

func main() {
	var content = os.Args[1:];
	
	if len(content) == 0 {
		fmt.Println("Give some content to be written in the file")
		return
	}

	sort.Strings(content)

	// content := []byte

	err := os.WriteFile("sorted.txt", []byte(strings.Join(content, "\n")), 0o400)

	if err!= nil {
		fmt.Println("content succesfully added and sorted to a file")
	}

	fmt.Println(content)
}
