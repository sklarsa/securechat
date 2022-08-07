package user

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"

	shell "github.com/ipfs/go-ipfs-api"
)

type User struct {
	Name           string
	PublicKeyBytes []byte
	PeerId         string
}

func (u *User) PublicKey() (*rsa.PublicKey, error) {
	return x509.ParsePKCS1PublicKey(u.PublicKeyBytes)
}

func CreateUser(name string, sh *shell.Shell) (*User, *rsa.PrivateKey, error) {
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
		Name:           name,
		PublicKeyBytes: x509.MarshalPKCS1PublicKey(&privateKey.PublicKey),
		PeerId:         idReq.ID,
	}, privateKey, nil

}

func (u *User) Encrypt(data []byte) ([]byte, error) {
	out := make([]byte, 0)
	pubKey, err := u.PublicKey()

	if err != nil {
		return out, err
	}

	encryptedBytes, err := rsa.EncryptOAEP(
		sha256.New(),
		rand.Reader,
		pubKey,
		data,
		nil,
	)
	return encryptedBytes, err
}
