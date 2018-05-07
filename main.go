package main

import (
	"fmt"
	"log"
	"net"

	"github.com/tonky/nen/player"
)

const serverAddr = "127.0.0.1:40000"

//Connect udp
func startServer(addr string) (*net.UDPConn, error) {
	sAddr, err := net.ResolveUDPAddr("udp", addr)

	if err != nil {
		return nil, err
	}

	return net.ListenUDP("udp", sAddr)
}

func main() {
	sConn, err := startServer(serverAddr)

	if err != nil {
		log.Fatal("Can't start server")
	}

	fmt.Println("listening on ", sConn.LocalAddr().String())

	defer sConn.Close()

	p1, err1 := player.New(serverAddr, "Player 1")
	p2, err2 := player.New(serverAddr, "Player 2")

	if err1 != nil {
		log.Fatal("Can't connect to server")
	}

	if err2 != nil {
		log.Fatal("Can't connect to server")
	}

	p1.Init()
	p2.Init()

	rb := make([]byte, 1024)

	for {
		i, addr, err := sConn.ReadFrom(rb)

		fmt.Println("Got from client: ", string(rb), i, addr, err)
	}

	//simple read
	// sConn.WriteTo([]byte("Hello from server"), addr)

	// fmt.Println("Sent to client")

	// time.Sleep(400 * time.Millisecond)

	// sConn.Close()

	//simple write
	// pc.WriteTo([]byte("Hello from server"), addr)
}
