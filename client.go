package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

type client struct {
	conn     net.Conn
	name     string
	room     *room
	commands chan<- command
}

func (c *client) readInput() {
	for {
		msg, err := bufio.NewReader(c.conn).ReadString('\n')
		if err != nil {
			return
		}

		msg = strings.Trim(msg, "\r\n")

		args := strings.Split(msg, " ")
		cmd := strings.TrimSpace(args[0])

		switch cmd {
		case "/name":
			c.commands <- command{
				id:     cmdName,
				client: c,
				args:   args,
			}
			fmt.Printf("%s call name", c.conn.RemoteAddr().String())
		case "/join":
			c.commands <- command{
				id:     cmdJoin,
				client: c,
				args:   args,
			}
			fmt.Printf("%s call join", c.conn.RemoteAddr().String())
		case "/rooms":
			c.commands <- command{
				id:     cmdRooms,
				client: c,
				args:   args,
			}
			fmt.Printf("%s call rooms", c.conn.RemoteAddr().String())
		case "/msg":
			c.commands <- command{
				id:     cmdMsg,
				client: c,
				args:   args,
			}
			fmt.Printf("%s call msg", c.conn.RemoteAddr().String())
		case "/quit":
			c.commands <- command{
				id:     cmdQuit,
				client: c,
				args:   args,
			}
			fmt.Printf("%s call quit", c.conn.RemoteAddr().String())
		default:
			c.err(fmt.Errorf("unknown command: %s", cmd))
		}
	}
}

func (c *client) err(err error) {
	c.conn.Write([]byte("ERR: " + err.Error() + "\n"))
}

func (c *client) msg(msg string) {
	c.conn.Write([]byte("> " + msg + "\n"))
}
