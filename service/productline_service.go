package service

import(
	"github.com/gin-gonic/gin"
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


func (productlineservice *ProductlineService) AddRouters(router *gin.Engine) {
	router.GET("/apl", productlineservice.GetAll)
}

	