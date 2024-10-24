# File Rename Tool

## Profile
This is a handy tool designed to facilitate the convenient renaming of your files.

I have just embarked on learning Golang. This represents my very first Golang project. If you have any suggestions, please do let me know. Thanks! 😊

## Author
**Sheyiyuan**

E-mail: sheyiyuantan90@qq.com

Github: @[Sheyiyuan](http://github.com/Sheyiyuan)

## Function Introduction

### (1) Modify file name and extension
1. Modify file name (cannot directly modify extension):
   `rename path/to/your/file.extension newName`
2. Modify file extension (add '.' before new extension):
   `rename path/to/your/file.extension.newExt`

### (2) Transform file name
1. Transform to upper case characters:
   `rename path/to/your/file :upper,:u`
2. Transform to lower case characters:
   `rename path/to/your/file :lower,:l`
3. Transform camel case to snake case:
   `rename path/to/your/file :camel,:c`
4. Transform snake case to snake case (possibly keeping snake case unchanged):
   `rename path/to/your/file :snake,:s`

### (3) Add prefix and suffix
1. Add prefix:
   `rename path/to/your/file ^pre`
2. Add suffix:
   `rename path/to/your/file +suf`

### (4) Use date or time as prefix or suffix
When using 'd'/'date' or 't'/'time' as prefix or suffix, add '\' before them.

### (5) Rename by folder
If you want to rename by folder, add '-f' as the first argument. Note that commands with "-f" will ignore all subfolders.

## Usage
1. View help document:
   `rename -h,--help`
2. Select the corresponding command according to your specific needs to perform file renaming operations.

## 4. License
This project follows the [MIT](license) license, allowing users to freely use, modify, and distribute this software under the terms of the license.