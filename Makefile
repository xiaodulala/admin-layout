
# build all by default
.DEFAULT_GOAL := all

.PHONY: all
all: clean tidy swagger  build   # lint暂时不能通过

# ==============================================================================
# Build options
ROOT_PACKAGE=github.com/xiaodulala/admin-layout
VERSION_PACKAGE=github.com/xiaodulala/admin-layout/pkg/version


# ==============================================================================
# Includes

include scripts/make-rules/common.mk # make sure include common.mk at the first include line
include scripts/make-rules/golang.mk
include scripts/make-rules/tools.mk
include scripts/make-rules/swagger.mk
include scripts/make-rules/image.mk

# ==============================================================================
# Usage

define USAGE_OPTIONS

Options:
  DEBUG            Whether to generate debug symbols. Default is 0.
  BINS             The binaries to build. Default is all of cmd.
                   This option is available when using: make build/build.multiarch
                   Example: make build BINS="appctl application"
  IMAGES           Backend images to make.
                   This option is available when using: make image/image.multiarch/push/push.multiarch
                   Example: make image.multiarch IMAGES="appctl application"
  REGISTRY_PREFIX  Docker registry prefix.
                   Example:
  PLATFORMS        The multiple platforms to build. Default is linux_amd64 and linux_arm64.
                   This option is available when using: make build.multiarch/image.multiarch/push.multiarch
                   Example: make image.multiarch IMAGES="appctl application" PLATFORMS="linux_amd64 linux_arm64"
  VERSION          The version information compiled into binaries.
                   The default is obtained from gsemver or git.
  V                Set to 1 enable verbose build. Default is 0.
endef
export USAGE_OPTIONS

# ==============================================================================
# Targets

## clean: 删除编译时产生的所有文件.
.PHONY: clean
clean:
	@echo "===========> Cleaning all build output"
	@-rm -vrf $(OUTPUT_DIR)

## build: 根据当前平台编译对应的application可执行文件.
.PHONY: build
build:
	@$(MAKE) go.build


## build.multiarch: 编译多平台的可执行文件. See option PLATFORMS.
.PHONY: build.multiarch
build.multiarch:
	@$(MAKE) go.build.multiarch


## lint: 代码静态检查.
.PHONY: lint
lint:
	@$(MAKE) go.lint

## install: 安装开发项目时的所有工具(网络不好请重试).
.PHONY: install
install:
	@$(MAKE) tools.install

## swagger: 生成swagger文档.
.PHONY: swagger
swagger:
	@$(MAKE) swagger.run
## serve-swagger: 启动swagger接口服务
.PHONY: serve-swagger
serve-swagger:
	@$(MAKE) swagger.serve

## tidy: 安装go mod 依赖
.PHONY: tidy
tidy:
	@$(GO) mod tidy -go=1.17

## image: 以当前平台制作镜像
.PHONY: image
image:
	@$(MAKE) image.build

## image.multiarch: 生成多平台镜像. See option PLATFORMS.
.PHONY: image.multiarch
image.multiarch:
	@$(MAKE) image.build.multiarch

## help: Show this help info.
.PHONY: help
help: Makefile
	@echo -e "\nUsage: make <TARGETS> <OPTIONS> ...\n\nTargets:"
	@sed -n 's/^##//p' $< | column -t -s ':' | sed -e 's/^/ /'
	@echo "$$USAGE_OPTIONS"


.PHONY: testmk
testmk:
	@echo $(wildcard ${ROOT_DIR}/build/docker/*)
	@echo $(foreach image,${IMAGES_DIR},$(notdir ${image}))
	@echo  $(filter-out tools,$(foreach image,${IMAGES_DIR},$(notdir ${image})))
	@echo $(addprefix image.build., $(addprefix $(IMAGE_PLAT)., $(IMAGES)))
	@echo $(addprefix go.build., $(addprefix $(PLATFORM)., $(ROOT_DIR)))
	@echo $(ROOT_PACKAGE)
	@echo $(filter-out %docs, $(wildcard ${ROOT_DIR}/cmd/*))
