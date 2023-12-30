package discovery

import (
	"fmt"
	mapset "github.com/deckarep/golang-set/v2"
	"github.com/hashicorp/raft"
	boltdb "github.com/hashicorp/raft-boltdb"
	"ha/pkg/config"
	_const "ha/pkg/const"
	"ha/pkg/log"
	"net"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type Store struct {
	raftDir                    string
	raftBind                   string
	logRetain                  int
	raft                       *raft.Raft
	applier                    CommandApplier
	snapshotDataCreatorApplier SnapshotCreatorApplier
}

type storeCommand struct {
	Op    string `json:"op,omitempty"`
	Value []byte `json:"value,omitempty"`
}

func NewStore(configuration config.Configuration, applier CommandApplier, snapshotCreatorApplier SnapshotCreatorApplier) *Store {
	return &Store{
		raftDir:                    configuration.RaftDataDir,
		raftBind:                   configuration.RaftBind,
		logRetain:                  configuration.LogRetain,
		applier:                    applier,
		snapshotDataCreatorApplier: snapshotCreatorApplier,
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
	fmt.Println(peers)

	err = makeRaftDir(s.raftDir)
	if err != nil {
		return err
	}

	ldb, err := boltdb.NewBoltStore(filepath.Join(s.raftDir, "logs.dat"))
	if err != nil {
		return err
	}
	sdb, err := boltdb.NewBoltStore(filepath.Join(s.raftDir, "stable.dat"))
	if err != nil {
		return err
	}
	snapshotStore, err := raft.NewFileSnapshotStore(s.raftDir, s.logRetain, os.Stderr)
	if err != nil {
		log.G(_const.TODO).Errorf("create snapshot store error=%+v", err)
		return err
	}

	raftInstance, err := raft.NewRaft(config, (*fsm)(s), ldb, sdb, snapshotStore, transport)
	if err != nil {
		return err
	}
	s.raft = raftInstance
	log.G(_const.TODO).Info("raft created")
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

func makeRaftDir(raftDir string) error {
	if _, err := os.Stat(raftDir); err != nil {
		if os.IsNotExist(err) {
			err := os.MkdirAll(raftDir, os.ModePerm)
			if err != nil {
				log.G(_const.TODO).Errorf("makedir raftdir (%s) error: %+v", raftDir, err)
				return err
			}
		} else {
			log.G(_const.TODO).Errorf("stat raftdir (%s) error: %+v", raftDir, err)
			return err
		}
	}
	return nil
}
