# Go语言学习框架

这是一个用于学习Go语言特性的交互式框架。通过运行不同的命令可以体验各种Go语言特性。

## 项目结构

```
LearnGo/
├── main.go                 # 主程序入口
├── go.mod                  # Go模块定义
├── readme.md               # 项目说明
└── lessons/                # 学习模块目录
    ├── variables/          # 变量学习模块
    ├── arrays/             # 数组学习模块
    ├── slices/             # 切片学习模块
    ├── maps/               # 映射学习模块
    ├── structs/            # 结构体学习模块
    ├── functions/          # 函数学习模块
    ├── interfaces/         # 接口学习模块
    ├── goroutines/         # 协程学习模块
    └── channels/           # 通道学习模块
```

## 使用方法

```bash
# 运行变量学习示例
go run main.go variables

# 运行数组学习示例
go run main.go arrays

# 运行切片学习示例
go run main.go slices

# 运行映射学习示例
go run main.go maps

# 运行结构体学习示例
go run main.go structs

# 运行函数学习示例
go run main.go functions

# 运行接口学习示例
go run main.go interfaces

# 运行协程学习示例
go run main.go goroutines

# 运行通道学习示例
go run main.go channels
```

## 学习模块

目前包含以下学习模块：

- **变量(Variables)**: 演示Go中不同类型的变量声明和使用
- **数组(Arrays)**: 演示固定大小数组的使用方法
- **切片(Slices)**: 演示动态数组（切片）的使用方法
- **映射(Maps)**: 演示键值对集合的使用方法
- **结构体(Structs)**: 演示自定义数据类型的方法
- **函数(Functions)**: 演示函数定义、参数传递和返回值
- **接口(Interfaces)**: 演示Go的接口概念
- **协程(Goroutines)**: 演示并发编程基础
- **通道(Channels)**: 演示协程间通信机制

## 扩展性

此框架设计具有良好的扩展性，你可以轻松添加新的学习模块：

1. 创建新的学习模块目录和Go文件
2. 在main.go中导入新的模块
3. 在main函数的switch语句中添加新的case
4. 在帮助信息中更新可用模块列表

## 贡献

欢迎提交PR来增加更多的Go语言学习模块！