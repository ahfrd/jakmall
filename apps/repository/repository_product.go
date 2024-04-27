package repository

import (
	"jm/apps/entity/response"

	"github.com/gin-gonic/gin"
)

type ProductRepository interface {
	GetDataProduct(ctx *gin.Context) ([]response.DetailProductResponse, error)
	GetReviewProduct(ctx *gin.Context) ([]response.ReviewProductResponse, error)
}
