package pkg

import (
	"net/http"
)

type Exception struct {
	Code    int
	Message string
	Data    any
}

func NewException(code int, message string, data any) *Exception {
	return &Exception{
		Code:    code,
		Message: message,
		Data:    data,
	}
}

var exceptions = map[string]*Exception{
	"NOT_FOUND": NewException(
		http.StatusNotFound,
		http.StatusText(http.StatusNotFound),
		"Address {data} does not exist.",
	),
	"BAD_REQUEST": NewException(
		http.StatusBadRequest,
		http.StatusText(http.StatusBadRequest),
		"Request is not valid.",
	),
	"DONT_PARSE": NewException(
		http.StatusBadRequest,
		http.StatusText(http.StatusBadRequest),
		"Unable to parse the request body.",
	),
	"REPO_ERROR": NewException(
		http.StatusInternalServerError,
		http.StatusText(http.StatusInternalServerError),
		"{data}",
	),
	"USECASE_ERROR": NewException(
		http.StatusInternalServerError,
		http.StatusText(http.StatusInternalServerError),
		"{data}",
	),
	"ALREADY_EXIST": NewException(
		http.StatusInternalServerError,
		http.StatusText(http.StatusInternalServerError),
		"It already exists. ({data})",
	),
	"STRCONV_ERROR": NewException(
		http.StatusBadRequest,
		http.StatusText(http.StatusBadRequest),
		"{data}",
	),
	"SVG_PARSE_ERROR": NewException(
		http.StatusBadRequest,
		http.StatusText(http.StatusBadRequest),
		"Failed to parse SVG template",
	),
	"SVG_FILL_ERROR": NewException(
		http.StatusBadRequest,
		http.StatusText(http.StatusBadRequest),
		"Failed to fill SVG template. ({data})",
	),
	"OAUTH_ERROR": NewException(
		http.StatusInternalServerError,
		http.StatusText(http.StatusInternalServerError),
		"{data}",
	),
	"TOKEN_NOT_FOUND": NewException(
		http.StatusUnauthorized,
		http.StatusText(http.StatusUnauthorized),
		"The token could not be found.",
	),
	"TOKEN_PARSE_ERROR": NewException(
		http.StatusUnauthorized,
		http.StatusText(http.StatusUnauthorized),
		"Error parsing JWT token. {data}",
	),
	"TOKEN_VALID_ERROR": NewException(
		http.StatusUnauthorized,
		http.StatusText(http.StatusUnauthorized),
		"The JWT token has expired. {data}",
	),
	"TOKEN_GENERATE_ERROR": NewException(
		http.StatusInternalServerError,
		http.StatusText(http.StatusInternalServerError),
		"Error issuing JWT token. {data}",
	),
	"BASE64_ERROR": NewException(
		http.StatusInternalServerError,
		http.StatusText(http.StatusInternalServerError),
		"Image to base64 error. {data}",
	),
}

func GetException(key string) *Exception {
	return exceptions[key]
}