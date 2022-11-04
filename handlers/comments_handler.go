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

type CommentsHandler struct {
	db datastore.Database
}

func NewCommentsHandler(db datastore.Database) CommentsHandler {
	return CommentsHandler{db: db}
}

func (ch CommentsHandler) List(c *gin.Context) {
	var commentData ListCommentsRequest
	if errs := utils.BindUriVerifier(c, &commentData); errs != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": errs})
		return
	}

	if ch.db.GetPostByID(commentData.PostID) == nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "the provided post id is not found"})
		return
	}

	comments := ch.db.ListComments(commentData.PostID)
	if comments == nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, &ListCommentsResponse{Comments: comments})
}

func (ch CommentsHandler) Create(c *gin.Context) {
	userID := c.GetString("user_id")
	var commentData CreateCommentRequest
	if errs := utils.BindUriVerifier(c, &commentData.PostID); errs != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": errs})
		return
	}
	if errs := utils.BindJsonVerifier(c, &commentData.Text); errs != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": errs})
		return
	}

	if ch.db.GetPostByID(commentData.PostID.PostID) == nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "the provided post id is not found"})
		return
	}

	if err := ch.db.CreateComment(&types.Comment{ID: uuid.NewString(), UserID: userID, PostID: commentData.PostID.PostID, Text: commentData.Text.Text, PostedAt: time.Now().Unix()}); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusCreated)
}

func (ch CommentsHandler) Delete(c *gin.Context) {
	userID := c.GetString("user_id")
	var commentData DeleteCommentRequest
	if errs := utils.BindUriVerifier(c, &commentData); errs != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": errs})
		return
	}

	comment := ch.db.GetCommentByID(commentData.ID)
	if comment == nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "the provided comment id is not found"})
		return
	}

	if comment.UserID != userID {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	if err := ch.db.DeleteComment(comment.ID); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
}

func (ch CommentsHandler) Count(c *gin.Context) {
	var commentData CountCommentsRequest
	if errs := utils.BindUriVerifier(c, &commentData); errs != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": errs})
		return
	}

	if ch.db.GetPostByID(commentData.PostID) == nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "the provided post id is not found"})
		return
	}

	commentsCount, err := ch.db.CountComments(commentData.PostID)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, &CountCommentsResponse{Count: commentsCount})
}
