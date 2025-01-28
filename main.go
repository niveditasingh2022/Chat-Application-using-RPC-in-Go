package main

import (
	"../shared"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"time"
)

const (
	MAX_CLIENTS = 10

	CLIENT_NAME = "Anonymous"

	ERROR_PREFIX   = "Error: "
	ERROR_SEND     = ERROR_PREFIX + "First join the chat room and then start chat(Lobby chat is not allowed).\n"
	ERROR_CREATE   = ERROR_PREFIX + "Same chat room not allowed. Chat room already exists.\n"
	ERROR_JOIN     = ERROR_PREFIX + "Chat room not exist. Please create first and then join.\n"
	ERROR_LEAVE    = ERROR_PREFIX + "Sorry! leaving the lobby is not allowed.\n"
	ERROR_TOKEN    = ERROR_PREFIX + "Client already exists.\n"
	ERROR_NO_TOKEN = ERROR_PREFIX + "Sorry! no client exists with that token.\n"

	NOTICE_PREFIX          = "Notice: "
	NOTICE_ROOM_JOIN       = NOTICE_PREFIX + "\"%s\"successfully joined the chat room. WELCOME!\n"
	NOTICE_ROOM_LEAVE      = NOTICE_PREFIX + "\"%s\" left the chat room. See you soon.\n"
	NOTICE_ROOM_NAME       = NOTICE_PREFIX + "\"%s\" changed successfully their name to \"%s\".\n"
	NOTICE_ROOM_DELETE     = NOTICE_PREFIX + "Sorry! Chat room is inactive and will be deleted shortly.\n"
	NOTICE_PERSONAL_CREATE = NOTICE_PREFIX + "Successfully created chat room \"%s\".\n"
	NOTICE_PERSONAL_NAME   = NOTICE_PREFIX + "Changed name to \"\".\n"

	MSG_CONNECT = "Welcome to the server room! Please type \"/help\" to get a list of actions to proceed.\n"
	MSG_FULL    = "Sorry! server is full. Kindly try reconnecting later."

	EXPIRY_TIME time.Duration = 7 * 24 * time.Hour
)

func main() {
	rcvr := new(Receiver)
	rpc.Register(rcvr)
	rpc.HandleHTTP()
	l, e := net.Listen(shared.CONN_TYPE, shared.CONN_PORT)
	if e != nil {
		log.Fatal("listen error:", e)
	}
	http.Serve(l, nil)
}
