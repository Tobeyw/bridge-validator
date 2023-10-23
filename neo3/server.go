package neo3

import (
	"github.com/bane-labs/bridge-validator/config"
	"github.com/joeqian10/neo3-gogogo/rpc"
	"github.com/sirupsen/logrus"
	"strconv"
	"strings"
	"time"
)

var lastBlockNumber = 0

func Server(cfg config.Config, logger *logrus.Logger) {
	neoRpc := rpc.NewClient(*cfg.NeoN3NodeURL)
	if neoRpc == nil {
		logger.Fatal("Client connection failed")
	}

	blockCountResponse := neoRpc.GetBlockCount()
	if blockCountResponse.HasError() {
		logger.Error("neoSdk.GetBlockCount error:", blockCountResponse.GetErrorInfo())
	}

	blockCount := blockCountResponse.Result
	startBlockNum := lastBlockNumber
	lastBlockNumber = blockCount //update blockNumber

	for i := startBlockNum; i < blockCount; i++ { //Unknown block
		blockResponse := neoRpc.GetBlock(strconv.Itoa(int(i)))
		if blockResponse.HasError() {
			logger.Error("neoSdk.GetBlock error:", blockResponse.GetErrorInfo())
		}

		time.Sleep(200 * time.Millisecond) // rpc too many request error
		txs := blockResponse.Result.Tx
		for _, tx := range txs {
			logResponse := neoRpc.GetApplicationLog(tx.Hash)
			if logResponse.HasError() {
				if strings.Contains(logResponse.GetErrorInfo(), "Unknown transaction") {
					continue
				}
				logger.Error("neoSdk.GetApplicationLog error:", logResponse.GetErrorInfo())
			}

			executions := logResponse.Result.Executions
			for _, execution := range executions {
				notifications := execution.Notifications
				for _, notification := range notifications {
					contract := notification.Contract
					eventName := notification.EventName

					if contract == *cfg.NeoN3BridgeContractAddr && eventName == "OnDeposit" {
						handleDeposit(notification, cfg, logger)
					} else if contract == *cfg.NeoN3BridgeContractAddr && eventName == "OnDeposit" {
						handleWithdrawal(notification, cfg, logger)
					}
				}

			} // notification
		}

	}

}
