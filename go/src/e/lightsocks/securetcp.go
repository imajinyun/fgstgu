package lightsocks

import (
	"io"
	"log"
	"net"
	"sync"
)

const bufSize = 1024

var bpool sync.Pool

func init() {
	bpool.New = func() interface{} {
		return make([]byte, bufSize)
	}
}
func bufferPoolGet() []byte {
	return bpool.Get().([]byte)
}

func bufferPoolPut(b []byte) {
	bpool.Put(b)
}

// SecureTCPConn struct.
type SecureTCPConn struct {
	io.ReadWriteCloser
	Cipher *Cipher
}

// DecodeRead method for SecureTCPConn.
func (s *SecureTCPConn) DecodeRead(bs []byte) (n int, err error) {
	n, err = s.Read(bs)
	if err != nil {
		return
	}
	s.Cipher.Decode(bs[:])
	return
}

// EncodeWrite method for SecureTCPConn.
func (s *SecureTCPConn) EncodeWrite(bs []byte) (int, error) {
	s.Cipher.Encode(bs)
	return s.Write(bs)
}

// EncodeCopy method for SecureTCPConn.
func (s *SecureTCPConn) EncodeCopy(rwc io.ReadWriteCloser) error {
	buf := bufferPoolGet()
	defer bufferPoolPut(buf)
	for {
		readCount, errRead := s.Read(buf)
		if errRead != nil {
			if errRead != io.EOF {
				return errRead
			}
			return nil
		}
		if readCount > 0 {
			writeCount, errWrite := (&SecureTCPConn{
				ReadWriteCloser: rwc,
				Cipher:          s.Cipher,
			}).EncodeWrite(buf[:readCount])
			if errWrite != nil {
				return errWrite
			}
			if readCount != writeCount {
				return io.ErrShortWrite
			}
		}
	}
}

// DecodeCopy method for SecureTCPConn.
func (s *SecureTCPConn) DecodeCopy(writer io.Writer) error {
	buf := bufferPoolGet()
	defer bufferPoolPut(buf)
	for {
		readCount, errRead := s.DecodeRead(buf)
		if errRead != nil {
			if errRead != io.EOF {
				return errRead
			}
			return nil
		}
		if readCount > 0 {
			writeCount, errWrite := writer.Write(buf[0:readCount])
			if errWrite != nil {
				return errWrite
			}
			if readCount != writeCount {
				return io.ErrShortWrite
			}
		}
	}
}

// DialEncryptedTCP func.
func DialEncryptedTCP(addr *net.TCPAddr, cipher *Cipher) (*SecureTCPConn, error) {
	conn, err := net.DialTCP("tcp", nil, addr)
	if err != nil {
		return nil, err
	}
	conn.SetLinger(0)
	return &SecureTCPConn{
		ReadWriteCloser: conn,
		Cipher:          cipher,
	}, nil
}

// ListenEncryptedTCP func.
func ListenEncryptedTCP(tcpaddr *net.TCPAddr, cipher *Cipher, conn func(s *SecureTCPConn), listen func(addr *net.TCPAddr)) error {
	listener, err := net.ListenTCP("tcp", tcpaddr)
	if err != nil {
		return err
	}
	defer listener.Close()

	if listen != nil {
		go listen(listener.Addr().(*net.TCPAddr))
	}
	for {
		tcpconn, err := listener.AcceptTCP()
		if err != nil {
			log.Println(err)
			continue
		}
		tcpconn.SetLinger(0)
		go conn(&SecureTCPConn{
			ReadWriteCloser: tcpconn,
			Cipher:          cipher,
		})
	}
}
