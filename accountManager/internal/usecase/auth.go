package usecase

import (
	"github.com/ansxy/golang-boilerplate-gin/internal/model"
	"github.com/ansxy/golang-boilerplate-gin/pkg/constant"
	"github.com/gin-gonic/gin"
	"github.com/supabase-community/gotrue-go/types"
)

// Login implements IUsecase.
func (u *Usecase) Login(c *gin.Context, data types.SignupRequest) (*types.TokenResponse, error) {
	user, err := u.SupaBase.Auth.SignInWithEmailPassword(data.Email, data.Password)
	if err != nil {
		return nil, err
	}

	return user, err
}

// Register implements IUsecase.
func (u *Usecase) Register(c *gin.Context, data types.SignupRequest) error {
	user, err := u.SupaBase.Auth.Signup(data)
	if err != nil {
		return err
	}

	err = u.Repo.CreateUser(c, &model.User{
		Email:  user.Email,
		UserID: user.ID,
		Accounts: []*model.Account{
			{
				Balance: 0,
				Type:    string(constant.Default),
				UserID:  user.ID,
			},
		},
	})

	if err != nil {
		return err
	}

	return nil
}
