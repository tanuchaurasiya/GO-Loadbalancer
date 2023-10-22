package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

type Server interface {
	Address() string
	IsAlive() bool
	Serve(rw http.ResponseWriter, r *http.Request)
}

type simpleServer struct {
	addr  string
	proxy httputil.ReverseProxy
}

type LoadBalancer struct {
	port            string
	roundRobinCount int
	servers         []Server
}

func handleError(err error) {
	if err != nil {
		fmt.Printf("error:%v", err)
		os.Exit(1)
	}
}

func newSimpleServer(addr string) *simpleServer {
	serverUrl, err := url.Parse(addr)
	fmt.Println(serverUrl)
	handleError(err)
	return &simpleServer{
		addr:  addr,
		proxy: *httputil.NewSingleHostReverseProxy(serverUrl),
	}
}

func newLoadBalancer(port string, servers []Server) *LoadBalancer {
	return &LoadBalancer{
		port:            port,
		roundRobinCount: 0,
		servers:         servers,
	}
}

func main() {
	// s := newSimpleServer("192.@")
	// fmt.Println(s.addr)

}
