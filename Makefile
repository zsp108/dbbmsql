# 项目名称
APP_NAME = dbbmsql

# Go 命令
GO = go
GOBUILD = $(GO) build
GOMOD = $(GO) mod

# 路径
BIN_DIR = bin
PKG_DIR = pkg
CMD_DIR = cmd
CLIENT_MAIN = $(CMD_DIR)/dbbmsql/main.go

# 输出文件
OUTPUT = $(BIN_DIR)/$(APP_NAME)

# 默认任务：编译项目
.PHONY: all
all: build

# 下载依赖到 vendor 目录
.PHONY: vendor
vendor:
	$(GOMOD) vendor

# 清理项目的二进制文件和缓存
.PHONY: clean
clean:
	$(GO) clean -modcache
	rm -rf $(BIN_DIR)/* 

# 编译项目
.PHONY: build
build: vendor
	$(GOBUILD) -o $(OUTPUT) $(CLIENT_MAIN)

# 运行项目
.PHONY: run
run: build
	@./$(OUTPUT)

# 更新依赖
.PHONY: tidy
tidy:
	$(GOMOD) tidy

# 测试项目
.PHONY: test
test:
	$(GO) test ./...

。PHONY: gitlint
gitlint:
	- ./scripts/gitlint --subject-regex "^(feat|fix|docs|style|refactor|test|chore|ci|perf)(\([a-zA-Z0-9-_/]+\))?: .+"

# 帮助信息
.PHONY: help
help:
	@echo "可用命令："
	@echo "  make build    编译项目"
	@echo "  make          编译项目"
	@echo "  make vendor   下载依赖到 vendor 目录"
	@echo "  make clean    清理项目的二进制文件和缓存"
	@echo "  make run      运行项目"
	@echo "  make tidy     更新依赖"
	@echo "  make test     运行测试"
	@echo "  make help     显示帮助信息"
	@echo "  make gitlint  检查提交信息是否符合规范"