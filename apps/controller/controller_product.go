package controller

import (
	"encoding/json"
	"fmt"
	"jm/apps/service"
	"jm/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
	guuid "github.com/google/uuid"
)

type ProductController struct {
	ProductService service.ProductService
}

func NewProductController(productService *service.ProductService) ProductController {
	return ProductController{
		ProductService: *productService,
	}
}

func (controller *ProductController) SummaryProduct(ctx *gin.Context) {
	requestId := guuid.New()

	logStart := helpers.LogRequest(ctx, "get-summary-product", requestId.String())
	fmt.Println(logStart)

	response := controller.ProductService.SummaryProduct(ctx, requestId.String())

	responseData, err := json.Marshal(response)
	if err != nil {
		helpers.LogError(ctx, err.Error(), requestId.String())
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	logStop := helpers.LogResponse(ctx, string(responseData), requestId.String())
	fmt.Println(logStop)
	ctx.JSON(http.StatusOK, &response)
}

func (controller *ProductController) GetSpesificSummary(ctx *gin.Context) {

	requestId := guuid.New()

	logStart := helpers.LogRequest(ctx, fmt.Sprintf("get-summary-product-%s", ctx.Param("id")), requestId.String())
	fmt.Println(logStart)

	response := controller.ProductService.GetSpesificSummary(ctx, ctx.Param("id"), requestId.String())

	responseData, err := json.Marshal(response)
	if err != nil {
		helpers.LogError(ctx, err.Error(), requestId.String())
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	logStop := helpers.LogResponse(ctx, string(responseData), requestId.String())
	fmt.Println(logStop)
	ctx.JSON(http.StatusOK, &response)
}
