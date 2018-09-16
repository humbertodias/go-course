fmt:
	go fmt ./...

dep:
	go get github.com/tools/godep
	godep get *