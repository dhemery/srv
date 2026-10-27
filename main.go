package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"path"
	"strings"
)

var (
	host = "localhost"
	port = 0
	dir  = "."
)

func init() {
	flag.StringVar(&host, "host", host, "host address to serve")
	flag.IntVar(&port, "port", port, "port to serve (default any available port)")
	flag.StringVar(&dir, "dir", dir, "directory to serve")
}

func main() {
	flag.Usage = usage
	flag.Parse()

	address := fmt.Sprintf("%s:%d", host, port)
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

const usageText = `
srv serves a directory of static files on the specified host and port.

Usage: srv [flags]

Flags:
`

func usage() {
	fmt.Fprintln(flag.CommandLine.Output(), strings.TrimSpace(usageText))
	flag.PrintDefaults()

}
