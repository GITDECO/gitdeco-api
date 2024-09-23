package pkg

import (
	"fmt"
	"mime/multipart"
	"path/filepath"
	"strings"

	"github.com/go-playground/validator/v10"
)

var v *validator.Validate

var allowExtentions = []string{".svg"}

func init() {
	v = validator.New()

	v.RegisterValidation("image", func(fl validator.FieldLevel) bool {
		field := fl.Field()

		if fileHeader, ok := field.Interface().(multipart.FileHeader); ok {
			ext := strings.ToLower(filepath.Ext(fileHeader.Filename))
			for _, allowedExt := range allowExtentions {
				if ext == allowedExt {
					return true
				}
			}
		}
		return false
	})
}

func Validator(request any) error {
	if err := v.Struct(request); err != nil {
		var messages []string
		for _, err := range err.(validator.ValidationErrors) {
			message := fmt.Sprintf("%s: %s", err.Tag(), err.Error())
			messages = append(messages, message)
		}
		errMessage := strings.Join(messages, ", ")
		return fmt.Errorf("validator: %s", errMessage)
	}

	return nil
}
