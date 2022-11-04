package handlers

import (
	"net/http"
	"time"

	"github.com/OmarTariq612/codersquare-go/datastore"
	"github.com/OmarTariq612/codersquare-go/types"
	"github.com/OmarTariq612/codersquare-go/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type PostsHandler struct {
	db datastore.Database
}

func NewPostsHandler(db datastore.Database) PostsHandler {
	return PostsHandler{db: db}
}

func (ph PostsHandler) List(c *gin.Context) {
	userID := c.GetString("user_id")
	posts := ph.db.ListPosts(userID)
	if posts == nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "could not list posts from the database"})
		return
	}

	c.JSON(http.StatusOK, &ListPostsResponse{posts})
}

func (ph PostsHandler) Create(c *gin.Context) {
	// userID := c.MustGet("user_id").(string)
	userID := c.GetString("user_id")
	var post CreatePostRequest
	if errs := utils.BindJsonVerifier(c, &post); errs != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": errs})
		return
	}

	if post := ph.db.GetPostByURL(post.URL); post != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "duplicate post url"})
		return
	}

	if err := ph.db.CreatePost(&types.Post{ID: uuid.NewString(), Title: post.Title, URL: post.URL, UserID: userID, PostedAt: time.Now().Unix()}); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusCreated)
}

func (ph PostsHandler) Delete(c *gin.Context) {
	userID := c.GetString("user_id")
	var postData DeletePostRequest
	if errs := utils.BindUriVerifier(c, &postData); errs != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": errs})
		return
	}

	post := ph.db.GetPost(postData.PostID, "")
	if post == nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "the provided post id is not found"})
		return
	}

	if post.UserID != userID {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	if err := ph.db.DeletePost(post.ID); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
}

func (ph PostsHandler) Get(c *gin.Context) {
	userID := c.GetString("user_id")
	var postData GetPostRequest
	if errs := utils.BindUriVerifier(c, &postData); errs != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": errs})
		return
	}

	post := ph.db.GetPost(postData.PostID, userID)
	if post == nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, &GetPostResponse{Post: post})
}
