package log

import (
	
)

type Plog struct {
	Level int
}

const (
	INFO int = iota
	DEBUG
	WARNING
	ERROR
	FATAL
)


func NewPlog() *Plog {
	return &Plog{
		Level: 1,
	}
}

func (p *Plog) Info() {

}

func (p *Plog) InfoF() {

}