package concord

import (
	"git.apache.org/thrift.git/lib/go/thrift"
	bolt "github.com/concord/concord-go/thrift"
	"os"
)

// Serve starts service for the given Computation.
//
// Must be called from main() function of worker.
func Serve(comp Computation) error {
	bindAddr := os.Getenv(bolt.KConcordEnvKeyClientListenAddr)
	proxyAddr := os.Getenv(bolt.KConcordEnvKeyClientProxyAddr)

	// Init transport
	transport, err := thrift.NewTServerSocket(bindAddr)
	if err != nil {
		panic("failed to create server")
	}
	factory := thrift.NewTTransportFactory()
	transportF := thrift.NewTFramedTransportFactory(factory)

	protocolF := thrift.NewTBinaryProtocolFactoryDefault()

	proxy, err := newProxy(proxyAddr, comp.Metadata())
	if err != nil {
		panic("failed to initialize proxy")
	}

	service := newComputationService(comp, proxy)

	processor := bolt.NewComputationServiceProcessor(service)

	srv := thrift.NewTSimpleServer4(processor, transport, transportF, protocolF)
	return srv.Serve()
}
