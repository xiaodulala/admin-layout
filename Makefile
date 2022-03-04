
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
                   Example: make build BINS="dss iam-authz-server"
  IMAGES           Backend images to make. Default is all of cmd starting with iam-.
                   This option is available when using: make image/image.multiarch/push/push.multiarch
                   Example: make image.multiarch IMAGES="iam-apiserver iam-authz-server"
  REGISTRY_PREFIX  Docker registry prefix. Default is marmotedu.
                   Example: make push REGISTRY_PREFIX=ccr.ccs.tencentyun.com/marmotedu VERSION=v1.6.2
  PLATFORMS        The multiple platforms to build. Default is linux_amd64 and linux_arm64.
                   This option is available when using: make build.multiarch/image.multiarch/push.multiarch
                   Example: make image.multiarch IMAGES="iam-apiserver iam-pump" PLATFORMS="linux_amd64 linux_arm64"
  VERSION          The version information compiled into binaries.
                   The default is obtained from gsemver or git.
  V                Set to 1 enable verbose build. Default is 0.
endef
export USAGE_OPTIONS

# ==============================================================================
# Targets

## clean: Remove all files that are created by building.
.PHONY: clean
clean:
	@echo "===========> Cleaning all build output"
	@-rm -vrf $(OUTPUT_DIR)

## build: Build source code for host platform.
.PHONY: build
build:
	@$(MAKE) go.build


## build.multiarch: Build source code for multiple platforms. See option PLATFORMS.
.PHONY: build.multiarch
build.multiarch:
	@$(MAKE) go.build.multiarch


## lint: Check syntax and styling of go sources.
.PHONY: lint
lint:
	@$(MAKE) go.lint

## install: Install iam system with all its components.
.PHONY: install
install:
	@$(MAKE) tools.install

## swagger: Generate swagger document.
.PHONY: swagger
swagger:
	@$(MAKE) swagger.run
## serve-swagger: Serve swagger spec and docs.
.PHONY: swagger.serve
serve-swagger:
	@$(MAKE) swagger.serve

## tidy
.PHONY: tidy
tidy:
	@$(GO) mod tidy -go=1.17

## image: Build docker images for host arch
.PHONY: image
image:
	@$(MAKE) image.build

## image.multiarch: Build docker images for multiple platforms. See option PLATFORMS.
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
