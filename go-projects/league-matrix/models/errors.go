package models

import (
	"encoding/csv"
	"errors"
	"net/http"
	"strconv"
)

// contains errors used in error api response
var (
	ErrNumError           = func(err error) bool { var numErr *strconv.NumError; return errors.As(err, &numErr) }
	ErrParseError         = func(err error) bool { var parseErr *csv.ParseError; return errors.As(err, &parseErr) }
	ErrNoSuchFile         = http.ErrMissingFile
	ErrMissingArgument    = errors.New("missing argument")
	ErrInvalidMessageType = errors.New("invalid message-type")
	ErrNotFound           = errors.New("not found")
)
