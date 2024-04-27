package repository

import (
	"encoding/json"
	"io"
	"jm/apps/entity/response"
	"os"

	"github.com/gin-gonic/gin"
)

type productRepositoryImpl struct{}

func NewProductRepositoryImpl() ProductRepository {
	return &productRepositoryImpl{}
}

func (repository *productRepositoryImpl) GetDataProduct(ctx *gin.Context) ([]response.DetailProductResponse, error) {
	var dataProduct []response.DetailProductResponse
	productFile, err := os.Open("./apps/files/products.json")
	if err != nil {
		return nil, err
	}

	byteValProduct, err := io.ReadAll(productFile)
	if err != nil {
		return nil, err
	}
	_ = json.Unmarshal(byteValProduct, &dataProduct)

	return dataProduct, nil
}

func (repository *productRepositoryImpl) GetReviewProduct(ctx *gin.Context) ([]response.ReviewProductResponse, error) {
	var dataReview []response.ReviewProductResponse
	reviewFile, err := os.Open("./apps/files/reviews.json")
	if err != nil {
		return nil, err
	}

	byteValReview, err := io.ReadAll(reviewFile)
	if err != nil {
		return nil, err
	}
	_ = json.Unmarshal(byteValReview, &dataReview)

	return dataReview, nil
}
