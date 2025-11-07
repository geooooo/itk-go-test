package models

import "github.com/geooooo/itk-go-test/internal/db"

type WalletRequest struct {
	Id        string           `json:"valletId"`
	Operation db.OperationType `json:"operationType"`
	Amount    uint             `json:"amount"`
}

func (wr *WalletRequest) IsValid() bool {
	// TODO: нужно валидировать в соответствии с uuid
	if len(wr.Id) == 0 {
		return false
	}

	if wr.Operation != db.DepositOperation && wr.Operation != db.WithdrawOperation {
		return false
	}

	return true
}
