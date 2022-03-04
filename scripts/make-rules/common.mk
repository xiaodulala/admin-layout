SHELL := /bin/bash

# 当前makefile所在的目录
COMMON_SELF_DIR := $(dir $(lastword $(MAKEFILE_LIST)))

# 设置项目根路径
ifeq ($(origin ROOT_DIR),undefined)
ROOT_DIR :=  $(abspath $(shell cd $(COMMON_SELF_DIR)/../../ && pwd -P))
endif

ifeq ($(origin OUTPUT_DIR),undefined)
OUTPUT_DIR := $(ROOT_DIR)/_output
$(shell mkdir -p $(OUTPUT_DIR))
endif

ifeq ($(origin TOOLS_DIR),undefined)
TOOLS_DIR := $(OUTPUT_DIR)/tools
$(shell mkdir -p $(TOOLS_DIR))
endif

ifeq ($(origin TMP_DIR),undefined)
TMP_DIR := $(OUTPUT_DIR)/tmp
$(shell mkdir -p $(TMP_DIR))
endif

# 设置版本编号
ifeq ($(origin VERSION),undefined)
VERSION := $(shell git describe --tags --always --match='v*')
endif

GIT_TREE_STATE := "dirty"
ifeq (,$(shell git status --procelain 2>/dev/null))
	GIT_TREE_STATE = "clean"
endif
GIT_COMMIT := $(shell git rev-parse HEAD)

# 最小的测试覆盖率
ifeq ($(origin COVERAGE),undefined)
COVERAGE := 60
endif

# The OS must be linux when building docker images
PLATFORMS ?= linux_amd64 linux_arm64
# The OS can be linux/windows/darwin when building binaries
# PLATFORMS ?= darwin_amd64 windows_amd64 linux_amd64 linux_arm64

# 设置指定平台
ifeq ($(origin PLATFORM),undefined)
	ifeq ($(origin GOOS),undefined)
		GOOS := $(shell go env GOOS)
	endif
	ifeq ($(origin GOARCH),undefined)
		GOARCH := $(shell go env GOARCH)
	endif
	PLATFORM := $(GOOS)_$(GOARCH)
# 使用linux作为镜像的操作系统
	IMAGE_PLAT := linux_$(GOARCH)
else
	GOOS := $(word 1,$(subst _, ,$(PLATFORM)))
	GOARCH := $(word 2,$(subst _, ,$(PLATFORM)))
	IMAGE_PLAT := $(PLATFORM)
endif

# linux命令设置
FIND := find . ! -path './third_party/*' ! -path './vendor/*'
XARGS := xargs --no-run-if-empty


# Makefile settings
ifndef V
MAKEFLAGS += --no-print-directory
endif

# 每次执行makefile时拷贝githook
# COPY_GITHOOK := $(shell cp -f githooks/* .git/hooks/)


# 指定工具的严重程度, include: BLOCKER_TOOLS(拦截类工具), CRITICAL_TOOLS(重要类工具), TRIVIAL_TOOLS(辅助类工具).
# Missing BLOCKER_TOOLS can cause the CI flow execution failed, i.e. `make all` failed.
# Missing CRITICAL_TOOLS can lead to some necessary operations failed. i.e. `make release` failed.
# TRIVIAL_TOOLS are Optional tools, missing these tool have no affect.
BLOCKER_TOOLS ?= gsemver golines go-junit-report golangci-lint  goimports codegen
CRITICAL_TOOLS ?= swagger mockgen gotests git-chglog github-release coscmd go-mod-outdated protoc-gen-go
TRIVIAL_TOOLS ?= depth go-callvis gothanks richgo rts kube-score coscli

COMMA := ,
SPACE :=
SPACE +=
