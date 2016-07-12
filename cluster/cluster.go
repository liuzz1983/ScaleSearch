package cluster

import (
	"net"
	"github.com/hashicorp/raft"
	"github.com/hashicorp/raft-boltdb"
)


var (
	// ErrNotLeader is returned when a node attempts to execute a leader-only
	// operation.
	ErrNotLeader = errors.New("not leader")
)

