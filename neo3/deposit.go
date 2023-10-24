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

func handlDeposit(notification models.RpcNotification, cfg config.Config, logger *logrus.Logger) {
	if notification.State.Type != "Array" {
		logger.Error("notification.State.Type error: Type is not Array")
	}
	state := models.ConvertInvokeStackArray(notification.State) // convert Value to []InvokeStack
	if len(state) != 6 {
		logger.Error("notification.State.Value error: Wrong length of states")
	}
	// get nonce
	nonce, err := strconv.Atoi(state[0].Value.(string))
	if err != nil {
		logger.Error("strconv.Atoi nonce error: ", err)
	}
	// get from
	fromBytes, err := crypto.Base64Decode(state[1].Value.(string))
	if err != nil {
		logger.Error("crypto.Base64Decode fromAddress error:", err)
	}
	fromScriptHash := "0x" + helper.UInt160FromBytes(fromBytes).String()

	// get to
	toBytes, err := crypto.Base64Decode(state[2].Value.(string))
	if err != nil {
		logger.Error("crypto.Base64Decode toAddress error:", err)
	}
	toScriptHash := "0x" + helper.UInt160FromBytes(toBytes).String()

	// get amount
	//	amount, err := strconv.Atoi(state[3].Value.(string))
	amount := new(big.Int)
	amount, ok := amount.SetString(state[3].Value.(string), 10)
	if !ok {
		logger.Error("big.Int.SetString amount error: ", err)
	}

	// get depositHash
	depositBytes, err := crypto.Base64Decode(state[4].Value.(string))
	if err != nil {
		logger.Error("crypto.Base64Decode fromAssetHash error: %v", err)

	}
	depositHash := string(depositBytes)

	// get root
	rootBytes, err := crypto.Base64Decode(state[5].Value.(string))
	if err != nil {
		logger.Error("crypto.Base64Decode fromAssetHash error: %v", err)

	}
	root := string(rootBytes)

	eventData := common.DepositEvent{
		Nonce:       &nonce,
		FromAddress: &fromScriptHash,
		ToAddress:   &toScriptHash,
		Amount:      amount,
		DepositHash: &depositHash,
		Root:        &root,
	}
	fmt.Println(eventData)

	// sign  data
	signature, err := common.SignEventData(cfg, eventData)
	if err != nil {
		logger.Fatal(err)
	}
	fmt.Println(signature)

	// send data to relayer
	//TODO

}
