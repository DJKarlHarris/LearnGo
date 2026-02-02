.PHONY: help variables arrays slices maps structs functions interfaces goroutines channels all

help:
	@echo "learn go"
	@echo "usage"
	@echo "  make variables    - run variables learning example"
	@echo "  make arrays       - run arrays learning example"
	@echo "  make slices       - run slices learning example"
	@echo "  make maps         - 运行映射学习示例"
	@echo "  make structs      - 运行结构体学习示例"
	@echo "  make functions    - 运行函数学习示例"
	@echo "  make interfaces   - 运行接口学习示例"
	@echo "  make goroutines   - 运行协程学习示例"
	@echo "  make channels     - 运行通道学习示例"
	@echo "  make all          - 运行所有学习示例"

variables:
	go run main.go variables

arrays:
	go run main.go arrays

slices:
	go run main.go slices

maps:
	go run main.go maps

structs:
	go run main.go structs

functions:
	go run main.go functions

interfaces:
	go run main.go interfaces

goroutines:
	go run main.go goroutines

channels:
	go run main.go channels

all: variables arrays slices maps structs functions interfaces goroutines channels