package greg_test

import (
	"net"
	"testing"

	"github.com/f01c33/greg"
)

func TestScanner(T *testing.T) {
	ports := []string{}
	go func() {
		listener, err := net.Listen("tcp", "127.0.0.1:8666")
		if err != nil {
			T.Error(err)
		}
		conn, err := listener.Accept()
		if err != nil {
			T.Error(err)
		}
		_, err = conn.Write([]byte("the game"))
		if err != nil {
			T.Error(err)
		}
		conn.Close()
	}()
	for addr := range greg.PortScan([]string{"127.0.0.1"}, ports) {
		if addr != "127.0.0.1:8666" {
			T.Errorf("failed to scan open port")
		}
	}
}
