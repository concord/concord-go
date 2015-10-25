package concord

import (
	"errors"
	"git.apache.org/thrift.git/lib/go/thrift"
	bolt "github.com/concord/concord-go/thrift"
	"log"
	"net"
	"strconv"
)

// Proxy represents client to Concord proxy.
type Proxy struct {
	*bolt.BoltProxyServiceClient
}

// NewProxy inits and connects to new Proxy.
func NewProxy(hostport string, md *Metadata) (*Proxy, error) {
	socket, err := thrift.NewTSocket(hostport)
	if err != nil {
		log.Println("[ERROR] Failed to create proxy socket:", err)
		return nil, errors.New("proxy socket error")
	}
	transport := thrift.NewTFramedTransport(socket)
	protocol := thrift.NewTBinaryProtocolFactoryDefault()
	client := bolt.NewBoltProxyServiceClientFactory(transport, protocol)

	proxy := &Proxy{client}
	err = transport.Open()
	if err != nil {
		log.Println("[ERROR] failed to open transport")
		return nil, err
	}

	if err := proxy.Register(hostport, md); err != nil {
		log.Println("[ERROR] wrong hostport", err)
		return nil, err
	}

	return proxy, nil
}

// Register registers proxy instance with the scheduler, update endpoint info.
func (p *Proxy) Register(hostport string, metadata *Metadata) error {
	host, port, err := net.SplitHostPort(hostport)
	if err != nil {
		return err
	}
	md := metadata.ToBoltMetadata()

	endpoint := bolt.NewEndpoint()
	endpoint.Ip = host
	portI, err := strconv.ParseInt(port, 10, 0)
	if err != nil {
		log.Println("[ERROR] wrong port for proxy:", err)
		return err
	}
	endpoint.Port = int16(portI)

	md.ProxyEndpoint = endpoint
	return p.RegisterWithScheduler(md)
}
