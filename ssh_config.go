package main

type config struct {
	hosts []host
}

type host struct {
	id       string
	hostname string
	port     int
}
