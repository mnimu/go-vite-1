package p2p

import "net"

// Protocol is a abstract communication layer above connection, there can be multi protocol on a connection.
// A Protocol usually has many different message codes, each code has a handler to handle the message from peer.
type Protocol interface {
	// Name is the protocol name, should better be unique, for readability
	Name() string

	// ID MUST be unique.
	ID() ProtocolID

	// ProtoData return the data to handshake, will transmit to peer with HandshakeMsg.
	ProtoData() []byte

	// ReceiveHandshake handle the HandshakeMsg and protoData from peer.
	// The connection will be disconnected if err is not nil.
	// Level MUST small than 255, each level has it`s strategy, Eg, access by count-constraint, broadcast priority
	// If peer support multiple protocol, the highest protocol level will be the peer`s level.
	// As default, there are 4 levels from low to high: Inbound - Outbound - Trusted - Superior
	ReceiveHandshake(msg HandshakeMsg, protoData []byte, sender net.Addr) (state interface{}, level Level, err error)

	// Handle message from sender, if the return error is not nil, will disconnect with peer
	Handle(msg Msg) error

	// State get the Protocol state, will be sent to peers by heartbeat
	State() []byte

	// SetState handle the state data from sender
	SetState(state []byte, peer Peer)

	// OnPeerAdded will be invoked after Peer run
	// peer will be closed if return error is not nil
	OnPeerAdded(peer Peer) error

	// OnPeerRemoved will be invoked after Peer closed
	OnPeerRemoved(peer Peer) error
}

type MsgHandler = func(msg Msg) error

type ProtocolMap = map[ProtocolID]Protocol

type Level = byte

const (
	Inbound  Level = 0
	Outbound Level = 1
	Trusted  Level = 2
	Superior Level = 3
)