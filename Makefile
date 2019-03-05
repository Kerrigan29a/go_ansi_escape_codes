VERSION_FLAGS=-ldflags "-X main.version=`cat VERSION` -X main.date=`date -u +%Y/%m/%d-%H:%M:%S`"

all: build

build: build_extended_colors_demo build_rgb_colors_demo

build_extended_colors_demo:
	go build $(VERSION_FLAGS) -o extended_colors_demo cmd/extended_colors_demo/main.go

build_rgb_colors_demo:
	go build $(VERSION_FLAGS) -o rgb_colors_demo cmd/rgb_colors_demo/main.go

clean:
	go clean
	-rm extended_colors_demo
	-rm rgb_colors_demo

run_extended_colors_demo:
	go run $(VERSION_FLAGS) cmd/extended_colors_demo/main.go

run_rgb_colors_demo:
	go run $(VERSION_FLAGS) cmd/rgb_colors_demo/main.go

test:
	go test

update:
	go mod vendor

format:
	find . -path ./vendor -prune -o -name '*.go' -exec gofmt -s -w {} \;
