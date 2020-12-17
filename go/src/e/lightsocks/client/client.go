package client

import (
	"e/lightsocks"
	"net"
	"log"
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

// Listen method for Client.
func (c *Client) Listen(listen func(listenAddr *net.TCPAddr)) error {
	return lightsocks.ListenEncryptedTCP(c.ListenAddr, c.Cipher, c.handleConn, listen)
}

func (c *Client) handleConn(conn *lightsocks.SecureTCPConn) {
	defer conn.Close()

	srv, err := lightsocks.DialEncryptedTCP(c.RemoteAddr, c.Cipher)
	if err != nil {
		log.Println(err)
		return
	}
	defer srv.Close()

	go func() {
		if err := srv.DecodeCopy(conn); err != nil {
			conn.Close()
			srv.Close()
		}
	}()
	conn.EncodeCopy(srv)
}
