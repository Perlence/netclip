package main

import (
	"bytes"
	"flag"
	"io/ioutil"
	"log"
	"net"

	"github.com/Perlence/netclip/Godeps/_workspace/src/github.com/atotto/clipboard"

	"github.com/Perlence/netclip/config"
)

func init() {
	flag.StringVar(&config.Addr, "a", config.Addr, "start server on this addr")
	flag.BoolVar(&config.Unix, "u", config.Unix, "convert Windows line endings to Unix on paste")
}

func main() {
	flag.Parse()

	l, err := net.Listen("tcp", config.Addr)
	fatalOnError(err)
	defer l.Close()
	log.Println("netclip server is listening on", config.Addr)

	for {
		conn, err := l.Accept()
		fatalOnError(err)

		go handle(conn)
	}
}

func handle(c net.Conn) {
	defer c.Close()
	data, err := ioutil.ReadAll(c)
	if err != nil {
		log.Println(err)
		return
	}

	if len(data) == 0 {
		var clip string
		clip, err = clipboard.ReadAll()
		if err != nil {
			log.Println(err)
			return
		}
		bclip := []byte(clip)
		if config.Unix {
			bclip = bytes.Replace(bclip, []byte{'\r', '\n'}, []byte{'\n'}, -1)
		}
		log.Printf("sending %d bytes", len(bclip))
		_, err = c.Write(bclip)
	} else {
		log.Printf("received %d bytes", len(data))
		err = clipboard.WriteAll(string(data))
	}
	if err != nil {
		log.Println(err)
		return
	}
}

func fatalOnError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
