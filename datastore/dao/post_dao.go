package dao

import "github.com/OmarTariq612/codersquare-go/types"

type PostDAO interface {
	ListPosts() []*types.Post
	CreatePost(port *types.Post) error
	GetPostByURL(url string) *types.Post // added
	GetPostByID(id string) *types.Post
	DeletePost(id string) error
}
