package handlers

import (
	"net/http"

	"github.com/OmarTariq612/codersquare-go/datastore"
	"github.com/OmarTariq612/codersquare-go/types"
	"github.com/OmarTariq612/codersquare-go/utils"
	"github.com/gin-gonic/gin"
)

type LikesHandler struct {
	db datastore.Database
}

func NewLikesHandler(db datastore.Database) LikesHandler {
	return LikesHandler{db: db}
}

func (l LikesHandler) List(c *gin.Context) {
	var likeData ListLikesRequest
	if errs := utils.BindUriVerifier(c, &likeData); errs != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": errs})
		return
	}

	likesCount, err := l.db.GetLikes(likeData.PostID)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, &ListLikesResponse{Likes: likesCount})
}

func (l LikesHandler) Create(c *gin.Context) {
	userID := c.GetString("user_id")
	var likeData CreateLikeRequest
	if errs := utils.BindUriVerifier(c, &likeData); errs != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": errs})
		return
	}

	if l.db.GetPostByID(likeData.PostID) == nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "the provided post id is not found"})
		return
	}

	like := &types.Like{UserID: userID, PostID: likeData.PostID}

	if l.db.Exists(like) {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "duplicate like"})
		return
	}

	if err := l.db.CreateLike(like); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusCreated)
}

func (l LikesHandler) Delete(c *gin.Context) {
	userID := c.GetString("user_id")
	var likeData DeleteLikeRequest
	if errs := utils.BindUriVerifier(c, &likeData); errs != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": errs})
		return
	}

	if l.db.GetPostByID(likeData.PostID) == nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "the provided post id is not found"})
		return
	}

	if err := l.db.DeleteLike(&types.Like{UserID: userID, PostID: likeData.PostID}); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
}
