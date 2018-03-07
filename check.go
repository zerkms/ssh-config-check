package main

import (
	"fmt"
	"io"
	"net"
	"sync"
	"time"
)

type result struct {
	accessible bool
	host       host
}

type checker func(host) bool

func check(sshConfigReader io.Reader, checkerF checker) (<-chan result, error) {
	config, err := parse(sshConfigReader)
	if err != nil {
		return nil, err
	}

	results := make(chan result, len(config.hosts))

	wg := sync.WaitGroup{}
	wg.Add(len(config.hosts))

	for _, h := range config.hosts {
		go func(h host) {
			results <- result{
				accessible: checkerF(h),
				host:       h,
			}
			wg.Done()
		}(h)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	return results, nil
}

func tcpTryConnect(h host) bool {
	address := fmt.Sprintf("%s:%d", h.hostname, h.port)
	_, err := net.DialTimeout("tcp", address, 1*time.Second)

	return err == nil
}
