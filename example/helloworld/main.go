package main

import (
	"fmt"
	"net"

	nfs "github.com/willscott/go-nfs"
	"github.com/willscott/go-nfs/filesystem/basefs"
	nfshelper "github.com/willscott/go-nfs/helpers"
)

func main() {

	listener, err := net.Listen("tcp", ":5566")
	if err != nil {
		fmt.Printf("Failed to listen: %v\n", err)
		return
	}
	fmt.Printf("Server running at %s\n", listener.Addr())

	// fs := basefs.NewOsFS("/path/to/dir")
	fs := basefs.NewMemMapFS()
	fileName := "hello.txt"
	f, err := fs.Create(fileName)
	if err != nil {
		fmt.Printf("Failed to create file: %v\n", err)
		return
	}
	fs.Chown(fileName, 502, 20)
	fs.Chmod(fileName, 0644)
	fs.MkdirAll("test_dir", 0755)
	_, _ = f.Write([]byte("hello world"))
	_ = f.Close()

	handler := nfshelper.NewNullAuthHandler(fs)
	cacheHelper := nfshelper.NewCachingHandler(handler, 1024)
	fmt.Printf("%v", nfs.Serve(listener, cacheHelper))
}
