package main

import (
	"jm/apps/controller"
	"jm/apps/repository"
	"jm/apps/service"
	"jm/helpers"
	"jm/routes"

	"github.com/gin-gonic/gin"
)

func init() {
	env := helpers.Env{}
	env.StartingCheck()

}
func main() {
	router := gin.Default()
	productRepository := repository.NewProductRepositoryImpl()
	//Service
	productService := service.NewProductServiceImpl(&productRepository)
	//Controller
	productController := controller.NewProductController(&productService)
	routes.SetUpProductRoute(router, &productController)
	router.Run(":8080")
}
