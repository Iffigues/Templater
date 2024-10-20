package main

import (
	"flag"
	"fmt"
)

func main() {
	fileName := flag.String("filename", "", "a name to greet")
	var varFileArray arrayFlag
	flag.Var(&varFileArray, "fileVar", "Add an item to the list. Can be used multiple times.")
	var myArray arrayFlag
	flag.Var(&myArray, "var", "Add an item to the list. Can be used multiple times.")
	flag.Parse()
	fmt.Println(*fileName, myArray, varFileArray)

}
