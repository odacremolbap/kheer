# PLEASE: when relevant changes are introduced to this file, document at docs/building.md accordingly

REPO ?= github.com/kheer/kheer
# list of output binaries
ALL_BIN ?= kheer

# CMD_SRC_DIR where subfolders by the binaries names are located, containing entry source code (main)
CMD_SRC_DIR := cmd
ALL_SRC_DIRS := cmd pkg

# list of all supported architectures
ALL_ARCH := amd64 arm64
ARCH ?= amd64

# list of all supported OSes
ALL_OS := linux darwin
OS ?= $(shell go env GOOS)

# image tag for go
GO_IMAGE := golang:1.12.5
COMPILER_IMAGE := kheer-compile-$(GO_IMAGE)

GCO_ENABLED := 0
OUTPUT_DIR := _output


export GO111MODULE=on

all: \
		build

.PHONY: build
build: \
		prebuild-bin-$(ARCH)-$(OS)

.PHONY: build-all
build-all: prebuild-arch

# There targets will iterate in a nested way: architectures, OSes and output binaries.
# All iterations will end up calling launch-build after setting variables for compilation
prebuild-arch: $(foreach arch, $(ALL_ARCH), prebuild-os-$(arch))
	$(NOOP)
prebuild-os-%: $(foreach os, $(ALL_OS), prebuild-bin-%-$(os))
	$(NOOP)
prebuild-bin-%: $(foreach bin, $(ALL_BIN), prebuild-launch-%-$(bin))
		$(NOOP)
prebuild-launch-%:
	$(eval STR = $(subst -, ,$@))
	$(eval ARCH = $(word 3, $(STR)))
	$(eval OS = $(word 4, $(STR)))
	$(eval BINARY = $(word 5, $(STR)))

	@$(MAKE) --no-print-directory BINARY=$(BINARY) ARCH=$(ARCH) OS=$(OS) launch-build

launch-build: compiler-image
	$(eval OUTPUT_BIN_DIR = $(OUTPUT_DIR)/$(OS)/$(ARCH))

ifneq ($(ARCH)-$(OS),arm64-darwin)
	$(info ****** launch-build $(BINARY) for $(OS)/$(ARCH))
	mkdir -p $(OUTPUT_BIN_DIR)

	docker run -ti --rm  \
					-v "$$(pwd):/kheer" \
					-e "CGO_ENABLED=$(CGO_ENABLED)" \
					-e "GOOS=$(OS)" \
					-e "GOARCH=$(ARCH)" \
					-e "GO111MODULE=on" \
					$(COMPILER_IMAGE) \
					go build $(GO_FLAGS) -ldflags "$(GO_LDFLAGS)" -o ./$(OUTPUT_BIN_DIR)/$(BINARY) ./$(CMD_SRC_DIR)/$(BINARY)

else
	@rm -rf $(OUTPUT_BIN_DIR)
endif

# compiler-image creates a go based image that already contains the cached dependences.
# This will speed up compiling at the launch-build step
.PHONY: compiler-image
compiler-image:
	$(info ****** building cached dependencies container for compiling)
	docker build -f build/Dockerfile . -t $(COMPILER_IMAGE) \
		--build-arg GO_IMAGE=$(GO_IMAGE)


.PHONY: check
check: test test-race vet gofmt

.PHONY: test
test:
	$(info ****** test)
	go test -mod=readonly ./...

.PHONY: test-race
test-race: | test
	$(info ****** test-race)
	go test -race -mod=readonly ./...

.PHONY: vet
vet: | test
	$(info ****** vet)
	go vet ./...

.PHONY: clean
clean:
	$(info ****** clean)
	rm -rf _output

.PHONY: gofmt
gofmt:
	$(info ****** gofmt)
	test -z "$(shell gofmt -s -l -d -e $(ALL_SRC_DIRS) | tee /dev/stderr)"
