package client

import (
	"e/lightsocks"
	"net"
)

// Client struct.
type Client struct {
	Cipher     *lightsocks.Cipher
	ListenAddr *net.TCPAddr
	RemoteAddr *net.TCPAddr
}

// NewClient returns a Client.
func NewClient(password string, laddr, raddr string) (*Client, error) {
	pass, err := lightsocks.ParsePassword(password)
	if err != nil {
		return nil, err
	}
	listenAddr, err := net.ResolveTCPAddr("tcp", laddr)
	if err != nil {
		return nil, err
	}
	remoteAddr, err := net.ResolveTCPAddr("tcp", raddr)
	if err != nil {
		return nil, err
	}
	return &Client{
		Cipher:     lightsocks.NewCipher(pass),
		ListenAddr: listenAddr,
		RemoteAddr: remoteAddr,
	}, nil
}
