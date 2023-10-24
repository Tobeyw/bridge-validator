package neo3

import "math/big"

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
