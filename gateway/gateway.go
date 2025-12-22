package gateway

import (
	"github.com/kolosys/atomic/collection"
	"github.com/kolosys/axon"
	"github.com/kolosys/nova/bus"
)

type Gateway struct {
	client    *axon.Client[GatewayPayload]
	session   *Session
	heartbeat *HeartbeatManager
	events    *bus.EventBus
	shards    *collection.Collection[int, *Shard]
}
