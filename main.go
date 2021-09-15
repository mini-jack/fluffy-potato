package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	out := new(bytes.Buffer)
	if !(len(os.Args) == 2 || len(os.Args) == 3) {
		panic("usage go run main.go . [-f]")
	}
	//path := os.Args[0]
	path := filepath.Dir("mian.go")
	printFiles := len(os.Args) == 2 && os.Args[1] == "-f"
	err := dirTree(out, path, printFiles)
	if err != nil {
		fmt.Println("outside dirtree panic")
		panic(err.Error())
	}
}

func dirTree(out *bytes.Buffer, path string, printFiles bool) error {
	return dirTree1(path, printFiles, 1)
}

func dirTree1(path string, printFiles bool, level int) error {
	contents, err := ioutil.ReadDir(filepath.Dir(path))
	spacer := string(strings.Repeat("│\t", (level - 1)))
	if err != nil {
		fmt.Println("inside dirtree panic")
		return err
	}
	if printFiles == false {
		return nil
	}
	for _, file := range contents {
		if strings.HasPrefix(file.Name(), ".") {
			continue
		}
		if file.IsDir() == false {
			fmt.Printf("%s├─%s(%v b)\n", spacer, file.Name(), file.Size())
			continue
		}
		fmt.Printf("%s└%s\n", (spacer), file.Name())
		dirTree1((path + "/" + file.Name() + "/"), printFiles, (level + 1))
	}
	return nil
}
