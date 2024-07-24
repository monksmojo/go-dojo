package matrix

import (
	"encoding/csv"
	"fmt"
	"log"
	"mime/multipart"
	"strconv"
	"strings"
)

func NewMatrixService() matrixService {
	return &matrixServiceImpl{}
}

type matrixService interface {
	Echo() (string, error)
	Invert() (string, error)
	Flatten() (string, error)
	Sum() (int, error)
	Multiply() (float64, error)
	SetFile(file *multipart.FileHeader)
}

type matrixServiceImpl struct {
	file *multipart.FileHeader
}

func (ms *matrixServiceImpl) SetFile(file *multipart.FileHeader) {
	ms.file = file
}

// The `Echo` method format all the elements in a matrix 
// is registered to handler
func (ms *matrixServiceImpl) Echo() (string, error) {
	records, err := ms.parseCsvFile(ms.file)
	if err != nil {
		return "", err
	}
	return ms.formatMatrix(records), nil
}

// The `Invert` method invert all the elements in a matrix 
// is registered to handler
func (ms *matrixServiceImpl) Invert() (string, error) {
	records, err := ms.parseCsvFile(ms.file)
	if err != nil {
		return "", err
	}
	invertedRecord := ms.invertMatrix(records)
	return ms.formatMatrix(invertedRecord), nil
}

// The `FLatten` method flatten all the elements in a matrix 
// is registered to handler
func (ms *matrixServiceImpl) Flatten() (string, error) {
	records, err := ms.parseCsvFile(ms.file)
	if err != nil {
		return "", err
	}
	return ms.flattenMatrix(records), nil
}

// The `Sum` method calculates the sum of all the elements in a matrix 
// is registered to handler
func (ms *matrixServiceImpl) Sum() (int, error) {
	records, err := ms.parseCsvFile(ms.file)
	if err != nil {
		return 0, err
	}
	value, err := ms.matrixSum(records)
	if err != nil {
		return 0, err
	}
	return value, nil

}

// The `Multiply` method calculates the product of all the elements in a matrix 
// is registered to handler
func (ms *matrixServiceImpl) Multiply() (float64, error) {
	records, err := ms.parseCsvFile(ms.file)
	if err != nil {
		return 0.0, err
	}
	value, err := ms.matrixProduct(records)
	if err != nil {
		return 0.0, err
	}
	return value, nil
}

// This `parseCsvFile` function is responsible for 
// parsing a CSV file provided as a `multipart.FileHeader`
func (ms *matrixServiceImpl) parseCsvFile(file *multipart.FileHeader) ([][]string, error) {
	fileObject, err := file.Open()
	if err != nil {
		log.Printf("failed to open the file %v \n", err.Error())
		return make([][]string, 0), err
	}
	records, err := csv.NewReader(fileObject).ReadAll()
	defer fileObject.Close()
	if err != nil {
		log.Printf("failed to read content from the file %v \n", err.Error())
		return make([][]string, 0), err
	}
	return records, nil
}

// The `formatMatrix` function is responsible for formatting of a matrix as a string
func (ms *matrixServiceImpl) formatMatrix(records [][]string) (response string) {
	if len(records) == 0 || len(records[0]) == 0 {
		return
	}
	for _, row := range records {
		response = fmt.Sprintf("%s%s\n", response, strings.Join(row, ","))
	}
	return response
}

// The `invertMatrix` function is responsible for transpose of a matrix
func (ms *matrixServiceImpl) invertMatrix(records [][]string) (invertedRecords [][]string) {
	if len(records) == 0 || len(records[0]) == 0 {
		return
	}
	rows := len(records)
	cols := len(records[0])

	invertedRecords = make([][]string, cols)
	for i := range invertedRecords {
		invertedRecords[i] = make([]string, rows)
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			invertedRecords[j][i] = records[i][j]
		}
	}
	return
}

// The `flattenMatrix` function is responsible for flattening a matrix 
func (ms *matrixServiceImpl) flattenMatrix(records [][]string) (response string) {
	if len(records) == 0 || len(records[0]) == 0 {
		return
	} else {
		for _, row := range records {
			response = response + "," + strings.Join(row, ",")
		}
		return response[1:]
	}
}

// This `matrixSum` function in the `matrixServiceImpl` struct calculates the sum of all the elements
func (ms *matrixServiceImpl) matrixSum(records [][]string) (int, error) {
	if len(records) == 0 || len(records[0]) == 0 {
		return 0, nil
	} else {
		sum := 0
		for _, row := range records {
			for _, col := range row {
				val, err := strconv.ParseFloat(col, 64)
				if err != nil {
					return 0, err
				}
				sum += int(val)
			}
		}
		return sum, nil
	}
}

// This function `matrixProduct` calculates the product of all the elements in a matrix 
func (ms *matrixServiceImpl) matrixProduct(records [][]string) (float64, error) {
	if len(records) == 0 || len(records[0]) == 0 {
		return 0.0, nil
	} else {
		product := 1.0
		for _, row := range records {
			for _, col := range row {
				val, err := strconv.ParseFloat(col, 64)
				if err != nil {
					return 0.0, err
				}
				product *= val
			}
		}
		return float64(product), nil
	}
}
