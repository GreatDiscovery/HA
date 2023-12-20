package raft

import (
	"context"
	"fmt"
	"ha/pkg/log"
	"ha/pkg/runtime"
	"net"
	"strings"
)

type Config struct {
	RaftBind string // host:port
}

func SetUp(ctx context.Context, config Config) error {
	log.G(ctx).Info("setting up raft")
	hostname := runtime.HostName
	println(hostname)
	node, err := normalizedNode(config.RaftBind)
	fmt.Println(node)
	if err != nil {
		return err
	}
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
