package handler

import "github.com/monksmojo/go-dojo/go-project/go-gin-api-square-box/internal/domain/matrix"

type Handler struct {
	MatrixHandler matrix.MatrixHandler
}

// The function `FetchHandlers` returns a new instance of `Handler` with a `MatrixHandler` initialized.
func FetchHandlers() *Handler {
	return &Handler{
		MatrixHandler: matrix.NewMatrixHandler(),
	}
}
