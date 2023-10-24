package bane

import (
	"fmt"
	common2 "github.com/bane-labs/bridge-validator/common"
	"github.com/bane-labs/bridge-validator/config"
	"github.com/bane-labs/bridge-validator/contracts/bridge"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"

	"github.com/sirupsen/logrus"
)

func handleDeposit(abi abi.ABI, data []byte, cfg config.Config, logger *logrus.Logger) {
	//Validate the values
	deposit, err := abi.Unpack("Deposit", data)
	if err != nil {
		logger.Error(err)
	}

	if len(deposit) != 3 {
		logger.Error("Deposit event of bane error")
		return
	}
	var eventData = bridge.BridgeEventDeposit{
		Nonce:  deposit[0].(*big.Int),
		To:     deposit[1].(common.Address),
		Amount: deposit[2].(*big.Int),
		Raw:    types.Log{},
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
