package models

import (
	"encoding/csv"
	"errors"
	"net/http"
	"strconv"
)

// contains errors used in error api response

type ErrNumError func(err error) bool
type ErrParseError func(err error) bool
type ErrNoSuchFile error
type ErrMissingArgument error
type ErrInvalidMessageType error
type ErrNotFound error

type HttpErrors struct {
	ErrNumError           ErrNumError
	ErrParseError         ErrParseError
	ErrNoSuchFile         ErrNoSuchFile
	ErrMissingArgument    ErrMissingArgument
	ErrInvalidMessageType ErrInvalidMessageType
	ErrNotFound           ErrNotFound
}

func GetHttpError() *HttpErrors {
	return &HttpErrors{
		ErrNumError:           func(err error) bool { var numErr *strconv.NumError; return errors.As(err, &numErr) },
		ErrParseError:         func(err error) bool { var parseErr *csv.ParseError; return errors.As(err, &parseErr) },
		ErrNoSuchFile:         http.ErrMissingFile,
		ErrMissingArgument:    errors.New("missing argument"),
		ErrInvalidMessageType: errors.New("invalid message-type"),
		ErrNotFound:           errors.New("not found"),
	}
}
