package discovery

import (
	"fmt"
	mapset "github.com/deckarep/golang-set/v2"
	"github.com/hashicorp/raft"
	"net"
	"os"
	"strings"
	"time"
)

type Store struct {
	raftDir  string
	raftBind string
	raft     *raft.Raft
}

func NewStore(raftDir string, raftBind string) *Store {
	return &Store{
		raftDir:  raftDir,
		raftBind: raftBind,
	}
}

// SetUpRaft raft core ini method
func (s *Store) SetUpRaft(peerNodes []string) error {
	config := raft.DefaultConfig()

	addr, err := net.ResolveTCPAddr("tcp", s.raftBind)
	if err != nil {
		return err
	}
	transport, err := raft.NewTCPTransport(s.raftBind, addr, 3, 10*time.Second, os.Stderr)
	if err != nil {
		return err
	}
	peers := uniquePeer(peerNodes)

	// todo delete
	fmt.Println(config)
	fmt.Println(transport)
	fmt.Println(peers)
	return nil
}

func uniquePeer(peerNodes []string) []string {
	uniquePeers := make([]string, 0)
	set := mapset.NewSet[string]()
	for _, node := range peerNodes {
		noSpace := strings.TrimSpace(node)
		if set.Contains(noSpace) {
			continue
		}
		set.Add(noSpace)
		uniquePeers = append(uniquePeers, noSpace)
	}
	return uniquePeers
}
