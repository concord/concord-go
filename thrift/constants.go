// Autogenerated by Thrift Compiler (1.0.0-dev)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package bolt

import (
	"bytes"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = bytes.Equal

const KBoltEnvKeyBasePath = "BOLT_BASE_PATH"
const KBoltDefaultEnvBasePath = "/tmp/"
const KBoltEnvKeyPathPrefix = "BOLT"
const KDefaultThriftServiceIOThreads = 2
const KConcordEnvKeyClientListenAddr = "CONCORD_client_listen_address"
const KConcordEnvKeyClientProxyAddr = "CONCORD_client_proxy_address"
const KDatabasePath = "/tmp"
const KDatabaseEntryTTL = 43200
const KDefaultBatchSize = 2048
const KDefaultTraceSampleEveryN = 1024
const KPrincipalComputationName = "principal_computation"
const KIncomingMessageQueueTopic = "incoming"
const KPrincipalTimerQueueTopic = "principal_timers"
const KOutgoingMessageQueueTopic = "outgoing"
const KQueueStreamNameToIdMapTopic = "stream_map"
const KMessageQueueWatermarkTopic = "watermarks"
const KMessageQueueBatchSize = 1024
const KMessageQueueTTL = 21600

func init() {
}
