package cluseter


// clusterMeta represents cluster meta which must be kept in consensus.
type Meta struct {
	Peers map[string]string // Map from Raft address to API address
}

func (c *Meta) AddrForPeer(addr string) string {
	if api, ok := c.Peers[addr]; ok && api != "" {
		return api
	}

	// Go through each entry, and see if any key resolves to addr.
	for k, v := range c.APIPeers {
		resv, err := net.ResolveTCPAddr("tcp", k)
		if err != nil {
			continue
		}
		if resv.String() == addr {
			return v
		}
	}

	return ""
}