package dao

import "github.com/OmarTariq612/codersquare-go/types"

type PostDAO interface {
	ListPosts(userID string) []*types.Post
	CreatePost(port *types.Post) error
	GetPostByURL(url string) *types.Post
	GetPost(postID, userID string) *types.Post
	DeletePost(id string) error
}
