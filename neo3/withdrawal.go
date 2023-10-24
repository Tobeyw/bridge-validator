package neo3

import (
	"fmt"
	"github.com/bane-labs/bridge-validator/common"
	"github.com/bane-labs/bridge-validator/config"
	"github.com/joeqian10/neo3-gogogo/crypto"
	"github.com/joeqian10/neo3-gogogo/helper"
	"github.com/joeqian10/neo3-gogogo/rpc/models"
	"github.com/sirupsen/logrus"
	"math/big"
	"strconv"
)

func handlWithdrawal(notification models.RpcNotification, cfg config.Config, logger *logrus.Logger) {
	if notification.State.Type != "Array" {
		logger.Error("notification.State.Type error: Type is not Array")
	}
	state := models.ConvertInvokeStackArray(notification.State) // convert Value to []InvokeStack
	if len(state) != 3 {
		logger.Error("notification.State.Value error: Wrong length of states")
	}
	// get merkleProof.nonce
	nonce, err := strconv.Atoi(state[0].Value.(string))
	if err != nil {
		logger.Error("strconv.Atoi nonce error: ", err)
	}

	// get toAddress
	toBytes, err := crypto.Base64Decode(state[2].Value.(string))
	if err != nil {
		logger.Error("crypto.Base64Decode toAddress error:", err)
	}
	toScriptHash := "0x" + helper.UInt160FromBytes(toBytes).String()

	// get amount
	amount := new(big.Int)
	amount, ok := amount.SetString(state[3].Value.(string), 10)
	if !ok {
		logger.Error("big.Int.SetString amount error: ", err)
	}

	eventData := common.WithdrawalEvent{
		Nonce:     &nonce,
		ToAddress: &toScriptHash,
		Amount:    amount,
	}
	// sign  data
	signature, err := common.SignEventData(cfg, eventData)
	if err != nil {
		logger.Fatal(err)
	}
	fmt.Println(signature)

	// send data to relayer
	//TODO

}
