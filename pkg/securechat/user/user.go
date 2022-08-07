package user

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"

	shell "github.com/ipfs/go-ipfs-api"
)

type User struct {
	Id             string
	PublicKeyBytes []byte
	PeerId         string
}

func (u *User) PublicKey() (*rsa.PublicKey, error) {
	return x509.ParsePKCS1PublicKey(u.PublicKeyBytes)
}

func CreateUser(id string, sh *shell.Shell) (*User, *rsa.PrivateKey, error) {
	// Get current user node id
	idReq, err := sh.ID()
	if err != nil {
		return nil, nil, err
	}

	// Generate public private keypair
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, nil, err
	}

	// Return user and private key
	return &User{
		Id:             id,
		PublicKeyBytes: x509.MarshalPKCS1PrivateKey(privateKey),
		PeerId:         idReq.ID,
	}, privateKey, nil

}
