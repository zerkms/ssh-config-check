package main

import (
	"strings"
	"testing"
)

func TestParseParamHost(t *testing.T) {
	h, c := parseParam("Host foo", host{}, config{})

	if h.id != "foo" {
		t.Error("Expected foo, got", h.id)
	}

	if h.port != 22 {
		t.Error("The default port should be 22, got", h.port)
	}

	if h.hostname != "foo" {
		t.Error("The default hostname should be foo, got", h.hostname)
	}

	if len(c.hosts) != 0 {
		t.Error("Expected 0 hosts in config, got", len(c.hosts))
	}
}

func TestParseParamHostAppend(t *testing.T) {
	h, c := parseParam("Host foo", host{id: "bar"}, config{})

	if h.id != "foo" {
		t.Error("Expected foo, got", h.id)
	}

	if len(c.hosts) != 1 {
		t.Error("Expected 1 host in config, got", len(c.hosts))
	}

	if c.hosts[0].id != "bar" {
		t.Error("Expected bar, got", c.hosts[0].id)
	}
}

func TestParseParamPort(t *testing.T) {
	h, _ := parseParam("Port 33", host{}, config{})

	if h.port != 33 {
		t.Error("Expected port 33, got", h.port)
	}
}

func TestParseParamPortInvalid(t *testing.T) {
	h, _ := parseParam("Port foo", host{port: 42}, config{})

	if h.port != 42 {
		t.Error("Expected port to be left untouched, got", h.port)
	}
}

func TestParse(t *testing.T) {
	strConfig := `Host foo
Host bar
	Hostname bar_hostname
	Port 42
`

	cfg, err := parse(strings.NewReader(strConfig))

	if err != nil {
		t.Error("Unexpected error", err)
	}

	if len(cfg.hosts) != 2 {
		t.Error("Expected 2 hosts, got", len(cfg.hosts))
	}

	foo := host{
		id:       "foo",
		hostname: "foo",
		port:     22,
	}

	if cfg.hosts[0] != foo {
		t.Error("Expected first host", foo, "got", cfg.hosts[0])
	}

	bar := host{
		id:       "bar",
		hostname: "bar_hostname",
		port:     42,
	}

	if cfg.hosts[1] != bar {
		t.Error("Expected first host", bar, "got", cfg.hosts[1])
	}
}
