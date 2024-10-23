package main

import (
	"fmt"
	"os"
	"path"
	"strings"
	"time"
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
	if to[0] == '.' {
		renameExt(from, to, dir)
	} else if to[0] == ':' {
		nameTransform(from, to, dir, ext)
	} else if to[0] == '^' {
		namePre(from, to, dir, ext)
	} else if to[0] == '+' {
		nameSuf(from, to, dir, ext)
	} else {
		renameFile(from, to, dir, ext)
	}
}

func showHelp() {
	fmt.Printf("help:")
	fmt.Printf("Usage: %s [OPTION]... [FILE]...\n", os.Args[0])
}
func renameFile(from string, to string, dir string, ext string) {
	newName := dir + "/" + to + ext
	err := os.Rename(from, newName)
	if err != nil {
		fmt.Println("ERROR:", err)
		os.Exit(1)
	}
	fmt.Printf("Renamed %s to %s\n", from, newName)
}

func renameExt(from string, to string, dir string) {
	basename := path.Base(from)
	part := strings.Split(basename, ".")[0]
	newExt := to
	newName := dir + "/" + part + newExt
	err := os.Rename(from, newName)
	if err != nil {
		fmt.Println("ERROR:", err)
		os.Exit(2)
	}
	fmt.Printf("Renamed %s to %s\n", from, newName)
}

func nameTransform(from string, to string, dir string, ext string) {
	basename := path.Base(from)
	part := strings.Split(basename, ".")[0]
	if to == ":upper" {
		part = strings.ToUpper(part)
	} else if to == ":lower" {
		part = strings.ToLower(part)
	}
	newName := dir + "/" + part + ext
	err := os.Rename(from, newName)
	if err != nil {
		fmt.Println("ERROR:", err)
		os.Exit(3)
	}
	fmt.Printf("Renamed %s to %s\n", from, newName)
}

func namePre(from string, to string, dir string, ext string) {
	basename := path.Base(from)
	part := strings.Split(basename, ".")[0]
	if to == "^d" || to == "^date" {
		date := fmt.Sprintf("%v", time.Now())
		date = strings.Split(date, " ")[0]
		datePre := strings.Split(date, "-")[0] + strings.Split(date, "-")[1] + strings.Split(date, "-")[2]
		newName := dir + "/" + datePre + "-" + part + ext
		err := os.Rename(from, newName)
		if err != nil {
			fmt.Println("ERROR:", err)
			os.Exit(4)
		}
		fmt.Printf("Renamed %s to %s\n", from, newName)
	} else if to == "^t" || to == "^time" {
		timeInCmd := fmt.Sprintf("%v", time.Now())
		timeInCmd = strings.Split(timeInCmd, ".")[0]
		timeInCmd = strings.Split(timeInCmd, " ")[0] + strings.Split(timeInCmd, " ")[1]
		timeInCmd = strings.Split(timeInCmd, "-")[0] + strings.Split(timeInCmd, "-")[1] + strings.Split(timeInCmd, "-")[2]
		timePre := strings.Split(timeInCmd, ":")[0] + strings.Split(timeInCmd, ":")[1] + strings.Split(timeInCmd, ":")[2]
		newName := dir + "/" + timePre + "-" + part + ext
		err := os.Rename(from, newName)
		if err != nil {
			fmt.Println("ERROR:", err)
			os.Exit(4)
		}
		fmt.Printf("Renamed %s to %s\n", from, newName)
	} else {
		if to[1] == '\\' {
			to = to[0:1] + to[2:]
		}
		newName := dir + "/" + to[1:] + part + ext
		err := os.Rename(from, newName)
		if err != nil {
			fmt.Println("ERROR:", err)
			os.Exit(4)
		}
		fmt.Printf("Renamed %s to %s\n", from, newName)
	}
}

func nameSuf(from string, to string, dir string, ext string) {
	basename := path.Base(from)
	part := strings.Split(basename, ".")[0]
	if to == "+d" || to == "+date" {
		date := fmt.Sprintf("%v", time.Now())
		date = strings.Split(date, " ")[0]
		dateSuf := strings.Split(date, "-")[0] + strings.Split(date, "-")[1] + strings.Split(date, "-")[2]
		newName := dir + "/" + part + "-" + dateSuf + ext
		err := os.Rename(from, newName)
		if err != nil {
			fmt.Println("ERROR:", err)
			os.Exit(4)
		}
		fmt.Printf("Renamed %s to %s\n", from, newName)
	} else if to == "+t" || to == "+time" {
		timeInCmd := fmt.Sprintf("%v", time.Now())
		timeInCmd = strings.Split(timeInCmd, ".")[0]
		timeInCmd = strings.Split(timeInCmd, " ")[0] + strings.Split(timeInCmd, " ")[1]
		timeInCmd = strings.Split(timeInCmd, "-")[0] + strings.Split(timeInCmd, "-")[1] + strings.Split(timeInCmd, "-")[2]
		timeSuf := strings.Split(timeInCmd, ":")[0] + strings.Split(timeInCmd, ":")[1] + strings.Split(timeInCmd, ":")[2]
		newName := dir + "/" + part + "-" + timeSuf + ext
		err := os.Rename(from, newName)
		if err != nil {
			fmt.Println("ERROR:", err)
			os.Exit(4)
		}
		fmt.Printf("Renamed %s to %s\n", from, newName)
	} else {
		if to[1] == '\\' {
			to = to[0:1] + to[2:]
		}
		newName := dir + "/" + part + to[1:] + ext
		err := os.Rename(from, newName)
		if err != nil {
			fmt.Println("ERROR:", err)
			os.Exit(4)
		}
		fmt.Printf("Renamed %s to %s\n", from, newName)
	}
}
