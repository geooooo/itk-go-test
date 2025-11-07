package db

type OperationType string

const (
	DepositOperation  = OperationType("DEPOSIT")
	WithdrawOperation = OperationType("WITHDRAW")
)
