package matrix

import (
	"log"
	"mime/multipart"

	"github.com/gin-gonic/gin"

	"github.com/monksmojo/go-dojo/go-project/go-gin-api-square-box/api"
	_ "github.com/monksmojo/go-dojo/go-project/go-gin-api-square-box/docs"
)

func NewMatrixHandler() MatrixHandler {
	return &matrixHandlerImpl{
		matrixService: NewMatrixService(),
	}
}

type MatrixHandler interface {
	Echo(ginContext *gin.Context)
	Invert(ginContext *gin.Context)
	Flatten(ginContext *gin.Context)
	Sum(ginContext *gin.Context)
	Multiply(ginContext *gin.Context)
}

type matrixHandlerImpl struct {
	matrixService matrixService
}

// @Summary Echo the contents of a CSV file
// @Description Receives a CSV file and echoes its contents in formatted form
// @Tags echo
// @Accept multipart/form-data
// @Produce json
// @Param file formData file true "CSV file to echo"
// @Success 200 {object} map[string]any
// @Failure 400 {object} map[string]any
// @Failure 500 {object} map[string]any
// @Router /echo [post]
func (mh *matrixHandlerImpl) Echo(ginContext *gin.Context) {
	file, err := getMultipartFile(ginContext)
	if err != nil {
		api.ErrorResponse(ginContext, err)
		return
	}
	mh.matrixService.SetFile(file)
	response, err := mh.matrixService.Echo()
	if err != nil {
		api.ErrorResponse(ginContext, err)
		return
	}
	api.SuccessResponse(ginContext, response)
}

// @Summary Invert the contents of a CSV file
// @Description Receives a CSV file and inverts its contents in formatted form
// @Tags invert
// @Accept multipart/form-data
// @Produce json
// @Param file formData file true "CSV file to invert"
// @Success 200 {object} map[string]any
// @Failure 400 {object} map[string]any
// @Failure 500 {object} map[string]any
// @Router /invert [post]
func (mh *matrixHandlerImpl) Invert(ginContext *gin.Context) {
	file, err := getMultipartFile(ginContext)
	if err != nil {
		api.ErrorResponse(ginContext, err)
		return
	}
	mh.matrixService.SetFile(file)
	response, err := mh.matrixService.Invert()
	if err != nil {
		api.ErrorResponse(ginContext, err)
		return
	}
	api.SuccessResponse(ginContext, response)

}

// @Summary Flatten the contents of a CSV file
// @Description Receives a CSV file and flatten its contents in formatted form
// @Tags flatten
// @Accept multipart/form-data
// @Produce json
// @Param file formData file true "CSV file to flatten"
// @Success 200 {object} map[string]any
// @Failure 400 {object} map[string]any
// @Failure 500 {object} map[string]any
// @Router /flatten [post]
func (mh *matrixHandlerImpl) Flatten(ginContext *gin.Context) {
	file, err := getMultipartFile(ginContext)
	if err != nil {
		api.ErrorResponse(ginContext, err)
		return
	}
	mh.matrixService.SetFile(file)
	response, err := mh.matrixService.Flatten()
	if err != nil {
		api.ErrorResponse(ginContext, err)
		return
	}
	api.SuccessResponse(ginContext, response)

}

// @Summary Sum the contents of a CSV file
// @Description Receives a CSV file and sum its contents in formatted form
// @Tags sum
// @Accept multipart/form-data
// @Produce json
// @Param file formData file true "CSV file to sum"
// @Success 200 {object} map[string]any
// @Failure 400 {object} map[string]any
// @Failure 500 {object} map[string]any
// @Router /Sum [post]
func (mh *matrixHandlerImpl) Sum(ginContext *gin.Context) {
	file, err := getMultipartFile(ginContext)
	if err != nil {
		api.ErrorResponse(ginContext, err)
		return
	}
	mh.matrixService.SetFile(file)
	response, err := mh.matrixService.Sum()
	if err != nil {
		api.ErrorResponse(ginContext, err)
		return
	}
	api.SuccessResponse(ginContext, response)

}

// @Summary multiply the contents of a CSV file
// @Description Receives a CSV file and multiply its contents in formatted form
// @Tags multiply
// @Accept multipart/form-data
// @Produce json
// @Param file formData file true "CSV file to multiply"
// @Success 200 {object} map[string]any
// @Failure 400 {object} map[string]any
// @Failure 500 {object} map[string]any
// @Router /multiply [post]
func (mh *matrixHandlerImpl) Multiply(ginContext *gin.Context) {
	file, err := getMultipartFile(ginContext)
	if err != nil {
		api.ErrorResponse(ginContext, err)
		return
	}
	mh.matrixService.SetFile(file)
	response, err := mh.matrixService.Multiply()
	if err != nil {
		api.ErrorResponse(ginContext, err)
		return
	}
	api.SuccessResponse(ginContext, response)

}

func getMultipartFile(ginContext *gin.Context) (*multipart.FileHeader, error) {
	file, err := ginContext.FormFile("file")
	if err != nil {
		log.Printf("failed to load the file %v \n", err.Error())
		return nil, err
	}
	return file, nil
}
