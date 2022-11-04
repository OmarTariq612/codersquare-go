package handlers

import "github.com/OmarTariq612/codersquare-go/types"

// Users
type SignupRequest struct {
	Firstname string `json:"firstname" binding:"required"`
	Lastname  string `json:"lastname" binding:"required"`
	Username  string `json:"username" binding:"required"`
	Email     string `json:"email" binding:"required,email"`
	Password  string `json:"password" binding:"required"`
}

type SignupResponse struct {
	JWT string `json:"jwt"`
}

type SigninRequest struct {
	Login    string `json:"login" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type User struct {
	ID        string `json:"id"`
	Email     string `json:"email"`
	Username  string `json:"username"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

type SigninResponse struct {
	User `json:"user"`
	JWT  string `json:"jwt"`
}

type GetUserRequest struct {
	ID string `uri:"id" binding:"required,uuid"`
}

type GetUserResponse struct {
	ID        string `json:"id"`
	Username  string `json:"username"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

type GetCurrentUserResponse struct {
	User
}

// =====================================

// Posts
type ListPostsRequest struct {
}

type ListPostsResponse struct {
	Posts []*types.Post `json:"posts"`
}

type CreatePostRequest struct {
	Title string `json:"title" binding:"required"`
	URL   string `json:"url" binding:"required,url"`
	// UserID string `json:"user_id" binding:"required,uuid"`
}

type CreatePostResponse struct {
}

type GetPostRequest struct {
	PostID string `uri:"id" binding:"required,uuid"`
}

type GetPostResponse struct {
	Post *types.Post `json:"post"`
}

type DeletePostRequest struct {
	PostID string `uri:"id" binding:"required,uuid"`
}

type DeletePostResponse struct {
}

// =====================================

// Comments
type CreateCommentRequest struct {
	// UserID string `json:"user_id" binding:"required,uuid"`
	PostID string `uri:"id" binding:"required,uuid"`
	Text   string `json:"text" binding:"required"`
}

type CreateCommentResponse struct {
}

type DeleteCommentRequest struct {
	ID string `uri:"id" binding:"required,uuid"`
}

type DeleteCommentResponse struct {
}

type CountCommentsRequest struct {
	PostID string `uri:"id" binding:"required,uuid"`
}

type CountCommentsResponse struct {
	Count uint64 `json:"count"`
}

type ListCommentsRequest struct {
	PostID string `uri:"id" binding:"required,uuid"`
}

type ListCommentsResponse struct {
	Comments []*types.Comment `json:"comments"`
}

// Likes
type CreateLikeRequest struct {
	// UserID string `json:"user_id" binding:"required,uuid"`
	PostID string `uri:"post_id" binding:"required,uuid"`
}

type CreateLikeResponse struct {
}

type DeleteLikeRequest struct {
	PostID string `uri:"post_id" binding:"required,uuid"`
}

type DeleteLikeResponse struct {
}

type ListLikesRequest struct {
	PostID string `uri:"post_id" binding:"required,uuid"`
}

type ListLikesResponse struct {
	Likes uint64 `json:"likes"`
}
