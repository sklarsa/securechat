package main

import (
	"crypto"
	"crypto/rsa"

	shell "github.com/ipfs/go-ipfs-api"
	"github.com/sklarsa/securechat/pkg/user"
)

func main() {
	//cmd.Execute()
	sh := shell.NewShell("http://localhost:5001")
	u, privKey, err := user.CreateUser("hello", sh)
	if err != nil {
		panic(err)
	}

	data, err := u.Encrypt([]byte("hello"))
	if err != nil {
		panic(err)
	}

	decrypted, err := privKey.Decrypt(nil, data, &rsa.OAEPOptions{Hash: crypto.SHA256})
	if err != nil {
		panic(err)
	}
	println(string(decrypted))

}
