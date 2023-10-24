package bane

import (
	"fmt"
	common2 "github.com/bane-labs/bridge-validator/common"
	"github.com/bane-labs/bridge-validator/config"
	"github.com/bane-labs/bridge-validator/contracts/bane"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"

	"github.com/sirupsen/logrus"
)

func handleWithdrawal(abi abi.ABI, data []byte, cfg config.Config, logger *logrus.Logger) {
	//Validate the values
	withdrawal, err := abi.Unpack("Deposit", data)
	if err != nil {
		logger.Error(err)
	}

	if len(withdrawal) != 6 {
		logger.Error("Withdrawal event of bane error")
		return
	}

	var eventData = bane.BridgeEventWithdrawal{
		Nonce:          withdrawal[0].(*big.Int),
		From:           withdrawal[1].(common.Address),
		To:             withdrawal[2].(common.Address),
		Amount:         withdrawal[3].(*big.Int),
		WithdrawalHash: withdrawal[5].([]uint8),
		Root:           withdrawal[5].([]uint8),
		Raw:            types.Log{},
	}
	// sign  data
	signature, err := common2.SignEventData(cfg, eventData)
	if err != nil {
		logger.Fatal(err)
	}
	fmt.Println(signature)

	// send data to relayer
	//TODO

}
