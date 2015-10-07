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

	str := string(data)
	log.Printf("%#v", str)

	if len(data) == 0 {
		return
	}
	switch data[0] {
	case 'p':
		var clip string
		clip, err = clipboard.ReadAll()
		if err != nil {
			log.Println(err)
			return
		}
		_, err = c.Write([]byte(clip))
	case 'y':
		err = clipboard.WriteAll(str[1:])
	default:
		log.Println("command is missing")
		return
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
