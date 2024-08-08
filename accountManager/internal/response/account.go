package response

import "github.com/ansxy/golang-boilerplate-gin/internal/model"

type AccountResponse struct {
	UserID string `json:"user_id"`
	Email  string `json:"email"`

	Account *model.Account `json:"account"`
}
