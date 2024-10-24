// Copyright [2024] [Sheyiyuan]
//
// Licensed under the MIT License.
// See the LICENSE file in the project root for more information.
package main

import (
	"fmt"
	"os"
	"path"
	"strings"
	"time"
)

func main() {
	doCommand(os.Args[1:])
}

func doCommand(args []string) {
	if len(args) < 2 || args[0] == "-h" || args[0] == "--help" {
		showHelp()
		return
	}
	if args[0] == "-f" {
		from := args[1]
		to := args[2]
		dir := path.Dir(from) + "/" + strings.Split(from, "/")[len(strings.Split(from, "/"))-1]
		ext := ""
		list, _ := listFilesInDirectory(dir)
		if to[0] == '.' {
			for i := 0; i < len(list); i++ {
				from = dir + "/" + list[i]
				ext = path.Ext(from)
				renameExt(from, to, dir)
			}
		} else if to[0] == ':' {
			for i := 0; i < len(list); i++ {
				from = dir + "/" + list[i]
				ext = path.Ext(from)
				nameTransform(from, to, dir, ext)
			}
		} else if to[0] == '^' {
			for i := 0; i < len(list); i++ {
				from = dir + "/" + list[i]
				ext = path.Ext(from)
				namePre(from, to, dir, ext)
			}
		} else if to[0] == '+' {
			for i := 0; i < len(list); i++ {
				from = dir + "/" + list[i]
				ext = path.Ext(from)
				nameSuf(from, to, dir, ext)
			}
		} else {
			fmt.Printf("You can't rename all file name in a folder.")
		}
	} else {
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
}

func showHelp() {
	helpDoc := `
File Rename Tool	version:1.0.0
====Sheyiyuan====

rename -h,--help
read help document

Modify Name:
rename path/to/your/file.extension newName
	Modify file name (cannot directly modify extension)
rename path/to/your/file.extension.newExt
	Modify file extension (add '.' before new extension)

Transform Name:
rename path/to/your/file :upper,:u
	transform to upper character
rename path/to/your/file :lower,:l
	transform to lower character
rename path/to/your/file :camel,:c
	transform camel to snake
rename path/to/your/file :snake,:s
	transform snake to snake

Add prefix/suffix
rename path/to/your/file ^pre
	add prefix
rename path/to/your/file +suf
	add suffix
Use date or time as prefix/suffix:
	Use 'd'/'date' or 't'/'time'. 
If adding these words directly, add '\\' before them.

if you want to change by folder,add -f as the first argument.
Tip:Command with "-f" will ignore all subfolder.
`
	fmt.Printf(helpDoc)
}
func renameFile(from string, to string, dir string, ext string) {
	newName := dir + "/" + to + ext
	doRename(from, newName, 1)
}

func renameExt(from string, to string, dir string) {
	basename := path.Base(from)
	part := strings.Split(basename, ".")[0]
	newExt := to
	newName := dir + "/" + part + newExt
	doRename(from, newName, 2)
}

func nameTransform(from string, to string, dir string, ext string) {
	basename := path.Base(from)
	part := strings.Split(basename, ".")[0]
	if to == ":upper" || to == ":u" {
		part = strings.ToUpper(part)
	} else if to == ":lower" || to == ":l" {
		part = strings.ToLower(part)
	} else if to == ":camel" || to == ":c" {
		part = snakeToCamel(part)
	} else if to == ":snake" || to == ":s" {
		part = camelToSnake(part)
	}
	newName := dir + "/" + part + ext
	doRename(from, newName, 3)

}

func namePre(from string, to string, dir string, ext string) {
	basename := path.Base(from)
	part := strings.Split(basename, ".")[0]
	if to == "^d" || to == "^date" {
		date := fmt.Sprintf("%v", time.Now())
		date = strings.Split(date, " ")[0]
		datePre := strings.Split(date, "-")[0] + strings.Split(date, "-")[1] + strings.Split(date, "-")[2]
		newName := dir + "/" + datePre + "-" + part + ext
		doRename(from, newName, 4)
	} else if to == "^t" || to == "^time" {
		timeInCmd := fmt.Sprintf("%v", time.Now())
		timeInCmd = strings.Split(timeInCmd, ".")[0]
		timeInCmd = strings.Split(timeInCmd, " ")[0] + strings.Split(timeInCmd, " ")[1]
		timeInCmd = strings.Split(timeInCmd, "-")[0] + strings.Split(timeInCmd, "-")[1] + strings.Split(timeInCmd, "-")[2]
		timePre := strings.Split(timeInCmd, ":")[0] + strings.Split(timeInCmd, ":")[1] + strings.Split(timeInCmd, ":")[2]
		newName := dir + "/" + timePre + "-" + part + ext
		doRename(from, newName, 4)
	} else {
		if to[1] == '\\' {
			to = to[0:1] + to[2:]
		}
		newName := dir + "/" + to[1:] + part + ext
		doRename(from, newName, 4)
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
		doRename(from, newName, 4)
	} else if to == "+t" || to == "+time" {
		timeInCmd := fmt.Sprintf("%v", time.Now())
		timeInCmd = strings.Split(timeInCmd, ".")[0]
		timeInCmd = strings.Split(timeInCmd, " ")[0] + strings.Split(timeInCmd, " ")[1]
		timeInCmd = strings.Split(timeInCmd, "-")[0] + strings.Split(timeInCmd, "-")[1] + strings.Split(timeInCmd, "-")[2]
		timeSuf := strings.Split(timeInCmd, ":")[0] + strings.Split(timeInCmd, ":")[1] + strings.Split(timeInCmd, ":")[2]
		newName := dir + "/" + part + "-" + timeSuf + ext
		doRename(from, newName, 4)
	} else {
		if to[1] == '\\' {
			to = to[0:1] + to[2:]
		}
		newName := dir + "/" + part + to[1:] + ext
		doRename(from, newName, 4)
	}
}

func doRename(from string, newName string, exitCode int) {
	err := os.Rename(from, newName)
	if err != nil {
		fmt.Println("ERROR:", err)
		os.Exit(exitCode)
	}
	fmt.Printf("Renamed %s to %s\n", from, newName)
}

func snakeToCamel(s string) (sNew string) {
	sliceS := strings.Split(s, "_")
	for i := 0; i < len(sliceS); i++ {
		sliceS[i] = capitalizeFirstLetter(sliceS[i])
	}
	sNew = strings.Join(sliceS, "")
	return
}

func capitalizeFirstLetter(s string) string {
	if len(s) == 0 {
		return s
	}
	firstChar := s[0]
	if firstChar >= 'a' && firstChar <= 'z' {
		return string(firstChar-'a'+'A') + s[1:]
	}
	return s
}

func camelToSnake(s string) string {
	var result string
	for i, r := range s {
		if i > 0 && r >= 'A' && r <= 'Z' {
			result += "_"
		}
		result += string(r)
	}
	return strings.ToLower(result)
}

func listFilesInDirectory(dir string) ([]string, error) {
	files, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	var fileNames []string
	for _, file := range files {
		if !file.IsDir() {
			fileNames = append(fileNames, file.Name())
		}
	}

	return fileNames, nil
}
