package model

type InternalTransferRequest struct {
	SourceAccountNumber string `json:"sourceAccountNumber" binding:"required"`
	TargetAccountNumber string `json:"targetAccountNumber" binding:"required"`
	Amount              int64  `json:"amount" binding:"required,min=0"`
	IsSourceFee         *bool  `json:"isSourceFee" binding:"required"`
}
