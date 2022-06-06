# TWReporter's Golang Backend API

## Environment 
### Development

#### Go module
After go-api@5.0.0 is released, go-api no longer needs to be developed within $GOPATH/src directory thanks to the go module support. Make sure your go version is go@1.11 or higher to have full compatibility. You can clone to anywhere outside the $GOPATH as you wish.

```golang
$ git clone github.com/twreporter/go-api
$ cd go-api

// Run test
$ go test ./...

// Use makefile
make start
// Or
$ go run main.go

// Build server binaries
$ go build -o go-api
$ ./go-api