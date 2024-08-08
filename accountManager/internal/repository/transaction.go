package repository

import (
	"github.com/ansxy/golang-boilerplate-gin/internal/model"
	"github.com/ansxy/golang-boilerplate-gin/internal/request"
	"github.com/gin-gonic/gin"
)

// GetListTransaction implements IFaceRepository.
func (r *Repository) GetListTransaction(c *gin.Context, req request.ReqListTransaction) ([]model.Transaction, int64, error) {
	var transactions []model.Transaction
	var total int64

	query := r.db.WithContext(c).Model(&model.Transaction{}).Where("user_id = ?", req.UserID)

	if req.Status != "" {
		query = query.Where("status = ?", req.Status)
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := query.Limit(req.PerPage).Offset((req.Page - 1) * req.PerPage).Find(&transactions).Error; err != nil {
		return transactions, total, err
	}

	return transactions, total, nil

}
