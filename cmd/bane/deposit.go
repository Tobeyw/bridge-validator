package bane

import (
	"fmt"
	"github.com/bane-labs/bridge-validator/cmd/common"
	"github.com/bane-labs/bridge-validator/cmd/config"

	"github.com/sirupsen/logrus"
)

func handleDeposit(data []byte, cfg config.Config, logger *logrus.Logger) {
	//Validate the values

	// sign  data
	signature, err := common.SignEventData(cfg, data)
	if err != nil {
		logger.Fatal(err)
	}
	fmt.Println(signature)

	// send data to relayer
	//TODO

}
