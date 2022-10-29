package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/OmarTariq612/codersquare-go/datastore"
	"github.com/OmarTariq612/codersquare-go/types"
	"github.com/OmarTariq612/codersquare-go/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// func ErrorHandlingMiddleware(typ reflect.Type) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		obj := reflect.New(typ)
// 		objPtr := reflect.PointerTo(obj.Type())

// 		if err := c.ShouldBindJSON(objPtr); err != nil {
// 			var errors []*APIError
// 			if errs, ok := err.(validator.ValidationErrors); ok {
// 				t := reflect.TypeOf(obj)
// 				for _, er := range errs {
// 					sf, _ := t.FieldByName(er.Field())
// 					errors = append(errors, &APIError{Field: sf.Tag.Get("json"), Reason: er.Tag()})
// 				}
// 				c.JSON(http.StatusBadRequest, gin.H{"errors": errors})
// 				return
// 			}
// 			if err, ok := err.(*json.UnmarshalTypeError); ok {
// 				errors = append(errors, &APIError{Field: err.Field, Reason: fmt.Sprintf("%s required (passed %s)", err.Type, err.Value)})
// 				c.JSON(http.StatusBadRequest, gin.H{"errors": errors})
// 				return
// 			}

// 			fmt.Printf("%T\n", err)
// 			c.AbortWithStatus(http.StatusBadRequest)
// 			return
// 		}

// 		c.Set("obj", obj)
// 		c.Next()
// 	}
// }

const validPeriod = 15 * 24 * time.Hour // 15 days

func SignupHandler(c *gin.Context) {
	var userData SignupRequest
	if err := c.ShouldBindJSON(&userData); err != nil {
		var errors []*APIError
		if errs, ok := err.(validator.ValidationErrors); ok {
			t := reflect.TypeOf(userData)
			for _, er := range errs {
				sf, _ := t.FieldByName(er.Field())
				errors = append(errors, &APIError{Field: sf.Tag.Get("json"), Reason: er.Tag()})
			}
			c.JSON(http.StatusBadRequest, gin.H{"errors": errors})
			return
		}
		if err, ok := err.(*json.UnmarshalTypeError); ok {
			errors = append(errors, &APIError{Field: err.Field, Reason: fmt.Sprintf("%s required (passed %s)", err.Type, err.Value)})
			c.JSON(http.StatusBadRequest, gin.H{"errors": errors})
			return
		}
		fmt.Printf("%T\n", err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userData.Password), bcrypt.DefaultCost)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	user := &types.User{ID: uuid.NewString(), Firstname: userData.Firstname, Lastname: userData.Lastname, Username: userData.Username, Email: userData.Email, Password: string(hashedPassword), CreatedAt: time.Now().Unix()}

	if err := datastore.DB.CreateUser(user); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	token, err := utils.SignJWT(&types.JWTObject{RegisteredClaims: jwt.RegisteredClaims{ExpiresAt: jwt.NewNumericDate(time.Now().Add(validPeriod))}, UserID: user.ID})
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusCreated, &SignupResponse{JWT: token})
}

func SigninHandler(c *gin.Context) {
	var userData SigninRequest
	if err := c.ShouldBindJSON(&userData); err != nil {
		var errors []*APIError
		if errs, ok := err.(validator.ValidationErrors); ok {
			t := reflect.TypeOf(userData)
			for _, er := range errs {
				sf, _ := t.FieldByName(er.Field())
				errors = append(errors, &APIError{Field: sf.Tag.Get("json"), Reason: er.Tag()})
			}
			c.JSON(http.StatusBadRequest, gin.H{"errors": errors})
			return
		}
		if err, ok := err.(*json.UnmarshalTypeError); ok {
			errors = append(errors, &APIError{Field: err.Field, Reason: fmt.Sprintf("%s required (passed %s)", err.Type, err.Value)})
			c.JSON(http.StatusBadRequest, gin.H{"errors": errors})
			return
		}
		fmt.Printf("%T\n", err)
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	user := datastore.DB.GetUserByUsername(userData.Login)
	if user == nil {
		user = datastore.DB.GetUserByEmail(userData.Login)
	}

	if user == nil || bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userData.Password)) != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	token, err := utils.SignJWT(&types.JWTObject{RegisteredClaims: jwt.RegisteredClaims{ExpiresAt: jwt.NewNumericDate(time.Now().Add(validPeriod))}, UserID: user.ID})
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	c.JSON(http.StatusOK, &SigninResponse{User: struct {
		ID        string `json:"id"`
		Email     string `json:"email"`
		Username  string `json:"username"`
		Firstname string `json:"firstname"`
		Lastname  string `json:"lastname"`
		CreatedAt int64  `json:"created_at"`
	}{
		ID: user.ID, Email: user.Email, Username: user.Username, Firstname: user.Firstname, Lastname: user.Lastname, CreatedAt: user.CreatedAt,
	}, JWT: token})
}
