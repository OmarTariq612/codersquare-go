package handlers

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func BindJsonErrorHandler(c *gin.Context, obj any) []*APIError {
	var errors []*APIError
	if err := c.ShouldBindJSON(obj); err != nil {
		switch err := err.(type) {
		case validator.ValidationErrors:
			t := reflect.TypeOf(obj).Elem()
			for _, er := range err {
				sf, _ := t.FieldByName(er.Field())
				errors = append(errors, &APIError{Field: sf.Tag.Get("json"), Reason: er.Tag()})
			}
		case *json.UnmarshalTypeError:
			errors = append(errors, &APIError{Field: err.Field, Reason: fmt.Sprintf("%s required (passed %s)", err.Type, err.Value)})
		case *json.SyntaxError:
			errors = append(errors, &APIError{Field: "json syntax error", Reason: err.Error()})
		default:
			errors = append(errors, &APIError{Field: "unspecified", Reason: err.Error()})
		}
	}

	return errors
}
