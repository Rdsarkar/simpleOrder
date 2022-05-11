package repository

import (
	"github.com/tools/simple/model"
	"github.com/tools/simple/model/custommodel"
	"github.com/tools/simple/util"
	"net/http"
)

type ProductlineRepository struct{}

//Get
func (productlineRepository *ProductlineRepository) GetAllProductlines() custommodel.ResponseDto {
	var output custommodel.ResponseDto

	//database connection
	db := util.CreateConnectionUsingGormToCommonSchema()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()
	var c []model.Productline
	result := db.Find(&c)
	//if no action then close the connection

	if result.RowsAffected == 0 {
		output.Message = "Server Error So no data found"
		output.IsSuccess = false
		output.Payload = nil
		output.StatusCode = http.StatusInternalServerError
		return output
	}

	type tempOutput struct {
		Output      []model.Productline `json:"output"`
		OutputCount int                 `json:"outputCount"`
	}
	var tOutput tempOutput
	tOutput.Output = c
	tOutput.OutputCount = len(c)
	output.Message = "Success"
	output.IsSuccess = true
	output.Payload = tOutput
	output.StatusCode = http.StatusOK
	return output
}
