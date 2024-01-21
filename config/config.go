package config

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"io/ioutil"
	"os"

	"github.com/a-fandy/finan/exception"
	"github.com/joho/godotenv"
)

type Config interface {
	Get(key string) string
	GetPrivateKey() *rsa.PrivateKey
	GetPublicKey() *rsa.PublicKey
}

type ConfigImpl struct {
	PrivateKey *rsa.PrivateKey
	PublicKey  *rsa.PublicKey
}

func (config ConfigImpl) Get(key string) string {
	return os.Getenv(key)
}

func (config ConfigImpl) GetPrivateKey() *rsa.PrivateKey {
	return config.PrivateKey
}

func (config ConfigImpl) GetPublicKey() *rsa.PublicKey {
	return config.PublicKey
}

func New(filenames ...string) Config {
	err := godotenv.Load(filenames...)
	exception.PanicIfError(err)

	var configImpl ConfigImpl
	privateKeyBytes, err := os.ReadFile("private.pem")
	exception.PanicIfError(err)
	privateKeyBlock, _ := pem.Decode(privateKeyBytes)
	if privateKeyBlock == nil || privateKeyBlock.Type != "RSA PRIVATE KEY" {
		panic(errors.New("Error decoding private key"))
	}
	configImpl.PrivateKey, err = x509.ParsePKCS1PrivateKey(privateKeyBlock.Bytes)
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
	configImpl.PublicKey = pub

	return &configImpl
}
