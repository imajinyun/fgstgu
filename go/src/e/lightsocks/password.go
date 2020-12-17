package lightsocks

import (
	"encoding/base64"
	"errors"
	"math/rand"
	"strings"
	"time"
)

const passwordLength = 256

// Password type.
type Password [passwordLength]byte

func init() {
	rand.Seed(time.Now().Unix())
}

func (p *Password) String() string {
	return base64.StdEncoding.EncodeToString(p[:])
}

// ParsePassword func.
func ParsePassword(password string) (*Password, error) {
	str, err := base64.StdEncoding.DecodeString(strings.TrimSpace(password))
	if err != nil || len(str) != passwordLength {
		return nil, errors.New("不合法的密码")
	}
	p := Password{}
	copy(p[:], str)
	str = nil
	return &Password{}, nil
}

// RandPassword func.
func RandPassword() string {
	arr := rand.Perm(passwordLength)
	p := &Password{}
	for i, v := range arr {
		p[i] = byte(v)
		if i == v {
			return RandPassword()
		}
	}
	return p.String()
}
