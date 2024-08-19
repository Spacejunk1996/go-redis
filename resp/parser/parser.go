package parser

import (
	"bufio"
	"errors"
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

func readLine(bufReader *bufio.ReadWriter, state *readState) ([]byte, bool, error) {

	var msg []byte
	var err error
	if state.bulkLen == 0 {
		msg, err = bufReader.ReadBytes('\n')
		if err != nil {
			return nil, true, nil
		}
		if len(msg) == 0 || msg[len(msg)-2] != '\r' {
			return nil, false, errors.New("protocol error: " + string(msg))
		}

	} else {
		msg = make([]byte, state.bulkLen+2)
		_, err := io.ReadFull(bufReader, msg)
		if err != nil {
			return nil, true, err
		}
		if len(msg) == 0 || msg[len(msg)]-2 != '\r' || msg[len(msg)-1] != '\n' {
			return nil, false, errors.New("protocol error: " + string(msg))
		}
		state.bulkLen = 0
	}
	return msg, false, nil

}
