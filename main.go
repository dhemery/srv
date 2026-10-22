package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"path"
)

var (
	address = ":0"
	dir     = "."
)

func init() {
	flag.StringVar(&address, "a", address, "`address` to serve")
	flag.StringVar(&dir, "d", dir, "`dir` to serve")
}

func main() {
	flag.Parse()

	listener, err := net.Listen("tcp4", address)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	http.Handle("/", http.FileServer(http.Dir(dir)))

	p, _ := os.Getwd()
	fmt.Println("Serving", path.Join(p, dir), "on http://"+listener.Addr().String())

	log.Fatal(http.Serve(listener, nil))
}
