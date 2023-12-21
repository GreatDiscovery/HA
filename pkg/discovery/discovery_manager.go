package discovery

import (
	"context"
	"fmt"
	"ha/pkg/config"
	"ha/pkg/log"
	"ha/pkg/runtime"
	"ha/pkg/service"
	"net"
	"strings"
)

type Discovery struct {
	RaftEnabled bool
	RaftBind    string
	RaftDataDir string
	RaftNodes   []string
}

func NewDiscoveryManager(configuration config.Configuration) service.DiscoveryManager {
	return &Discovery{
		RaftEnabled: configuration.RaftEnabled,
		RaftBind:    configuration.RaftBind,
		RaftDataDir: configuration.RaftDataDir,
		RaftNodes:   configuration.RaftNodes,
	}
}

func (d *Discovery) Discovery(ctx context.Context) error {
	log.G(ctx).Info("starting auto discovery")
	if !d.RaftEnabled {
		log.G(ctx).Info("raft disabled")
		return nil
	}
	err := SetUp(ctx, *d)
	if err != nil {
		return err
	}
	return nil
}

func SetUp(ctx context.Context, config Discovery) error {
	log.G(ctx).Info("setting up raft")
	hostname := runtime.HostName
	println(hostname)
	node, err := normalizedNode(config.RaftBind)
	fmt.Println(node)
	if err != nil {
		return err
	}
	store := NewStore(config.RaftDataDir, config.RaftBind)
	peerNode := make([]string, 0)
	for _, raftNode := range config.RaftNodes {
		node, err := normalizedNode(raftNode)
		if err != nil {
			return err
		}
		peerNode = append(peerNode, node)
	}
	store.SetUpRaft(peerNode)
	return nil
}

func normalizedNode(node string) (string, error) {
	hostPort := strings.Split(node, ":")
	if len(hostPort) != 2 {
		return node, fmt.Errorf("node format is not ip:port")
	}
	host, err := normalizedIp(hostPort[0])
	if err != nil {
		return host, err
	}
	return fmt.Sprintf("%s:%s", host, hostPort[1]), nil
}

func normalizedIp(host string) (string, error) {
	if ip := net.ParseIP(host); ip != nil {
		return ip.String(), nil
	}

	ips, err := net.LookupIP(host)
	if err != nil {
		log.G(context.TODO()).Error("resolve failed", err)
		// ignore
		return host, nil
	}
	for _, ip := range ips {
		if ip.To4() == nil {
			continue
		}
		return ip.String(), nil
	}
	return host, fmt.Errorf("%+v resolved but no IP found", host)
}
