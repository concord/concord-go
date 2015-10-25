package concord

import (
	"errors"
	"git.apache.org/thrift.git/lib/go/thrift"
	bolt "github.com/concord/concord-go/thrift"
	"log"
	"net"
	"strconv"
)

// proxy represents client to Concord proxy.
type proxy struct {
	*bolt.BoltProxyServiceClient
}

// nePProxy inits and connects to new proxy.
func newProxy(hostport string, md *Metadata) (*proxy, error) {
	socket, err := thrift.NewTSocket(hostport)
	if err != nil {
		log.Println("[ERROR] Failed to create proxy socket:", err)
		return nil, errors.New("proxy socket error")
	}
	transport := thrift.NewTFramedTransport(socket)
	protocol := thrift.NewTBinaryProtocolFactoryDefault()
	client := bolt.NewBoltProxyServiceClientFactory(transport, protocol)

	pr := &proxy{client}
	err = transport.Open()
	if err != nil {
		log.Println("[ERROR] failed to open transport")
		return nil, err
	}

	if err := pr.register(hostport, md); err != nil {
		log.Println("[ERROR] wrong hostport", err)
		return nil, err
	}

	return pr, nil
}

// register registers proxy instance with the scheduler, update endpoint info.
func (p *proxy) register(hostport string, metadata *Metadata) error {
	host, port, err := net.SplitHostPort(hostport)
	if err != nil {
		return err
	}
	md := metadata.toBoltMetadata()

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
