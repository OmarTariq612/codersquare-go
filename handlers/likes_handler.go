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

func (lh LikesHandler) List(c *gin.Context) {
	var likeData ListLikesRequest
	if errs := utils.BindUriVerifier(c, &likeData); errs != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": errs})
		return
	}

	if lh.db.GetPost(likeData.PostID, "") == nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": types.ErrPostNotFound})
		return
	}

	likesCount, err := lh.db.GetLikes(likeData.PostID)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, &ListLikesResponse{Likes: likesCount})
}

func (lh LikesHandler) Create(c *gin.Context) {
	userID := c.GetString("user_id")
	var likeData CreateLikeRequest
	if errs := utils.BindUriVerifier(c, &likeData); errs != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": errs})
		return
	}

	if lh.db.GetPost(likeData.PostID, "") == nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": types.ErrPostNotFound})
		return
	}

	like := &types.Like{UserID: userID, PostID: likeData.PostID}

	if lh.db.Exists(like) {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": types.ErrDuplicateLike})
		return
	}

	if err := lh.db.CreateLike(like); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusCreated)
}

func (lh LikesHandler) Delete(c *gin.Context) {
	userID := c.GetString("user_id")
	var likeData DeleteLikeRequest
	if errs := utils.BindUriVerifier(c, &likeData); errs != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": errs})
		return
	}

	if lh.db.GetPost(likeData.PostID, "") == nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": types.ErrPostNotFound})
		return
	}

	like := &types.Like{UserID: userID, PostID: likeData.PostID}

	if !lh.db.Exists(like) {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": types.ErrLikeNotFound})
		return
	}

	if err := lh.db.DeleteLike(like); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
}
