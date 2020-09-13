package main

type commandID int

const (
	cmdName commandID = iota
	cmdJoin
	cmdRooms
	cmdMsg
	cmdQuit
)

type command struct {
	id     commandID
	client *client
	args   []string
}
