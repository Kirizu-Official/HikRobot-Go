# HikRobot-Go
HikRobot-Go是一个使用Go语言编写的海康机器人SDK，旨在提供一个简单易用的接口来控制海康机器人的功能。
## 安装
```
go get -u github.com/Kirizu-Official/HikRobot-Go
```
### 安装MVS
请查看当前Release版本号，通常来说前几位对应MVS的版本号，如`4.5.1.1`对应MVS`4.5.1`版本，`4.6.0.1`对应MVS`4.6.0`版本，然后在 [海康官网](https://www.hikrobotics.com/cn/machinevision/service/download/?module=0) 下载对应版本的MVS客户端安装，通常开发环境只需要安装MVS客户端即可，无需安装 MVS Runtime 组件包。
### 配置CGO
本项目采用CGO的方式调用SDK，使用前请确保CGO已正确配置，环境中必须包含gcc编译器
```bash
gcc --version
```
```text
gcc.exe (Rev5, Built by MSYS2 project) 15.1.0
Copyright (C) 2025 Free Software Foundation, Inc.
This is free software; see the source for copying conditions.  There is NO
warranty; not even for MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.
```
### 配置pkg-config
您需要下载 [hik-mvs.pc](https://github.com/Kirizu-Official/HikRobot-Go/blob/main/hik-mvs.pc) 文件放到您的pkg-config库目录中，打开这个文件并修改第一行`prefix`设置为MVS安装目录，注意目录分隔符必须是`/`而不是`\`。

设置完成后通过以下命令检查pkg-config是否已识别：
```bash
pkg-config --modversion hik-mvs
```
```text
1.0
```

## 使用
请参考 [test.go](https://github.com/Kirizu-Official/HikRobot-Go/blob/main/test.go)

