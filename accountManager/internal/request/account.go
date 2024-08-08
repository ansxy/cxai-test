package request

type ReqCreateAccount struct {
	UserID   string  `json:"-" validate:"required"`
	Type     string  `json:"type" validate:"required"`
	Balance  float64 `json:"balance" validate:"required"`
	Schedule *string `json:"schedule" `
	PoT      int     `json:"pot"` // PoT = Period of Time
}
