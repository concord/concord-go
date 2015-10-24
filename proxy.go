package concord

import (
	"errors"
	"git.apache.org/thrift.git/lib/go/thrift"
	bolt "github.com/concord/concord-go/thrift"
	"log"
)

type Proxy struct {
	*bolt.BoltProxyServiceClient
}

func NewProxy(hostport string) (*Proxy, error) {
	socket, err := thrift.NewTSocket(hostport)
	if err != nil {
		log.Println("[ERROR] Failed to create proxy socket:", err)
		return nil, errors.New("proxy socket error")
	}
	transport := thrift.NewTFramedTransport(socket)
	protocol := thrift.NewTBinaryProtocolFactoryDefault()
	client := bolt.NewBoltProxyServiceClientFactory(transport, protocol)
	proxy := &Proxy{client}
	go func() {
		err := transport.Open()
		if err != nil {
			log.Println("[ERROR] failed to open transport")
		}
	}()
	return proxy, nil
}
