package main

import (
	"bufio"
	"encoding/gob"
	"fmt"
	"log"
	"net"
	"shabashov/tic-tac-toe/internal/board"
	"shabashov/tic-tac-toe/internal/player"
)

type Message struct {
	Player player.Player
	X      int
	Y      int
}

func ReadFromConn(conn net.Conn) {
	r := bufio.NewReader(conn)
	b := board.Board{}
	b.Init()
	for {
		if err := gob.NewDecoder(r).Decode(&b); err != nil {
			log.Println(err)
			conn.Close()
			break
		} else {
			b.Print()
		}
	}
}

func main() {
	c, e := net.Dial("tcp4", "127.0.0.1:8234")
	if e != nil {
		log.Println(e)
	}
	defer c.Close()

	go ReadFromConn(c)
	w := bufio.NewWriter(c)
	var Mark string
	fmt.Scanln(&Mark)
	p := player.Player{Name: "Some name", Mark: Mark}

	fmt.Println("Game is launched")

	for {
		var x, y int
		fmt.Scan(&x, &y)
		gob.NewEncoder(w).Encode(&Message{X: x, Y: y, Player: p})
		w.Flush()
	}
}
