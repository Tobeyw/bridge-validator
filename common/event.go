package common

import (
	"encoding/json"
	"github.com/bane-labs/bridge-validator/config"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"

	"math/big"
)

type DepositEvent struct {
	Nonce       *int
	FromAddress *string
	ToAddress   *string
	Amount      *big.Int
	DepositHash *string
	Root        *string
}

type WithdrawalEvent struct {
	Nonce     *int
	ToAddress *string
	Amount    *big.Int
}

func ToByteArray(event interface{}) ([]byte, error) {
	data, err := json.Marshal(event)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func GetDataHash(event interface{}) (common.Hash, error) {

	data, err := ToByteArray(event)
	if err != nil {
		return common.Hash{}, err
	}
	hash := crypto.Keccak256Hash(data)

	return hash, nil
}

func SignEventData(cfg config.Config, event interface{}) ([]byte, error) {
	hash, err := GetDataHash(event)
	if err != nil {
		return nil, err
	}
	privateKey, err := cfg.GetKeyPair()
	if err != nil {
		return nil, err
	}
	signature, err := crypto.Sign(hash.Bytes(), privateKey)
	if err != nil {
		return nil, err
	}
	return signature, nil
}
