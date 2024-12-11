package model

type InternalAccountRequest struct {
	AccountNumber string `json:"accountNumber" binding:"required"`
}

type InternalAccountResponse struct {
	CustomerName  string `json:"customerName" binding:"required"`
	AccountNumber string `json:"accountNumber" binding:"required"`
}
