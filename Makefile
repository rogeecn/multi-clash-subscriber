.PHONEY: build
build:
	CGO_ENABLE=0 go build