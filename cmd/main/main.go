package main

import (
	"crypto/tls"
	"fmt"
	"net"

	"github.com/quic-go/quic-go"
)

func handleError(err error) {
    fmt.Errorf("Error detected: %w \naborting", err)
}

func main() {
    udpConn, err := net.ListenUDP("udp4", &net.UDPAddr{Port: 1234})
    if err != nil {
        handleError(err)
    }

    tr := quic.Transport{
        Conn: udpConn,
    }

    ln, err := tr.Listen(&tls.Config{}, &quic.Config{})
    if err != nil {
        handleError(err)
    }

    for {
        conn, err := ln.Accept()
        // ... error handling
        // handle the connection, usually in a new Go routine
    }
}
