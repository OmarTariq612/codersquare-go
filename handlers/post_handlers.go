package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/OmarTariq612/codersquare-go/datastore"
	"github.com/OmarTariq612/codersquare-go/types"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func ListPostsHandler(c *gin.Context) {
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
	if errs := BindJsonErrorHandler(c, &post); errs != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": errs})
		return
	}

	if err := datastore.DB.CreatePost(&types.Post{ID: uuid.NewString(), Title: post.Title, URL: post.URL, UserID: userID, PostedAt: time.Now().Unix()}); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusCreated)
}
