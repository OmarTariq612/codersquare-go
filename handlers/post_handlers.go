package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/OmarTariq612/codersquare-go/datastore"
	"github.com/OmarTariq612/codersquare-go/types"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

func ListPostsHandler(c *gin.Context) {
	// posts, err := datastore.DB.ListPosts()
	// if err != nil {
	// 	c.AbortWithStatus(http.StatusInternalServerError)
	// 	fmt.Println("database list posts error:", err)
	// 	return
	// }
	posts := datastore.DB.ListPosts()
	if posts == nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		fmt.Println("database list posts error")
		return
	}

	c.JSON(http.StatusOK, ListPostsResponse{posts})
}

func CreatePostHandler(c *gin.Context) {
	// userID := c.MustGet("user_id").(string)
	userID := c.GetString("user_id")
	post := CreatePostRequest{}
	if err := c.ShouldBindJSON(&post); err != nil {
		var errors []*APIError
		if errs, ok := err.(validator.ValidationErrors); ok {
			t := reflect.TypeOf(post)
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

	// if err := datastore.DB.CreatePost(&types.Post{ID: uuid.NewString(), Title: post.Title, URL: post.URL, UserID: post.UserID, PostedAt: time.Now().Unix()}); err != nil {
	if err := datastore.DB.CreatePost(&types.Post{ID: uuid.NewString(), Title: post.Title, URL: post.URL, UserID: userID, PostedAt: time.Now().Unix()}); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusCreated)
}
