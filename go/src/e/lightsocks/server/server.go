package server

import (
	"e/lightsocks"
	"encoding/binary"
	"net"
)

// Server struct.
type Server struct {
	Cipher     *lightsocks.Cipher
	ListenAddr *net.TCPAddr
}

// NewServer returns a Server.
func NewServer(password string, addr string) (*Server, error) {
	pass, err := lightsocks.ParsePassword(password)
	if err != nil {
		return nil, err
	}
	listenAddr, err := net.ResolveTCPAddr("tcp", addr)
	if err != nil {
		return nil, err
	}
	return &Server{
		Cipher:     lightsocks.NewCipher(pass),
		ListenAddr: listenAddr,
	}, nil
}

// Listen method for Server.
func (s *Server) Listen(listen func(addr *net.TCPAddr)) error {
	return lightsocks.ListenEncryptedTCP(s.ListenAddr, s.Cipher, s.handleConn, listen)
}

func (s *Server) handleConn(conn *lightsocks.SecureTCPConn) {
	defer conn.Close()
	buf := make([]byte, 256)
	_, err := conn.DecodeRead(buf)
	if err != nil || buf[0] != 0x05 {
		return
	}
	conn.EncodeWrite([]byte{0x05, 0x00})
	n, err := conn.DecodeRead(buf)
	if err != nil || n < 7 {
		return
	}
	if buf[1] != 0x01 {
		return
	}

	var dIP []byte
	switch buf[3] {
	case 0x01:
		dIP = buf[4 : 4+net.IPv4len]
	case 0x03:
		ipAddr, err := net.ResolveIPAddr("ip", string(buf[5:n-2]))
		if err != nil {
			return
		}
		dIP = ipAddr.IP
	case 0x04:
		dIP = buf[4 : 4+net.IPv6len]
	default:
		return
	}
	dPort := buf[n-2:]
	dstAddr := &net.TCPAddr{
		IP: dIP,
		Port: int(binary.BigEndian.Uint16(dPort)),
	}
	dstsrv, err := net.DialTCP("tcp", nil, dstAddr)
	if err != nil {
		return
	} else {
		defer dstsrv.Close()
		dstsrv.SetLinger(0)
		conn.EncodeWrite([]byte{0x05, 0x00, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00})
	}
	go func() {
		if err := conn.DecodeCopy(dstsrv); err != nil {
			conn.Close()
			dstsrv.Close()
		}
	}()
	(&lightsocks.SecureTCPConn{
		ReadWriteCloser: dstsrv,
		Cipher: conn.Cipher,
	}).EncodeCopy(conn)
}
