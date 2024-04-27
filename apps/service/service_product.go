package service

import (
	"jm/apps/entity"

	"github.com/gin-gonic/gin"
)

type ProductService interface {
	SummaryProduct(ctx *gin.Context, requestId string) entity.GenericResponse
	GetSpesificSummary(ctx *gin.Context, idProduct string, requestId string) entity.GenericResponse
}
