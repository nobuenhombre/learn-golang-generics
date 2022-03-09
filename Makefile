PROJECT_NAME="learn-generics-golang"

GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTOOL=$(GOCMD) tool
GOTEST=$(GOCMD) test
GOMOD=$(GOCMD) mod
GOGET=$(GOCMD) get
GORUN=$(GOCMD) run

BIN_WINDOWS_PATH=bin/windows
BIN_LINUX_PATH=bin/linux

LOCAL_DOCKER_GROUP=go-lang
PATH_BUILDS = build/pkg
PATH_DEPLOYMENTS = deployments


help: Makefile
	@echo "Выберите опцию сборки "$(BINARY_NAME)":"
	@sed -n 's/^##//p' $< | column -s ':' |  sed -e 's/^/ /'

## all: Удалить старые сборки, скачать необходимые пакеты, протестировать, скомпилировать
all: clean deps test service/build

## clean: Удалить старые сборки
clean:
	$(GOCLEAN)
	rm -f cpu.out
	rm -f mem.out
	rm -f cover.out

## deps: Инициализация модулей, скачать все необходимые програме модули
deps:
	rm -f go.mod
	rm -f go.sum
	$(GOMOD) init $(PROJECT_NAME)
	$(GOGET) -u ./...

## build: Скомпилировать все
build: build-validator build-publisher

##build-httpserver

## fmt: fmt
fmt:
	go fmt ./...

## imp: imp
imp:
	./go-imports.sh

## fix-style: fix code style
fix-style: fmt imp

## before-git-push: Подготовка перед пушем
before-git-push: fmt imp

#=========================================
#
#   LINTERs
#
#=========================================
## lint: Проверка кода линтерами
lint: fmt imp lint-standart lint-bugs lint-complexity lint-format lint-performance lint-style lint-unused

## lint-standart: Проверка кода стандартным набором линтереров
lint-standart:
	golangci-lint run ./...

## lint-bugs: Проверка кода линтерами bugs
lint-bugs:
	golangci-lint run -p=bugs ./...

## lint-complexity: Проверка кода линтерами complexity
lint-complexity:
	golangci-lint run -p=complexity ./...

## lint-format: Проверка кода линтерами format
lint-format:
	golangci-lint run -p=format ./...

## lint-performance: Проверка кода линтерами performance
lint-performance:
	golangci-lint run -p=performance ./...

## lint-style: Проверка кода линтерами style
lint-style:
	golangci-lint run -p=style ./...

## lint-unused: Проверка кода линтерами unused
lint-unused:
	golangci-lint run -p=unused ./...

#=========================================
#
#   TESTS
#
#=========================================
## test: Запустить тесты
test:
	$(GOTEST) -v ./...

## coverage: Получить информацию о покрытии тестами кода
cover:
	$(GOTEST) -coverprofile=cover.out ./...
	$(GOTOOL) cover -html=cover.out -o cover.html
	rm -f cover.out

## profile-cpu: Профильрование CPU
profile-cpu:
	$(GOTEST) -bench . -benchmem -cpuprofile=cpu.out .
	$(GOTOOL) pprof $(BINARY_NAME) cpu.out
	rm -f cpu.out

## profile-mem: Профилирование памяти
profile-mem:
	$(GOTEST) -bench . -benchmem -memprofile=mem.out -memprofilerate=1 .
	$(GOTOOL) pprof $(BINARY_NAME) mem.out
	rm -f mem.out
