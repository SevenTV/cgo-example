linux:
	CGO_ENABLED=1 go build

windows:
	CGO_ENABLED=1 CC=/usr/bin/x86_64-w64-mingw32-gcc CXX=x86_64-w64-mingw32-g++ GOOS=windows GOARCH=amd64 go build
