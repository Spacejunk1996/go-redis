package tcp

import (
	"context"
	"github.com/Spacejunk1996/go-redis/lib/sync/atomic"
	"github.com/Spacejunk1996/go-redis/lib/sync/wait"
	"net"
	"sync"
	"time"
)

type EchoHandler struct {
	activeConn sync.Map
	closing    atomic.Boolean
}

type EchoClient struct {
	Conn    net.Conn
	Waiting wait.Wait
}

func (e *EchoClient) Close() error {
	e.Waiting.WaitWithTimeout(10 * time.Second)
	_ = e.Conn.Close()
	return nil
}

func (handler *EchoHandler) Handle(ctx context.Context, conn net.Conn) {
	if handler.closing.Get() {
		_ = conn.Close()
	}
	client := &EchoClient{
		Conn: conn,
	}
	handler.activeConn.Store(client, struct{}{})
}
