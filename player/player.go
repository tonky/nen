package player

import (
	"fmt"
	"net"
)

type Position struct {
	x, y, z int
}

func (pos Position) String() string {
	return fmt.Sprintf("Position: %d %d %d", pos.x, pos.y, pos.z)
}

type Player struct {
	server net.Conn
	Name   string
	Position
}

func (p Player) String() string {
	return fmt.Sprintf("Player| Name: %s, position: %s", p.Name, p.Position)
}

func New(addr, name string) (*Player, error) {
	fmt.Println("Connecting UDP Client ", name)

	conn, err := net.Dial("udp", addr)

	if err != nil {
		return nil, err
	}

	fmt.Println("Client dialled: ", conn)

	return &Player{conn, name, Position{1, 1, 1}}, nil

	//simple write
}

func (p Player) Init() error {
	fmt.Println("Player init")

	//simple write
	i, err := p.server.Write([]byte(p.String()))

	if err != nil {
		return err
	}

	fmt.Printf("Written %d bytes to server\n", i)

	return nil
}

func (p Player) Move(pos Position) error {
	fmt.Printf("Moving player %s to new position %s", p, pos)

	p.Position = pos
}

/*
func clientUDP(message string) error {
	fmt.Println("Starting UDP Client")

	conn, err := net.Dial("udp", serverAddr)

	if err != nil {
		return err
	}

	fmt.Println("Client dialled: ", conn)

	//simple write
	conn.Write([]byte(message))

	fmt.Println("Written to server")

	time.Sleep(200 * time.Millisecond)

	//simple Read
	buffer := make([]byte, 1024)
	i, err := conn.Read(buffer)

	fmt.Println("Got from server: ", string(buffer), i, err)

	conn.Close()

	return nil
}

*/
