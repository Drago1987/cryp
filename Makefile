.PHONY: all clean

GO_BIN ?= "go"

all: polygon_x86 yocto

yocto: polygon_arm

clean:
	rm -rf build


builddir:
	mkdir -p build

polygon: builddir
	cd ./go/polygon

polygon_x86: export GOARCH=amd64
polygon_x86: polygon
	cd ./go/polygon; ${GO_BIN} build -v -o ../../build/polygon_x86

polygon_arm: export GOARCH=arm
polygon_arm: export GOOS=linux
polygon_arm: export GOARM=7
polygon_arm: export CGO_ENABLED=0
polygon_arm: polygon
	cd ./go/polygon; ${GO_BIN} build -v -o ../../build/polygon_arm