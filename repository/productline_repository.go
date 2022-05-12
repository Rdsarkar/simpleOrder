package repository

import (
	"net/http"

	"github.com/tools/simple/model"
	"github.com/tools/simple/model/custommodel"
	"github.com/tools/simple/util"
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

//GetByID
func (productlineRepository *ProductlineRepository) GetProductlineByID(c1 model.Productline) custommodel.ResponseDto {
	var output custommodel.ResponseDto
	if c1.Id2 <= 0 {
		output.Message = "Invalid ID"
		output.IsSuccess = false
		output.Payload = nil
		output.StatusCode = http.StatusBadRequest
		return output
	}

	//database connection
	db := util.CreateConnectionUsingGormToCommonSchema()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()
	var o model.Productline
	result := db.Where(&model.Productline{Id2: c1.Id2}).First(&o)

	if result.RowsAffected == 0 {
		output.Message = "No Data found for this ID"
		output.IsSuccess = false
		output.Payload = nil
		output.StatusCode = http.StatusNotFound
		return output
	}
	type tempOutput struct {
		Output model.Productline `json:"output"`
	}
	var toutput tempOutput
	toutput.Output = o
	output.Message = "GetProductLineBYID"
	output.IsSuccess = true
	output.Payload = toutput
	output.StatusCode = http.StatusOK
	return output
}
