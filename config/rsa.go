package config

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"io/ioutil"
	"os"

	"github.com/a-fandy/finan/exception"
)

type RsaKey interface {
	GetPrivateKey() *rsa.PrivateKey
	GetPublicKey() *rsa.PublicKey
}

type Key struct {
	PrivateKey *rsa.PrivateKey
	PublicKey  *rsa.PublicKey
}

func NewRsaKey() RsaKey {
	var rsaKey Key
	privateKeyBytes, err := os.ReadFile("private.pem")
	exception.PanicIfError(err)
	privateKeyBlock, _ := pem.Decode(privateKeyBytes)
	if privateKeyBlock == nil || privateKeyBlock.Type != "RSA PRIVATE KEY" {
		panic(errors.New("Error decoding private key"))
	}
	rsaKey.PrivateKey, err = x509.ParsePKCS1PrivateKey(privateKeyBlock.Bytes)
	exception.PanicIfError(err)

	// Load RSA public key
	publicKeyBytes, err := ioutil.ReadFile("public.pem")
	exception.PanicIfError(err)
	publicKeyBlock, _ := pem.Decode(publicKeyBytes)
	if publicKeyBlock == nil || publicKeyBlock.Type != "RSA PUBLIC KEY" {
		panic(errors.New("Error decoding public key"))
	}
	pub, err := x509.ParsePKCS1PublicKey(publicKeyBlock.Bytes)
	exception.PanicIfError(err)
	rsaKey.PublicKey = pub
	return &rsaKey
}

func (key Key) GetPrivateKey() *rsa.PrivateKey {
	return key.PrivateKey
}

func (key Key) GetPublicKey() *rsa.PublicKey {
	return key.PublicKey
}
