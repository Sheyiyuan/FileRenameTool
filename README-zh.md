# 文件重命名工具

## 概述
这是一个能帮助你方便地重命名文件的工具。

我刚刚开始学习 Go 语言。这是我的第一个 Go 语言项目。如果你有任何建议，请告诉我，谢谢！😊

## 作者
**Sheyiyuan**

电子邮箱：sheyiyuantan90@qq.com

Github：@[Sheyiyuan](http://github.com/Sheyiyuan)

## 功能介绍

### （一）修改文件名和扩展名
1. 修改文件名（不能直接修改扩展名）：
   `rename path/to/your/file.extension newName`
2. 修改文件扩展名（在新扩展名前添加'.'）：
   `rename path/to/your/file.extension.newExt`

### （二）转换文件名
1. 转换为大写字符：
   `rename path/to/your/file :upper,:u`
2. 转换为小写字符：
   `rename path/to/your/file :lower,:l`
3. 驼峰命名转换为蛇形命名：
   `rename path/to/your/file :camel,:c`
4. 蛇形命名转换为蛇形命名（可能是保持蛇形命名不变）：
   `rename path/to/your/file :snake,:s`

### （三）添加前缀后缀
1. 添加前缀：
   `rename path/to/your/file ^pre`
2. 添加后缀：
   `rename path/to/your/file +suf`

### （四）使用日期或时间作为前缀后缀
使用'd'/'date'或't'/'time'作为前缀后缀时，需在前面添加'\\'。

### （五）按文件夹重命名
如果要按文件夹进行重命名操作，可在命令开头添加'-f'参数。但需注意，带有'-f'参数的命令将忽略所有子文件夹。

## 使用方法
1. 查看帮助文档：
   `rename -h,--help`
2. 根据具体需求选择上述相应的命令进行文件重命名操作。

## 许可证
本项目遵循 [MIT](license) 许可证，允许用户在遵循许可证条款的前提下自由使用、修改和分发本软件。