package service

import (
	"github.com/gin-gonic/gin"
	"github.com/tools/simple/model"
	"github.com/tools/simple/repository"
)

type ProductlineService struct{}

//GetALL
func (productlineservice *ProductlineService) GetAll(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	productlineRepository := new(repository.ProductlineRepository)
	myOutput := productlineRepository.GetAllProductlines()

	c.JSON(myOutput.StatusCode, myOutput)
}

//GetbyID
func (productlineservice *ProductlineService) SingleGet(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	var input model.Productline
	c.ShouldBind(&input)
	productlineRepository := new(repository.ProductlineRepository)
	myOutput := productlineRepository.GetProductlineByID(input)

	c.JSON(myOutput.StatusCode, myOutput)
}

//create
func(productlineservice *ProductlineService) CreateProductlineservice(c *gin.Context){
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	var input model.Productline
	c.ShouldBind(&input)
	productlineRepository := new(repository.ProductlineRepository)
	myOutput := productlineRepository.CreateProductline(input)

	c.JSON(myOutput.StatusCode, myOutput)
}

func (productlineservice *ProductlineService) AddRouters(router *gin.Engine) {
	router.GET("/apl", productlineservice.GetAll)
	router.POST("/spl", productlineservice.SingleGet)
	router.POST("/ipl", productlineservice.CreateProductlineservice)
}



	