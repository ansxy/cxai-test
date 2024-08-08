package usecase

import (
	"github.com/ansxy/golang-boilerplate-gin/internal/repository"
	"github.com/ansxy/golang-boilerplate-gin/internal/service"
	"github.com/supabase-community/supabase-go"
)

type Usecase struct {
	Repo     repository.IFaceRepository
	SupaBase *supabase.Client
	Svc      service.IService
}

func NewUsecase(u *Usecase) IUsecase {
	return &Usecase{
		Repo:     u.Repo,
		SupaBase: u.SupaBase,
		Svc:      u.Svc,
	}
}
