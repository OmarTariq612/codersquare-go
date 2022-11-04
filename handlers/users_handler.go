package handlers

import (
	"net/http"
	"time"

	"github.com/OmarTariq612/codersquare-go/datastore"
	"github.com/OmarTariq612/codersquare-go/types"
	"github.com/OmarTariq612/codersquare-go/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

const validPeriod = 15 * 24 * time.Hour // 15 days

type UsersHandler struct {
	db datastore.Database
}

func NewUsersHandler(db datastore.Database) UsersHandler {
	return UsersHandler{db: db}
}

func (u UsersHandler) Signup(c *gin.Context) {
	var userData SignupRequest
	if errs := utils.BindJsonVerifier(c, &userData); errs != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": errs})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userData.Password), bcrypt.DefaultCost)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	user := &types.User{ID: uuid.NewString(), Firstname: userData.Firstname, Lastname: userData.Lastname, Username: userData.Username, Email: userData.Email, Password: string(hashedPassword), CreatedAt: time.Now().Unix()}

	if err := u.db.CreateUser(user); err != nil {
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

func (u UsersHandler) Signin(c *gin.Context) {
	var userData SigninRequest
	if errs := utils.BindJsonVerifier(c, &userData); errs != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": errs})
		return
	}

	user := u.db.GetUserByUsername(userData.Login)
	if user == nil {
		user = u.db.GetUserByEmail(userData.Login)
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

	c.JSON(http.StatusOK, &SigninResponse{User: User{ID: user.ID, Email: user.Email, Username: user.Username, Firstname: user.Firstname, Lastname: user.Lastname}, JWT: token})
}

func (u UsersHandler) GetUser(c *gin.Context) {
	var userData GetUserRequest
	if errs := utils.BindUriVerifier(c, &userData); errs != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": errs})
		return
	}

	user := u.db.GetUserByID(userData.ID)
	if user == nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "could not find user with the specified id"})
		return
	}

	c.JSON(http.StatusOK, &GetUserResponse{ID: user.ID, Username: user.Username, Firstname: user.Firstname, Lastname: user.Lastname})
}

func (u UsersHandler) GetCurrentUser(c *gin.Context) {
	userID := c.GetString("user_id")
	user := u.db.GetUserByID(userID)
	if user == nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "could not find user with the provided id from the middleware"})
		return
	}

	c.JSON(http.StatusOK, &GetCurrentUserResponse{User: User{ID: user.ID, Email: user.Email, Username: user.Username, Firstname: user.Firstname, Lastname: user.Lastname}})
}
