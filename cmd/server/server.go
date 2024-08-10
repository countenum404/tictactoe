package main

import (
	"bufio"
	"encoding/gob"
	"log"
	"net"
	"shabashov/tic-tac-toe/internal/board"
	"shabashov/tic-tac-toe/internal/player"
	"sync"
)

type Message struct {
	Player player.Player
	X      int
	Y      int
}

type TcpServer struct {
	Ip       string
	Port     string
	Messages []string
	Clients  []net.Conn
	b        board.Board
}

func (t *TcpServer) Run(port string) {
	t.Ip = "127.0.0.1"
	t.Port = port
	l, e := net.Listen("tcp4", t.Ip+":"+t.Port)
	t.b = board.Board{}
	t.b.Init()
	if e != nil {
		log.Println(e)
	}
	log.Println("Server started")
	for {
		c, err := l.Accept()
		if err != nil {
			log.Println(err)
		}
		t.Clients = append(t.Clients, c)
		go t.HandleConnection(c)
	}
}

func (t *TcpServer) addMessage(m Message, conn net.Conn) {
	if err := t.b.SetCell(m.X, m.Y, m.Player.Mark); err != nil {
		log.Println(err)
	}

	wg := sync.WaitGroup{}
	wg.Add(len(t.Clients))
	for _, c := range t.Clients {
		go func() {
			w := bufio.NewWriter(c)
			gob.NewEncoder(w).Encode(&t.b)
			w.Flush()
			wg.Done()
		}()
	}
	wg.Wait()
}

func (t *TcpServer) HandleConnection(conn net.Conn) {
	defer conn.Close()

	r := bufio.NewReader(conn)
	for {
		var msg Message
		if err := gob.NewDecoder(r).Decode(&msg); err != nil {
			log.Println(err)
			conn.Close()
			break
		} else {
			t.addMessage(msg, conn)
		}
	}
}

func main() {
	s := TcpServer{}
	s.Run("8234")
}
