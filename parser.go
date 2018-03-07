package main

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

func parse(r io.Reader) (config, error) {
	scanner := bufio.NewScanner(r)

	var h host

	cfg := config{}

	for scanner.Scan() {
		line := scanner.Text()
		h, cfg = parseParam(line, h, cfg)
	}

	if err := scanner.Err(); err != nil {
		return config{}, err
	}

	h, cfg = appendHostIfNotEmpty(h, cfg)

	return cfg, nil
}

func parseParam(line string, h host, cfg config) (host, config) {
	tokens := strings.SplitN(line, " ", 2)

	if len(tokens) == 2 {
		param := strings.ToLower(strings.TrimSpace(tokens[0]))
		value := strings.ToLower(strings.TrimSpace(tokens[1]))

		switch param {
		case "host":
			h, cfg = appendHostIfNotEmpty(h, cfg)
			h.id = value
			h.port = 22

			if h.hostname == "" {
				h.hostname = value
			}

		case "hostname":
			h.hostname = value

		case "port":
			if port, err := strconv.Atoi(value); err == nil {
				h.port = port
			}
		}
	}

	return h, cfg
}

func appendHostIfNotEmpty(h host, cfg config) (host, config) {
	if h.id != "" {
		cfg.hosts = append(cfg.hosts, h)
		h = host{}
	}

	return h, cfg
}
