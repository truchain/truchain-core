package networking

import "github.com/coreos/etcd/raft"

func StartConsensusServer() {
	storage := raft.NewMemoryStorage()
	c := &raft.Config{
		ID:              0x01,
		ElectionTick:    10,
		HeartbeatTick:   1,
		Storage:         storage,
		MaxSizePerMsg:   4096,
		MaxInflightMsgs: 256,
	}
	// Set peer list to the other nodes in the cluster.
	// Note that they need to be started separately as well.
	raft.StartNode(c, []raft.Peer{{ID: 0x02}, {ID: 0x03}})
}
