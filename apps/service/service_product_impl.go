package service

import (
	"errors"
	"fmt"
	"jm/apps/entity"
	"jm/apps/entity/response"
	"jm/apps/repository"
	"jm/helpers"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type productServiceImpl struct {
	ProductRepository repository.ProductRepository
}

func NewProductServiceImpl(productRepository *repository.ProductRepository) ProductService {
	return &productServiceImpl{
		ProductRepository: *productRepository,
	}
}

func (service *productServiceImpl) SummaryProduct(ctx *gin.Context, requestId string) entity.GenericResponse {
	var resData entity.GenericResponse
	var responseSummary response.ProductResponse

	getReviewProduct, err := service.ProductRepository.GetReviewProduct(ctx)
	if err != nil {
		resData.Code = http.StatusBadRequest
		resData.Message = err.Error()
		resData.Status = "error get db"
		helpers.LogError(ctx, err.Error(), requestId)

		return resData
	}

	responseSummary = service.findStar(getReviewProduct)

	responseSummary.TotalReview = len(getReviewProduct)
	avarageRating := service.calculateAverageRating(responseSummary)
	responseSummary.AverageRatings = avarageRating
	resData.Code = http.StatusOK
	resData.Message = "sukses get summary"
	resData.Status = "sukses"
	resData.Data = responseSummary
	return resData
}
func (service *productServiceImpl) GetSpesificSummary(ctx *gin.Context, idProduct string, requestId string) entity.GenericResponse {
	var resData entity.GenericResponse
	var responseSummary response.ProductResponse
	getDataProduct, err := service.ProductRepository.GetDataProduct(ctx)
	if err != nil {
		resData.Code = http.StatusBadRequest
		resData.Message = err.Error()
		resData.Status = "error get db"
		helpers.LogError(ctx, err.Error(), requestId)

		return resData
	}
	getReviewProduct, err := service.ProductRepository.GetReviewProduct(ctx)
	if err != nil {
		resData.Code = http.StatusBadRequest
		resData.Message = err.Error()
		resData.Status = "error get db"
		helpers.LogError(ctx, err.Error(), requestId)

		return resData
	}
	intIdProd, _ := strconv.Atoi(idProduct)

	findDataProduct, err := service.findDataByID(getDataProduct, intIdProd)
	if err != nil {
		resData.Code = http.StatusBadRequest
		resData.Message = err.Error()
		resData.Status = "error data not found"
		helpers.LogError(ctx, err.Error(), requestId)

		return resData
	}
	responseSummary = service.findStarSpesificProduct(getReviewProduct, findDataProduct.Id)

	avarageRating := service.calculateAverageRating(responseSummary)
	responseSummary.AverageRatings = avarageRating
	resData.Code = http.StatusOK
	resData.Message = "sukses get summary"
	resData.Status = "sukses"
	resData.Data = responseSummary
	return resData
}

func (service *productServiceImpl) findStar(reviewData []response.ReviewProductResponse) response.ProductResponse {
	var starCounts [5]int // Array to store counts for each star rating

	var summaryResponse response.ProductResponse
	for _, item := range reviewData {
		fmt.Println(item)
		if item.Rating >= 1 && item.Rating <= 5 {
			starCounts[item.Rating-1]++
		}
	}

	summaryResponse.OneStar = starCounts[0]
	summaryResponse.TwoStar = starCounts[1]
	summaryResponse.ThreeStar = starCounts[2]
	summaryResponse.FourStar = starCounts[3]
	summaryResponse.FiveStar = starCounts[4]
	return summaryResponse
}
func (service *productServiceImpl) findDataByID(data []response.DetailProductResponse, id int) (*response.DetailProductResponse, error) {
	var res response.DetailProductResponse
	for _, item := range data {
		if item.Id == id {
			return &item, nil
		}
	}
	return &res, errors.New("data not found")
}
func (service *productServiceImpl) findStarSpesificProduct(reviewData []response.ReviewProductResponse, productId int) response.ProductResponse {
	var starCounts [5]int // Array to store counts for each star rating
	var totalReview int
	var summaryResponse response.ProductResponse
	for _, item := range reviewData {
		fmt.Println(item)
		if item.ProductId == productId && item.Rating >= 1 && item.Rating <= 5 {
			totalReview++
			starCounts[item.Rating-1]++
		}
	}
	summaryResponse.TotalReview = totalReview
	summaryResponse.OneStar = starCounts[0]
	summaryResponse.TwoStar = starCounts[1]
	summaryResponse.ThreeStar = starCounts[2]
	summaryResponse.FourStar = starCounts[3]
	summaryResponse.FiveStar = starCounts[4]
	return summaryResponse
}

func (service *productServiceImpl) calculateAverageRating(productResponse response.ProductResponse) float64 {
	totalStars := (5 * productResponse.FiveStar) + (4 * productResponse.FourStar) + (3 * productResponse.ThreeStar) + (2 * productResponse.TwoStar) + (1 * productResponse.OneStar)
	totalReviews := productResponse.FiveStar + productResponse.FourStar + productResponse.ThreeStar + productResponse.TwoStar + productResponse.OneStar

	if totalReviews == 0 {
		return 0
	}
	fmt.Println(float64(totalStars))
	fmt.Println(float64(totalReviews))

	// Calculate the average rating with one decimal place
	averageRating := float64(totalStars) / float64(totalReviews)
	roundedRating := float64(int(averageRating*10)) / 10
	return roundedRating
}
