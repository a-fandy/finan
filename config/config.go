package config

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"io/ioutil"
	"os"

	"github.com/gofiber/fiber/v2/log"
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
	if err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}

	var configImpl ConfigImpl
	privateKeyBytes, err := os.ReadFile("private.pem")
	if err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}
	privateKeyBlock, _ := pem.Decode(privateKeyBytes)
	if privateKeyBlock == nil || privateKeyBlock.Type != "RSA PRIVATE KEY" {
		log.Error("Error decoding private key")
		os.Exit(1)
	}
	configImpl.PrivateKey, err = x509.ParsePKCS1PrivateKey(privateKeyBlock.Bytes)
	if err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}

	// Load RSA public key
	publicKeyBytes, err := ioutil.ReadFile("public.pem")
	if err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}
	publicKeyBlock, _ := pem.Decode(publicKeyBytes)
	if publicKeyBlock == nil || publicKeyBlock.Type != "RSA PUBLIC KEY" {
		log.Error("Error decoding public key")
		os.Exit(1)
	}
	pub, err := x509.ParsePKCS1PublicKey(publicKeyBlock.Bytes)
	if err != nil {
		log.Error(err.Error())
		os.Exit(1)
	}
	configImpl.PublicKey = pub

	return &configImpl
}
