package repository

import (
	"github.com/ansxy/golang-boilerplate-gin/internal/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
	BaseRepository
}

// GetListAccount implements IFaceRepository.
func (r *Repository) GetListAccount(c *gin.Context, userID string) ([]model.Account, error) {
	panic("unimplemented")
}

func NewRepository(db *gorm.DB) IFaceRepository {
	return &Repository{
		db: db,
	}
}
