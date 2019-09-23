GOPATH=`realpath ../../../..`

.PHONY=clean

ddctl: main.go
	@GOPATH=$(GOPATH) go build -o ddctl

clean:
	@rm -f ./ddctl ./main
