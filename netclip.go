package main

import (
	"io/ioutil"
	"log"
	"net"

	"github.com/Perlence/netclip/Godeps/_workspace/src/github.com/atotto/clipboard"

	"github.com/Perlence/netclip/config"
)

func main() {
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
