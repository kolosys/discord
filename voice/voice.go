package voice

import (
	"net"
	"sync/atomic"

	"github.com/kolosys/axon"
	"github.com/kolosys/ion/workerpool"
)

type VoicePayload struct{}
type VoiceConnection struct {
	gateway  *axon.Client[VoicePayload]
	udp      *net.UDPConn
	encoder  *OpusEncoder
	speaking atomic.Bool
	pool     *workerpool.Pool // Worker pool for audio processing
}
