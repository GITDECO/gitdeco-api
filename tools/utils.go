package tools

import (
	"bytes"
	"encoding/base64"
	"gitdeco-api/internal/exception"
	"image"
	"image/jpeg"
	"image/png"
	"net/http"
	"strconv"
	"strings"
)

func Ternary(condition bool, trueValue, falseValue interface{}) interface{} {
	if condition {
		return trueValue
	}
	return falseValue
}

func UintParseX(str string) uint {
	ID, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		panic(&exception.Error{Key: "BAD_REQUEST", Data: ""})
	}
	return uint(ID)
}

func ErrorParse(err error) (key, data string) {
	for _, property := range strings.Split(err.Error(), "\n") {
		parts := strings.Split(property, ":")
		if len(parts) != 2 {
			continue
		}
		switch strings.TrimSpace(parts[0]) {
		case "KEY":
			key = strings.TrimSpace(parts[1])
		case "DATA":
			data = strings.TrimSpace(parts[1])
		}
	}
	return key, data
}

func FormatErrorData(err string, data string) string {
	return strings.Replace(err, "{data}", data, 1)
}

func ImageToBase64(url string) string {
    resp, err := http.Get(url)
    if err != nil {
        panic(&exception.Error{Key: "BASE64_ERROR", Data: err.Error()})
    }
    defer resp.Body.Close()

    img, format, err := image.Decode(resp.Body)
    if err != nil {
        panic(&exception.Error{Key: "BASE64_ERROR", Data: err.Error()})
    }

    var buf bytes.Buffer
    switch format {
    case "jpeg":
        err = jpeg.Encode(&buf, img, nil)
    case "png":
        err = png.Encode(&buf, img)
    default:
        panic(&exception.Error{Key: "BASE64_ERROR", Data: err.Error()})
    }
    if err != nil {
        panic(&exception.Error{Key: "BASE64_ERROR", Data: err.Error()})
    }

    base64Str := base64.StdEncoding.EncodeToString(buf.Bytes())
    return base64Str
}