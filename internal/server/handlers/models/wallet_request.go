package models

type operationType string

const (
	depositOperation  = operationType("DEPOSIT")
	withdrawOperation = operationType("WITHDRAW")
)

type WalletRequest struct {
	Id        string        `json:"valletId"`
	Operation operationType `json:"operationType"`
	Amount    uint          `json:"amount"`
}

func (wr *WalletRequest) IsValid() bool {
	// TODO: нужно валидировать в соответствии с uuid
	if len(wr.Id) == 0 {
		return false
	}

	if wr.Operation != depositOperation && wr.Operation != withdrawOperation {
		return false
	}

	return true
}
