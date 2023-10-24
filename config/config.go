package config

import (
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/crypto"
	"os"
)

type Config struct {
	KeyPairFile             *string
	MessageBrokerAddress    *string
	NeoN3NodeURL            *string
	NeoN3NodeWSS            *string
	BaneNodeURL             *string
	BaneNodeWSS             *string
	BaneBridgeContractAddr  *string
	NeoN3BridgeContractAddr *string
}

func (cfg *Config) LoadKeyFile() (string, error) {
	keypairPath := cfg.KeyPairFile
	file, err := os.Open(*keypairPath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	byteSlice := make([]byte, 64)
	_, err = file.Read(byteSlice)
	if err != nil {
		return "", err
	}
	return string(byteSlice), nil
}

func (cfg Config) GetKeyPair() (*ecdsa.PrivateKey, error) {
	keypair, err := cfg.LoadKeyFile()
	if err != nil {
		return nil, err
	}
	privateKey, err := crypto.HexToECDSA(keypair)
	if err != nil {
		return nil, err
	}
	return privateKey, nil
}
