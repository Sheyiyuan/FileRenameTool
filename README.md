以下是修改后的文档：

# File Rename Tool

## Profile
This is a handy tool designed to facilitate the convenient renaming of your files.

I have just embarked on learning Golang. This represents my very first Golang project. If you have any suggestions, please do let me know. Thanks! 😊

## Author
**Sheyiyuan**

E-mail: sheyiyuantan90@qq.com

Github: @[Sheyiyuan](http://github.com/Sheyiyuan)

## Introduction

Welcome to utilize this tool!

### Help
```shell
rename -h
rename --help
```
To access the help document.

### Modify File Name and Extension
```shell
rename path/to/your/file.extension newName
```
This command can modify your file to path/to/your/newName.extension.

**Tip**: You cannot modify the extension using this command. For instance, if you run "rename a.txt b.md", you will obtain "b.md.txt" instead of "b.md". If you wish to modify the extension, execute the following command.
```shell
rename path/to/your/file.extension.newExt
```
This command enables you to modify the extension. Remember to input a '.' before the new extension.

## Add Prefix and Suffix
```shell
rename path/to/your/file ^pre
rename path/to/your/file +suf
```
You can add a prefix and suffix by running these commands. You can use 'd'/'date' and 't'/'time' to utilize the date or time as a prefix or suffix. If you merely want to add these words, input '\\\\' before them.

## License
MIT[README.md](README.md)