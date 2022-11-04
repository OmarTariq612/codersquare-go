package dao

import "github.com/OmarTariq612/codersquare-go/types"

type CommentDAO interface {
	CreateComment(c *types.Comment) error
	ListComments(postID string) []*types.Comment
	DeleteComment(id string) error
	CountComments(postID string) (uint64, error)
	GetCommentByID(id string) *types.Comment // added
}
