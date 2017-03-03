package main

import (
	"fmt"
	"time"
	"github.com/SvenWiltink/dispatcher/dispatcher"
	"net"
	"log"
	"io"
)

type ConnectionJob struct {
	conn net.Conn
}

func (cj *ConnectionJob) Execute() {
	defer cj.conn.Close()

	tmp := make([]byte, 256) // using small tmo buffer for demonstrating
	_, err := cj.conn.Read(tmp)
	if err != nil {
		if err != io.EOF {
			fmt.Println("read error:", err)
		}
	}

	cj.conn.Write(tmp)

	fmt.Printf("%s", tmp)
}

func main() {

	jobDispatcher := dispatcher.NewJobDispatcher(10, 10)
	jobDispatcher.Start()

	ln, err := net.ListenTCP("tcp", &net.TCPAddr{Port: 9000})

	if err != nil {
		panic(err)
	}

	for {
		con, err := ln.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		con.SetReadDeadline(time.Now().Add(10 * time.Second))
		conJob := &ConnectionJob{conn: con}
		jobDispatcher.Jobs <- conJob
	}

}