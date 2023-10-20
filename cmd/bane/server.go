package bane

import (
	"context"
	"fmt"
	"github.com/bane-labs/bridge-validator/cmd/config"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/sirupsen/logrus"
	"math/big"
)

func Server(cfg config.Config, logger *logrus.Logger) {
	ethcli, err := ethclient.Dial(*cfg.BaneNodeWSS)
	if err != nil {
		logger.Fatal("Client connection failed:", err)
	}
	contractAddress := common.HexToAddress(*cfg.BaneBridgeContractAddr)

	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(4499448),
		//ToBlock:   big.NewInt(4499460),
		Addresses: []common.Address{contractAddress},
	}

	logs := make(chan types.Log)
	sub, err := ethcli.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		logger.Fatal(err)
	}

	for {
		select {
		case err := <-sub.Err():
			logger.Fatal(err)
		case vLog := <-logs:
			fmt.Println(vLog) // pointer to common log

			topics := vLog.Topics
			data := vLog.Data
			event := topics[0]
			depositEvent := []byte("Deposit(uint,address,uint)")
			withdrawalEvent := []byte("Withdrawal(uint,address,uint)")
			depositHash := crypto.Keccak256Hash(depositEvent)
			withdrawalHash := crypto.Keccak256Hash(withdrawalEvent)

			if event == depositHash {
				handleDeposit(data, cfg, logger)
			} else if event == withdrawalHash {
				handleWithdrawal(data, cfg, logger)
			}
		}
	}
}
