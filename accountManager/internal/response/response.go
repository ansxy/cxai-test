package response

import (
	"encoding/json"
	"errors"
	"math"
	"net/http"

	"github.com/ansxy/golang-boilerplate-gin/pkg/constant"
	"gorm.io/gorm"
)

type JSONResponse struct {
	Success bool               `json:"success"`
	Data    interface{}        `json:"data"`
	Paging  *PaginatedResponse `json:"paging"`
	Error   *JSONErrorResponse `json:"error"`
}

type JSONErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type PaginatedResponse struct {
	Page       int   `json:"page"`
	PerPage    int   `json:"per_page"`
	Count      int64 `json:"count"`
	TotalPages int   `json:"total_pages"`
}

func Success(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(JSONResponse{Success: true, Data: data})
}

func Pagination(w http.ResponseWriter, list interface{}, page int, perPage int, count int64) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	var paging *PaginatedResponse
	total := math.Ceil(float64(count) / float64(perPage))
	if page > 0 {
		paging = &PaginatedResponse{
			Page:       page,
			PerPage:    perPage,
			Count:      count,
			TotalPages: int(total),
		}
	}

	json.NewEncoder(w).Encode(JSONResponse{Success: true, Data: list, Paging: paging})

}

func Error(w http.ResponseWriter, err error) {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(JSONResponse{Success: false, Error: &JSONErrorResponse{
			Code:    constant.DefaultNotFoundErrorCode,
			Message: constant.ErrorMessageMap[constant.DefaultNotFoundErrorCode],
		}})
		return
	}

	// fmt.Printf("%s \n", err.(*custom_error.CustomeError).ErrorContext.Function)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(JSONResponse{Success: false, Error: &JSONErrorResponse{
		Code:    constant.DefaultUnhandledErrorCode,
		Message: constant.ErrorMessageMap[constant.DefaultUnhandledErrorCode],
	}})

}
