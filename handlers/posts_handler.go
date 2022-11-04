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

func (p PostsHandler) List(c *gin.Context) {
	posts := p.db.ListPosts()
	if posts == nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "could not list posts from the database"})
		return
	}

	c.JSON(http.StatusOK, &ListPostsResponse{posts})
}

func (p PostsHandler) Create(c *gin.Context) {
	// userID := c.MustGet("user_id").(string)
	userID := c.GetString("user_id")
	var post CreatePostRequest
	if errs := utils.BindJsonVerifier(c, &post); errs != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": errs})
		return
	}

	if post := p.db.GetPostByURL(post.URL); post != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "duplicate post url"})
		return
	}

	if err := p.db.CreatePost(&types.Post{ID: uuid.NewString(), Title: post.Title, URL: post.URL, UserID: userID, PostedAt: time.Now().Unix()}); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusCreated)
}

func (p PostsHandler) Delete(c *gin.Context) {

}

func (p PostsHandler) Get(c *gin.Context) {
	var postData GetPostRequest
	if errs := utils.BindUriVerifier(c, &postData); errs != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": errs})
		return
	}

	post := p.db.GetPostByID(postData.PostID)
	if post == nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, &GetPostResponse{Post: post})
}