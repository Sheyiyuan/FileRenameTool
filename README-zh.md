# 文件重命名工具

## 概述
这是一个能帮助你方便地重命名文件的工具。

我刚刚开始学习 Go 语言。这是我的第一个 Go 语言项目。如果你有任何建议，请告诉我，谢谢！😊

## 作者
**Sheyiyuan**

电子邮箱：sheyiyuantan90@qq.com

Github：@[Sheyiyuan](http://github.com/Sheyiyuan)

## 介绍

欢迎使用这个工具！

### 帮助
```shell
rename -h
rename --help
```
查看帮助文档。

### 修改文件名和扩展名
```shell
rename path/to/your/file.extension newName
```
它可以将你的文件修改为 path/to/your/newName.extension。

**提示**：你不能通过这个命令修改扩展名。例如，如果你运行“rename a.txt b.md”，你将得到“b.md.txt”而不是“b.md”。如果你想修改扩展名，应该运行下一个命令。
```shell
rename path/to/your/file.extension.newExt
```
这个命令可以修改扩展名。你应该在新扩展名前加上一个“.”。

## 添加前缀和后缀
```shell
rename path/to/your/file ^pre
rename path/to/your/file +suf
```
你可以通过运行这个命令添加前缀和后缀。

你可以使用“d”/“date”和“t”/“time”来使用日期或时间作为前缀/后缀。如果你只是想添加这些单词，在它们前面输入“\\”。

## 许可证
MIT 许可证。