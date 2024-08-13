package parser

import (
	"github.com/Spacejunk1996/go-redis/interface/resp"
	"io"
)

type Payload struct {
	Data resp.Reply
	Err  error
}

type readState struct {
	readingMultiLine bool
	expectedArgCount int
	msgType          byte
	args             [][]byte
	bulkLen          int64
}

func (s *readState) finish() bool {
	return s.expectedArgCount > 0 && len(s.args) == s.expectedArgCount
}

func ParseStream(reader io.Reader) <-chan *Payload {
	ch := make(chan *Payload)
	go parse0(reader, ch)
	return ch
}

func parse0(reader io.Reader, ch chan<- *Payload) {

}
