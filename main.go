package main

import (
	"fmt"
	"os"
	"path"
)

func main() {
	//fmt.Println("----------------")
	//fmt.Println(os.Args)
	//fmt.Println(os.Args[1:])
	//fmt.Println("----------------")
	doCommand(os.Args[1:])
}

func doCommand(args []string) {
	if len(args) < 2 || args[0] == "-h" || args[0] == "--help" {
		showHelp()
		return
	}
	from := args[0]
	to := args[1]
	dir := path.Dir(from)
	ext := path.Ext(from)
	newName := dir + "/" + to + ext
	err := os.Rename(from, newName)
	if err != nil {
		fmt.Println("ERROR:", err)
		os.Exit(1)
	}
	fmt.Printf("Renamed %s to %s\n", from, newName)
}

func showHelp() {
	fmt.Printf("help:")
	fmt.Printf("Usage: %s [OPTION]... [FILE]...\n", os.Args[0])
}
