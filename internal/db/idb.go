package db

type IDb interface {
	UpdateWalletBalance(uuid string, amount uint, operation OperationType) error
	GetWalletBalance(uuid string) (uint, error)
}
