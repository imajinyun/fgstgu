package lightsocks

// Cipher struct.
type Cipher struct {
	encodePassword *Password
	decodePassword *Password
}

// Encode method for Cipher.
func (c *Cipher) Encode(bs []byte) {
	for i, v := range bs {
		bs[i] = c.encodePassword[v]
	}
}

// Decode method for Cipher.
func (c *Cipher) Decode(bs []byte) {
	for i, v := range bs {
		bs[i] = c.decodePassword[v]
	}
}

// NewCipher returns a Chiper.
func NewCipher(encodePassword *Password) *Cipher {
	decodePassword := &Password{}
	for i, v := range encodePassword {
		encodePassword[i] = v
		decodePassword[v] = byte(i)
	}
	return &Cipher{
		encodePassword: encodePassword,
		decodePassword: decodePassword,
	}
}
