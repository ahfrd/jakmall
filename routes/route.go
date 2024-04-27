package routes

import (
	"jm/apps/controller"

	"github.com/gin-gonic/gin"
)

func SetUpProductRoute(router *gin.Engine, productController *controller.ProductController) {
	router.GET("review/summary", productController.SummaryProduct)
	router.GET("review/product/:id", productController.GetSpesificSummary)

}
