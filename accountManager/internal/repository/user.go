package repository

import (
	"log"

	"github.com/ansxy/golang-boilerplate-gin/internal/model"
	"github.com/gin-gonic/gin"
)

// CreateUser implements IFaceRepository.
func (r *Repository) CreateUser(c *gin.Context, data *model.User) error {
	return r.BaseRepository.Create(r.db, data)
}

// GetUser implements IFaceRepository.
func (r *Repository) GetUser(c *gin.Context, query ...interface{}) (*model.User, error) {
	var res *model.User

	if err := r.BaseRepository.FindOne(r.db.WithContext(c).Where(query[0], query[1:]...).Preload("Accounts"), &res); err != nil {
		log.Println("err", err)
		return nil, err
	}

	return res, nil
}
