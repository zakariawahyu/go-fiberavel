package request

type CreateGiftRequest struct {
	Bank          string `json:"bank" form:"bank" validate:"required,max=255"`
	AccountName   string `json:"account_name" form:"account_name" validate:"required,max=255"`
	AccountNumber string `json:"account_number" form:"account_number" validate:"required,max=255"`
}

type UpdateGiftRequest struct {
	ID            int64  `json:"id" form:"id" validate:"required"`
	Bank          string `json:"bank" form:"bank" validate:"required,max=255"`
	AccountName   string `json:"account_name" form:"account_name" validate:"required,max=255"`
	AccountNumber string `json:"account_number" form:"account_number" validate:"required,max=255"`
}
