package main

import (
	"fmt"
	"io"
	"net"
	"time"
)

var conn net.Conn

func Dial(netw, addr string) (net.Conn, error) {
	if conn != nil {
		e := conn.Close()
		if e != nil {
			fmt.Println("close twitter dial error", e)
		}
		conn = nil
	}
	netc, err := net.DialTimeout(netw, addr, 5*time.Second)
	if err != nil {
		return nil, err
	}
	conn = netc
	return netc, nil
}

var reader io.ReadCloser

func closeConn() {
	if conn != nil {
		_ = conn.Close()
	}
	if reader != nil {
		_ = reader.Close()
	}
}
