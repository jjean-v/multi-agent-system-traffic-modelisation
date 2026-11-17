package com

type messageType int

const (
	Mess1 messageType = iota
	Mess2
)

type Message struct {
	Type messageType
	Txt  string
}
