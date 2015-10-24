package concord

import (
	"git.apache.org/thrift.git/lib/go/thrift"
	bolt "github.com/concord/concord-go/thrift"
	"log"
	"os"
)

// Serve starts service for the given Computation.
func Serve(comp *Computation) error {
	bind := os.Getenv(bolt.KConcordEnvKeyClientListenAddr)

	// Init transport
	transport, err := thrift.NewTServerSocket(bind)
	if err != nil {
		log.Println("[ERROR] failed to bind:", err)
		return err
	}
	factory := thrift.NewTTransportFactory()
	transportF := thrift.NewTFramedTransportFactory(factory)

	protocolF := thrift.NewTBinaryProtocolFactoryDefault()
	service := NewComputationService(comp, transportF.GetTransport(nil), protocolF)
	processor := bolt.NewComputationServiceProcessor(service)

	srv := thrift.NewTSimpleServer4(processor, transport, transportF, protocolF)
	return srv.Serve()
}
