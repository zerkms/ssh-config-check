package main

import (
	"net"
	"strconv"
	"strings"
	"testing"
	"time"
)

func trueChecker(h host) bool {
	return true
}

func TestEmptyConfig(t *testing.T) {
	cfg := ``

	r, err := check(strings.NewReader(cfg), trueChecker)

	if err != nil {
		t.Error("Should not be error, got", err)
	}

	select {
	case _, ok := <-r:
		if ok {
			t.Error("Should be closed")
		}
	case <-time.After(5 * time.Second):
		t.Error("Should be closed, but is blocked")
	}
}

func TestAllAccessible(t *testing.T) {
	cfg := `Host foo
Host bar`

	hosts := []result{
		{accessible: true, host: host{id: "foo"}},
		{accessible: false, host: host{id: "bar"}},
	}

	mapChecker := func(h host) bool {
		for _, v := range hosts {
			if v.host.id == h.id {
				return v.accessible
			}
		}
		return false
	}

	r, err := check(strings.NewReader(cfg), mapChecker)

	if err != nil {
		t.Error("Should not be error, got", err)
	}

	for _, h := range hosts {
		select {
		case result := <-r:
			if result.host.id != h.host.id {
				t.Error("Unexpected host", result.host)
			}

			if result.accessible != h.accessible {
				t.Error("Should be accessible, got", result.accessible)
			}
		case <-time.After(5 * time.Second):
			t.Error("Should be closed, but is blocked")
		}
	}
}

func TestConnect(t *testing.T) {
	l, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Error("Could not create a tcp server", err)
	}

	defer l.Close()

	maybeHostPort := strings.Split(l.Addr().String(), ":")

	if len(maybeHostPort) != 2 {
		t.Error("Unexpected address", l.Addr().String())
	}

	port, err := strconv.Atoi(maybeHostPort[1])
	if err != nil {
		t.Error("Unexpected port value", maybeHostPort[1])
	}

	h := host{
		id:   maybeHostPort[0],
		port: port,
	}

	result := tcpTryConnect(h)

	if !result {
		t.Error("Should be able to connect")
	}
}

func TestConnectUnavailable(t *testing.T) {
	h := host{
		id:   "definitely.not.existing.host",
		port: 999,
	}

	result := tcpTryConnect(h)

	if result {
		t.Error("Should not connect")
	}
}
